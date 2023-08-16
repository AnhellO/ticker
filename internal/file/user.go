package file

import (
	"encoding/json"
	"os"
)

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains the
// information belonging to a user
type User struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IpAddress string `json:"ip_address"`
}

// GetData simulates a SELECT query to a DB
// but here this happens for a local file
func GetData(filename string) (Users, error) {
	var users Users

	contents, err := os.ReadFile(filename)
	if err != nil {
		return users, err
	}

	err = json.Unmarshal(contents, &users)
	return users, err
}
