package core

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfoPassword struct {
	Password string `json:"password"`
}

func UpdateUserPassword(endpoint string, accessToken string, userId string, newPassword string) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	changePasswordDataBytes, err := json.Marshal(UserInfoPassword{
		Password: newPassword,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", endpoint, bytes.NewReader(changePasswordDataBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	req.Header.Set("Content-Type", "application/json")

	response, requestErr := client.Do(req)

	if requestErr != nil {
		return requestErr
	}

	if response.StatusCode == 401 {
		return ErrorUnAuthorized
	}

	defer response.Body.Close()

	return nil
}
