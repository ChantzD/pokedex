package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"io"
	"net/url"
	"strconv"
)

type locationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(config *config) error {
	res, err := http.Get(config.Next)	
	if err != nil {
		return err	
	}
	defer res.Body.Close()	

	areas := locationArea{}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, &areas)
	if err != nil {
		return err
	}

	for _, result := range areas.Results{	
		fmt.Println(result.Name)	
	}

	incOffset(config)

	return nil
}

func commandMapb(config *config) error {
	if config.Previous == "https://pokeapi.co/api/v2/location-area/?offset=0" {
		fmt.Println("you're on the first page")	
		return nil
	}

	decOffset(config)	

	res, err := http.Get(config.Previous)	
	if err != nil {
		return err	
	}
	defer res.Body.Close()	

	areas := locationArea{}
	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, &areas)
	if err != nil {
		return err
	}

	for _, result := range areas.Results{	
		fmt.Println(result.Name)	
	}

	return nil
}

func incOffset(config *config) error {
	config.Previous = config.Next 
	u, err := url.Parse(config.Next)
	if err != nil {
		return err
	}

	currentOffset := u.Query().Get("offset")
	offset := 0
	if currentOffset != "" {
		offset, _ = strconv.Atoi(currentOffset)
	}

	offset += 20

	q := u.Query()
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()

	config.Next = u.String()
	return nil
}

func decOffset(config *config) error {
	config.Next = config.Previous 
	u, err := url.Parse(config.Previous)
	if err != nil {
		return err
	}

	currentOffset := u.Query().Get("offset")
	offset := 0
	if currentOffset != "" {
		offset, _ = strconv.Atoi(currentOffset)
	}

	offset -= 20

	q := u.Query()
	q.Set("offset", strconv.Itoa(offset))
	u.RawQuery = q.Encode()

	config.Previous = u.String()

	return nil
}
