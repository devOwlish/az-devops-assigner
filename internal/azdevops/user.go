package azdevops

import (
	"encoding/json"
	"fmt"
)

// userList is a struct for unmarshalling the response from the API;
// we only need several fields from the response, so it doesn't make any sense to parse the whole response.
type userList struct {
	Members []struct {
		ID   string `json:"id"`
		User struct {
			MailAddress string `json:"mailAddress"`
		} `json:"user"`
	} `json:"members"`
}

// GetUserIDByEmail returns a user ID by email address.
func GetUserIDByEmail(email string) (string, error) {
	users, err := SendRequest("userentitlements", "user", "", "GET", nil)
	if err != nil {
		return "", fmt.Errorf("failed to get users: %w", err)
	}

	var list userList

	err = json.Unmarshal(users.Body(), &list)
	if err != nil {
		return "", fmt.Errorf("failed to unmarshal users: %w", err)
	}

	for _, user := range list.Members {
		if user.User.MailAddress == email {
			return user.ID, nil
		}
	}

	return "", fmt.Errorf("user with email %s not found", email)
}
