package unsafe

import (
	"reflect"
	"unsafe"
)

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

func str2byte(str string) []byte {
	tmp := (*[2]uintptr)(unsafe.Pointer(&str))
	bs := [3]uintptr{tmp[0], tmp[1], tmp[1]}
	return *(*[]byte)(unsafe.Pointer(&bs))
}

func Str2byte(str string) (bs []byte) {
	strHdr := (*reflect.StringHeader)(unsafe.Pointer(&str))
	sliceHdr := (*reflect.SliceHeader)(unsafe.Pointer(&bs))
	sliceHdr.Data = strHdr.Data
	sliceHdr.Cap = strHdr.Len
	sliceHdr.Len = strHdr.Len
	return
}

func Byte2str(bt []byte) string {
	return *(*string)(unsafe.Pointer(&bt))
}
