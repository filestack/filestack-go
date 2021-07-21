package url

import (
	"encoding/json"
	"sort"
	"strings"
)

// Task represents a transformation task.
type Task struct {
	Name   string
	Params map[string]interface{}
}

// NewTask constructor.
func NewTask(name string, params map[string]interface{}) (task Task) {
	return Task{
		Name:   name,
		Params: params,
	}
}

// AsString converts a task to a string.
func (t *Task) AsString() (task string) {
	if t.Params == nil || len(t.Params) == 0 {
		return t.Name
	}

	var params []string
	for key, value := range t.Params {
		params = append(params, key+":"+paramValueToString(value))
	}
	sort.Strings(params)

	return t.Name + "=" + strings.Join(params, ",")
}

func paramValueToString(i interface{}) (strValue string) {
	switch v := i.(type) {
	case string:
		return v
	case *string:
		if v == nil {
			return ""
		}
		return *v
	case []string:
		strValue = "["
		for key, value := range v {
			if key > 0 {
				strValue = strValue + ","
			}
			strValue = strValue + value

		}
		return strValue + "]"
	default:
		jsonVal, err := json.Marshal(v)
		if err != nil {
			return ""
		}
		return string(jsonVal)
	}
}
