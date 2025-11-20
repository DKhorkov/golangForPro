package filepath

import (
	"fmt"
	"strings"
)

const (
	csvSource  = "csv"
	jsonSource = "json"

	CSVExtension  = ".csv"
	JSONExtension = ".json"

	csvDefaultFilepath  = "phonebook.csv"
	jsonDefaultFilepath = "phonebook.json"
)

var (
	validExtensions = []string{CSVExtension, JSONExtension}
	validSources    = []string{csvSource, jsonSource}
)

func Get(filepath, source string) (string, error) {
	switch {
	case strings.HasSuffix(filepath, CSVExtension), strings.HasSuffix(filepath, JSONExtension):
		return filepath, nil
	case len(filepath) == 0 && source == csvSource:
		return csvDefaultFilepath, nil
	case len(filepath) == 0 && source == jsonSource:
		return jsonDefaultFilepath, nil
	}

	return "", fmt.Errorf(
		"invalid path (%s) and/or unknown source (%s). Valid path extensions: %v. Valid Sources: %v",
		filepath,
		source,
		validExtensions,
		validSources,
	)
}
