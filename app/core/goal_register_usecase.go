package core

func NewGoalRegisterUsecase(repository GoalRepository) *GoalRegisterUsecase {
	return &GoalRegisterUsecase{}
}

type GoalRegisterUsecase struct{}

func (u GoalRegisterUsecase) Execute(title string, endDate string) {

}

type GoalRepository interface {
}
