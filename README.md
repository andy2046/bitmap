

# bitmap
`import "github.com/andy2046/bitmap"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package bitmap implements Bitmap in Go.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type Bitmap](#Bitmap)
  * [func New(size uint64) *Bitmap](#New)
  * [func (b *Bitmap) GetBit(offset uint64) bool](#Bitmap.GetBit)
  * [func (b *Bitmap) SetBit(offset uint64, v bool) bool](#Bitmap.SetBit)
  * [func (b *Bitmap) Size() uint64](#Bitmap.Size)


#### <a name="pkg-files">Package files</a>
[bitmap.go](./bitmap.go) 


## <a name="pkg-constants">Constants</a>
``` go
const MaxBitmapSize uint64 = 0x01 << 40
```
MaxBitmapSize is the maximum bitmap size (in bits).





## <a name="Bitmap">type</a> [Bitmap](./bitmap.go?s=186:240#L8)
``` go
type Bitmap struct {
    // contains filtered or unexported fields
}
```
Bitmap represents a bitmap.







### <a name="New">func</a> [New](./bitmap.go?s=271:300#L14)
``` go
func New(size uint64) *Bitmap
```
New creates a new Bitmap.





### <a name="Bitmap.GetBit">func</a> (\*Bitmap) [GetBit](./bitmap.go?s=848:891#L43)
``` go
func (b *Bitmap) GetBit(offset uint64) bool
```
GetBit returns the value of bit at `offset`.




### <a name="Bitmap.SetBit">func</a> (\*Bitmap) [SetBit](./bitmap.go?s=530:581#L29)
``` go
func (b *Bitmap) SetBit(offset uint64, v bool) bool
```
SetBit sets bit at `offset` to value `v`.




### <a name="Bitmap.Size">func</a> (\*Bitmap) [Size](./bitmap.go?s=1065:1095#L52)
``` go
func (b *Bitmap) Size() uint64
```
Size returns the bitmap size (in bits).



