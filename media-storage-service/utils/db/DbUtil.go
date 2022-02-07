package db

import (
	"fmt"
	"reflect"
)

func GenerateCreateQuery(collectionName string, data interface{}) (string, map[string]interface{}) {
	var query = fmt.Sprintf("Create  (a:%s {", collectionName)

	val := reflect.ValueOf(data)
	t := reflect.TypeOf(data)
	num := val.NumField()

	res := make(map[string]interface{})
	for i := 0; i < num; i++ {
		key := t.Field(i).Name
		value := val.Field(i).Interface()

		query = fmt.Sprintf("%s %s: $%s,", query, key, key)
		res[key] = value
	}
	query = query[:len(query)-1]

	query = fmt.Sprintf("%s})", query)

	return query, res
}
