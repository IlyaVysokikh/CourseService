package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	ierrors "CourseService/pkg/errors"
	"errors"

	"context"
	"log/slog"
	"time"

	"github.com/google/uuid"
)

type CourseServiceImpl struct {
	repo       repositories.CourseRepository
	dateFormat string
}

func NewCourseServiceImpl(repo repositories.CourseRepository) CourseService {
	return &CourseServiceImpl{
		repo:       repo,
		dateFormat: "2006-01-02T15:04:05Z07:00",
	}
}

func (c *CourseServiceImpl) GetAllCourses(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error) {

	// todo На user service фильтр по роли. Студент получает только свои курсы, а препод все
	courses, err := c.repo.GetAllCourses(filter)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error getting all courses", "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get courses", err)
		}

		slog.Error("Error getting all courses", "error", err)
		return nil, err
	}

	var result []dto.CourseList
	for _, course := range courses {
		isArchived := false
		parsedDateEnd, err := time.Parse(c.dateFormat, course.DateEnd)
		if err != nil {
			slog.Error("Error parsing course end date", "error", err)
			continue
		}

		if parsedDateEnd.Before(time.Now()) || parsedDateEnd.Equal(time.Now()) {
			isArchived = true
		}

		result = append(result, dto.CourseList{
			Id:         course.ID,
			Name:       course.Name,
			IsArchived: isArchived,
			ImagePath:  course.ImagePath,
		})
	}

	return result, nil
}

func (c *CourseServiceImpl) GetCourse(ctx context.Context, id uuid.UUID) (*dto.Course, error) {
	course, err := c.repo.GetCourse(id)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Course not found", "courseID", id)
			return nil, ierrors.New(ierrors.ErrNotFound, "course not found", err)
		} else if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error getting course", "error", err)
			return nil, ierrors.New(ierrors.ErrInternal, "failed to get course", err)
		}

		slog.Error("Error getting course", "error", err)
		return nil, err
	}

	parsedDateStart, err := time.Parse(c.dateFormat, course.DateStart)
	if err != nil {
		slog.Error("Error parsing course start date", "error", err)
		return nil, ierrors.ErrInvalidInput
	}

	parsedDateEnd, err := time.Parse(c.dateFormat, course.DateEnd)
	if err != nil {
		slog.Error("Error parsing course end date", "error", err)
		return nil, ierrors.ErrInvalidInput
	}

	courseDto := dto.Course{
		Id:          course.ID,
		Name:        course.Name,
		AuthorID:    course.AuthorID,
		Description: course.Description,
		DateStart:   parsedDateStart.Format(c.dateFormat),
		DateEnd:     parsedDateEnd.Format(c.dateFormat),
		ImagePath:   course.ImagePath,
		IsArchived:  false,
	}

	return &courseDto, nil
}

func (c *CourseServiceImpl) CreateCourse(ctx context.Context, course *dto.CreateCourse) (uuid.UUID, error) {
	createdCourseId, err := c.repo.Create(course)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error creating course", "error", err)
			return uuid.Nil, ierrors.New(ierrors.ErrInternal, "failed to create course", err)
		}

		slog.Error("Error creating course", "error", err)
		return uuid.Nil, err
	}

	return *createdCourseId, nil
}

func (c *CourseServiceImpl) CloneCourse(ctx context.Context, course *dto.CloneCourseRequest) (uuid.UUID, error) {
	createdCourseId, err := c.repo.Clone(course)
	if err != nil {
		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error cloning course", "error", err)
			return uuid.Nil, ierrors.New(ierrors.ErrInternal, "failed to clone course", err)
		}

		slog.Error("Error cloning course", "error", err)
		return uuid.Nil, err
	}

	return *createdCourseId, nil
}

func (c *CourseServiceImpl) DeleteCourse(ctx context.Context, id uuid.UUID) error {
	if err := c.repo.Delete(id); err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			slog.Warn("Course not found", "courseID", id)
			return ierrors.New(ierrors.ErrNotFound, "course not found", err)
		}

		if errors.Is(err, ierrors.ErrInternal) {
			slog.Error("Error deleting course", "error", err)
			return ierrors.New(ierrors.ErrInternal, "failed to delete course", err)
		}
	}

	return nil
}
