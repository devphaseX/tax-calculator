package filemanager

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManger struct {
	InputFilePath  string
	OutputFilePath string
}

func (fm FileManger) ReadLines() ([]string, error) {
	file, err := os.Open(fm.InputFilePath)

	if err != nil {
		return nil, errors.New("failed to open file")
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()

	if err != nil {
		return nil, errors.New("failed to read line in file")
	}

	return lines, nil
}

func (fm FileManger) WriteResult(data any) error {
	file, err := os.Create(fm.OutputFilePath)

	if err != nil {
		return errors.New("failed to create file")
	}

	defer file.Close()

	encoder := json.NewEncoder(file)

	err = encoder.Encode(data)

	if err != nil {
		return errors.New("failed to convert data to JSON")
	}

	return nil
}

func New(inputPath string, outputPath string) FileManger {
	return FileManger{InputFilePath: inputPath, OutputFilePath: outputPath}
}
