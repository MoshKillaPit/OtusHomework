package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Stats struct {
	Level         string
	CountPerLevel int
}

func AnalyzeLogs(lines []string, logLevel string) Stats {
	var stats Stats
	stats.Level = logLevel
	count := 0

	logLevelUpper := strings.ToUpper(logLevel)
	pattern := fmt.Sprintf("[%s]", logLevelUpper)

	for _, line := range lines {
		if strings.Contains(strings.ToUpper(line), pattern) {
			count++
		}
	}

	stats.CountPerLevel = count
	return stats
}

func main() {
	// Объявляем переменные для флагов
	var filePathFlag string
	var logLevelFlag string
	var outputFlag string

	flag.StringVar(&filePathFlag, "file", "", "Path to the log file (required)")
	flag.StringVar(&logLevelFlag, "level", "info", "Log level to analyze (optional, default: info)")
	flag.StringVar(&outputFlag, "output", "", "Output file path for statistics (optional)")
	flag.Parse()

	if filePathFlag == "" {
		filePathFlag = os.Getenv("LOG_ANALYZER_FILE")
	}
	if logLevelFlag == "info" {
		envLevel := os.Getenv("LOG_ANALYZER_LEVEL")
		if envLevel != "" {
			logLevelFlag = envLevel
		}
	}
	if outputFlag == "" {
		outputFlag = os.Getenv("LOG_ANALYZER_OUTPUT")
	}

	if filePathFlag == "" {
		fmt.Println("Error: path to the log file is required (use -file or LOG_ANALYZER_FILE).")
		os.Exit(1)
	}

	if err := processFile(filePathFlag, logLevelFlag, outputFlag); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func processFile(filePathFlag, logLevelFlag, outputFlag string) error {
	file, err := os.Open(filePathFlag)
	if err != nil {
		return fmt.Errorf("cannot open file %s: %w", filePathFlag, err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		return fmt.Errorf("error reading file: %w", err)
	}

	stats := AnalyzeLogs(lines, logLevelFlag)
	resultStr := fmt.Sprintf("Total lines with level [%s]: %d\n", stats.Level, stats.CountPerLevel)

	if outputFlag != "" {
		if err = writeToFile(outputFlag, resultStr); err != nil {
			return err
		}
	} else {
		fmt.Print(resultStr)
	}

	return nil
}

func writeToFile(outputFlag, resultStr string) error {
	outFile, err := os.Create(outputFlag)
	if err != nil {
		return fmt.Errorf("cannot create output file %s: %w", outputFlag, err)
	}
	defer outFile.Close()

	_, err = outFile.WriteString(resultStr)
	if err != nil {
		return fmt.Errorf("cannot write to output file %s: %w", outputFlag, err)
	}

	return nil
}
