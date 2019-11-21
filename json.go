package tool

import (
	"encoding/json"
)

/*map 转 object*/
func MapToObject(data map[string]interface{}, v interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(b, v); err != nil {
		return err
	}
	return nil
}
