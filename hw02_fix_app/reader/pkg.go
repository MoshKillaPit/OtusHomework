package reader

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/fixme_my_friend/hw02_fix_app/types"
)

func ReadJSON(filePath string) ([]types.Employee, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("could not open file: %v", err)
	}
	defer f.Close()

	byteJSON, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("could not read file: %v", err)
	}

	var data []types.Employee
	err = json.Unmarshal(byteJSON, &data)
	if err != nil {
		return nil, fmt.Errorf("could not unmarshal JSON: %v", err)
	}

	return data, nil
}
