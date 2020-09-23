package utils

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"image/draw"
	_"image/gif"
	"image/jpeg"
	_"image/jpeg"
	_"image/png"
	"io"
)

func LoadImg(r io.Reader) (image.Image, error) {
	src, _, err := image.Decode(r)
	if err != nil {
		return nil, errors.New("image decode err :" + err.Error())
	}
	return src, nil
}
func ImgToBase64(img image.Image) string {
	emptyBuff := bytes.NewBuffer(nil)
	jpeg.Encode(emptyBuff, img, nil)
	return base64.StdEncoding.EncodeToString(emptyBuff.Bytes())
}

func TrimmingImage(src image.Image, x, y, w, h int) (image.Image, error) {
	var subImg image.Image
	if rgbImg, ok := src.(*image.YCbCr); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.YCbCr) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.RGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.RGBA) //图片裁剪x0 y0 x1 y1
	} else if rgbImg, ok := src.(*image.NRGBA); ok {
		subImg = rgbImg.SubImage(image.Rect(x, y, x+w, y+h)).(*image.NRGBA) //图片裁剪x0 y0 x1 y1
	} else {
		return subImg, errors.New("图片解码失败")
	}
	return subImg, nil
}

func Draw(img image.Image,r image.Image,x,y,w,h int) image.Image {
	r1 := image.Rectangle{
		image.Point{x,y},
		image.Point{x+w,y+10},
	}
	r2 := image.Rectangle{
		image.Point{x,y},
		image.Point{x+10,y+h},
	}
	r3 := image.Rectangle{
		image.Point{x+w,y},
		image.Point{x+w+10,y+h},
	}
	r4 := image.Rectangle{
		image.Point{x,y+h},
		image.Point{x+w+10,y+h+10},
	}
	draw.Draw(img.(draw.Image),r1,r,r1.Min,draw.Src)
	draw.Draw(img.(draw.Image),r2,r,r2.Min,draw.Src)
	draw.Draw(img.(draw.Image),r3,r,r3.Min,draw.Src)
	draw.Draw(img.(draw.Image),r4,r,r4.Min,draw.Src)
	return img
}