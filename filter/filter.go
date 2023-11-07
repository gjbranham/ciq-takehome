package filter

import (
	"time"

	"github.com/gjbranham/ciq-takehome/args"
	c "github.com/gjbranham/ciq-takehome/csv"
)

func FilterData(data []c.AccessInfo, args *args.Arguments) ([]c.AccessInfo, error) {
	filteredData := data
	var err error

	if args.AllResults {
		return filteredData, nil
	}

	if args.Username != "" {
		filteredData = filterByUsername(data, args.Username)
	}

	if args.Date != "" {
		filteredData, err = filterByDate(filteredData, args.Date)
		if err != nil {
			return nil, err
		}
	}

	if args.GreaterThanSize > 0 || args.LessThanSize > 0 {
		filteredData, err = filterBySize(filteredData, args.GreaterThanSize, args.LessThanSize)
		if err != nil {
			return nil, err
		}
	}
	return filteredData, nil
}

func filterByUsername(data []c.AccessInfo, username string) []c.AccessInfo {
	filtered := []c.AccessInfo{}
	for _, item := range data {
		if item.Username == username {
			filtered = append(filtered, item)
		}
	}
	return filtered
}

func filterByDate(data []c.AccessInfo, filterDate string) ([]c.AccessInfo, error) {
	filtered := []c.AccessInfo{}

	for _, item := range data {
		parsedT, err := time.Parse(time.UnixDate, item.Timestamp)
		if err != nil {
			return nil, err
		}
		formattedT := parsedT.Format("02/01/2006")

		if formattedT == filterDate {
			filtered = append(filtered, item)
		}
	}
	return filtered, nil
}

func filterBySize(data []c.AccessInfo, greaterThan, lessThan int) ([]c.AccessInfo, error) {
	filtered := []c.AccessInfo{}

	for _, item := range data {
		if greaterThan > 0 && lessThan > 0 {
			if greaterThan < item.Size && item.Size < lessThan {
				filtered = append(filtered, item)
			}
		} else if greaterThan > 0 && item.Size > greaterThan || lessThan > 0 && item.Size < lessThan {
			filtered = append(filtered, item)
		}
	}
	return filtered, nil
}
