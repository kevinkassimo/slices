package slices

import (
	//"reflect"
	"log"
)

func crash(errString string) {
	log.Fatalf("slices: %s\n", errString)
}
//
//func checkSlice(s interface{}) {
//	s_val := reflect.ValueOf(s);
//
//	if s_val.Kind() != reflect.Slice {
//		crash("the given type is not a slice")
//	}
//}
//
//type _SliceDispatcher interface {
//	Append(args ...interface{}) []interface{}
//}
//
//func On(s interface{}) _SliceDispatcher {
//	if reflect.ValueOf(s).Kind() != reflect.Slice {
//		crash("given s is NOT a slice")
//	}
//	switch reflect.TypeOf(s).Elem().Kind() {
//	case reflect.Int:
//		return _IntSlice{s.([]int)}
//	case reflect.Int8:
//	case reflect.Int16:
//	case reflect.Int32:
//	case reflect.Int64:
//
//	case reflect.Uint:
//	case reflect.Uint8:
//	case reflect.Uint16:
//	case reflect.Uint32:
//	case reflect.Uint64:
//
//	case reflect.Float32:
//	case reflect.Float64:
//
//	case reflect.Bool:
//
//	case reflect.String:
//
//	default:
//		crash("Given slice type is not implemented, or not implementable")
//	}
//
//	return struct{}{}
//}
//
//type Dispatcher struct {
//	Type reflect.Kind
//	Int *_IntSlice
//}
//
//type IntDispatcher struct {
//	Body []int
//}