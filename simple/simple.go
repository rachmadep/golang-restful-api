package simple

type SimpleRepository struct {

}

type SimpleService struct {
	*SimpleRepository
}

func NewSimpleRepository() *SimpleRepository {
	return &SimpleRepository{}
}

func NewSimpleService(simpleRepository *SimpleRepository) *SimpleService {
	return &SimpleService{
		SimpleRepository: simpleRepository,
	}
}