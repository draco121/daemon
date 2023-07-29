package utils

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// func Archive(BundlePath string, BundleArchivePath string) error {

// 	// tar + gzip
// 	var buf bytes.Buffer
// 	err := compress(BundlePath, &buf)
// 	if err != nil {
// 		return err
// 	}
// 	// write the .tar.gzip
// 	fileToWrite, err := os.OpenFile(BundleArchivePath, os.O_CREATE|os.O_RDWR, os.FileMode(0600))
// 	if err != nil {
// 		return err
// 	}
// 	if _, err := io.Copy(fileToWrite, &buf); err != nil {
// 		return err
// 	}
// 	return nil
// }

func Archive(BundlePath string, BundleArchivePath string) error {
	// Create the target file
	targetFile, err := os.Create(BundleArchivePath)
	if err != nil {
		return fmt.Errorf("failed to create target file: %v", err)
	}
	defer targetFile.Close()
	// Create a gzip writer
	gzipWriter := gzip.NewWriter(targetFile)
	defer gzipWriter.Close()

	// Create a tar writer
	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	// Walk through the source directory recursively
	err = filepath.Walk(BundlePath, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip the sourcePath itself
		if filePath == BundlePath {
			return nil
		}

		// Get the relative path of the file/directory within the source directory
		relPath, err := filepath.Rel(BundlePath, filePath)
		if err != nil {
			return fmt.Errorf("failed to get relative path: %v", err)
		}

		// Create a tar header for the current file/directory
		tarHeader, err := tar.FileInfoHeader(info, "")
		if err != nil {
			return fmt.Errorf("failed to create tar header: %v", err)
		}

		// Replace Windows backslashes with forward slashes in the tar header
		tarHeader.Name = strings.ReplaceAll(filepath.ToSlash(relPath), "\\", "/")

		// Write the tar header
		if err := tarWriter.WriteHeader(tarHeader); err != nil {
			return fmt.Errorf("failed to write tar header: %v", err)
		}

		// For files, copy the file data into the tar
		if !info.IsDir() {
			file, err := os.Open(filePath)
			if err != nil {
				return fmt.Errorf("failed to open file: %v", err)
			}
			defer file.Close()

			_, err = io.Copy(tarWriter, file)
			if err != nil {
				return fmt.Errorf("failed to copy data to tar: %v", err)
			}
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to walk through source directory: %v", err)
	}

	return nil
}

// func compress(src string, buf io.Writer) error {
// 	// tar > gzip > buf
// 	zr := gzip.NewWriter(buf)
// 	tw := tar.NewWriter(zr)

// 	// walk through every file in the folder
// 	filepath.Walk(src, func(file string, fi os.FileInfo, err error) error {
// 		// generate tar header

// 		header, err := tar.FileInfoHeader(fi, file)
// 		if err != nil {
// 			return err
// 		}

// 		relativePath := strings.TrimPrefix(filepath.ToSlash(file), RootDir+"/")
// 		header.Name = filepath.ToSlash(relativePath)

// 		// write
// 		if err := tw.WriteHeader(header); err != nil {
// 			return err
// 		}
// 		// if not a dir, write file content
// 		if !fi.IsDir() {
// 			data, err := os.Open(file)
// 			if err != nil {
// 				return err
// 			}
// 			if _, err := io.Copy(tw, data); err != nil {
// 				return err
// 			}
// 		}
// 		return nil
// 	})

// 	// produce tar
// 	if err := tw.Close(); err != nil {
// 		return err
// 	}
// 	// produce gzip
// 	if err := zr.Close(); err != nil {
// 		return err
// 	}
// 	//
// 	return nil
// }
