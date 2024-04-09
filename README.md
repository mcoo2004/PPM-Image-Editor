# PPM-Image-Editor
A Go-based image processing tool that applies user-specified filters to PPM ASCII format images, including flipping, grayscale conversion, color inversion, color flattening, and extreme contrast adjustments.

This project is a Go language application designed to apply various filters to images in the PPM (Portable Pixmap) ASCII format, outputting the results to a new file. The application manipulates the ASCII text representation of PPM images, which encode RGB values for each pixel, to implement the transformations.

The application supports the following image filters, which can be applied individually or in combination as specified by the user through command-line arguments:

  Flip Image Horizontally: Mirrors the image across a vertical axis.

  Flip Image Vertically: Mirrors the image across a horizontal axis.
  
  Convert to Greyscale: Averages the RGB values of each pixel to remove color.
  
  Invert Image Colors: Subtracts each RGB value from 255 to negate the color.
  
  Flatten Color (Red, Green, Blue): Removes all intensity of a specified color channel by setting it to 0.
  
  Extreme Contrast: Enhances contrast by setting RGB values above half the maximum (127.5) to 255 and all others to 0.
  
  User-specified filters are implemented via command-line arguments, leveraging Go's flag module for argument parsing. Argument flags include help text to guide users on proper usage. The ability to view and manipulate     command-line inputs allows for a highly customizable image processing experience.

Images are outputted in the same PPM ASCII format, and viewing the modified images requires a PPM-compatible image viewer, such as GIMP, which is freely available and supports multiple platforms.


Requirements:

Go programming language setup on your machine.

Input images in PPM ASCII format (P3 type).


Usage:

First, compile the program using Go's build command:

"go build program1.go"
Then, run the compiled program with desired flags. For example, to apply a grayscale filter to input.ppm and save the result:

"./program1 -file=input.ppm -g"


Here are the flags you can use:

-file: Specify the input image file (required).

-h: Flip the image horizontally.

-v: Flip the image vertically.

-g: Convert the image to grayscale.

-i: Invert the image colors.

-f: Flatten specified colors (e.g., -f rb for red and blue).

-x: Apply extreme contrast.

The output file will be saved in the same directory with "_transformed.ppm" appended to the original file name.

Viewing the Result:

To view the resulting PPM image, use an image viewer that supports the PPM format. GIMP is a free option that works across multiple platforms.
