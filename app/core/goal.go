package core

import "errors"

func NewGoal(title string, endDate string) (Goal, error) {
	if title == "" {
		return Goal{}, errors.New("title is required")
	}

	return Goal{
		endDate: endDate,
	}, nil
}

type Goal struct {
	endDate string
}

func (g Goal) SameEndDate(endDate string) bool {
	return g.endDate == endDate
}
