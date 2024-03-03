// Package api aaa
package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

// CreateUser aaa
func CreateUser(c echo.Context) error {
	firstName := c.FormValue("firstName")
	lastName := c.FormValue("lastName")
	email := c.FormValue("email")
	age := c.FormValue("age")
	experimentID := c.FormValue("experimentID")

	intAge, err := strconv.Atoi(age)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	}

	intExperimentID, err := strconv.Atoi(experimentID)
	if err != nil {
		fmt.Println("Error converting string to integer:", err)
	}

	data := map[string]interface{}{
		"firstName":    firstName,
		"lastName":     lastName,
		"email":        email,
		"age":          intAge,
		"experimentId": intExperimentID,
	}

	fmt.Println("First Name:", data["firstName"])
	fmt.Println("Last Name:", data["lastName"])
	fmt.Println("Email:", data["email"])
	fmt.Println("Age:", data["age"])
	fmt.Println("Experiment ID:", data["experimentId"])

	content := map[string]interface{}{
		"member": data,
	}

	jsonValue, err := json.Marshal(content)
	if err != nil {
		return err
	}

	// Make the POST request
	resp, err := http.Post("http://localhost:8181/api/members", "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		bodyString := string(bodyBytes)
		fmt.Println(bodyString)
		return errors.New("failed to create user")
	}

	return nil
}

// ReadAllUsers aaa
func ReadAllUsers() ([]map[string]interface{}, error) {
	response, err := http.Get("http://localhost:8181/api/members")
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %w", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("Failed to read response body: %w", err)
	}

	var responseData map[string][]map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		return nil, fmt.Errorf("Failed to decode JSON: %w", err)
	}

	members, ok := responseData["members"]
	if !ok {
		return nil, fmt.Errorf("Response does not contain 'members' key")
	}

	return members, nil
}

// ReadUser aaa
func ReadUser(id int) error {
	response, err := http.Get("http://localhost:8181/api/members/" + fmt.Sprint(id))
	if err != nil {
		return err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// print the response
	fmt.Println(string(body))
	return nil
}

// UpdateUser aaa
func UpdateUser(c echo.Context) error {
	return nil
}

// DeleteUser aaa
func DeleteUser(c echo.Context) error {
	return nil
}
