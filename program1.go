// Program Name: program1.go
// Author: Michael Cooper
// Date: 3/8/2024
// Class and Section: CS 424-01
// Description:
// This program applies user-specified filters to an image in PPM ASCII format and saves the modified image to a new file.
// Supported filters include flipping horizontally and vertically, converting to grayscale, inverting colors,
// flattening colors, and applying extreme contrast.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Command-line flags
var (
	file      string
	flipH     bool
	flipV     bool
	grayscale bool
	invert    bool
	flatten   string
	extreme   bool
)

func init() {
	flag.StringVar(&file, "file", "", "Specify the name of the PPM image file to process.")
	flag.BoolVar(&flipH, "h", false, "Flip image horizontally")
	flag.BoolVar(&flipV, "v", false, "Flip image vertically")
	flag.BoolVar(&grayscale, "g", false, "Convert file to grayscale")
	flag.BoolVar(&invert, "i", false, "Invert image colors")
	flag.StringVar(&flatten, "f", "", "Flatten the specified colors. Enter one or more colors (e.g., -f gb or -f b).")
	flag.BoolVar(&extreme, "x", false, "Apply extreme contrast filter")
}

func main() {
	flag.Parse()

	if file == "" {
		fmt.Println("You must specify the input file using the -file flag.")
		flag.Usage()
		return
	}

	pixels, width, height, err := readPPMFile(file)
	if err != nil {
		log.Fatalf("Error reading PPM file: %v", err)
	}

	if flipH {
		pixels = flipHorizontally(pixels, width, height)
	}
	if flipV {
		pixels = flipVertically(pixels, width, height)
	}
	if grayscale {
		pixels = convertToGrayscale(pixels)
	}
	if invert {
		pixels = invertColors(pixels)
	}
	if flatten != "" {
		pixels = flattenColors(pixels, flatten)
	}
	if extreme {
		pixels = applyExtremeContrast(pixels)
	}

	outputFileName := strings.TrimSuffix(file, ".ppm") + "_transformed.ppm"
	if err := writePPMFile(outputFileName, pixels, width, height); err != nil {
		log.Fatalf("Error writing transformed PPM file: %v", err)
	}

	fmt.Println("Transformed image saved to", outputFileName)
}

func readPPMFile(fileName string) ([]Pixel, int, int, error) {
	inFile, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, err
	}
	defer inFile.Close()

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanWords)

	// Read PPM header (P3 format, dimensions, and max color value)
	if !scanner.Scan() || scanner.Text() != "P3" {
		return nil, 0, 0, fmt.Errorf("file is not in PPM P3 format")
	}
	scanner.Scan()
	ppmCols, err := strconv.Atoi(scanner.Text())
	checkError(err)
	scanner.Scan()
	ppmRows, err := strconv.Atoi(scanner.Text())
	checkError(err)
	scanner.Scan() // Skip max color value line

	// Set up a two-dimensional slice of each pixel's RGB values
	inputRGBValues := make([][]Pixel, ppmRows)
	for i := range inputRGBValues {
		inputRGBValues[i] = make([]Pixel, ppmCols)
		for j := range inputRGBValues[i] {
			var pixel Pixel
			for k := 0; k < 3; k++ {
				scanner.Scan()
				val, err := strconv.Atoi(scanner.Text())
				checkError(err)
				switch k {
				case 0:
					pixel.R = val
				case 1:
					pixel.G = val
				case 2:
					pixel.B = val
				}
			}
			inputRGBValues[i][j] = pixel
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}
	// Convert the 2D Pixel slice to a 1D Pixel slice for further processing
	pixels := make([]Pixel, 0, ppmRows*ppmCols)
	for _, row := range inputRGBValues {
		pixels = append(pixels, row...)
	}

	return pixels, ppmCols, ppmRows, nil
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func writePPMFile(fileName string, pixels []Pixel, width, height int) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	header := fmt.Sprintf("P3\n%d %d\n255\n", width, height)
	if _, err := file.WriteString(header); err != nil {
		return err
	}

	for i, pixel := range pixels {
		if _, err := fmt.Fprintf(file, "%d %d %d ", pixel.R, pixel.G, pixel.B); err != nil {
			return err
		}
		if (i+1)%width == 0 {
			if _, err := file.WriteString("\n"); err != nil {
				return err
			}
		}
	}

	return nil
}

func flipHorizontally(pixels []Pixel, width, height int) []Pixel {
	for y := 0; y < height; y++ {
		for x := 0; x < width/2; x++ {
			opposite := width - 1 - x
			leftIndex := y*width + x
			rightIndex := y*width + opposite
			pixels[leftIndex], pixels[rightIndex] = pixels[rightIndex], pixels[leftIndex]
		}
	}
	return pixels
}

func flipVertically(pixels []Pixel, width, height int) []Pixel {
	for y := 0; y < height/2; y++ {
		for x := 0; x < width; x++ {
			opposite := (height - 1 - y) * width
			current := y * width
			pixels[current+x], pixels[opposite+x] = pixels[opposite+x], pixels[current+x]
		}
	}
	return pixels
}

func convertToGrayscale(pixels []Pixel) []Pixel {
	for i := range pixels {
		gray := (pixels[i].R + pixels[i].G + pixels[i].B) / 3
		pixels[i].R, pixels[i].G, pixels[i].B = gray, gray, gray
	}
	return pixels
}

func invertColors(pixels []Pixel) []Pixel {
	for i := range pixels {
		pixels[i].R = 255 - pixels[i].R
		pixels[i].G = 255 - pixels[i].G
		pixels[i].B = 255 - pixels[i].B
	}
	return pixels
}

func flattenColors(pixels []Pixel, colors string) []Pixel {
	for i := range pixels {
		if strings.Contains(colors, "r") {
			pixels[i].R = 0
		}
		if strings.Contains(colors, "g") {
			pixels[i].G = 0
		}
		if strings.Contains(colors, "b") {
			pixels[i].B = 0
		}
	}
	return pixels
}

func applyExtremeContrast(pixels []Pixel) []Pixel {
	for i := range pixels {
		pixels[i].R = contrastValue(pixels[i].R)
		pixels[i].G = contrastValue(pixels[i].G)
		pixels[i].B = contrastValue(pixels[i].B)
	}
	return pixels
}

func contrastValue(color int) int {
	if color > 127 {
		return 255
	}
	return 0
}

type Pixel struct {
	R, G, B int
}
