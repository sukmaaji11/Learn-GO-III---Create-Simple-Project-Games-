package main

import (
	"encoding/json"
	"html/template"
	"learngo/rps"
	"log"
	"net/http"
	"strconv"
)

func main() {
	//Create a Basic Hello World
	//fmt.Println("Hello World")

	//VARIABLE

	//Boolean
	//var temp bool
	//fmt.Println(temp)

	//Integer
	//var temp int = 5
	//var temp int16 = 5
	//fmt.Println(temp)

	//Float
	//var temp float32 = 1.65
	//fmt.Println(temp)

	//String
	//var temp string = "Hello"
	//fmt.Println(temp)

	//Multiple Variable
	//x, y := 5, 10
	//fmt.Println(x)
	//fmt.Println(y)
	//var (
	//	x = 5
	//	y = 10.5
	//	z = "Hai"
	//)
	//fmt.Println(x)
	//fmt.Println(y)
	//fmt.Println(z)

	//Constant
	//const temp = 5
	//temp := 4
	//fmt.Println(temp)

	//Array
	//var temp = [3]int{1, 2, 3}
	//fmt.Println(temp)

	//Slices
	//var temp []int
	//var num = []int{10, 20, 30}

	//temp = append(temp, num...)
	//fmt.Println(temp)

	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
 	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("assets"))))
	
	//Starting Web Server
	log.Println("Starting web server on port 8080")
	http.ListenAndServe(":8080", nil)
}

// Render HTML File
func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Println(err)
		return
	}
}

// RPS Games Play
func playRound(w http.ResponseWriter, r *http.Request) {
	playerChoice, _ := strconv.Atoi(r.URL.Query().Get("c"))
	result := rps.PlayRound(playerChoice)

	out, err := json.MarshalIndent(result, "", "    ")

	if err != nil {
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}
