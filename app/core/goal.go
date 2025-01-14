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
		return Goal{}, errors.New(ErrorTitleRequired)
	}

	YMDGolangDateFormat := "2006-01-02"
	_, err := time.Parse(YMDGolangDateFormat, endDate)

	if err != nil {
		return Goal{}, errors.New(ErrorInvalidEndDate)
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
