package pic

import (
	"image"
	"image/color"
)

// Picture is an alternative to
// the official *image.RGBA.
type Picture struct {
	Orig     *Picture
	Rect     image.Rectangle
	Pix      []byte
	Stride   int
	relative bool
}

// Equals returns 'true' if the picture equals
// another one, and 'false' otherwise.
func (pic *Picture) Equals(other *Picture) bool {
	equals := true

	for j := 0; j < pic.Rect.Dy(); j++ {
		for i := 0; i < pic.Rect.Dx(); i++ {
			if pic.At(i, j) != other.At(i, j) {
				equals = false
				break
			}
		}
	}

	return equals
}

// At returns the color of the pixel
// located at the specified position.
func (pic *Picture) At(x, y int) color.RGBA {
	if pic.relative {
		return pic.Orig.At(pic.Rect.Min.X+x, pic.Rect.Min.Y+y)
	}

	pos := y*pic.Rect.Dx()*4 + x*4

	return color.RGBA{
		R: pic.Pix[pos],
		G: pic.Pix[pos+1],
		B: pic.Pix[pos+2],
		A: pic.Pix[pos+3],
	}
}

// SubPicture returns a new picture
// based on the given picture.
//
// Copies the data of the original
// picture to the derivative one.
//
// The new picture will not have
// any connection to the base one.
func (pic *Picture) SubPicture(rect image.Rectangle) *Picture {
	sub := &Picture{
		Rect:   image.Rect(0, 0, rect.Dx(), rect.Dy()),
		Pix:    make([]byte, rect.Dx()*rect.Dy()*4),
		Stride: 4 * rect.Dx(),
	}

	for j := rect.Min.Y; j < rect.Max.Y; j++ {
		for i := rect.Min.X; i < rect.Max.X; i++ {
			col := pic.At(i, j)
			pos := (j-rect.Min.Y)*4*
				rect.Dx() + (i-rect.Min.X)*4

			sub.Pix[pos] = col.R
			sub.Pix[pos+1] = col.G
			sub.Pix[pos+2] = col.B
			sub.Pix[pos+3] = col.A
		}
	}

	return sub
}

// SubImageRelative creates a new picture based
// on the given one but without copying the data.
func (pic *Picture) SubImageRelative(rect image.Rectangle) *Picture {
	sub := &Picture{
		Rect:     rect,
		Pix:      pic.Pix,
		relative: true,
		Orig:     pic,
		Stride:   pic.Stride,
	}

	return sub
}

// CreateImageRGBAFromPicture creates a new
// instance of *image.RGBA and copies the data
// of the original picture to the new image.
func (pic *Picture) CreateImageRGBAFromPicture() *image.RGBA {
	img := &image.RGBA{
		Rect:   pic.Rect,
		Stride: pic.Stride,
		Pix:    make([]byte, pic.Rect.Dx()*pic.Rect.Dy()*4),
	}
	copy(img.Pix, pic.Pix)

	return img
}

// NewImageRGBAFromPicture creates a new
// instance of *image.RGBA without copying
// the data of the original picture.
func (pic *Picture) NewImageRGBAFromPicture() *image.RGBA {
	return &image.RGBA{
		Rect:   pic.Rect,
		Pix:    pic.Pix,
		Stride: pic.Stride,
	}
}

// CreatePictureFromRGBA creates a new picture
// from the given RGBA image with copying the data.
func CreatePictureFromRGBA(imgRGBA *image.RGBA) *Picture {
	pic := &Picture{
		Rect:   imgRGBA.Rect,
		Pix:    make([]byte, imgRGBA.Rect.Dx()*imgRGBA.Rect.Dy()*4),
		Stride: imgRGBA.Stride,
	}
	copy(pic.Pix, imgRGBA.Pix)

	return pic
}

// NewPictureFromRGBA creates a new picture
// from the given RGBA image without copying the data.
func NewPictureFromRGBA(imgRGBA *image.RGBA) *Picture {
	return &Picture{
		Rect:   imgRGBA.Rect,
		Pix:    imgRGBA.Pix,
		Stride: imgRGBA.Stride,
	}
}
