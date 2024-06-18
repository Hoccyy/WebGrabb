package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var choice int
	var webURL string

	for {
		fmt.Print("\033[H\033[2J")
		fmt.Print("1. Save Website\n2. Quit\n\nEnter Choice\n> ")

		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("\033[H\033[2J")
			fmt.Print("Enter Web URL\n> ")

			if scanner.Scan() {
				webURL = scanner.Text()
			}

			data, err := fetchData(webURL)
			if err != nil {
				log.Fatalf("Error fetching data: %v", err)
			}

			file, err := os.Create("view.html")
			if err != nil {
				fmt.Println("File error")
			}

			defer file.Close()
			_, err = file.Write([]byte(data))
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}

			// Launch the downloaded HTML file in default browser
			openSavedPage := exec.Command("open", "view.html").Start()
			if openSavedPage != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("File opened successfully")
			}
		case 2:
			fmt.Print("\033[H\033[2J")
			return
		}

	}
}

func fetchData(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	fmt.Println("Response code: ", response.StatusCode)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
