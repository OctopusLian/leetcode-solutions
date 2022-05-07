package main

import (
	"fmt"
	"unsafe"
)

func main() {
	short := int16(1)
	int1 := int32(1)
	long := int64(1)
	longlong := int(1)
	fmt.Printf("The size of short is %d bytes.\n", unsafe.Sizeof(short))
	fmt.Printf("The size of int is %d bytes.\n", unsafe.Sizeof(int1))
	fmt.Printf("The size of long is %d bytes.\n", unsafe.Sizeof(long))
	fmt.Printf("The size of long long is %d bytes.\n", unsafe.Sizeof(longlong))
}
