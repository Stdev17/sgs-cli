package main

import (
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func DownloadKubeconfig() {
	kubeconfigURL := "https://raw.githubusercontent.com/bacchus-snu/sgs-cli/main/config.yaml"
	destinationPath := filepath.Join(os.Getenv("HOME"), ".sgs", "config.yaml")

	// Check if the file already exists
	if _, err := os.Stat(destinationPath); os.IsNotExist(err) {
		// Create the directory if it doesn't exist
		err := os.MkdirAll(filepath.Dir(destinationPath), 0755)
		if err != nil {
			log.Fatalf("Failed to create directory: %v", err)
		}

		// Download the kubeconfig file
		response, err := http.Get(kubeconfigURL)
		if err != nil {
			log.Fatalf("Failed to download kubeconfig file: %v", err)
		}
		defer response.Body.Close()

		// Create the file
		file, err := os.Create(destinationPath)
		if err != nil {
			log.Fatalf("Failed to create file: %v", err)
		}
		defer file.Close()

		// Copy the response body to the file
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatalf("Failed to save kubeconfig file: %v", err)
		}

		log.Printf("Kubeconfig file downloaded successfully")
	} else {
		log.Printf("Kubeconfig file already exists")
	}
}
