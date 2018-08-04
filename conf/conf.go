package conf

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	"lheinrich.de/secpass/shorts"
)

var (
	// DB PostgreSQL
	DB *sql.DB

	// Config Map
	Config map[string]map[string]string

	// Lang Languages
	Lang = map[string]*map[string]string{}
)

// ReadConfig Unmarshal file to map
func ReadConfig(jsonFile string) map[string]map[string]string {
	// open config file and check for error
	file, err := os.Open(jsonFile)
	shorts.Check(err, true)

	// read config file and check for error
	jsonBytes, err := ioutil.ReadAll(file)
	shorts.Check(err, true)

	// unmarshal config to map and check for error
	var jsonConfig map[string]map[string]string
	shorts.Check(json.Unmarshal(jsonBytes, &jsonConfig), true)

	// return the config map
	return jsonConfig
}

// ReadLanguage Unmarshal file to map
func ReadLanguage(jsonFile string) *map[string]string {
	// open language file and check for error
	file, err := os.Open(jsonFile)
	shorts.Check(err, true)

	// read language file and check for error
	jsonBytes, err := ioutil.ReadAll(file)
	shorts.Check(err, true)

	// unmarshal language to map and check for error
	var jsonConfig map[string]string
	shorts.Check(json.Unmarshal(jsonBytes, &jsonConfig), true)

	// return the language map
	return &jsonConfig
}
