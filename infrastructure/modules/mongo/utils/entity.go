package mongoUtils

import (
	"encoding/json"
)

func ConvertSchema[t any, r any](toJson *r) *t {
	marshal := new(t)
	js, _ := json.Marshal(toJson)
	_ = json.Unmarshal(js, &marshal)

	return marshal
}
