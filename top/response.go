package top

import (
	"encoding/json"

	"github.com/iszhange/lubansdk/top/requests"
)

func ErrorResponse(data []byte) (requests.ErrorResponse, error) {
	var result requests.ErrorResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return requests.ErrorResponse{}, err
	}

	return result, nil
}
