package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const ApiUri = "http://api.openweathermap.org/data/2.5/weather"

type Result struct {
	Name string
	Main Main
	Sys  Sys
}

type Main struct {
	Temp     float64
	Humidity float64
}

type Sys struct {
	Country string
}

func main() {
	client := &http.Client{}

	apiKey := flag.String("key", "", "Provide a valid https://openweathermap.org/api API key")
	flag.Parse()

	// Validating inputs presence. Present -h if required argument is not given.
	if *apiKey == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	input := strings.Join(flag.Args(), " ")

	if input == "" {
		fmt.Println("Please, provide the name of the city.")
		os.Exit(1)
	}

	request, err := http.NewRequest("GET", ApiUri, nil)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	// Building query params
	query := request.URL.Query()
	query.Add("appid", *apiKey)
	query.Add("q", input)
	query.Add("units", "metric")

	// Encoding silly data coming from user input as well
	request.URL.RawQuery = query.Encode()

	// Explicitly building headers
	request.Header.Add("Accept", "application/json")

	// Making the actual request and "scheduling" body closure
	response, err := client.Do(request)
	defer response.Body.Close()

	responseStruct := Result{}

	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {
		body, _ := ioutil.ReadAll(response.Body)

		// Changes &responseBody struct with body data
		jsonErr := json.Unmarshal(body, &responseStruct)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}
	}

	message := fmt.Sprintf("It's %v°C right now in %v, %v!", responseStruct.Main.Temp, responseStruct.Name, responseStruct.Sys.Country)
	fmt.Println(message)
}
