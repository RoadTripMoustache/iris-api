// Package utils contains all the utils methods for controllers
package utils

import (
	"encoding/json"
	"io"
)

func BodyFormatter(body io.ReadCloser, data interface{}) error {
	defer body.Close()
	bodyBytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyBytes, data)
	return err
}
