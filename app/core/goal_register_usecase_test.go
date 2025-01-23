package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/vir-api/core"
	"testing"
)

func Test_目標を登録する(t *testing.T) {
	db := newMockGoalRepository()
	usecase := core.NewGoalRegisterUsecase(&db)
	usecase.Execute("Title", "2025-01-01")

	assert.True(t, db.calledRegister(1, "Title", "2025-01-01"), true)
}

func newMockGoalRepository() *mockGoalRepository {
	return &mockGoalRepository{}
}

type mockGoalRepository struct {
	calledRegisterArgs []struct {
		title   string
		endDate string
	}
}

func (a *mockGoalRepository) calledRegister(count int, title string, endDate string) bool {
	return false
}

func (a *mockGoalRepository) get(index int) inMemoryGoal {
	return inMemoryGoal{}
}

type inMemoryGoal struct {
	title   string
	endDate string
}
