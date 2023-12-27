package usecase

type MasterClientUseCaseInterface interface {
	CreateMasterClient() error
}

type masterClientUseCase struct {
}

func MasterClientUseCase() *masterClientUseCase {
	return &masterClientUseCase{}
}

func (useCase *masterClientUseCase) CreateMasterClient() error {
	return nil
}
