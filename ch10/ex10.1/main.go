package main

import (
	"flag"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
)

var filePathFrom = flag.String("pf", "", "path of file to convert")
var filePathTo = flag.String("pt", "", "path of file to convert to")
var flagFormatTo = flag.String("f", "jpeg", "image format to convert to")

// go build main.go
// ./ex10.1.exe -pf cat.jpg -pt a.jpg -f png
func main() {
	flag.Parse()

	pathFrom := *filePathFrom
	pathTo := *filePathTo
	format := *flagFormatTo

	if len(pathFrom) == 0 {
		fmt.Fprintln(os.Stderr, "error: file path must not be empty")
		os.Exit(1)
	}

	if len(format) == 0 {
		fmt.Fprintln(os.Stderr, "error: format must not be empty")
		os.Exit(1)
	}

	fileFrom, err := os.Open(pathFrom)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: file from: %v\n", err)
		os.Exit(1)
	}

	fileTo, err := os.Create(pathTo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: file to: %v\n", err)
		os.Exit(1)
	}

	if err := convert(fileFrom, fileTo, format); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func convert(in io.Reader, out io.Writer, format string) error {
	img, kind, err := image.Decode(in)
	if err != nil {
		return err
	}

	fmt.Fprintln(os.Stderr, "Input format: ", kind)

	switch format {
	case "jpeg":
		return jpeg.Encode(out, img, &jpeg.Options{Quality: 95})
	case "png":
		return png.Encode(out, img)
	case "gif":
		return gif.Encode(out, img, &gif.Options{})
	default:
		return fmt.Errorf("unkown format")
	}
}
