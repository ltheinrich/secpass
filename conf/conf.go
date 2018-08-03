package conf

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"os"

	"lheinrich.de/extgo/shorts"
)

var (
	// DB PostgreSQL
	DB *sql.DB
	// Config Map
	Config map[string]map[string]string
)

// ReadConfig Unmarshal file to JSONConfig
func ReadConfig(jsonFile string) map[string]map[string]string {
	// open config file and check for error
	file, err := os.Open(jsonFile)
	shorts.Check(err)

	// read config file and check for error
	jsonBytes, err := ioutil.ReadAll(file)
	shorts.Check(err)

	// unmarshal config to map and check for error
	var jsonConfig map[string]map[string]string
	shorts.Check(json.Unmarshal(jsonBytes, &jsonConfig))

	// return the config map
	return jsonConfig
}
