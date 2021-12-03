package internal

import (
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

const cookieSessionFile = ".cookie-session.txt"

func getSessionCookie(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return strings.TrimSuffix(string(content), "\n"), nil
}

func GetInput(url string) ([]byte, error) {
	session, err := getSessionCookie(cookieSessionFile)
	if err != nil {
		return []byte{}, err
	}
	client := &http.Client{
		Timeout: time.Second * 30,
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
