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
//
// func index_handler(w http.ResponseWriter, r * http.Request){
//   tempL.Execute(w,nil)
// }
//
// func search_handler(w http.ResponseWriter, r *http.Request){
//   u, err := url.Parse(r.URL.String())
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte("Internal server error"))
// 		return
// 	}
//
// 	params := u.Query()
// 	searchKey := params.Get("q")
// 	page := params.Get("page")
// 	if page == "" {
// 		page = "1"
// 	}
//
// 	//fmt.Println("Search Query is: ", searchKey)
// 	//fmt.Println("Results page is: ", page)
//   search := &Search{}
// 	search.SearchKey = searchKey
//
// 	if err != nil {
// 		http.Error(w, "Unexpected server error", http.StatusInternalServerError)
// 		return
// 	}
//
// 	endpoint := fmt.Sprintf("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/pokedex.json", url.QueryEscape(search.SearchKey))
// 	resp, err := http.Get(endpoint)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	defer resp.Body.Close()
//
// 	if resp.StatusCode != 200 {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
// 	err = json.NewDecoder(resp.Body).Decode(&search.SearchResults)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
//
//
//
//   err = tempL.Execute(w, search)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 	}
// }


func init(){
  fmt.Println("Website start")

  doc, _ := goquery.NewDocument("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/pokedex.json")

  //var metaDesc string
  //var pageTitle string

  pageBody := doc.Find("body").Contents().Text()

  fmt.Println(pageBody)

  json.Unmarshal([]byte(pageBody), &Pokedex)




  /*infile, _ := os.Open("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/sprites/005MS.png")

  defer infile.Close()


  // Decode will figure out what type of image is in the file on its own.
  // We just have to be sure all the image packages we want are imported.
  src, _, _ := image.Decode(infile)

  fmt.Println(infile)
  fmt.Println(src)

  Pokedex[0].Sprite = infile*/


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


    //sprt, _ := goquery.NewDocument("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/sprites/005MS.png")
    //spriteBody := sprt.Find("body").Contents().Text()

    //Pokedex[i].Sprite = []byte(spriteBody)
    //fmt.Println(spriteBody)
    //buffer := bufio.NewReader(sprt)
    //Pokedex[i].Sprite = buffer
  }

/*  resp, _ := http.Get("https://golangcode.com")
  fmt.Println(resp)

  	// Convert HTML into goquery document
  	doc, _ := goquery.NewDocumentFromReader(resp.Body)
    fmt.Println(doc)

  	// Save each .post-title as a list
  	titles := ""
  	doc.Find(".post-title").Each(func(i int, s *goquery.Selection) {
  		titles += "- " + s.Text() + "\n"
  	})


  fmt.Println(titles)*/

  /*byteValue, _ := ioutil.ReadAll("https://raw.githubusercontent.com/fanzeyi/pokemon.json/master/pokedex.json")

  json.Unmarshal(byteValue, Pokedex)*/

  fmt.Println("Website loaded")


}

// main.go

// beginning of the file

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
