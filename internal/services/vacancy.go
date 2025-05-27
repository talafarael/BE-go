package services

type (
	vacancyService struct{}
	VacancyService interface{}
)

func NewVacancyService() VacancyService {
	return &vacancyService{}
}

func CreateVacancy() {
}
