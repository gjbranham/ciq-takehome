package csv

import (
	"encoding/csv"
	"io"
	"strconv"
)

type AccessInfo struct {
	Timestamp string
	Username  string
	Operation string
	Size      int
}

func ReadCsv(reader io.Reader) ([]AccessInfo, error) {
	csvReader := csv.NewReader(reader)

	// skip header
	_, err := csvReader.Read()
	if err != nil {
		return nil, err
	}

	rawCsvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	// convert the raw csv data into a data structure we can use
	parsedCsvData := make([]AccessInfo, 0, len(rawCsvData)) // use cap here since same len
	for _, record := range rawCsvData {
		size, _ := strconv.Atoi(record[3])
		entry := AccessInfo{
			Timestamp: record[0],
			Username:  record[1],
			Operation: record[2],
			Size:      size,
		}
		parsedCsvData = append(parsedCsvData, entry)
	}
	return parsedCsvData, nil
}
