//go:build js && wasm

package gl

import (
	"fmt"
	"runtime"
	"syscall/js"
	"unsafe"
)

func int8ToJSArray(s []int8) js.Value {
	a := js.Global().Get("Int8Array").New(len(s))
	js.CopyBytesToJS(a, int8toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func int8toBytes(s []int8) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s))
}

func int16ToJSArray(s []int16) js.Value {
	a := js.Global().Get("Int16Array").New(len(s))
	js.CopyBytesToJS(a, int16toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func int16toBytes(s []int16) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*2)
}

func int32ToJSArray(s []int32) js.Value {
	a := js.Global().Get("Int32Array").New(len(s))
	js.CopyBytesToJS(a, int32toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func int32toBytes(s []int32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

func int64ToJSArray(s []int64) js.Value {
	a := js.Global().Get("Int64Array").New(len(s))
	js.CopyBytesToJS(a, int64toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func int64toBytes(s []int64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*8)
}

func uint8ToJSArray(s []uint8) js.Value {
	a := js.Global().Get("Uint8Array").New(len(s))
	js.CopyBytesToJS(a, s)
	runtime.KeepAlive(s)
	return a
}

func uint16ToJSArray(s []uint16) js.Value {
	a := js.Global().Get("Uint16Array").New(len(s))
	js.CopyBytesToJS(a, uint16toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func uint16toBytes(s []uint16) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*2)
}

func uint32ToJSArray(s []uint32) js.Value {
	a := js.Global().Get("Uint32Array").New(len(s))
	js.CopyBytesToJS(a, uint32toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func uint32toBytes(s []uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

func uint64ToJSArray(s []uint64) js.Value {
	a := js.Global().Get("Uint64Array").New(len(s))
	js.CopyBytesToJS(a, uint64toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func uint64toBytes(s []uint64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*8)
}

func float32ToJSArray(s []float32) js.Value {
	a := js.Global().Get("Float32Array").New(len(s))
	js.CopyBytesToJS(a, float32toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func float32toBytes(s []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

func float64ToJSArray(s []float64) js.Value {
	a := js.Global().Get("Float64Array").New(len(s))
	js.CopyBytesToJS(a, float64toBytes(s))
	runtime.KeepAlive(s)
	return a
}

func float64toBytes(s []float64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*8)
}

func SliceToTypedArray(s any) js.Value {
	if s == nil {
		return js.Null()
	}

	switch s := s.(type) {
	case []int8:
		return int8ToJSArray(s)
	case []int16:
		return int16ToJSArray(s)
	case []int32:
		return int32ToJSArray(s)
	case []uint8:
		return uint8ToJSArray(s)
	case []uint16:
		return uint16ToJSArray(s)
	case []uint32:
		return uint32ToJSArray(s)
	case []float32:
		return float32ToJSArray(s)
	case []float64:
		return float64ToJSArray(s)
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}
}
