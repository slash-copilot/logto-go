package core

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserInfo struct {
	customData interface{}
}

func UpdateUserCustomData(userInfoEndpoint, accessToken string, customData interface{}) error {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	customDataBytes, err := json.Marshal(UserInfo{
		customData: customData,
	})
	if err != nil {
		return err
	}

	req, err := http.NewRequest("PATCH", userInfoEndpoint, bytes.NewReader(customDataBytes))
	if err != nil {
		return err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", accessToken))
	response, requestErr := client.Do(req)

	if requestErr != nil {
		return requestErr
	}

	defer response.Body.Close()

	return nil
}
