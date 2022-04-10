package pic_test

import (
	"image"
	"testing"

	"github.com/zergon321/pic"
)

func BenchmarkImageAt(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	for i := 0; i < b.N; i++ {
		img.At(250, 250)
	}
}

func BenchmarkPictureAt(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	picture := pic.CreatePictureFromRGBA(img)

	for i := 0; i < b.N; i++ {
		picture.At(250, 250)
	}
}

func BenchmarkImageSubimage(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))

	for i := 0; i < b.N; i++ {
		img.SubImage(image.Rect(30, 30, 150, 150))
	}
}

func BenchmarkPictureSubpicture(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	picture := pic.CreatePictureFromRGBA(img)

	for i := 0; i < b.N; i++ {
		picture.SubPicture(image.Rect(30, 30, 150, 150))
	}
}

func BenchmarkPictureSubpictureRelative(b *testing.B) {
	img := image.NewRGBA(image.Rect(0, 0, 300, 300))
	picture := pic.CreatePictureFromRGBA(img)

	for i := 0; i < b.N; i++ {
		picture.SubPictureRelative(image.Rect(30, 30, 150, 150))
	}
}
