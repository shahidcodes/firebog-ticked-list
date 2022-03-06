package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const adFileName = "ads.txt"
const tmpFileName = "ads.tmp.txt"

func getExisitngUrlList() string {
	exisitingAdFileBuf, err := ioutil.ReadFile(adFileName)

	if err != nil {
		log.Fatal(err)
	}

	exisitingAdFile := string(exisitingAdFileBuf)
	return exisitingAdFile
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

			fmt.Println(url, listItem, urlSplit)

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
		fmt.Println(k)
		isSame = oldSet[k]

		if !isSame {
			break
		}
	}
	fmt.Println("is Same", isSame)
	return isSame
}
