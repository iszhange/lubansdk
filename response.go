package lubantop

import (
	"encoding/json"
	"errors"
	"lubantop/requests"
)

func ErrorResponse(data []byte) (requests.ErrorResponse, error) {
	var result requests.ErrorResponse
	err := json.Unmarshal(data, &result)
	if err != nil {
		return requests.ErrorResponse{}, err
	}
	if result.ErrorResponse.Code != 0 {
		return requests.ErrorResponse{}, errors.New("响应为空")
	}

	return result, nil
}
