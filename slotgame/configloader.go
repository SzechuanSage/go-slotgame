package slotgame

import (
	"encoding/json"
	"log"
	"os"
)

func ConfigLoader(file string) (map[string]interface{}, error) {
	content, err := os.ReadFile(file)
	if err != nil {
		log.Printf("Error when opening file: %v", err)
		return nil, err
	}

	var config map[string]interface{}
	err = json.Unmarshal(content, &config)
	if err != nil {
		log.Printf("Error during Unmarshal(): %v", err)
		return nil, err
	}

	return config, err
}
