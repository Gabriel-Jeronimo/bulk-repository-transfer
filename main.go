package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v56/github"
	"github.com/joho/godotenv"
)

type RequestBody struct {
	NewOwner string `json:"new_owner"`
}

func main() {
	args := os.Args[1:]

	owner := args[0]
	newOwner := args[1]

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Loading .env file failed: %v", err)
	}

	ownerPersonalToken := os.Getenv("OWNER_PERSONAL_TOKEN")

	data, err := os.ReadFile("repositories.txt")

	if err != nil {
		log.Fatal("Reading file failed: %v", err)
	}

	repositoriesArray := strings.Split(string(data), "\n")
	client := github.NewClient(nil).WithAuthToken(ownerPersonalToken)
	for _, repoName := range repositoriesArray {
		repo, _, err := client.Repositories.Transfer(
			context.Background(),
			owner,
			repoName,
			github.TransferRequest{NewOwner: newOwner},
		)

		if err != nil {
			log.Fatal("Transfering repository failed: %v", err)
		}

		fmt.Printf("%v", repo)
	}

}
