// Package bitmap implements Bitmap in Go.
package bitmap

// MaxBitmapSize is the maximum bitmap size (in bits).
const MaxBitmapSize int64 = 0x01 << 40

// Bitmap represents a bitmap.
type Bitmap struct {
	data    []byte
	bitsize int64
}

// New creates a new Bitmap.
func New(size int64) *Bitmap {
	if size <= 0 || size > MaxBitmapSize {
		size = MaxBitmapSize
	}
	r := size & 7 // size % 8
	if r != 0 {
		r = 1
	}
	return &Bitmap{
		data:    make([]byte, (size>>3)+r), // size/8+r
		bitsize: size,
	}
}

// SetBit sets bit at offset to value v.
func (b *Bitmap) SetBit(offset int64, v bool) bool {
	if offset > b.bitsize {
		return false
	}
	index, bit := offset>>3, offset&7 // offset/8, offset%8
	if v {
		b.data[index] |= 0x01 << uint64(bit)
	} else {
		b.data[index] &^= 0x01 << uint64(bit)
	}
	return true
}

// GetBit returns the value of bit at offset.
func (b *Bitmap) GetBit(offset int64) bool {
	if offset > b.bitsize {
		return false
	}
	index, bit := offset>>3, offset&7 // offset/8, offset%8
	return (b.data[index]>>uint64(bit))&0x01 != 0
}

// Size returns the bitmap size (in bits).
func (b *Bitmap) Size() int64 {
	return b.bitsize
}
