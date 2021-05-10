package utils

import (
	"bytes"
	"path/filepath"
	"os"
	"image/png"
)

func GetStayHydratedPath() string {
	path, err := filepath.Abs("./StayHydrated.exe")
	if err != nil {
		os.Exit(1)
	}
	
	return path
}

func ImageToByteArray() []byte {
	imagePath, _ := filepath.Abs("../misc/logo.png")
	image, _ := os.Open(imagePath)

	defer image.Close()

	img, _ := png.Decode(image)
	
	buf := new(bytes.Buffer)
	png.Encode(buf, img)

	imageBytes := buf.Bytes()

	return imageBytes
}