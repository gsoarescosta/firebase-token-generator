package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gsoarescosta/firebase-token-generator/internal/firebase"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	ctx := context.Background()
	token := ""
	client := firebase.Client{Credential: os.Getenv("FIREBASE_ADMIN_SDK_BASE64_CREDENTIAL")}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Type 1 for phone number, 2 for email: ")
	for {
		scanner.Scan()

		switch scanner.Text() {
		case "1":
			fmt.Print("Type the user phone number: ")
			scanner.Scan()
			phoneNumber := scanner.Text()
			token, _ = client.GenerateCustomToken(ctx, phoneNumber, "phoneNumber")
		case "2":
			fmt.Print("Type the user email: ")
			scanner.Scan()
			email := scanner.Text()
			token, _ = client.GenerateCustomToken(ctx, email, "email")
		default:
			fmt.Print("Invalid input. Type 1 for phone number, 2 for email: ")
		}

		if token != "" {
			fmt.Println("Token: ", token)
			fmt.Println("----------------------------------------------")
			fmt.Print("Do you want to generate another token?\nType 1 for phone number, 2 for email: ")
		}
	}

}
func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
}
