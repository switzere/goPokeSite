package main

import (
  "fmt"
  //"net/http"
  //"html/template"
  //"os"
  //"net/url"
  //"time"
  //"flag"
  //"log"
  "strconv"
  "encoding/json"
  //"math"
  //"strings"
  // "image"
  // "image/png"
  // "bytes"
  // "encoding/base64"

  //"github.com/PuerkitoBio/goquery"
  "github.com/mtslzr/pokeapi-go"
)

var (
  Pokedex []Pokemon
)

type temp struct{
  Pokedex []Pokemon
  Name string
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Forms          []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices            []interface{} `json:"game_indices"`
	Height                 int           `json:"height"`
	HeldItems              []interface{} `json:"held_items"`
	ID                     int           `json:"id"`
	IsDefault              bool          `json:"is_default"`
	LocationAreaEncounters string        `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name    string `json:"name"`
	Order   int    `json:"order"`
	Species struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      interface{} `json:"back_default"`
		BackFemale       interface{} `json:"back_female"`
		BackShiny        interface{} `json:"back_shiny"`
		BackShinyFemale  interface{} `json:"back_shiny_female"`
		FrontDefault     interface{} `json:"front_default"`
		FrontFemale      interface{} `json:"front_female"`
		FrontShiny       interface{} `json:"front_shiny"`
		FrontShinyFemale interface{} `json:"front_shiny_female"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func main(){
  fmt.Println("Starting...")

  for i := 1; i <= 807; i++{
    Pokedex[i-1].Name = " "
    temp, _ := pokeapi.Pokemon(strconv.Itoa(i))
    tempString, _ := json.Marshal(temp)
    json.Unmarshal([]byte(tempString), &Pokedex[i-1])
    fmt.Println(Pokedex[i-1].Name)
  }
}
