package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/vir-api/core"
	"testing"
)

func Test_目標にタイトルが設定されていない場合にエラー(t *testing.T) {
	emptyTitle := ""
	_, err := core.NewGoal(emptyTitle, "")

	assert.Error(t, err)
}

func Test_目標の日付を登録できる(t *testing.T) {
	endDate := "2025-01-01"
	title := "Test Title"

	goal, _ := core.NewGoal(title, endDate)

	assert.True(t, goal.SameEndDate("2025-01-01"))
}

func Test_目標の日付を確認できる(t *testing.T) {
	endDate := "2025-01-01"

	goal, _ := core.NewGoal("Test Title", endDate)

	assert.Equal(t, true, goal.SameEndDate("2025-01-01"))
}

func Test_目標の日付の形式が正しくない場合にエラー(t *testing.T) {
	cases := []struct {
		name              string
		invalidDateFormat string
	}{
		{
			name:              "文字列",
			invalidDateFormat: "文字列形式",
		},
		{
			name:              "年のみ",
			invalidDateFormat: "2025",
		},
		{
			name:              "年月",
			invalidDateFormat: "2025-01",
		},
		{
			name:              "年月日時",
			invalidDateFormat: "2025-01-01 00",
		},
		{
			name:              "年月日時分",
			invalidDateFormat: "2025-01-01 00:00",
		},
		{
			name:              "年月日時分秒",
			invalidDateFormat: "2025-01-01 00:00:00",
		},
		{
			name:              "形式が変",
			invalidDateFormat: "2025/01/01",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			_, err := core.NewGoal("Title", tt.invalidDateFormat)

			assert.Error(t, err)
		})
	}
}
