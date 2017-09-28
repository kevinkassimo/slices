package slices

import (
	"reflect"
	"log"
)

func crash(errString string) {
	log.Fatalf("slices: %s\n", errString)
}

func checkSlice(s interface{}) {
	s_val := reflect.ValueOf(s);

	if s_val.Kind() != reflect.Slice {
		crash("the given type is not a slice")
	}
}


//func Insert(s interface{}, i int, args ...interface{}) interface{} {
//	checkSlice(s)
//	return
//}

func Append(s interface{}, args ...interface{}) []interface{} {
	checkSlice(s)
	return append(s.([]interface{}), args...);
}

func Remove() {

}