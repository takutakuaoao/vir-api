package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_目標にタイトルが設定されていない場合にエラー(t *testing.T) {
	emptyTitle := ""
	_, err := NewGoal(emptyTitle)

	assert.Error(t, err)
}
