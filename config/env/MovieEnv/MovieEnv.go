package MovieEnv

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

//Parameter models
type Parameter struct {
	Trxid  string
	URL    string `yaml:"url"`
	Apikey string `yaml:"apikey"`
}

//function to get env paramater
func (e *Parameter) GetParameter() error {

	yamlFile, err := ioutil.ReadFile("../config/env/MovieEnv/Movie.yaml")
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(yamlFile, e)
	if err != nil {
		return err
	}

	return nil
}
