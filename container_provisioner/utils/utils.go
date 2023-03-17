package utils

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
)

// ErrorHandler is a generic error handler
func ErrorHandler(err error) {
	if err != nil {
		formattedError := fmt.Errorf("Error: %w", err)
		fmt.Println(formattedError)
		os.Exit(0)
	}
}

func WriteToFile(filename string, tarF io.ReadCloser) error {

	// Create the file
	out, err := os.Create(filename)
	ErrorHandler(err)
	defer out.Close()

	// Untar the file
	// Note: This is not a generic untar function. It only works for a single file
	/**
		A tar file is a collection of binary data segments (usually sourced from files). Each segment starts with a header that contains metadata about the binary data, that follows it, and how to reconstruct it as a file.

	+---------------------------+
	| [name][mode][uid][guild]  |
	| ...                       |
	+---------------------------+
	| XXXXXXXXXXXXXXXXXXXXXXXXX |
	| XXXXXXXXXXXXXXXXXXXXXXXXX |
	| XXXXXXXXXXXXXXXXXXXXXXXXX |
	+---------------------------+
	| [name][mode][uid][guild]  |
	| ...                       |
	+---------------------------+
	| XXXXXXXXXXXXXXXXXXXXXXXXX |
	| XXXXXXXXXXXXXXXXXXXXXXXXX |
	+---------------------------+
		**/

	// Read the tar file
	tarReader := tar.NewReader(tarF)

	// Go to the next entry in the tar file
	_, err = tarReader.Next()
	ErrorHandler(err)

	// Write the file to disk
	_, err = io.Copy(out, tarReader)
	ErrorHandler(err)

	return nil
}
