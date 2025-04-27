package usecase

import (
	"CourseService/internal/services"
	"CourseService/internal/usecase/course"
	"CourseService/internal/usecase/module"
	"CourseService/internal/usecase/shared"
	"CourseService/internal/usecase/task"
	"CourseService/internal/usecase/test_data"
)

type UseCase struct {
	GetAllCourseUseCase    shared.GetAllCourseUseCase
	GetCourseUseCase       shared.GetCourseUseCase
	CreateCourseUseCase    shared.CreateCourseUseCase
	CloneCourseUseCase     shared.CloneCourseUseCase
	CreateModulesUseCase   shared.CreateModulesUseCase
	GetModuleUseCase       shared.GetModuleUseCase
	GetTaskUseCase         shared.GetTaskUseCase
	DeleteCourseUseCase    shared.DeleteCourseUseCase
	UpdateCourseUseCase    shared.UpdateCourseUseCase
	DeleteModuleUseCase    shared.DeleteModuleUseCase
	DeleteTaskUseCase      shared.DeleteTaskUseCase
	CreateModuleAttachment shared.CreateModuleAttachmentUseCase

	CreateTestDataUseCase shared.CreateTestDataUseCase
	DeleteTestDataUseCase shared.DeleteTestDataUseCase
	GetAllTestDataUseCase shared.GetAllTestDataUseCase
	GetTestDataUseCasa    shared.GetTestDataUseCase
	UpdateTestDataUseCase shared.UpdateTestDataUseCase
}

func NewUseCase(services *services.Service) *UseCase {
	return &UseCase{
		GetAllCourseUseCase: course.NewGetAllCourseUseCase(services.CourseService),
		GetCourseUseCase:    course.NewGetCourseUseCase(services.CourseService, services.ModuleService, services.TaskService),
		CreateCourseUseCase: course.NewCreateCourseUseCase(services.CourseService),
		CloneCourseUseCase:  course.NewCloneCourseUseCase(services.CourseService),
		UpdateCourseUseCase: course.NewUpdateCourseUseCase(services.CourseService),
		DeleteCourseUseCase: course.NewDeleteCourseUseCase(services.CourseService),

		CreateModulesUseCase:   module.NewCreateModuleUseCase(services.ModuleService),
		GetModuleUseCase:       module.NewGetModuleUseCase(services.ModuleService, services.TaskService, services.ModuleAttachmentService),
		DeleteModuleUseCase:    module.NewDeleteModuleUseCase(services.ModuleService),
		CreateModuleAttachment: module.NewCreateModuleAttachmentUseCase(services.ModuleAttachmentService),

		GetTaskUseCase:    task.NewGetTaskUseCase(services.TaskService),
		DeleteTaskUseCase: task.NewDeleteTaskUseCase(services.TaskService),

		CreateTestDataUseCase: test_data.NewCreateTestDataUseCase(services.TestDataService),
		DeleteTestDataUseCase: test_data.NewDeleteTestDataUseCase(services.TestDataService),
		GetAllTestDataUseCase: test_data.NewGetAllTestDataUseCase(services.TestDataService),
		GetTestDataUseCasa:    test_data.NewGetTestDataUseCase(services.TestDataService),
		UpdateTestDataUseCase: test_data.NewUpdateTestDataUseCase(services.TestDataService),
	}
}
