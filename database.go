package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"strconv"
)

// Database of the API
type Database struct {
	jsonMap map[string]string
}

// Initialize starts the database
func (d *Database) Initialize() {
	d.jsonMap = make(map[string]string)
	data, err := ioutil.ReadFile("database.json")

	if err != nil {
		log.Println(err)
	}

	json.Unmarshal(data, &d.jsonMap)
}

// GetAppVersion returns the application version
func (d *Database) GetAppVersion(appType string) (string, error) {
	switch appType {
	case "ios":
		return d.jsonMap["iosVersion"], nil
	case "android":
		return d.jsonMap["androidVersion"], nil
	default:
		return "", errors.New("Application not available")
	}
}

// GetForceUpdate returns the application force update
func (d *Database) GetForceUpdate(appType string) (bool, error) {
	switch appType {
	case "ios":
		b := d.jsonMap["iosForceUpdate"] == "true"
		return b, nil
	case "android":
		b, err := strconv.ParseBool(d.jsonMap["androidForceUpdate"])
		return b, err
	default:
		return false, errors.New("Application not available")
	}
}

// GetURL returns the application URL
func (d *Database) GetURL(appType string) (string, error) {
	switch appType {
	case "ios":
		return d.jsonMap["iosURL"], nil
	case "android":
		return d.jsonMap["androidURL"], nil
	default:
		return "", errors.New("Application not available")
	}
}
