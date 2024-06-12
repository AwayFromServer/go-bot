package config

import (
	"encoding/json"
	"log"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

const BT = "botToken"
const BP = "botPrefix"

func GetConfig(filename string) ([]map[string]interface{}, error) {
	log.Printf("Working with config file: %s", filename)
	s := strings.Split(filename, ".")
	ext := s[len(s)-1] // last element should be the extension
	var m []map[string]interface{}

	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	switch {
	case ext == "yaml" || ext == "yml":
		var temp map[string]interface{}
		err = yaml.Unmarshal(f, &temp)
		m = append(m, temp)
	case ext == "json":
		err = json.Unmarshal(f, &m)
	}

	if err != nil {
		return nil, err
	}
	return m, nil
}

func GetOverrides(in []map[string]interface{}) []map[string]interface{} {
	bt, _ := os.LookupEnv(BT)
	bp, _ := os.LookupEnv(BP)
	out := in // just for legibility

	switch {
	case bt != "" && bp != "":
		log.Printf("Override detected for %s. New value hidden for safekeeping", BT)
		log.Printf("Override detected for %s. New value: %s", BP, bp)
		out[0][BT] = bt
		out[0][BP] = bp
	case bt != "":
		log.Printf("Override detected for %s. New value hidden for safekeeping", BT)
		out[0][BT] = bt
	case bp != "":
		log.Printf("Override detected for %s. New value: %s", BP, bp)
		out[0][BP] = bp
	}
	return out
}
