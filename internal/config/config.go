package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url string 	`json:"db_url"`
	Current_user_name string 	`json:"current_user_name"`
}

func Read() Config {
	file, err := getConfigFilePath()
	if err != nil {
		return Config{}
	}

	confFile, err := os.Open(file)
	if err != nil {
		fmt.Println("Could not open file")
		return Config{}
	}
	defer confFile.Close()
	var conf Config
	
	decoder := json.NewDecoder(confFile)
	err = decoder.Decode(&conf)
	if err != nil {
		fmt.Println("Error decoding JSON")
		return Config{}
	}
	return conf
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Problem getting home directory: %v", err)
		return "", err
	}

	fileName := filepath.Join(home, configFileName)
	return fileName, nil

}

func (conf Config) SetUser(user string) error {
	conf.Current_user_name = user

	// write conf to json file
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}

	confFile, err := os.Open(file)
	if err != nil {
		return err
	}
	defer confFile.Close()

	data, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	err = os.WriteFile(file, data, 0644)
	return nil
}
