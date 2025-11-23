package cell

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Cell struct {
	bitsSize uint
	data     []byte
	refs     []*Cell
}

type Builder struct {
	bitsSize uint
	data     []byte
	refs     []*Cell
}

func BeginCell() *Builder {
	return &Builder{
		bitsSize: 0,
		data:     []byte{},
		refs:     []*Cell{},
	}
}

func (b *Builder) StoreUInt(value uint64, bitsSize uint) error {
	if b.bitsSize+bitsSize > 1023 {
		return fmt.Errorf("cell would exceed 1023 bits")
	}

	bytesNeeded := (bitsSize + 7) / 8

	buf := make([]byte, 8)
	for i := 0; i < 8; i++ {
		buf[7-i] = byte(value & 0xFF)
		value >>= 8
	}

	buf = buf[8-bytesNeeded:]

	bitOffset := b.bitsSize % 8
	if bitOffset > 0 {
		for i := len(b.data) - 1; i >= 0; i-- {
			if i < len(b.data)-1 {
				b.data[i+1] |= b.data[i] >> (8 - bitOffset)
			}
			b.data[i] <<= bitOffset
		}
	}

	b.data = append(b.data, buf...)
	b.bitsSize += bitsSize

	if b.bitsSize%8 != 0 {
		lastByteIdx := len(b.data) - 1
		unusedBits := 8 - (b.bitsSize % 8)
		b.data[lastByteIdx] &= 0xFF << unusedBits
	}

	return nil
}

func (b *Builder) MustStoreUInt(value uint64, bitsSize uint) *Builder {
	if err := b.StoreUInt(value, bitsSize); err != nil {
		panic(err)
	}
	return b
}

func (b *Builder) EndCell() *Cell {
	return &Cell{
		bitsSize: b.bitsSize,
		data:     append([]byte{}, b.data...),
		refs:     b.refs,
	}
}

func (c *Cell) Hash() []byte {
	h := sha256.New()

	descriptor := byte(len(c.refs))
	if c.bitsSize%8 != 0 {
		descriptor |= 0x08
	}
	h.Write([]byte{descriptor})

	dataLen := (c.bitsSize + 7) / 8
	h.Write([]byte{byte(dataLen)})

	h.Write(c.data)

	for _, ref := range c.refs {
		refHash := ref.Hash()
		h.Write(refHash)
	}

	return h.Sum(nil)
}

func (c *Cell) Dump() string {
	return hex.EncodeToString(c.data)
}

func (b *Builder) StoreRef(ref *Cell) error {
	if len(b.refs) >= 4 {
		return fmt.Errorf("cell can have max 4 references")
	}
	if ref == nil {
		return fmt.Errorf("reference cannot be nil")
	}
	b.refs = append(b.refs, ref)
	return nil
}

func (b *Builder) MustStoreRef(ref *Cell) *Builder {
	if err := b.StoreRef(ref); err != nil {
		panic(err)
	}
	return b
}
