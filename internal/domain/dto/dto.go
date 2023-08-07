package dto

import (
	"reflect"
	"encoding/json"
	"strconv"
)

func ToMap(s interface{}) map[string]any {
	r := make(map[string]any)
			t := reflect.TypeOf(s);if t.Kind()==reflect.Ptr { t=t.Elem() };//fmt.Println(t.Kind())
	//if t, is := s.([]interface{}); is {
	if t.Kind() == reflect.Slice||t.Kind() == reflect.Array {
			var q []interface{}; b,e := json.Marshal(s); if e != nil { return nil }
			e = json.Unmarshal(b, &q) 
			for k,v := range q { r[strconv.Itoa(k)] = v }
	} else {
			b,e := json.Marshal(s); if e != nil { return nil }; e = json.Unmarshal(b, &r)
	}
	return r
}

//

func init () {
	//
}