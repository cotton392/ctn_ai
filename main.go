package main

import (
	"fmt"

	"github.com/cotton392/ctn_ai/TwitterBot"
	"github.com/joho/godotenv"
)

func dotenvLoad(){
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file. err: %v" ,err)
	}
}

func main(){
	dotenvLoad()
	TwitterBot.TweetText()
}