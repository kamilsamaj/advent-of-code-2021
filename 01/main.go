package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const cookieSessionFile = "../.cookie-session.txt"
const inputUrl = "https://adventofcode.com/2021/day/1/input"

func getSessionCookie(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(content), "\n"), nil
}

func getInput(url string) ([]byte, error) {
	session, err := getSessionCookie(cookieSessionFile)
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{
		Timeout: time.Second * 10,
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Cookie", session)
	resp, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	return io.ReadAll(resp.Body)
}

func countIncreases(input string) (increases int64, err error) {
	var prevNumber, currNumber int64
	for i, s := range strings.Split(input, "\n") {
		if s == "" {
			continue
		}
		if i == 0 {
			prevNumber, err = strconv.ParseInt(s, 10, 64)
			if err != nil {
				return 0, err
			}
			continue
		}
		currNumber, err = strconv.ParseInt(s, 10, 64)
		if err != nil {
			log.Fatalln(err)
		}
		if currNumber > prevNumber {
			increases++
		}
		prevNumber = currNumber
	}
	return increases, nil
}

func main() {
	var err error
	input, err := getInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	increases, err := countIncreases(string(input))

	fmt.Println("Increases: ", increases)
}
