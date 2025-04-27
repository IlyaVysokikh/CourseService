package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
)

type TestDataServiceImpl struct {
	repo repositories.TestDataRepository
}

func NewTestDataService(repo repositories.TestDataRepository) TestDataService {
	return &TestDataServiceImpl{repo: repo}
}

func (s *TestDataServiceImpl) GetAll(ctx context.Context, taskId uuid.UUID) ([]dto.TestDataResponse, error) {
	testData, err := s.repo.GetAll(taskId)
	if err != nil {
		return nil, err
	}

	var res []dto.TestDataResponse
	for _, t := range testData {
		res = append(res, dto.TestDataResponse{
			Id:     t.Id,
			TaskId: t.Id,
			Input:  t.Input,
			Output: t.Output,
		})
	}

	return res, nil
}

func (s *TestDataServiceImpl) Get(ctx context.Context, id uuid.UUID) (dto.TestDataResponse, error) {
	testData, err := s.repo.Get(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return dto.TestDataResponse{}, errors.New("not found")
		}

		return dto.TestDataResponse{}, err
	}

	return dto.TestDataResponse{
		Id:     testData.Id,
		TaskId: testData.TaskId,
		Input:  testData.Input,
		Output: testData.Output,
	}, nil
}

func (s *TestDataServiceImpl) Create(ctx context.Context, request dto.CreateTestDataRequest) (uuid.UUID, error) {
	return s.repo.Create(ctx, request)
}

func (s *TestDataServiceImpl) Update(ctx context.Context, id uuid.UUID, request dto.UpdateTestDataRequest) error {
	return s.repo.Update(ctx, id, request)
}

func (s *TestDataServiceImpl) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(id)
}
