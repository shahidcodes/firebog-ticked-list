package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
	"time"
)

var BASE_PATH = fmt.Sprintf("https://api.github.com/repos/%s/%s", GH_OWNER, GH_REPO_NAME)
var CREATE_RELEASE_PATH = BASE_PATH + "/releases"

type CreateReleaseResponse struct {
	Url       string `json:"url"`
	UploadUrl string `json:"upload_url"`
}

func CreateTmpFile(content string) {
	err := os.WriteFile("./tmp.txt", []byte(content), 0644)
	CheckError(err)
}

func UploadFile(fileContent string, uploadUrl string) {
	url := uploadUrl + "?name=ads.txt&label=ads.txt"
	method := "POST"
	payload := &bytes.Buffer{}
	writer := multipart.NewWriter(payload)
	// file, errFile1 := os.Open("./tmp.txt")
	// CheckError(errFile1)
	// defer file.Close()
	part1, errFile1 := writer.CreateFormFile("tmp.txt", "tmp.txt")
	CheckError(errFile1)
	_, errFile1 = io.Copy(part1, strings.NewReader(fileContent))
	CheckError(errFile1)
	err := writer.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	CheckError(err)
	req.Header.Add("Authorization", "token "+GH_API_TOKEN)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	res, err := client.Do(req)
	CheckError(err)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	CheckError(err)
	fmt.Println(string(body))
}

func UploadToGithub(fileContent string) {
	now := time.Now()
	values := map[string]string{"tag_name": fmt.Sprintf("%d", now.Unix())}
	jsonValue, _ := json.Marshal(values)
	fmt.Println(CREATE_RELEASE_PATH)
	req, err := http.NewRequest("POST", CREATE_RELEASE_PATH, bytes.NewBuffer(jsonValue))
	CheckError(err)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "token "+GH_API_TOKEN)
	resp, err := http.DefaultClient.Do(req)
	CheckError(err)
	body, err := ioutil.ReadAll(resp.Body)
	CheckError(err)
	var data CreateReleaseResponse
	err = json.Unmarshal(body, &data)
	CheckError(err)
	uploadUrl := data.UploadUrl
	// uploadUrl := "https://uploads.github.com/repos/shahidcodes/firebog-ticked-list/releases/62197065/assets{?name,label}"
	uploadUrl = strings.Replace(uploadUrl, "{?name,label}", "", 1)
	CreateTmpFile(fileContent)
	UploadFile(fileContent, uploadUrl)
}
