package utils

import (
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/go-gota/gota/dataframe"
)

const (
	DateFormat = "20060102"
)

var (
	DEG_LIST    = []float64{-40, 0, 6, 12, 18, 24, 30, 50}
	COLOR_CODES = []string{"donkerblauw", "blauw", "lichtblauw", "creme", "lightrood", "rood", "donkerrood"}
)

func GetColor(temperature float64) (string, error) {
	for i, deg := range DEG_LIST[1:] {
		if temperature <= deg {
			return COLOR_CODES[i], nil
		}
	}
	return COLOR_CODES[len(COLOR_CODES)-1], nil
}

func ParseIntSlice(s string) ([]int, error) {
	var result []int
	for _, v := range strings.Split(s, ",") {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func ParseDate(s string) (time.Time, error) {
	return time.Parse("20060102", s)
}

func ParseRow(row map[string]interface{}) (string, float64, float64, error) {
	dateStr := row["date"].(string)
	date, err := time.Parse(time.RFC3339, dateStr)
	if err != nil {
		return "", 0, 0, err
	}
	formattedDate := date.Format("2006-01-02")
	minTemp := float64(row["TN"].(int)) / 10
	maxTemp := float64(row["TX"].(int)) / 10
	return formattedDate, minTemp, maxTemp, nil
}

func SortDataFrame(df dataframe.DataFrame, reverse bool) dataframe.DataFrame {
	// Extract dates and sort them
	dates := df.Col("date").Records()
	sort.SliceStable(dates, func(i, j int) bool {
		date1, _ := time.Parse(time.RFC3339, dates[i])
		date2, _ := time.Parse(time.RFC3339, dates[j])
		if reverse {
			return date1.After(date2)
		}
		return date1.Before(date2)
	})

	// Create a new dataframe with sorted dates
	sortedMaps := make([]map[string]interface{}, len(dates))
	for i, date := range dates {
		for _, row := range df.Maps() {
			if row["date"] == date {
				sortedMaps[i] = row
				break
			}
		}
	}

	return dataframe.LoadMaps(sortedMaps)
}
