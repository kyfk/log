package format

import (
	"encoding/json"
)

// JSONPretty is format of message output.
func JSONPretty(v map[string]interface{}) (string, error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return "", err
	}
	return string(b), nil
}
