package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Error loading .env file: %s\n", err)
		return
	}

	hostsEnv := os.Getenv("HOSTS")
	params := os.Getenv("PARAMS")
	methodEnv := os.Getenv("METHOD")
	if hostsEnv == "" || params == "" {
		fmt.Println("Check the env!")
		return
	}

	hosts := strings.Split(hostsEnv, ",")

	for {
		fmt.Println("============Starting to hit endpoints============")

		for _, host := range hosts {
			url := fmt.Sprintf("http://%s/channel/empty?%s", host, params)

			client := &http.Client{}
			req, err := http.NewRequest(methodEnv, url, nil)

			if err != nil {
				fmt.Println(err)
				return
			}
			res, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer res.Body.Close()

			fmt.Printf("\nResponse Code from %s : %v \n\n", url, res.StatusCode)
			fmt.Println("====================================")
		}

		fmt.Println("All endpoints hit. Waiting for 10 seconds...")
		time.Sleep(10 * time.Second)
	}
}
