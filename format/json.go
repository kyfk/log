package format

import (
	"encoding/json"
)

// JSON is format of message output.
func JSON(v map[string]interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
