package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/vir-api/core"
	"testing"
)

func Test_目標を登録する(t *testing.T) {
	db := newArrayRepository()
	usecase := core.NewGoalRegisterUsecase(&db)
	usecase.Execute("Title", "2025-01-01")

	assert.Equal(t, "Title", db.get(0).title)
	assert.Equal(t, "2025-01-01", db.get(0).endDate)
}

func newArrayRepository() *arrayRepository {
	return &arrayRepository{}
}

type arrayRepository struct {
}

func (a *arrayRepository) get(index int) inMemoryGoal {
	return inMemoryGoal{}
}

type inMemoryGoal struct {
	title   string
	endDate string
}
