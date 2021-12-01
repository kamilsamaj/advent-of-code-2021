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

func deleteEmptyItems(src []string) []string {
	var result []string
	for _, str := range src {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func slidingWindowIncreases(input string) (increases int64, err error) {
	var prevWindowSum, currWindowSum int64
	lines := deleteEmptyItems(strings.Split(input, "\n"))

	if len(lines) < 4 {
		return -1, fmt.Errorf("cannot create at least 2 sliding windows")
	}

	for i := 3; i < len(lines); i++ {
		for j := 0; j < 3; j++ {
			a, err := strconv.ParseInt(lines[i-j], 10, 64)
			if err != nil {
				return -1, err
			}
			b, err := strconv.ParseInt(lines[i-j-1], 10, 64)
			if err != nil {
				return -1, err
			}
			currWindowSum += a
			prevWindowSum += b
		}
		if currWindowSum > prevWindowSum {
			increases++
		}
		currWindowSum = 0
		prevWindowSum = 0
	}
	return increases, nil
}

func main() {
	var err error
	input, err := getInput(inputUrl)
	if err != nil {
		log.Fatalln(err)
	}

	// count increases per each measurement
	increases, err := countIncreases(string(input))
	fmt.Println("Regular increases: ", increases)

	// count increases in the sliding windows of size 3
	slidingWinIncreases, err := slidingWindowIncreases(string(input))
	fmt.Println("Sliding windows (3 measurements) increases: ", slidingWinIncreases)
}
