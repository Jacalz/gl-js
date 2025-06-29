//go:build js && wasm

package gl

import (
	"fmt"
	"runtime"
	"syscall/js"
	"unsafe"
)

func int8toBytes(s []int8) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s))
}

func int16toBytes(s []int16) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*2)
}

func int32toBytes(s []int32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

func int64toBytes(s []int64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*8)
}

func uint16toBytes(s []uint16) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*2)
}

func uint32toBytes(s []uint32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
}

func uint64toBytes(s []uint64) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*8)
}

func float32toBytes(s []float32) []byte {
	return unsafe.Slice((*byte)(unsafe.Pointer(&s[0])), len(s)*4)
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
		a := js.Global().Get("Int8Array").New(len(s))
		js.CopyBytesToJS(a, int8toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []int16:
		a := js.Global().Get("Int16Array").New(len(s))
		js.CopyBytesToJS(a, int16toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []int32:
		a := js.Global().Get("Int32Array").New(len(s))
		js.CopyBytesToJS(a, int32toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []uint8:
		a := js.Global().Get("Uint8Array").New(len(s))
		js.CopyBytesToJS(a, s)
		runtime.KeepAlive(s)
		return a
	case []uint16:
		a := js.Global().Get("Uint16Array").New(len(s))
		js.CopyBytesToJS(a, uint16toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []uint32:
		a := js.Global().Get("Uint32Array").New(len(s))
		js.CopyBytesToJS(a, uint32toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []float32:
		a := js.Global().Get("Float32Array").New(len(s))
		js.CopyBytesToJS(a, float32toBytes(s))
		runtime.KeepAlive(s)
		return a
	case []float64:
		a := js.Global().Get("Float64Array").New(len(s))
		js.CopyBytesToJS(a, float64toBytes(s))
		runtime.KeepAlive(s)
		return a
	default:
		panic(fmt.Sprintf("jsutil: unexpected value at SliceToTypedArray: %T", s))
	}
}
