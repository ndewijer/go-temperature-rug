package handlers

import (
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func KnmiWeer(station_numbers []int, vars []string, startDate time.Time, endDate time.Time, filefmt string) (string, error) {
	var stations string
	for i, num := range station_numbers {
		if i > 0 {
			stations += ":"
		}
		stations += strconv.Itoa(num)
	}
	varsStr := strings.Join(vars, ":")

	payload := "start=" + startDate.Format("20060102") +
		"&end=" + endDate.Format("20060102") +
		"&vars=" + varsStr +
		"&stns=" + stations +
		"&fmt=" + filefmt

	payloadReader := strings.NewReader(payload)

	response, err := http.Post("https://www.daggegevens.knmi.nl/klimatologie/daggegevens", "application/x-www-form-urlencoded", payloadReader)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(responseData), nil
}
