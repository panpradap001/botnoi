package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// โครงสร้างข้อมูลที่รับจาก PokeAPI
type Stat struct {
	BaseStat int `json:"base_stat"`
	Effort   int `json:"effort"`
	StatInfo struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
}

type Sprites struct {
	BackDefault      *string `json:"back_default"`
	BackFemale       *string `json:"back_female"`
	BackShiny        *string `json:"back_shiny"`
	BackShinyFemale  *string `json:"back_shiny_female"`
	FrontDefault     *string `json:"front_default"`
	FrontFemale      *string `json:"front_female"`
	FrontShiny       *string `json:"front_shiny"`
	FrontShinyFemale *string `json:"front_shiny_female"`
}

// โครงสร้างข้อมูลที่รับจาก API
type PokemonAPIResponse struct {
	Stats   []Stat  `json:"stats"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

// ส่งไปclient
type Pokemon struct {
	Stats   []Stat  `json:"stats"`
	Name    string  `json:"name"`
	Sprites Sprites `json:"sprites"`
}

// PokeAPI
func fetchPokemon(id int) (*Pokemon, error) {
	pokemonURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%d/", id)
	formURL := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon-form/%d/", id)

	// Stats, Name
	resp, err := http.Get(pokemonURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	//ถอดรหัสjson
	var apiResponse PokemonAPIResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		return nil, err
	}

	var filteredStats []Stat
	for _, stat := range apiResponse.Stats {
		if stat.StatInfo.Name == "hp" || stat.StatInfo.Name == "attack" {
			filteredStats = append(filteredStats, stat)
		}
	}

	// Sprites
	resp, err = http.Get(formURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var spriteData struct {
		Sprites Sprites `json:"sprites"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&spriteData); err != nil {
		return nil, err
	}

	// Stats, Name, Sprites
	pokemon := Pokemon{
		Stats: filteredStats,
		Name:  apiResponse.Name,
		// URL:     pokemonURL,
		Sprites: spriteData.Sprites,
	}

	return &pokemon, nil
}

// API Handler
func getPokemonData(c *gin.Context) {
	var req struct {
		ID int `json:"id"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	pokemon, err := fetchPokemon(req.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch data"})
		return
	}

	c.JSON(http.StatusOK, pokemon)
}

func main() {
	r := gin.Default()
	r.POST("/pokemon", getPokemonData)

	fmt.Println("Server running at http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
