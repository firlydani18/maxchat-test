package utils

import (
	"bufio"
	"os"
	"strings"

	"maxchat-backend/model"
)

func LoadData(filename string) ([]model.Item, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []model.Item
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		tech := strings.Split(parts[4], ";")
		item := model.Item{
			Code:        parts[0],
			Name:        parts[1],
			Description: parts[2],
			Model:       parts[3],
			Tech:        tech,
			Status:      parts[5],
		}
		items = append(items, item)
	}
	return items, scanner.Err()
}

func Contains(slice []string, str string) bool {
	for _, v := range slice {
		if v == str {
			return true
		}
	}
	return false
}
