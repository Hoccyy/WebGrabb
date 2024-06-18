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
	var webURL string
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

	openRes := exec.Command("open", "view.html").Start()
	if openRes != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("File opened successfully")
	}
}

func fetchData(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	fmt.Println(response.StatusCode)
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
