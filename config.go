package comment

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

func ReadConfig(filename string, config interface{}) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	data = Trim(data)
	return json.Unmarshal(data, config)
}

func WriteConfig(filename string, config interface{}) error {
	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return err
	}
	data = Shield(data)
	return ioutil.WriteFile(filename, data, os.ModePerm)
}
