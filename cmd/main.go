package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
	"github.com/ndewijer/go-temperature-rug/internal/handlers"
	"github.com/ndewijer/go-temperature-rug/internal/utils"
)

const (
	ReverseOrder = false
	DateStart    = "20220101"
	DateEnd      = "20230321"
)

func main() {
	startTime := time.Now()

	// Define command-line flags
	stationsFlag := flag.String("stations", "240", "Comma-separated list of station IDs")
	variablesFlag := flag.String("variables", "TG,TN,TX", "Comma-separated list of variables")
	startDateFlag := flag.String("start", DateStart, "Start date (YYYYMMDD)")
	endDateFlag := flag.String("end", DateEnd, "End date (YYYYMMDD)")
	formatFlag := flag.String("format", "json", "Output format")
	reverseFlag := flag.Bool("reverse", ReverseOrder, "Reverse order")

	flag.Parse()

	// Parse stations
	stations, err := utils.ParseIntSlice(*stationsFlag)
	if err != nil {
		log.Fatalf("Error parsing stations: %v", err)
	}

	// Parse variables
	variables := strings.Split(*variablesFlag, ",")

	// Parse dates
	startDate, err := utils.ParseDate(*startDateFlag)
	if err != nil {
		log.Fatalf("Error parsing start date: %v", err)
	}
	endDate, err := utils.ParseDate(*endDateFlag)
	if err != nil {
		log.Fatalf("Error parsing end date: %v", err)
	}

	// Call the function with the parsed variables
	responseData, err := handlers.KnmiWeer(stations, variables, startDate, endDate, *formatFlag)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}

	var jsonData []map[string]interface{}
	err = json.Unmarshal([]byte(responseData), &jsonData)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	df := dataframe.LoadMaps(jsonData)

	// Sort the dataframe based on the ReverseOrder constant
	df = utils.SortDataFrame(df, *reverseFlag)

	var result strings.Builder
	result.WriteString("\n------------------------------------------------------------------\n")
	result.WriteString("| Date       | Min.Temp | Min.Kleur    | Max.Temp | Max.Kleur    |\n")
	result.WriteString("------------------------------------------------------------------\n")

	for _, row := range df.Maps() {
		formattedDate, minTemp, maxTemp, err := utils.ParseRow(row)
		if err != nil {
			log.Fatalf("Error parsing row: %v", err)
		}

		minColour, err := utils.GetColor(minTemp)
		if err != nil {
			log.Fatalf("Error getting color: %v", err)
		}

		maxColour, err := utils.GetColor(maxTemp)
		if err != nil {
			log.Fatalf("Error getting color: %v", err)
		}

		result.WriteString(fmt.Sprintf("| %s | %7.1f°C| %-12s | %7.1f°C| %-12s |\n", formattedDate, minTemp, minColour, maxTemp, maxColour))
	}

	result.WriteString("------------------------------------------------------------------\n")

	log.Println(result.String())

	elapsedTime := time.Since(startTime) // Calculate elapsed time
	fmt.Printf("Execution time: %s\n", elapsedTime)
}
