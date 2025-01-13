package core

import "errors"

func NewGoal(title string, endDate string) (Goal, error) {
	if title == "" {
		return Goal{}, errors.New("title is required")
	}

	return Goal{
		EndDate: endDate,
	}, nil
}

type Goal struct {
	EndDate string
}
