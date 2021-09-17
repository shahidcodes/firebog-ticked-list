package main

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

var GH_OWNER = os.Getenv("GH_OWNER")
var GH_REPO_NAME = os.Getenv("GH_REPO")
var GH_API_TOKEN = os.Getenv("GH_API_TOKEN")
