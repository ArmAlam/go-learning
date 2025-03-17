package filemanager

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)

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
