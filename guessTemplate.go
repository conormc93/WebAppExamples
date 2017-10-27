package main

import (
	"html/template"
	"net/http"
	"math/rand"
	"time"
	"strconv"
)
	
type myMessage struct {
	Message string
	Result bool
	YourGuess string
	
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	message :="Guess a Number between 1 and 20"
	Result:=false
	yourGuess:= r.FormValue("guess")

	rand.Seed(time.Now().UTC().UnixNano())

	targetNumber:=rand.Intn(20-1)
	count :=0

	var cookie, err = r.Cookie("targetNumber")
	var cookie2, err2 = r.Cookie("count")

	if err != nil{
		cookie = &http.Cookie{
			Name: "targetNumber",
			Value: strconv.Itoa(targetNumber),
			Expires: time.Now().Add(24 * time.Hour),
		}
		//set the cookie
		http.SetCookie(w,cookie)	
	}	
	if err2 == nil{
		//convert value to an int
		count, _ = strconv.Atoi(cookie2.Value)
	}
	//convert value to an int
	targetNumber, _ = strconv.Atoi(cookie.Value)

	userGuess, _ := strconv.Atoi(yourGuess)

	//compare YourGuess to random number
	if userGuess == targetNumber{
		count=0
		Result=true
		message ="Correct Guess "+ yourGuess+" was the answer"
		cookie = &http.Cookie{
			Name: "targetNumber",
			Value: strconv.Itoa(rand.Intn(20-1)),
			Expires: time.Now().Add(24 * time.Hour),
		}	
		//set the cookie
		http.SetCookie(w,cookie)
	
	}else if userGuess < targetNumber {
	   message="Try Again your guess  was  too low"
		count ++
		
	}else if userGuess > targetNumber {
		message="Try Again your guess was too high"
		count ++
	}

	cookie2 = &http.Cookie{
		Name: "count",
		Value: strconv.Itoa(count),
		Expires: time.Now().Add(24 * time.Hour),
	}
	http.SetCookie(w,cookie2)

	//read the contents of guess.html and return a template
	t, _ := template.ParseFiles("guess.tmpl")

	//execute template and pass pointer to myMessage struct
	t.Execute(w, &myMessage{Message:message,YourGuess:yourGuess,Result:Result})
}

func main() {

	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}
