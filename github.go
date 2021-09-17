package main

import (
	"bytes"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var FILE_PATH = fmt.Sprintf("/repos/%s/%s/contents/ads.txt", GH_OWNER, GH_REPO_NAME)

type GitFileReq struct {
	Message string `json:"message"`
	Content string `json:"content"`
	SHA     string `json:"sha"`
}

type GitTree struct {
	PATH string `json:"path"`
	MODE string `json:"mode"`
	TYPE string `json:"type"`
	SHA  string `json:"sha"`
	SIZE int32  `json:"size"`
	URL  string `json:"url"`
}

type GitTreeResp struct {
	SHA  string    `json:"sha"`
	URL  string    `json:"url"`
	Tree []GitTree `json:"tree"`
}

func getFileSHA() string {
	req, _ := http.NewRequest("GET", fmt.Sprintf("https://api.github.com/repos/%s/%s/git/trees/HEAD", GH_OWNER, GH_REPO_NAME), nil)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+GH_API_TOKEN)
	resp, _ := http.DefaultClient.Do(req)
	responseObj := &GitTreeResp{}
	err := json.NewDecoder(resp.Body).Decode(&responseObj)
	if err != nil {
		log.Fatal(err)
	}
	var sha string
	for _, tree := range responseObj.Tree {
		if tree.PATH == "ads.txt" {
			sha = tree.SHA
			break
		}
	}
	return sha
}

func UploadToGithub(fileContent string) (bool, *http.Response) {
	sha := getFileSHA()

	gitFile := GitFileReq{
		Content: b64.StdEncoding.EncodeToString([]byte(fileContent)),
		Message: "UPDATE FILE",
		SHA:     sha,
	}

	b, err := json.Marshal(gitFile)
	if err != nil {
		log.Fatal("JSON Error: ", err)
	}
	fmt.Println("OK")

	br := bytes.NewBuffer(b)
	fmt.Println("Uploading...")
	req, err := http.NewRequest("PUT", "https://api.github.com"+FILE_PATH, br)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+GH_API_TOKEN)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal("HTTP Error: ", err)
	}

	return resp.StatusCode == 200, resp

}
