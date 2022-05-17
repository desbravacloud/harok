package database

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/tidwall/gjson"
)

func getDBCredentials() (dsn string, err error) {

	home := os.Getenv("HOME")

	content, err := ioutil.ReadFile(home + "/.harok/config.json") // the file is inside the local directory
	if err != nil {
		fmt.Println("Err")
		return dsn, err
	}
	dbuser := gjson.Get(string(content), "credentials.db_user").String()
	dbpassword := gjson.Get(string(content), "credentials.db_password").String()
	dbendpoint := gjson.Get(string(content), "credentials.db_endpoint").String()
	dbdatabase := gjson.Get(string(content), "credentials.db_database").String()
	dbport := gjson.Get(string(content), "credentials.db_port").String()

	dsn = "host=" + dbendpoint + " port=" + dbport + " dbname=" + dbdatabase + " user=" + dbuser + " password=" + dbpassword
	return dsn, err
}
