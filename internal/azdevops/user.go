package azdevops

import (
	"encoding/json"
	"fmt"
)

type userList struct {
	Members []struct {
		ID   string `json:"id"`
		User struct {
			MailAddress string `json:"mailAddress"`
		} `json:"user"`
	} `json:"members"`
}

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
