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
		log.Fatalf("Reading file failed: %v", err)
	}

	repositoriesArray := strings.Split(string(data), "\n")
	client := github.NewClient(nil).WithAuthToken(ownerPersonalToken)

	for _, repoName := range repositoriesArray {
		_, response, _ := client.Repositories.Transfer(
			context.Background(),
			owner,
			repoName,
			github.TransferRequest{NewOwner: newOwner},
		)

		if response.StatusCode != 202 || response.StatusCode != 200 {
			// log.Fatalf("Transfering repository failed: %v", err)
		}

		fmt.Printf("Repository %s transfered successfuly\n", repoName)
	}

}
