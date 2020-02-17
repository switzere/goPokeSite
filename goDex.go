package main

import (
  "fmt"
  "net/http"
  "html/template"
  "os"
  //"net/url"
  //"time"
  //"flag"
  "log"
  // "strconv"
  "encoding/json"
  //"math"
  "strings"
  // "image"
  // "image/png"
  // "bytes"
  // "encoding/base64"

  "github.com/PuerkitoBio/goquery"
)

var (
  Pokedex []Pokemon
)

type temp struct{
  Pokedex []Pokemon
  Name string
}

type Pokemon struct {
  Sprite string `json:"sprite"`
  Image string `json:"image"`
  Thumb string `json:"thumb"`
	ID   int `json:"id"`
	Name struct {
		English  string `json:"english"`
		Japanese string `json:"japanese"`
		Chinese  string `json:"chinese"`
		French   string `json:"french"`
	} `json:"name"`
	Type []string `json:"type"`
  TypeString string `json:"typestring"`
  TypeA string `json:"typeA"`
  TypeB string `json:"typeB"`
	Base struct {
		HP        int `json:"HP"`
		Attack    int `json:"Attack"`
		Defense   int `json:"Defense"`
		SpAttack  int `json:"Sp. Attack"`
		SpDefense int `json:"Sp. Defense"`
		Speed     int `json:"Speed"`
	} `json:"base"`
}




type SearchResults struct {
  Status string `json:"status"`
  TotalResults int `json:"totalResults"`
  Pokemons []Pokemon `json:"pokemons"`
}

type Search struct{
  SearchKey string
  SearchResults SearchResults
}


func init(){
  fmt.Println("Website start")

  doc, _ := goquery.NewDocument("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/pokedex.json")

  //var metaDesc string
  //var pageTitle string

  pageBody := doc.Find("body").Contents().Text()

  fmt.Println(pageBody)

  json.Unmarshal([]byte(pageBody), &Pokedex)




  for i := 0; i < len(Pokedex); i++{

    pNum := fmt.Sprintf("%03d", i+1)
    fmt.Println(pNum)

    Pokedex[i].TypeString = strings.Join(Pokedex[i].Type, ", ")
    Pokedex[i].TypeA = Pokedex[i].Type[0]
    if(len(Pokedex[i].Type) > 1){
      Pokedex[i].TypeB = Pokedex[i].Type[1]
    }
    Pokedex[i].Sprite = "https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/sprites/" + pNum + "MS.png"
    Pokedex[i].Image = "https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/images/" + pNum + ".png"
    Pokedex[i].Thumb = "https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/thumbnails/" + pNum + ".png"


  }



  fmt.Println("Website loaded")


}



func main() {
  fmt.Println("Website starting")


	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := http.NewServeMux()

	 // Add the following two lines
	fs := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//mux.HandleFunc("/", index_handler)

  mux.HandleFunc("/", home_handler)
  mux.HandleFunc("/about/", about_handler)
  mux.HandleFunc("/pokedex/", dex_handler)
  mux.HandleFunc("/pokedex/Bulbasaur", b_handler)

	http.ListenAndServe(":"+port, mux)
}


func about_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Expert web design")
}

func b_handler(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Expert web design")
}


func home_handler(w http.ResponseWriter, r *http.Request){
  //now := time.Now() // find the time right now
//  HomePageVars := Pokemon{ //store the date and time in a struct
  //  ID: 1,
  //}

  t, err := template.ParseFiles("home.html") //parse the html file homepage.html

  for i := 0; i < 1; i++{

      if err != nil { // if there is an error
        log.Print("template parsing error: ", err) // log it
      }
      err = t.Execute(w, Pokedex[i]) //execute the template and pass it the HomePageVars struct to fill in the gaps
      if err != nil { // if there is an error
        log.Print("template executing error: ", err) //log it
      }
  }


}

func dex_handler(w http.ResponseWriter, r *http.Request){
  //fmt.Fprintf(w, "dex")

  t, err := template.ParseFiles("dex.html") //parse the html file homepage.html

  var T temp
  T.Pokedex = Pokedex

  for i := 0; i < 1; i++{

      if err != nil { // if there is an error
        log.Print("template parsing error: ", err) // log it
      }
      err = t.Execute(w, T) //execute the template and pass it the HomePageVars struct to fill in the gaps
      if err != nil { // if there is an error
        log.Print("template executing error: ", err) //log it
      }
  }

}
