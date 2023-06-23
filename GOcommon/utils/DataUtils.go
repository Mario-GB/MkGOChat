package gocomm_utils

import (
	"encoding/json"
	"log"
	"reflect"
)

// From object to json-string
func ToJson(toCase any) string {
	toCaseBytes, castErr := json.Marshal(toCase)
	if castErr != nil {
		log.Fatalln("Error casting to JSON\nError: " + castErr.Error())
	}

	return string(toCaseBytes)
}

// From json-string to object
func FromJson(toCaseValue string, toCaseType reflect.Type) any {
	toReturnObj := reflect.New(toCaseType).Interface()
	castErr := json.Unmarshal([]byte(toCaseValue), &toReturnObj)
	if castErr != nil {
		log.Fatalln("Error casting from JSON\nError: " + castErr.Error())
	}

	return toReturnObj
}

type array struct {
	value [][]any
}

func (arr array) find(toFind any) any {
	for _, e := range arr.value {
		if e[0] == toFind {
			return toFind
		}
	}
	return nil
}

func (arr array) add(toAdd any) array {
	arr.value = append(arr.value, []any{toAdd, reflect.TypeOf(toAdd)})
	return arr
}
