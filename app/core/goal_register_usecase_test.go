package core_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/takutakuaoao/vir-api/core"
	"testing"
)

func Test_目標を登録する(t *testing.T) {
	t.Skip("モックリポジトリのアサート関数が正しく動くテストが動いてから実装")
	db := newMockGoalRepository()
	usecase := core.NewGoalRegisterUsecase(&db)
	usecase.Execute("Title", "2025-01-01")

	assert.True(t, db.calledRegister(1, "Title", "2025-01-01"), true)
}

func Test_モックリポジトリのアサートが正しく動くかのテスト(t *testing.T) {
	mock := newMockGoalRepository()
	mock.Register("Title", "2025-01-01")

	assert.True(t, mock.calledRegister(1, "Title", "2025-01-01"))
}

func Test_モックリポジトリのアサートが失敗した時のテスト(t *testing.T) {
	mock := newMockGoalRepository()
	mock.Register("Title", "2025-01-01")

	assert.False(t, mock.calledRegister(2, "Title", "2025-01-01"))
}

func Test_モックリポジトリのアサートが回数があっているが引数が異なっている場合に失敗するテスト(t *testing.T) {
	mock := newMockGoalRepository()
	mock.Register("Title", "2025-01-01")

	assert.False(t, mock.calledRegister(1, "not match title", "2025-01-01"))
}

func Test_モックリポジトリの2回以上コールしたときに最後の引数が異なっている場合に失敗するテスト(t *testing.T) {
	mock := newMockGoalRepository()
	mock.Register("Title", "2025-01-01")
	mock.Register("Title2", "2025-01-02")

	assert.False(t, mock.calledRegister(2, "Title2", "2025-01-03"))
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
	if len(a.calledRegisterArgs) != count {
		return false
	}

	for i := 0; i < count; i++ {
		currentArg := a.calledRegisterArgs[i]

		if currentArg.title != title || currentArg.endDate != endDate {
			return false
		}
	}

	return true
}

func (a *mockGoalRepository) Register(title, endDate string) {
	a.calledRegisterArgs = append(a.calledRegisterArgs, struct {
		title   string
		endDate string
	}{
		title:   title,
		endDate: endDate,
	})
}
