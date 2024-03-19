package utils

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ZipDir(sourceDir string, archivePath string) error {
	// Create a new zip file
	archive, err := os.Create(archivePath)
	if err != nil {
		return fmt.Errorf("error creating archive: %w", err)
	}
	defer archive.Close()

	// Create a new zip writer
	writer := zip.NewWriter(archive)
	defer writer.Close()

	// Function to recursively add files to the zip archive
	var walkFunc filepath.WalkFunc = func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the source directory itself
		if path == sourceDir {
			return nil
		}

		// Relative path within the archive
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Create a new file header within the zip archive
		header := &zip.FileHeader{
			Name:   relPath,
			Method: zip.Deflate, // Optional compression method
		}

		// Set permissions based on file type (corrected field name)

		// Create the file entry in the zip archive
		f, err := writer.CreateHeader(header)
		if err != nil {
			return err
		}

		// Open the file to be added (if it's a regular file)
		if !info.IsDir() {
			sourceFile, err := os.Open(path)
			if err != nil {
				return err
			}
			defer sourceFile.Close()

			// Copy file contents to the zip archive
			_, err = io.Copy(f, sourceFile)
			return err
		}

		return nil // Continue walking the directory structure for subdirectories
	}

	// Walk through the directory structure
	err = filepath.Walk(sourceDir, walkFunc)
	if err != nil {
		return fmt.Errorf("error walking directory: %w", err)
	}

	return nil
}

func TarDir(sourceDir string, archiveFile string) error {
	// Create a tar archive file
	outFile, err := os.Create(archiveFile)
	if err != nil {
		return fmt.Errorf("failed to create tar archive: %w", err)
	}
	defer outFile.Close()

	// Create a new tar writer
	writer := tar.NewWriter(outFile)
	defer writer.Close()

	// Walk through the directory recursively
	return filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the source directory itself
		if path == sourceDir {
			return nil
		}

		// Get relative path within the archive
		relPath, err := filepath.Rel(sourceDir, path)
		if err != nil {
			return err
		}

		// Create a new tar header
		header := &tar.Header{
			Name: relPath,
			Mode: 0777, // Set appropriate permissions (adjust as needed)
			Size: info.Size(),
		}

		// If it's a directory, set the type accordingly
		if info.IsDir() {
			header.Typeflag = tar.TypeDir
		} else {
			header.Typeflag = tar.TypeReg
		}

		// Write the header to the archive
		err = writer.WriteHeader(header)
		if err != nil {
			return err
		}

		// If it's a regular file, copy contents to the archive
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()

			_, err = io.Copy(writer, file)
			return err
		}

		return nil
	})
}
