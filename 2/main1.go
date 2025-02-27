package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// โครงสร้างข้อมูลที่ใช้สำหรับ API Response
type PokemonResponse struct {
	Stats   []Stat  `json:"stats"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"prites"`
}

type Stat struct {
	BaseStat int      `json:"base_stat"`
	Effort   int      `json:"effort"`
	StatInfo StatInfo `json:"stat"`
}

type StatInfo struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Sprites struct {
	BackDefault      string `json:"back_default"`
	BackFemale       string `json:"back_female"`
	BackShiny        string `json:"back_shiny"`
	BackShinyFemale  string `json:"back_shiny_female"`
	FrontDefault     string `json:"front_default"`
	FrontFemale      string `json:"front_female"`
	FrontShiny       string `json:"front_shiny"`
	FrontShinyFemale string `json:"front_shiny_female"`
}

// ฟังก์ชันสำหรับดึงข้อมูลจาก API
func fetchJSON(url string, target interface{}) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("API returned status code: %d", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(target)
}

func getPokemonData(c *gin.Context) {
	var requestData struct {
		ID int `json:"id"`
	}

	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON request"})
		return
	}

	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", requestData.ID)
	formURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d/", requestData.ID)

	var pokemonData struct {
		Stats []Stat `json:"stats"`
	}
	var formData struct {
		Name    string  `json:"name"`
		Sprites Sprites `json:"sprites"`
	}

	// Fetch Data
	if err := fetchJSON(pokemonURL, &pokemonData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Pokemon data"})
		return
	}
	if err := fetchJSON(formURL, &formData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch Pokemon Form data"})
		return
	}

	// สร้าง JSON Response
	response := PokemonResponse{
		Stats:   pokemonData.Stats,
		Name:    formData.Name,
		Sprites: formData.Sprites,
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	r := gin.Default()
	r.POST("/pokemon", getPokemonData)

	// เรียกใช้ API ที่ Port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
