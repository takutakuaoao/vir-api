package core

import (
	"errors"
	"time"
)

const (
	ErrorTitleRequired  = "目標のタイトルは必須です。"
	ErrorInvalidEndDate = "完了予定日のフォーマットが不正です。"
)

func NewGoal(title string, endDate string) (Goal, error) {
	if title == "" {
		return failNewGoal(errors.New(ErrorTitleRequired))
	}

	if err := validateDateFormat(endDate); err != nil {
		return failNewGoal(errors.New(ErrorInvalidEndDate))
	}

	return Goal{
		title:   title,
		endDate: endDate,
	}, nil
}

func validateDateFormat(date string) error {
	if date == "" {
		return nil
	}

	YMDGolangDateFormat := "2006-01-02"
	_, err := time.Parse(YMDGolangDateFormat, date)

	return err
}

type Goal struct {
	title   string
	endDate string
}

func (g Goal) Equal(other Goal) bool {
	return g.title == other.title && g.endDate == other.endDate
}

func failNewGoal(err error) (Goal, error) {
	return Goal{}, err
}
