package shorts

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	// implement postgresql driver
	_ "github.com/lib/pq"
	"github.com/satori/go.uuid"
)

// LogInfo log only error or everything
var LogInfo = true

// InitLogging Enable logging to file with automatic error printing
func InitLogging(logInfo bool) {
	// set LogInfo
	LogInfo = logInfo

	// create directories
	os.MkdirAll("logs", os.ModePerm)

	// create new file and check for error
	file, err := os.Create("logs/" + time.Now().Format("2006-01-02.txt"))
	if err == nil {
		log.SetOutput(file)
	} else {
		fmt.Println(err)
	}
	file.Close()
}

// InitLoggingFile Enable logging to specified file
func InitLoggingFile(fileName string, logInfo bool) error {
	// set LogInfo
	LogInfo = logInfo

	// split directories from filename and create them
	directory, _ := path.Split(fileName)
	os.MkdirAll(directory, os.ModePerm)

	// create new file and check for error
	file, err := os.Create(fileName)
	if err == nil {
		log.SetOutput(file)
	}
	defer file.Close()

	return err
}

// UUID generate random uuid and return it as a string
func UUID() string {
	// generate uuid
	id, _ := uuid.NewV4()
	return id.String()
}

// Check Print error if exists
func Check(err error, isError bool) {
	// check whether error exists
	if err != nil {
		// check whether to print or not
		if isError || LogInfo == true {
			// print error
			log.Println(err)
		}
	}
}

// ConnectPostgreSQL Connection to postgresql database
func ConnectPostgreSQL(host, port, database, username, password string, ssl string) *sql.DB {
	// open database connection and check for errors
	db, err := sql.Open("postgres", "postgres://"+username+":"+password+"@"+host+"/"+database+"?sslmode="+ssl)
	Check(err, true)
	Check(db.Ping(), true)

	// return database connection
	return db
}
