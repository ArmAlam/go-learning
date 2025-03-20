package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
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

	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		// file.Close() //replaced by defer
		fmt.Println("An error occurred on reading the file")
		fmt.Println(err)

		return nil, errors.New("An error occurred on opening the file")

	}

	// file.Close() //replaced by defer

	return lines, nil

}

func (fm FileManager) WriteResult(data interface{}) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("An error occurred on creating the file")
	}

	defer file.Close()

	time.Sleep(3 * time.Second)

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		// file.Close() //replaced by defer

		return errors.New("An error occurred converting data to JSON")
	}

	// file.Close() //replaced by defer

	return nil
}

func New(inputPath, outputPath string) FileManager {

	return FileManager{
		InputFilePath:  inputPath,
		OutputFilePath: outputPath,
	}

}
