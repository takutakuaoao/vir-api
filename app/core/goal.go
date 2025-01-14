package core

import (
	"errors"
	"time"
)

func NewGoal(title string, endDate string) (Goal, error) {
	if title == "" {
		return Goal{}, errors.New("title is required")
	}

	YMDGolangDateFormat := "2006-01-02"
	_, err := time.Parse(YMDGolangDateFormat, endDate)

	if err != nil {
		return Goal{}, errors.New("endDate is invalid")
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
