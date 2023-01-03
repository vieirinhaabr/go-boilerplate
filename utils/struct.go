package utils

import "encoding/json"

func UnmarshalString[object any](data string, obj object) error {
	if err := json.Unmarshal([]byte(data), &obj); err != nil {
		return err
	}

	return nil
}

func TranslateStruct[f any, t any](from *f, to *t) error {
	fromJs, err := json.Marshal(from)
	if err != nil {
		return err
	}

	err = json.Unmarshal(fromJs, to)

	return err
}
