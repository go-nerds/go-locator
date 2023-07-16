package main

import (
	"bufio"
	"errors"
	"log"
	"net"
	"os"
	"regexp"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

var fileLines []string

func initReader(userInput UserInput) error {

	if userInput.Ip == "" && userInput.File == "" {
		return errors.New("Please choose an option!")
	}

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if userInput.Ip != "" {
		readFromDb(db, userInput)
	}

	if userInput.File != "" {

		readFile(userInput)

		for _, line := range fileLines {
			userInput.Ip = line
			readFromDb(db, userInput)
		}

	}

	return nil
}

func validateIP(ip string) (bool, string) {
	ip = strings.Trim(ip, " ")
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true, ip
	}
	return false, ip
}

func readFile(userInput UserInput) error {
	file := userInput.File

	f, err := os.Open(file)

	if err != nil {
		return errors.New("Something went wrong while reading the given file!")
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fileLines = append(fileLines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return errors.New("Error in scanner!")
	}

	return nil

}

func readFromDb(db *geoip2.Reader, userInput UserInput) error {
	status, trimedIp := validateIP(userInput.Ip)

	if !status {
		return errors.New("Invalid ip address!")
	}

	ip := net.ParseIP(trimedIp)
	record, err := db.City(ip)
	if err != nil {
		return errors.New("Something went wrong!")
	}

	output(userInput.Output, *record)

	return nil
}
