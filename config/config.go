package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Env  string `json:"env"`
	Port int    `json:"port"`
}

func New(path string) (Config, error) {
	file, err := os.Open(path)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}

	return cfg, nil
}
