package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/toelsiba/comment"
)

func main() {
	exampleTrimm()
	exampleShield()
	exampleLoadConfig()
	exampleReadConfig()
	exampleSaveConfig()
}

func exampleTrimm() {

	text := `[
	"Golang", # this is comment
	"rune sharp is ##; this is not comment",
	"#### - two sharp"
	# "Rust" - line commented
]`

	data := comment.Trim([]byte(text))
	fmt.Println(string(data))

	var lines []string

	err := json.Unmarshal(data, &lines)
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range lines {
		fmt.Println(line)
	}
}

func exampleShield() {

	text := `{
	"username": "a#b",
	"password": "8h#sdl2##jkasd4325"
}`

	data := comment.Shield([]byte(text))
	fmt.Println("Shilded data:")
	fmt.Println(string(data))

	res := comment.Trim(data)
	fmt.Println("Trimmed data:")
	fmt.Println(string(res))
}

type Config struct {
	LogDir      string
	Streams     int
	Description string
	Servers     []Server
	Emails      []string
}

type Server struct {
	Id       int
	Address  string
	Password string
}

func exampleLoadConfig() {
	data, err := ioutil.ReadFile("config.cjson")
	if err != nil {
		log.Fatal(err)
	}

	data = comment.Trim(data)
	fmt.Println(string(data))

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(config)
}

func exampleReadConfig() {
	var c Config
	err := comment.ReadConfig("config.cjson", &c)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}

func exampleSaveConfig() {

	config := Config{
		LogDir:      "log_dir",
		Streams:     5,
		Description: "three sharp ### or 6",
		Servers: []Server{
			{
				Id:       500,
				Address:  "127.0.0.1",
				Password: "123!@#$%^##*()###^",
			},
		},
		Emails: []string{
			"email1@serty.rty",
			"wqrwetw@weq.56r",
			"#@#@.#@",
		},
	}

	data, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		log.Fatal(err)
	}

	data = comment.Shield(data)

	err = ioutil.WriteFile("config_dest.cjson", data, 0664)
	if err != nil {
		log.Fatal(err)
	}
}
