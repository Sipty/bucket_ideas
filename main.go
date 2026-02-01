package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

const DATAPATH = "data.json"

func loadData() (buckets map[string][]string) {
	data, err := os.ReadFile("data.json")
	if err != nil {
		// File missing, create buckets manually:
		return make(map[string][]string)
	}

	err = json.Unmarshal(data, &buckets)
	if err != nil {
		fmt.Printf("Error unmarshalling data from disk-read-file. %v", err)
		os.Exit(0)
	}

	return buckets
}

func saveData(buckets map[string][]string) error {
	data, err := tidyBuckets(buckets)
	if err != nil {
		fmt.Printf("Well shite, we failed saving the data!")
	}

	printBuckets(buckets)

	os.WriteFile(DATAPATH, data, 0644)

	return nil
}

func tidyBuckets(buckets map[string][]string) ([]byte, error) {
	// Serialize data and return it nice and indented
	data, err := json.MarshalIndent(buckets, "", " ")
	if err != nil {
		fmt.Printf("Error trying to Marshal buckets in saveData! Error:%v", err)
		return nil, err
	}
	return data, nil
}

func printBuckets(buckets map[string][]string) error {
	data, err := tidyBuckets(buckets)
	if err != nil {
		fmt.Printf("err printing buckets")
		return err
	}
	fmt.Printf("Here's the data: \n%v\n", string(data))

	return nil
}

func handleInput(buckets map[string][]string) error {
	/*
		Organize input into buckets as follows:
			left arg is the bucket name
			right arg is the content.
		In case there is only one arg, idea goes in the default bucket.

		Happy path implementation:
		bi vid: talk about how awesome Go is

		TODO: reorg the arg handler to something cleaner.
	*/
	if len(os.Args) >= 3 {
		// NB: This is only the happy path, definitely Fix it up :)
		key := os.Args[1]
		val := strings.Join(os.Args[2:], " ")

		buckets[key] = append(buckets[key], val)
		return nil
	}

	if os.Args[1] == "view" {
		printBuckets(buckets)
		return errors.New("View invoked.")
	}
	// TODO add a help section :)

	return errors.New("Too few args king, supply at least one! :)")

}

func main() {
	buckets := loadData()

	err := handleInput(buckets)
	if err != nil {
		return // TODO: is this the best way to handle it?
	}

	saveData(buckets)

	fmt.Println("Nice, enjoy!")
}
