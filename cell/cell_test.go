package cell

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestCell(t *testing.T) {
	// Test 1: Simple cell
	cell1 := BeginCell().MustStoreUInt(42, 8).EndCell()
	fmt.Printf("Test 1 - Simple cell:\n")
	fmt.Printf("   Dump: %s\n", cell1.Dump())
	fmt.Printf("   Hash: %s\n\n", hex.EncodeToString(cell1.Hash()))

	// Test 2: Cell with reference
	inner := BeginCell().MustStoreUInt(111, 8).EndCell()
	outer := BeginCell().
		MustStoreUInt(777, 16).
		MustStoreRef(inner).
		EndCell()
	fmt.Printf("Test 2 - Cell with reference:\n")
	fmt.Printf("   Dump: %s\n", outer.Dump())
	fmt.Printf("   Hash: %s\n\n", hex.EncodeToString(outer.Hash()))
}
