package main

import (
	"bytes"
	"image"
	"image/jpeg"

	"github.com/jung-kurt/gofpdf"
)

var kCustomSizeMap = make(map[string]gofpdf.SizeType)

func init() {
	kCustomSizeMap["L"]  = gofpdf.SizeType{Wd: 127, Ht: 89}
	kCustomSizeMap["2L"] = gofpdf.SizeType{Wd: 178, Ht: 127}
	kCustomSizeMap["KG"] = gofpdf.SizeType{Wd: 152, Ht: 102}
}

// CreateDoc ...
func CreateDoc(docSizeName string, imageDataRaw []byte, width, height float64) ([]byte, error) {
	var err error
	img, imageType, err := image.Decode(bytes.NewReader(imageDataRaw))
	if err != nil {
		return nil, err
	}
	var pdf gofpdf.Pdf = nil
	if docSize, ok := kCustomSizeMap[docSizeName]; ok {
		pdf = gofpdf.NewCustom(&gofpdf.InitType{
			OrientationStr: "P",
			UnitStr: "mm",
			SizeStr: "",
			Size: docSize,
			FontDirStr: "",
		})
	} else {
		pdf = gofpdf.New("P", "mm", docSizeName, "")
	}
	pdf.AddPage()
	var imageData []byte
	if !(imageType == "jpeg" || imageType == "png") {
		var buff bytes.Buffer
		err = jpeg.Encode(&buff, img, &jpeg.Options{Quality: 95})
		if err != nil {
			return nil, err
		}
		imageData = buff.Bytes()
		imageType = "jpeg"
	} else {
		imageData = imageDataRaw
	}
	w, h, _ := pdf.PageSize(pdf.PageNo())
	pdf.RegisterImageOptionsReader("image", gofpdf.ImageOptions{ImageType: imageType}, bytes.NewReader(imageData))
	pdf.ImageOptions("image", (w-width)/2, (h-height)/2, width, height, false, gofpdf.ImageOptions{}, 0, "")
	var buff bytes.Buffer
	err = pdf.Output(&buff)
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}
