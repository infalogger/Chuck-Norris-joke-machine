package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Joke struct {
	Value string `json:"value"`
}

func main() {
	fmt.Println("CHUCK NORRIS JOKE MACHINE")
	isfirstjoke := true

	for {
		if isfirstjoke {
			fmt.Println("Would you like to hear a Chuck Norris joke? (y/n): ")
		} else {
			fmt.Println("Would you like to hear another? (y/n): ")
		}

		var input string
		fmt.Scanln(&input)

		if input == "y" || input == "yes" {
			isfirstjoke = false
			reqjoke, _ := http.Get("https://api.chucknorris.io/jokes/random")
			defer reqjoke.Body.Close()
			body, _ := ioutil.ReadAll(reqjoke.Body)
			var joke Joke
			json.Unmarshal(body, &joke)
			fmt.Println(joke.Value)
		} else {
			fmt.Println("Well thats too bad goodbye!")
			break
		}
	}
}
