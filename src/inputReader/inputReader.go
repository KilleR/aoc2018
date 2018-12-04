package inputReader

import (
	"io/ioutil"
	"log"
	"strings"
)

func GetInput(path string) []string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalln("Failed to open input:", err)
	}

	return strings.Split(string(file), "\r\n")
}
