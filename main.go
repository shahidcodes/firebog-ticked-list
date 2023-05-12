package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var LIST_URL = "https://v.firebog.net/hosts/lists.php?type=tick"

func readReader(r io.ReadCloser) []string {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanRunes)
	var buf bytes.Buffer
	for scanner.Scan() {
		buf.WriteString(scanner.Text())
	}

	listStr := buf.String()

	urlSlice := strings.Split(listStr, "\n")
	return urlSlice
}

func readUrl(url string) []string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	fmt.Println("fetch complete", url)

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("ioutils failed, trying to read from buffer")
		urlSlice2 := readReader(resp.Body)
		return urlSlice2
	}

	listStr := string(body)
	urlSlice := strings.Split(listStr, "\n")
	fmt.Println("parsed", url)

	return urlSlice

}

func main() {
	// UploadToGithub("this is me")
	urlSlice := readUrl(LIST_URL)
	urlSlice = urlSlice[:len(urlSlice)-1] // removes last empty line
	var urls []string
	for _, url := range urlSlice {
		fmt.Println("fetching", url)
		blockedUrls := readUrl(url)
		urls = append(urls, blockedUrls...)
	}
	fmt.Println("all urls are fetched")
	urlsJoined := strings.Join(urls, "\n")

	if !IsBlockListSame(urlsJoined) {
		UploadToGithub(urlsJoined)
	}
}
