package adapter

import "encoding/json"

type jsonObject map[string]interface{}

func newJSONObject() jsonObject {
	return make(map[string]interface{})
}

func (m jsonObject) unmarshal(v interface{}) error {
	j, err := json.Marshal(m)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(j, &v); err != nil {
		return err
	}

	return nil
}
