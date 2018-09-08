

# bitmap
`import "github.com/andy2046/bitmap"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package bitmap implements Bitmap in Go.




## <a name="pkg-index">Index</a>
* [Constants](#pkg-constants)
* [type Bitmap](#Bitmap)
  * [func New(size int64) *Bitmap](#New)
  * [func (b *Bitmap) GetBit(offset int64) bool](#Bitmap.GetBit)
  * [func (b *Bitmap) SetBit(offset int64, v bool) bool](#Bitmap.SetBit)
  * [func (b *Bitmap) Size() int64](#Bitmap.Size)


#### <a name="pkg-files">Package files</a>
[bitmap.go](./bitmap.go) 


## <a name="pkg-constants">Constants</a>
``` go
const MaxBitmapSize int64 = 0x01 << 40
```
MaxBitmapSize is the maximum bitmap size (in bits).





## <a name="Bitmap">type</a> [Bitmap](./bitmap.go?s=185:238#L8)
``` go
type Bitmap struct {
    // contains filtered or unexported fields
}
```
Bitmap represents a bitmap.







### <a name="New">func</a> [New](./bitmap.go?s=269:297#L14)
``` go
func New(size int64) *Bitmap
```
New creates a new Bitmap.





### <a name="Bitmap.GetBit">func</a> (\*Bitmap) [GetBit](./bitmap.go?s=813:855#L43)
``` go
func (b *Bitmap) GetBit(offset int64) bool
```
GetBit returns the value of bit at offset.




### <a name="Bitmap.SetBit">func</a> (\*Bitmap) [SetBit](./bitmap.go?s=521:571#L29)
``` go
func (b *Bitmap) SetBit(offset int64, v bool) bool
```
SetBit sets bit at offset to value v.




### <a name="Bitmap.Size">func</a> (\*Bitmap) [Size](./bitmap.go?s=1028:1057#L52)
``` go
func (b *Bitmap) Size() int64
```
Size returns the bitmap size (in bits).



