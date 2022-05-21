package fselib

// #include "fse.h"
import "C"
import (
	"errors"
	"unsafe"
)

func Encode(dst, src []byte) (int, error) {
	pdst := unsafe.Pointer(&dst[0])
	psrc := unsafe.Pointer(&src[0])
	cn := C.FSE_compress(pdst, C.size_t(len(dst)), psrc, C.size_t(len(src)))

	n := int(cn)
	if n == 0 {
		return 0, errors.New("src is incompressible")
	}
	if n == 1 {
		return 0, errors.New("src is the same byte symbol repeated")
	}
	if C.FSE_isError(cn) != 0 {
		pstr := C.FSE_getErrorName(cn)
		return 0, errors.New(C.GoString(pstr))
	}

	return n, nil
}

func Decode(dst, src []byte) (int, error) {
	pdst := unsafe.Pointer(&dst[0])
	psrc := unsafe.Pointer(&src[0])
	cn := C.FSE_decompress(pdst, C.size_t(len(dst)), psrc, C.size_t(len(src)))

	if C.FSE_isError(cn) != 0 {
		pstr := C.FSE_getErrorName(cn)
		return 0, errors.New(C.GoString(pstr))
	}

	return int(cn), nil
}
