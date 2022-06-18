package gopfmat

import "encoding/json"

func Json(v any) (string, error) {
	bytes, err := json.Marshal(v)
	return string(bytes), err
}

func JsonPretty(v any) (string, error) {
	bytes, err := json.MarshalIndent(v, "", "    ")
	return string(bytes), err
}

func JsonUnsafe(v any) string {
	bytes, _ := json.Marshal(v)
	return string(bytes)
}
