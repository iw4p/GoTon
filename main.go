package main

import (
	"GoTON/crc16"
	"fmt"
)

func main() {
	fmt.Println("Hello, Go!")
	data := []byte("hello")
	crc := crc16.Compute(data)
	fmt.Printf("CRC16 of 'hello': %X\n", crc)
}
