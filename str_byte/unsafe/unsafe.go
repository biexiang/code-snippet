package unsafe

import "unsafe"

//// https://github.com/golang/go/blob/master/src/runtime/slice.go#L15
//type slice struct {
//	array unsafe.Pointer
//	len   int
//	cap   int
//}

//// https://github.com/golang/go/blob/master/src/runtime/string.go#L238
//type stringStruct struct {
//	str unsafe.Pointer
//	len int
//}

func Str2byte(str string) []byte {
	tmp := (*[2]uintptr)(unsafe.Pointer(&str))
	bs := [3]uintptr{tmp[0], tmp[1], tmp[1]}
	return *(*[]byte)(unsafe.Pointer(&bs))
}

func Byte2str(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
}
