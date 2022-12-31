package utils

import "encoding/json"

func UnmarshalString[object any](data string, obj object) error {
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}

	return nil
}
