package httpclient

import (
	"encoding/json"
	"os"
	"shak-daemon/models"
)

func GetLatestSpec(configPath string, spec *models.Spec) {
	configJson, err := os.ReadFile(configPath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configJson, &spec)
	if err != nil {
		panic(err)
	}
}
