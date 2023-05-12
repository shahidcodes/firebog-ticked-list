package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

type GitRelease []struct {
	URL       string `json:"url"`
	AssetsURL string `json:"assets_url"`
	UploadURL string `json:"upload_url"`
	HTMLURL   string `json:"html_url"`
	ID        int    `json:"id"`
	Author    struct {
		Login             string `json:"login"`
		ID                int    `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"author"`
	NodeID          string      `json:"node_id"`
	TagName         string      `json:"tag_name"`
	TargetCommitish string      `json:"target_commitish"`
	Name            interface{} `json:"name"`
	Draft           bool        `json:"draft"`
	Prerelease      bool        `json:"prerelease"`
	CreatedAt       time.Time   `json:"created_at"`
	PublishedAt     time.Time   `json:"published_at"`
	Assets          []struct {
		URL      string `json:"url"`
		ID       int    `json:"id"`
		NodeID   string `json:"node_id"`
		Name     string `json:"name"`
		Label    string `json:"label"`
		Uploader struct {
			Login             string `json:"login"`
			ID                int    `json:"id"`
			NodeID            string `json:"node_id"`
			AvatarURL         string `json:"avatar_url"`
			GravatarID        string `json:"gravatar_id"`
			URL               string `json:"url"`
			HTMLURL           string `json:"html_url"`
			FollowersURL      string `json:"followers_url"`
			FollowingURL      string `json:"following_url"`
			GistsURL          string `json:"gists_url"`
			StarredURL        string `json:"starred_url"`
			SubscriptionsURL  string `json:"subscriptions_url"`
			OrganizationsURL  string `json:"organizations_url"`
			ReposURL          string `json:"repos_url"`
			EventsURL         string `json:"events_url"`
			ReceivedEventsURL string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"uploader"`
		ContentType        string    `json:"content_type"`
		State              string    `json:"state"`
		Size               int       `json:"size"`
		DownloadCount      int       `json:"download_count"`
		CreatedAt          time.Time `json:"created_at"`
		UpdatedAt          time.Time `json:"updated_at"`
		BrowserDownloadURL string    `json:"browser_download_url"`
	} `json:"assets"`
	TarballURL string      `json:"tarball_url"`
	ZipballURL string      `json:"zipball_url"`
	Body       interface{} `json:"body"`
}

func getExisitngUrlList() string {
	var release GitRelease
	r, err := http.Get("https://api.github.com/repos/shahidcodes/firebog-ticked-list/releases")
	if err != nil {
		log.Fatal(err)
	}
	defer r.Body.Close()
	jsonBytes, err := io.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(jsonBytes, &release)
	latestUpdate := release[0].Assets[0]
	existingAdsList := GetText(latestUpdate.BrowserDownloadURL)
	return existingAdsList
}

func makeSetFromList(list string) map[string]bool {
	set := make(map[string]bool)
	listSplit := strings.Split(list, "\n")

	for _, listItem := range listSplit {
		if listItem != "" && !strings.HasPrefix(listItem, "#") {
			var url string
			urlSplit := strings.Split(listItem, " ")
			// handle string - 0.0.0.0 somewebsite
			if strings.HasPrefix("0.0.0", listItem) {
				url = urlSplit[1]
			} else {
				// handle strings
				// 1- googel.com
				// 2- google.com # google
				url = urlSplit[0]
			}

			// fmt.Println(url, listItem, urlSplit)

			set[url] = true
		}
	}

	return set
}

func IsBlockListSame(newUrlList string) bool {
	oldSet := makeSetFromList(getExisitngUrlList())
	newSet := makeSetFromList(newUrlList)

	isSame := true

	for k, _ := range newSet {
		// fmt.Println(k)
		isSame = oldSet[k]

		if !isSame {
			break
		}
	}
	fmt.Println("is Same", isSame)
	return isSame
}
