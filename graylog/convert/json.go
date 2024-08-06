package convert

import (
	"encoding/json"
	"fmt"

	"github.com/suzuki-shunsuke/go-dataeq/dataeq"
)

func OneSizeListToJSON(data map[string]interface{}, keys ...string) error {
	for _, key := range keys {
		v := data[key].([]interface{})[0].(map[string]interface{})
		b, err := json.Marshal(v)
		if err != nil {
			return fmt.Errorf("failed to marshal attributes '%s' as JSON: %w", key, err)
		}
		data[key] = string(b)
	}
	return nil
}

func DataToJSON(data map[string]interface{}, keys ...string) error {
	if len(keys) == 0 {
		// all keys
		for key, a := range data {
			s, err := json.Marshal(a)
			if err != nil {
				return fmt.Errorf("failed to marshal the '%s' as JSON: %w", key, err)
			}
			data[key] = string(s)
		}
		return nil
	}
	for _, key := range keys {
		if data[key] == nil {
			continue
		}
		s, err := json.Marshal(data[key])
		if err != nil {
			return fmt.Errorf("failed to marshal the '%s' as JSON: %w", key, err)
		}
		data[key] = string(s)
	}
	return nil
}

func JSONToData(data map[string]interface{}, keys ...string) error {
	if len(keys) == 0 {
		// all keys
		for key, v := range data {
			attr, err := dataeq.JSON.ConvertByte([]byte(v.(string)))
			if err != nil {
				return fmt.Errorf("failed to parse the '%s'. '%s' must be a JSON string: %w", key, key, err)
			}
			data[key] = attr
		}
		return nil
	}
	for _, key := range keys {
		v, ok := data[key]
		if !ok {
			continue
		}
		if v.(string) == "" {
			data[key] = nil
			continue
		}
		attr, err := dataeq.JSON.ConvertByte([]byte(v.(string)))
		if err != nil {
			return fmt.Errorf("failed to parse the '%s'. '%s' must be a JSON string: %w", key, key, err)
		}
		data[key] = attr
	}
	return nil
}
