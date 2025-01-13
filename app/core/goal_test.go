package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_目標にタイトルが設定されていない場合にエラー(t *testing.T) {
	emptyTitle := ""
	_, err := NewGoal(emptyTitle, "")

	assert.Error(t, err)
}

func Test_目標の日付を登録できる(t *testing.T) {
	endDate := "2025-01-01"
	title := "Test Title"

	goal, _ := NewGoal(title, endDate)

	assert.Equal(t, goal.EndDate, "2025-01-01")
}
