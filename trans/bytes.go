package trans

import (
	"reflect"
	"unsafe"
)

// UnsafeBytesToString tips: 只有在原有bytes确保不会发生变化时可以使用
func UnsafeBytesToString(data []byte) string {
	return *(*string)(unsafe.Pointer(&data))
}

// UnsafeStringToBytes 利用反射转移字符串数据到bytes
func UnsafeStringToBytes(src string) (data []byte) {
	stringHandler := (*reflect.StringHeader)(unsafe.Pointer(&src))
	bytesHandler := (*reflect.SliceHeader)(unsafe.Pointer(&data))
	bytesHandler.Data = stringHandler.Data
	bytesHandler.Len = stringHandler.Len
	bytesHandler.Cap = stringHandler.Len
	return data
}
