package data_parser

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

func readCSVFile(filePath string) ([]string, error) {
	// Open the file
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Read the first line of the file
	csvReader := csv.NewReader(f)
	line, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	return line, nil
}

func trimStringSlice(s []string) []string {
	return strings.FieldsFunc(strings.Join(s, ","), func(r rune) bool {
		return r == '=' || r == '\n' || r == '\t'
	})
}

func splitStringSlice(s []string) ([]string, error) {
	var result []string
	for _, v := range s {
		tokens := strings.Split(strings.TrimSpace(v), ",")
		for _, token := range tokens {
			if token != "" {
				result = append(result, token)
			}
		}
	}
	return result, nil
}

func atoiSlice(s []string) ([]int, error) {
	var result []int
	for _, v := range s {
		i, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		result = append(result, i)
	}
	return result, nil
}

func float64Slice(s []string) ([]float64, error) {
	var result []float64
	for _, v := range s {
		f, err := strconv.ParseFloat(v, 64)
		if err != nil {
			return nil, err
		}
		result = append(result, f)
	}
	return result, nil
}