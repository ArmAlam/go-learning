package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type FileManager struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManager) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		fmt.Println("An error occurred on opening the file")
		fmt.Println(err)
		return nil, errors.New("An error occurred on opening the file")
	}

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		file.Close()
		fmt.Println("An error occurred on reading the file")
		fmt.Println(err)

		return nil, errors.New("An error occurred on opening the file")

	}

	file.Close()

	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("An error occurred on creating the file")
	}

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		file.Close()

		return errors.New("An error occurred converting data to JSON")
	}

	file.Close()

	return nil
}

func New(inputPath, outputPath string) FileManager {

	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}
