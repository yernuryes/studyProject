package services

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

const (
	KeycloakURL  = "http://keycloak:8080"
	Realm        = "lmsSystem"
	ClientID     = "study-client"
	ClientSecret = "LVImsxpdiosIx2FRAWzFSQ9pjfgTOM1q"
)

func Login(username, password string) (map[string]interface{}, error) {

	form := url.Values{}

	form.Add("client_id", ClientID)
	form.Add("client_secret", ClientSecret)
	form.Add("grant_type", "password")
	form.Add("username", username)
	form.Add("password", password)

	resp, err := http.Post(
		KeycloakURL+"/realms/"+Realm+"/protocol/openid-connect/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(form.Encode()),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func Refresh(refreshToken string) (map[string]interface{}, error) {

	form := url.Values{}

	form.Add("client_id", ClientID)
	form.Add("client_secret", ClientSecret)
	form.Add("grant_type", "refresh_token")
	form.Add("refresh_token", refreshToken)

	resp, err := http.Post(
		KeycloakURL+"/realms/"+Realm+"/protocol/openid-connect/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(form.Encode()),
	)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result map[string]interface{}

	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result, nil
}
