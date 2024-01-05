package config

import "os"

var (
	API_SECRET = os.Getenv("API_SECRET")
)
