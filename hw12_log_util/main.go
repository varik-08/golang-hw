package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/pflag"
)

func main() {
	filePath, level, outputPath, err := getVariables()
	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("ошибка при открытии файла: " + err.Error())

		return
	}
	defer file.Close()

	br := bufio.NewReader(file)

	for {
		line, readErr := br.ReadString('\n')
		if readErr != nil {
			if readErr.Error() == "EOF" {
				break
			}

			fmt.Println("Ошибка при чтении файла:" + readErr.Error())

			return
		}

		date, message := AnalyseLog(line, level)

		if message != "" {
			writeLogErr := writeLog(date, message, outputPath)
			if writeLogErr != nil {
				fmt.Println("Ошибка при записи лога в файл:" + writeLogErr.Error())

				return
			}
		}
	}
}

func getVariables() (string, string, string, error) {
	var file, level, output string
	pflag.StringVarP(&file, "file", "f", "", "path to log file")
	pflag.StringVarP(&level, "level", "l", "info", "level log")
	pflag.StringVarP(&output, "output", "o", "", "path to log output file")
	pflag.Parse()

	err := LoadEnvFile()
	if err != nil {
		wErr := fmt.Errorf("ошибка загрузки env файла: " + err.Error())
		return "", "", "", wErr
	}

	if file == "" {
		file = os.Getenv("LOG_ANALYZER_FILE")

		if file == "" {
			wErr := errors.New("не указан путь к лог-файлу")

			return "", "", "", wErr
		}
	}
	if level == "info" {
		envLevel := os.Getenv("LOG_ANALYZER_LEVEL")

		if envLevel != "" {
			level = envLevel
		}

		resCheckLevel := CheckLogLevel(level)
		if !resCheckLevel {
			levelErr := errors.New("недопустимый уровень логирования")

			return "", "", "", levelErr
		}
	}
	if output == "" {
		envOutput := os.Getenv("LOG_ANALYZER_OUTPUT")

		if envOutput != "" {
			output = envOutput
		}
	}

	return file, level, output, nil
}

func LoadEnvFile() error {
	if err := godotenv.Load(".env"); err != nil {
		return fmt.Errorf("parse config: %w", err)
	}

	return nil
}

func writeLog(date string, message, outputPath string) error {
	if outputPath != "" {
		//nolint
		file, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return fmt.Errorf("ошибка открытия файла для записи лога: %w", err)
		}
		defer file.Close()

		_, err = file.WriteString(date + ": " + message + "\n")

		if err != nil {
			return fmt.Errorf("ошибка при записи лога в файл: %w", err)
		}
	} else {
		fmt.Println(date + ": " + message)
	}

	return nil
}
