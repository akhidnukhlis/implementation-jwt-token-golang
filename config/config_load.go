package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
	"playground/implementation-jwt-token-golang/helpers"
)

type Config struct {
	PrivateKey string `yaml:"PRIVATE_KEY"`
	PublicKey  string `yaml:"PUBLIC_KEY"`
}

// ReadConfigFromFile is function to read file credential
func ReadConfigFromFile() (*Config, error) {
	var credentials Config

	dir := helpers.GetCurrentDirectory()
	filePath := filepath.Join(dir, "config/config.yml")

	// Read file YAML
	yamlFile, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("gagal membaca file YAML: %s", err)
	}

	// Parsing data YAML
	err = yaml.Unmarshal(yamlFile, &credentials)
	if err != nil {
		log.Fatalf("gagal membaca data file YAML: %s", err)
	}

	return &credentials, nil
}
