package task

type Service interface {
	CreateTask(task *Task) error
	GetAllTasks() ([]Task, error)
	GetTaskByID(id int) (Task, error)
	UpdateTask(task *Task) error
	DeleteTask(id int) error
}

type taskService struct {
	repo Repository
}

func NewTaskService(repo Repository) Service {
	return &taskService{repo}
}

func (s *taskService) CreateTask(task *Task) error {
	return s.repo.Create(task)
}

func (s *taskService) GetAllTasks() ([]Task, error) {
	return s.repo.FindAll()
}

func (s *taskService) GetTaskByID(id int) (Task, error) {
	return s.repo.FindOne(id)
}

func (s *taskService) UpdateTask(task *Task) error {
	return s.repo.Update(task)
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.Delete(id)
}
