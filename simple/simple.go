package simple

import "errors"

type SimpleRepository struct {
	Error bool
}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(simpleRepository *SimpleRepository) (*SimpleService, error) {
	if simpleRepository.Error {
		return nil, errors.New("error")
	} else {
		return &SimpleService{
			SimpleRepository: simpleRepository,
		}, nil
	}
}