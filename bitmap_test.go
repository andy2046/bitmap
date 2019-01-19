package bitmap_test

import (
	"math/rand"
	"testing"

	"github.com/andy2046/bitmap"
)

func TestBitmap(t *testing.T) {
	t.Logf("MaxBitmapSize -> %d\n", bitmap.MaxBitmapSize)
	size := uint64(rand.Int63n(int64(bitmap.MaxBitmapSize)) + 1)
	bm := bitmap.New(size)
	s := bm.Size()
	if s != size {
		t.Fatalf("wrong size %d", s)
	}
	r := bm.SetBit(1, true)
	if !r || !bm.GetBit(1) {
		t.Fatalf("wrong value at bit %d", 1)
	}
	r = bm.SetBit(size, true)
	if !r || !bm.GetBit(size) {
		t.Fatalf("wrong value at bit %d", size)
	}
	r = bm.SetBit(bitmap.MaxBitmapSize+1, true)
	if r || bm.GetBit(bitmap.MaxBitmapSize+1) {
		t.Fatal("wrong value")
	}
}

func BenchmarkBitmap(b *testing.B) {
	x := uint64(10000)
	bm := bitmap.New(x)
	var index uint64
	for i := 0; i < b.N; i++ {
		index++
		if index == x {
			index = 0
		}
		bm.SetBit(index, !bm.GetBit(index))
	}
}
