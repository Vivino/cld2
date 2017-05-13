//+build linux,go1.8

package main

// #include <stdlib.h>
// #include "cld2.h"
import "C"
import (
	"unsafe"

	"github.com/klauspost/cld2/internal/info"
)

func main() {
	panic("install me as plugin")
}

func PluginDetect(text string) string {
	cs := C.CString(text)
	res := C.DetectLang(cs, -1)
	C.free(unsafe.Pointer(cs))
	var lang string
	if res != nil {
		lang = C.GoString(res)
	}
	return lang
}

func PluginDetectLang(text string) uint16 {
	cs := C.CString(text)
	res := C.DetectLangCode(cs, -1)
	C.free(unsafe.Pointer(cs))
	return uint16(res)
}

func PluginDetectThree(text string) info.Languages {
	cs := C.CString(text)
	dst := new(C.struct__result)
	C.DetectThree(dst, cs, -1)
	C.free(unsafe.Pointer(cs))
	res := make([]info.Estimate, 0, 3)
	for i := range dst.language {
		var est info.Estimate
		est.Language = uint16(dst.language[i])
		est.Percent = int(dst.percent[i])
		est.NormScore = float64(dst.normalized_score[i])
		res = append(res, est)
	}
	rel := true
	if dst.reliable == 0 {
		rel = false
	}
	return info.Languages{Estimates: res, Reliable: rel, TextBytes: int(dst.text_bytes)}
}
