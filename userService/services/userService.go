package services

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
)

const (
	AdminUser     = "admin"
	AdminPassword = "123456"
)

func getAdminToken() (string, error) {

	form := url.Values{}

	form.Add("client_id", ClientID)
	form.Add("client_secret", ClientSecret)
	form.Add("username", AdminUser)
	form.Add("password", AdminPassword)
	form.Add("grant_type", "password")

	resp, err := http.Post(
		KeycloakURL+"/realms/"+Realm+"/protocol/openid-connect/token",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(form.Encode()),
	)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	var result map[string]interface{}

	body, _ := io.ReadAll(resp.Body)

	json.Unmarshal(body, &result)

	token, ok := result["access_token"].(string)

	if !ok {
		return "", errors.New("admin token not found")
	}

	return token, nil
}

func CreateUser(
	username string,
	email string,
	lastName string,
	firstName string,
	password string,
	role string,
) error {

	adminToken, err := getAdminToken()

	if err != nil {
		return err
	}

	userBody := map[string]interface{}{
		"username":  username,
		"email":     email,
		"lastName":  lastName,
		"firstName": firstName,
		"enabled":   true,
		"credentials": []map[string]interface{}{
			{
				"type":      "password",
				"value":     password,
				"temporary": false,
			},
		},
	}

	jsonBody, _ := json.Marshal(userBody)

	req, _ := http.NewRequest(
		"POST",
		KeycloakURL+"/admin/realms/"+Realm+"/users",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	println("STATUS:", resp.Status)
	println("BODY:", string(body))

	if resp.StatusCode != 201 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	userID, err := getUserID(username, adminToken)
	if err != nil {
		return err
	}

	roleData, err := getRole(role, adminToken)
	if err != nil {
		return err
	}

	err = assignRole(userID, roleData, adminToken)
	if err != nil {
		return err
	}

	return nil
}

func getUserID(username string, adminToken string) (string, error) {

	req, _ := http.NewRequest(
		"GET",
		KeycloakURL+"/admin/realms/"+Realm+"/users?username="+username,
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var users []map[string]interface{}

	err = json.Unmarshal(body, &users)
	if err != nil {
		return "", err
	}

	if len(users) == 0 {
		return "", errors.New("user not found")
	}

	id, _ := users[0]["id"].(string)

	return id, nil
}

func getRole(roleName string, adminToken string) (map[string]interface{}, error) {

	req, _ := http.NewRequest(
		"GET",
		KeycloakURL+"/admin/realms/"+Realm+"/roles/"+roleName,
		nil,
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var role map[string]interface{}

	err = json.Unmarshal(body, &role)
	if err != nil {
		return nil, err
	}

	return role, nil
}

func assignRole(userID string, role map[string]interface{}, adminToken string) error {

	roleBody, _ := json.Marshal([]map[string]interface{}{role})

	req, _ := http.NewRequest(
		"POST",
		KeycloakURL+"/admin/realms/"+Realm+"/users/"+userID+"/role-mappings/realm",
		bytes.NewBuffer(roleBody),
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}

func UpdateUser(
	username string,
	email string,
	firstName string,
	lastName string,
) error {

	adminToken, err := getAdminToken()
	if err != nil {
		return err
	}

	userID, err := getUserID(username, adminToken)
	if err != nil {
		return err
	}

	bodyMap := map[string]interface{}{
		"email":     email,
		"firstName": firstName,
		"lastName":  lastName,
	}

	jsonBody, _ := json.Marshal(bodyMap)

	req, _ := http.NewRequest(
		"PUT",
		KeycloakURL+"/admin/realms/"+Realm+"/users/"+userID,
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}

func ChangePassword(
	username string,
	password string,
) error {

	adminToken, err := getAdminToken()
	if err != nil {
		return err
	}

	userID, err := getUserID(username, adminToken)
	if err != nil {
		return err
	}

	bodyMap := map[string]interface{}{
		"type":      "password",
		"value":     password,
		"temporary": false,
	}

	jsonBody, _ := json.Marshal(bodyMap)

	req, _ := http.NewRequest(
		"PUT",
		KeycloakURL+"/admin/realms/"+Realm+"/users/"+userID+"/reset-password",
		bytes.NewBuffer(jsonBody),
	)

	req.Header.Set("Authorization", "Bearer "+adminToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		return errors.New(string(body))
	}

	return nil
}
