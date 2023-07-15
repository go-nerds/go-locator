package main

import (
	"errors"
	"log"
	"net"
	"regexp"
	"strings"

	"github.com/oschwald/geoip2-golang"
)

func initReader(userInput UserInput) error {

	status, trimedIp := validateIP(userInput.Ip)

	if !status {
		return errors.New("Invalid ip address!")
	}

	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	ip := net.ParseIP(trimedIp)
	record, err := db.City(ip)
	if err != nil {
		log.Fatal(err)
		return errors.New("Something went wrong!")
	}

	output(userInput.Output, *record)

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
