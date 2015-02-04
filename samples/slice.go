package main

import "fmt"

func AddOneToEachElement(slice []byte) {
	for i := range slice {
		slice[i]++
	}
}

func SubstractOneFromLength(slice []byte) []byte {
	slice = slice[0 : len(slice) -1]
	return slice
}

func PtrSubstractOneFromLength(slicePtr *[]byte) {
	slice := *slicePtr
	*slicePtr = slice[0 : len(slice)-1]
}

func main() {
	var buffer [256]byte
	slice := buffer[10:20]
	for i := 0; i < len(slice); i++ {
		slice[i] = byte(i)
	}
	fmt.Println("before", slice)
	AddOneToEachElement(slice)
	fmt.Println("after", slice)
	fmt.Println("Before: len(slice)=", len(slice))
	newSlice := SubstractOneFromLength(slice)
	fmt.Println("After: len(slice)=", len(slice))
	fmt.Println("After: len(newSlice)=", len(newSlice))
	
	fmt.Println("Before: len(slice)", len(slice))
	PtrSubstractOneFromLength(&slice)
	fmt.Println("After: len(slice)", len(slice))

	const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"

	fmt.Printf("%q\n", sample)
	fmt.Printf("%+q\n", sample)
}
