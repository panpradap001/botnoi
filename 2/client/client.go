package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// รับจาก API
type PokemonResponse struct {
	Stats []Stat `json:"stats"`
	Name  string `json:"name"`
	// URL     string  `json:"url"`
	Sprites Sprites `json:"sprites"`
}

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

func main() {
	// รับค่า id
	var id int
	fmt.Print("Enter Pokémon ID: ")
	fmt.Scan(&id)

	requestBody, _ := json.Marshal(map[string]int{"id": id})

	resp, err := http.Post("http://localhost:8080/pokemon", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	// อ่าน Response
	body, _ := io.ReadAll(resp.Body)

	var formattedJSON bytes.Buffer
	json.Indent(&formattedJSON, body, "", "    ")

	fileName := fmt.Sprintf("pokemon_%d.json", id)
	err = os.WriteFile(fileName, formattedJSON.Bytes(), 0644)
	if err != nil {
		fmt.Println("Error saving JSON:", err)
		return
	}

	fmt.Println("JSON saved as", fileName)
}
