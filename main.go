package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/google/go-github/v56/github"
	"github.com/joho/godotenv"
)

type RequestBody struct {
	NewOwner string `json:"new_owner"`
}

var wg sync.WaitGroup

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		log.Println("Insufficient arguments. Usage: go run main.go <owner> <new_owner>")
		return
	}

	owner := args[0]
	newOwner := args[1]

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Loading .env file failed: %v", err)
	}

	ownerPersonalToken := os.Getenv("OWNER_PERSONAL_TOKEN")

	client := github.NewClient(nil).WithAuthToken(ownerPersonalToken)

	data, err := os.ReadFile("repositories.txt")

	if err != nil {
		log.Fatalf("Reading file failed: %v", err)
	}

	repositoriesArray := strings.Split(string(data), "\n")

	for _, repoName := range repositoriesArray {
		go transferRepository(owner, repoName, newOwner, client)
		wg.Add(1)
	}

	wg.Wait()
}

func transferRepository(owner string, repoName string, newOwner string, client *github.Client) {
	defer wg.Done()
	_, response, err := client.Repositories.Transfer(
		context.Background(),
		owner,
		repoName,
		github.TransferRequest{NewOwner: newOwner},
	)

	if response.StatusCode != 202 && response.StatusCode != 200 {
		log.Printf("Failed to transfer %s repository: %v", repoName, err)
		return
	}

	fmt.Printf("Repository %s transfered successfuly\n", repoName)
}
