package main

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func readObject(path string) (string, error) {

	if len(path) != 40 {
		return "", fmt.Errorf("invalid git object")
	}

	dir := fmt.Sprintf(".git/objects/%s", path[:2])
	file := fmt.Sprintf("%s/%s", dir, path[2:])
	reader, err := os.ReadFile(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading file %s", err)
		os.Exit(1)
	}

	gitObject, err := zlib.NewReader(bytes.NewReader(reader))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading git object %s", err)
		os.Exit(1)
	}

	defer gitObject.Close()
	var data bytes.Buffer

	_, err = io.Copy(&data, gitObject)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading git object %s", err)
		os.Exit(1)
	}
	
	parts := strings.SplitN(data.String(), "\x00", 2)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid git object")
	}

	return parts[1], nil

}

func hashObject(path string) error {
	reader, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error in reading file: %s", err)
	}

	contents := fmt.Appendf(nil, "%s %d\x00%s", "blob", len(reader), reader)
	obj := sha1.Sum(contents)
	hash := hex.EncodeToString(obj[:])
	fmt.Println(len(hash))
	dir := fmt.Sprintf(".git/objects/%s", hash[:2])
	err = os.MkdirAll(dir, 0755)
	if err != nil {
		return fmt.Errorf("Error creating directory: %s", err)
	}
	filepath := fmt.Sprintf("%s/%s", dir, hash[2:])
	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("error in creating file: %s", err)
	}
	defer file.Close()

	var content bytes.Buffer
	writer := zlib.NewWriter(&content)
	_, err = writer.Write([]byte(contents))
	if err != nil {
		return fmt.Errorf("error in writing file: %s", err)
	}
	writer.Close()

	// Write the file content to the created file
	_, err = file.Write(content.Bytes())
	if err != nil {
		return fmt.Errorf("error in writing file: %s", err)
	}

	return nil
}
