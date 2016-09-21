package main

import (
	"github.com/nfnt/resize"
	"image/jpeg"
	"image/png"
	"os"
	"path"
	"strings"
)

func Resize(name string) {
	decoder := jpeg.Decode
	var JPG bool

	if !strings.HasSuffix(name, ".jpg") && !strings.HasSuffix(name, ".png") {
		return
	}

	JPG = strings.HasSuffix(name, ".jpg")
	if !JPG {
		decoder = png.Decode
	}

	fileinfo, _ := os.Stat(name)
	if fileinfo.Size()/1024 < 1024 {
		return
	}

	file, err := os.Open(name)
	if err != nil {
		return
	}

	img, err := decoder(file)
	if err != nil {
		return
	}
	file.Close()

	m := resize.Resize(600, 0, img, resize.Lanczos3)

	d := path.Dir(name)
	var temp string = path.Join(d, "_temp_")
	out, err := os.Create(temp)
	if err != nil {
		return
	}
	// write new image to file
	if JPG {
		jpeg.Encode(out, m, nil)
	} else {
		png.Encode(out, m)
	}
	out.Close()
	os.Rename(temp, name)
}
