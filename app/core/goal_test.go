package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/vir-api/core"
	"testing"
)

func Test_目標にタイトルを設定する(t *testing.T) {
	title := "Title Test"
	goal, _ := core.NewGoal(title, "")
	expect, _ := core.NewGoal(title, "")

	assert.True(t, goal.Equal(expect))
}

func Test_目標にタイトルが設定されていない場合にエラー(t *testing.T) {
	emptyTitle := ""
	_, err := core.NewGoal(emptyTitle, "")

	assert.Errorf(t, err, core.ErrorTitleRequired)
}

func Test_目標の日付を登録できる(t *testing.T) {
	endDate := "2025-01-01"
	title := "Test Title"

	goal, _ := core.NewGoal(title, endDate)
	expect, _ := core.NewGoal(title, endDate)

	assert.True(t, goal.Equal(expect))
}

func Test_目標の日付は空でもOK(t *testing.T) {
	emptyEndDate := ""

	_, err := core.NewGoal("Test Title", emptyEndDate)

	assert.Nil(t, err)
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

			assert.Errorf(t, err, core.ErrorInvalidEndDate)
		})
	}
}

func Test_equal(t *testing.T) {
	type goalParam struct {
		title   string
		endDate string
	}

	cases := []struct {
		name   string
		self   goalParam
		other  goalParam
		expect bool
	}{
		{
			name:   "equal",
			self:   goalParam{title: "Title", endDate: "2025-01-01"},
			other:  goalParam{title: "Title", endDate: "2025-01-01"},
			expect: true,
		},
		{
			name:   "not equal",
			self:   goalParam{title: "Title", endDate: "2025-01-01"},
			other:  goalParam{title: "Title2", endDate: "2025-01-02"},
			expect: false,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			self, _ := core.NewGoal(tt.self.title, tt.self.endDate)
			other, _ := core.NewGoal(tt.other.title, tt.other.endDate)

			assert.Equal(t, tt.expect, self.Equal(other))
		})
	}
}
