package main

/*
#cgo CFLAGS: -I./
#cgo LDFLAGS: -L. -lgo
#include "libgo.h"
*/
import "C"

func main() {
	C.mylog()
}

