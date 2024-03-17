package utils

import (
	"archive/tar"
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func ZipDir(sourceDir string, archiveFile string) error {
	// Create a zip file
	outFile, err := os.Create(archiveFile)
	if err != nil {
		return fmt.Errorf("failed to create zip file: %w", err)
	}
	defer outFile.Close()

	// Create a new zip writer
	writer := zip.NewWriter(outFile)
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

		// Create a new zip file entry
		f, err := writer.Create(relPath)
		if err != nil {
			return err
		}

		// If it's a directory, we don't need to add any content
		if info.IsDir() {
			return nil
		}

		// Open the file to be added
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		// Copy file contents to the zip archive
		_, err = io.Copy(f, file)
		return err
	})
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
			Mode: 0644, // Set appropriate permissions (adjust as needed)
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
