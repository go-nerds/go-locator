package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/oschwald/geoip2-golang"
)

func output(fileName string, record geoip2.City) {
	if fileName != "" {
		err := writeToFile(fileName, record)

		if err != nil {
			log.Fatal(err)
		}

	} else {
		fmt.Printf("Country name: %v\n", record.Country.Names["en"])
		fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
		fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
		fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
		fmt.Printf("Accuracy Radius: %v\n", record.Location.AccuracyRadius)
		fmt.Printf("Metro Code: %v\n", record.Location.MetroCode)
		fmt.Printf("Is Anonymous Proxy: %v\n", record.Traits.IsAnonymousProxy)
		fmt.Printf("Is Satellite Provider: %v\n", record.Traits.IsSatelliteProvider)
		fmt.Printf("Is In European Union: %v\n", record.Country.IsInEuropeanUnion)
		fmt.Printf("========================================\n")
	}
}

func writeToFile(filename string, record geoip2.City) error {

	file, err := os.OpenFile(filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		return errors.New("Error while creating the file!")
	}

	defer file.Close()

	_, err2 := file.WriteString(
		"Country name: " + record.Country.Names["en"] + "\n" +
			"ISO country code: " + record.Country.IsoCode + "\n" +
			"Time zone: " + record.Location.TimeZone + "\n" +
			"Coordinates: " + fmt.Sprintf("%v", record.Location.Latitude) + " " + fmt.Sprintf("%v", record.Location.Longitude) + "\n" +
			"Accuracy Radius: " + fmt.Sprintf("%v", record.Location.AccuracyRadius) + "\n" +
			"Metro Code: " + fmt.Sprintf("%v", record.Location.MetroCode) + "\n" +
			"Is Anonymous Proxy: " + fmt.Sprintf("%v", record.Traits.IsAnonymousProxy) + "\n" +
			"Is Satellite Provider: " + fmt.Sprintf("%v", record.Traits.IsSatelliteProvider) + "\n" +
			"Is In European Union: " + fmt.Sprintf("%v", record.Country.IsInEuropeanUnion) + "\n" +
			"========================================\n")

	if err2 != nil {
		return errors.New("Error in writing into file")
	}

	return nil
}
