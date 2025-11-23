package main

import (
	"GoTON/cell"
	"GoTON/crc16"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println("Hello, Go!")
	fmt.Println("CRC16")
	data := []byte("hello")
	crc := crc16.Compute(data)
	fmt.Printf("CRC16 of 'hello': %X\n", crc)

	fmt.Println("Cell")
	cell := cell.BeginCell().MustStoreUInt(42, 8).EndCell()
	fmt.Printf("Cell: %s\n", cell.Dump())
	fmt.Printf("Cell hash: %s\n", hex.EncodeToString(cell.Hash()))
}
