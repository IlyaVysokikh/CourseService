package services

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/repositories"
	"context"
	"log/slog"
	"time"
)

type CourseServiceImpl struct {
	repo repositories.CourseRepository
	dateFormat string
}

func NewCourseServiceImpl(repo repositories.CourseRepository) CourseService {
	return &CourseServiceImpl{
		repo: repo,
		dateFormat: "2006-01-02T15:04:05Z07:00",
	}
}

func (c *CourseServiceImpl) GetAllCourses(ctx context.Context, filter *dto.CourseFilter) ([]dto.CourseList, error) {

	// todo На user service фильтр по роли. Студент получает только свои курсы, а препод все
	courses, err := c.repo.GetAllCourses(filter)
	if err != nil {
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
			Id: course.ID,
			Name: course.Name,
			IsArchived: isArchived,
			ImagePath: course.ImagePath,
		})
	}
	
	return result, nil
}
