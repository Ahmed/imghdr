// imghdr project imghdr.go
package imghdr

import (
	"os"
	s "strings"
)

const (
	BMP     = 0
	EXIF    = 1
	GIF     = 2
	JPEG    = 3
	PBM     = 4
	PGM     = 5
	PNG     = 6
	PPM     = 7
	RAST    = 8
	RGB     = 9
	XBM     = 10
	OPENEXR = 11
)

//First 32 bytes.
func GetHeader(filename string) string {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	b := make([]byte, 32)
	_, err = f.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)

}
func IsJPEG(header string) bool {
	return header[6:10] == "JFIF"
}
func IsEXIF(header string) bool {
	return header[6:10] == "EXIF"
}
func IsGIF(header string) bool {
	return header[:6] == "GIF87a" || header[:6] == "GIF89a"
}
func IsPNG(header string) bool {
	return header[:8] == "\x89PNG\r\n\x1a\n"
}
func IsRGB(header string) bool {
	return header[:2] == "\x001\x332"
}
func IsPBM(header string) bool {
	return string(header[0]) == "P" && s.Contains("14", string(header[1])) && s.Contains("\t\n\r", string(header[2]))
}
func IsPGM(header string) bool {
	return string(header[0]) == "P" && s.Contains("25", string(header[1])) && s.Contains("\t\n\r", string(header[2]))
}
func IsPPM(header string) bool {
	return string(header[0]) == "P" && s.Contains("36", string(header[1])) && s.Contains("\t\n\r", string(header[2]))
}
func IsRAST(header string) bool {
	return header[:4] == "\x59\xA6\x6A\x95"
}
func IsXBM(header string) bool {
	return header[:7] == "#define"
}
func IsBMP(header string) bool {
	return header[:2] == "BM"
}
func IsOPENEXR(header string) bool {
	return header[:4] == "\x76\x2f\x31\x01"
}
func What(header string) int {
	if IsBMP(header) {
		return BMP
	} else if IsEXIF(header) {
		return EXIF
	} else if IsGIF(header) {
		return GIF
	} else if IsJPEG(header) {
		return JPEG
	} else if IsPBM(header) {
		return PBM
	} else if IsPGM(header) {
		return PGM
	} else if IsPNG(header) {
		return PNG
	} else if IsPPM(header) {
		return PPM
	} else if IsRAST(header) {
		return RAST
	} else if IsRGB(header) {
		return RGB
	} else if IsXBM(header) {
		return XBM
	} else if IsOPENEXR(header) {
		return OPENEXR
	} else {
		panic("unknown type.")
	}
}
func WhatToString(header string) string {
	names := []string{"bmp", "exif", "gif", "jpeg", "pbm", "pgm", "png", "ppm", "rast", "rgb", "xbm", "openexr"}
	return names[What(header)]
}
