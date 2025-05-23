package config

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Configuration struct {
	AppName               string `json:"appName"`
	Port                  string `json:"port"`
	GCPProjectID          string `json:"gcpProjectId"`
	PathToAnthropicAPIKey string `json:"pathToAnthropicApiKey"`
}

// dev or prod
func buildConfiguration() Configuration {
	return Configuration{
		AppName:               os.Getenv("APP_NAME"),
		Port:                  os.Getenv("PORT"),
		GCPProjectID:          os.Getenv("GCP_PROJECT_ID"),
		PathToAnthropicAPIKey: os.Getenv("ANTHROPIC_API_KEY"),
	}
}
func Load(envFile *os.File) (Configuration, error) {
	defer envFile.Close()

	scanner := bufio.NewScanner(envFile)
	isFileEmpty := true

	for scanner.Scan() {
		// read the next line
		line := strings.TrimSpace(scanner.Text())
		// skip commented lines
		if strings.HasPrefix(line, "#") {
			continue
		}
		// we need to skip blank lines
		if line != "" {
			// if we are here then we have a line
			// and we do not have an empty file
			isFileEmpty = false
			// we need to make sure that there is an
			// "=" in the line
			if strings.Contains(line, "=") {
				parts := strings.SplitN(line, "=", 2)
				// ["APP_NAME", "goapi"]
				key := strings.TrimSpace(parts[0])
				value := strings.TrimSpace(parts[1])

				// read these into os env vars
				os.Setenv(key, value)
			}
		}
	}

	if isFileEmpty {
		return Configuration{}, fmt.Errorf("file is empty: %s", envFile.Name())
	}

	// by the time you get here you need to be done with envFile
	return buildConfiguration(), nil
}
