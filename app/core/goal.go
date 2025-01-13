package core

import "errors"

func NewGoal(title string) (string, error) {
	if title == "" {
		return "", errors.New("title is required")
	}

	return "", nil
}
