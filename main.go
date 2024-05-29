package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/generative-ai-go/genai"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
	GeminiApiKey := os.Getenv("GEMINI_API_KEY")
	ctx := context.Background()
	
	client, err := genai.NewClient(ctx, option.WithAPIKey(GeminiApiKey))
	if err != nil {
	log.Fatal(err)
	}
	defer client.Close()

	// The Gemini 1.5 models are versatile and work with most use cases
	model := client.GenerativeModel("gemini-1.5-flash")

	resp, err := model.GenerateContent(ctx, genai.Text("Write a story about a magic backpack."))
	if err != nil {
	log.Fatal(err)
	}
	printResponse(resp)
}

func printResponse(resp *genai.GenerateContentResponse) {
	for _, cand := range resp.Candidates {
		if cand.Content != nil {
			for _, part := range cand.Content.Parts {
				fmt.Println(part)
			}
		}
	}
	fmt.Println("---")
}