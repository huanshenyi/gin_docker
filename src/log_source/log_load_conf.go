package log_source

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type LogConfig struct {
	LogDir          string `json:"log_dir"`
	LogLevel        string `json:"log_level"`
	LogOutPutToFile bool   `json:"log_out_put_to_file"`
}

func LoadLogConfig() *LogConfig {
	log_conf := LogConfig{}
	file, err := os.Open("conf/log_conf.json")
	if err != nil {
		file, err = os.Open("../../conf/log_conf.json")
		if err != nil {
			panic(err)
		}
	}
	defer file.Close()

	byte_data, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byte_data, &log_conf)
	if err != nil {
		panic(err)
	}

	return &log_conf
}
