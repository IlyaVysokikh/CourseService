package rest

import (
	"CourseService/internal/interfaces/rest/dto"
	"CourseService/internal/usecase"
	"CourseService/internal/usecase/shared"
	ierrors "CourseService/pkg/errors"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TestDataHandler struct {
	BaseHandler
	CreateTestDataUseCase shared.CreateTestDataUseCase
	GetTestDataUseCase    shared.GetTestDataUseCase
	GetAllTestDataUseCase shared.GetAllTestDataUseCase
	UpdateTestDataUseCase shared.UpdateTestDataUseCase
	DeleteTestDataUseCase shared.DeleteTestDataUseCase
}

func NewTestDataHandler(useCase *usecase.UseCase) *TestDataHandler {
	return &TestDataHandler{
		BaseHandler:           BaseHandler{},
		CreateTestDataUseCase: useCase.CreateTestDataUseCase,
		GetAllTestDataUseCase: useCase.GetAllTestDataUseCase,
		GetTestDataUseCase:    useCase.GetTestDataUseCasa,
		UpdateTestDataUseCase: useCase.UpdateTestDataUseCase,
		DeleteTestDataUseCase: useCase.DeleteTestDataUseCase,
	}
}

// GetTestData godoc
// @Summary Получить тестовые данные
// @Description Получение одного тестового набора данных по ID
// @Tags TestData
// @Accept json
// @Produce json
// @Param id path string true "UUID тестовых данных"
// @Success 200 {object} dto.TestDataResponse "Успешный ответ"
// @Failure 400 {object} dto.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} dto.ErrorResponse "Данные не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка"
// @Router /test-data/{id} [get]
func (h *TestDataHandler) GetTestData(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.badRequest(ctx, err)
		return
	}

	res, err := h.GetTestDataUseCase.Handle(ctx, id)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.ok(ctx, res)
}

// GetAllTestData godoc
// @Summary Получить все тестовые данные задачи
// @Description Получение всех тестовых данных по ID задачи
// @Tags TestData
// @Accept json
// @Produce json
// @Param id path string true "UUID задачи"
// @Success 200 {array} dto.TestDataResponse "Список тестовых данных"
// @Failure 400 {object} dto.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} dto.ErrorResponse "Данные не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка"
// @Router /test-data/task/{id} [get]
func (h *TestDataHandler) GetAllTestData(ctx *gin.Context) {
	taskId, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.badRequest(ctx, err)
		return
	}

	res, err := h.GetAllTestDataUseCase.Handle(ctx, taskId)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.ok(ctx, res)
	return
}

// CreateTestData godoc
// @Summary Создать тестовые данные
// @Description Создание нового тестового набора данных
// @Tags TestData
// @Accept json
// @Produce json
// @Param request body dto.CreateTestDataRequest true "Данные для создания"
// @Success 201 {object} dto.TestDataResponse "Созданные тестовые данные"
// @Failure 400 {object} dto.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} dto.ErrorResponse "Не найдена задача для привязки данных"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка"
// @Router /test-data [post]
func (h *TestDataHandler) CreateTestData(ctx *gin.Context) {
	var request dto.CreateTestDataRequest
	if err := ctx.ShouldBind(&request); err != nil {
		h.badRequest(ctx, err)
		return
	}

	res, err := h.CreateTestDataUseCase.Handle(ctx, request)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.created(ctx, res)
}

// UpdateTestData godoc
// @Summary Обновить тестовые данные
// @Description Обновление существующего тестового набора данных по ID
// @Tags TestData
// @Accept json
// @Produce json
// @Param id path string true "UUID тестовых данных"
// @Param request body dto.UpdateTestDataRequest true "Данные для обновления"
// @Success 200 {object} dto.ErrorResponse "Успешное обновление"
// @Failure 400 {object} dto.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} dto.ErrorResponse "Данные не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка"
// @Router /test-data/{id} [put]
func (h *TestDataHandler) UpdateTestData(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.badRequest(ctx, err)
		return
	}

	var request dto.UpdateTestDataRequest
	if err := ctx.ShouldBind(&request); err != nil {
		h.badRequest(ctx, err)
		return
	}

	err = h.UpdateTestDataUseCase.Handle(ctx, id, request)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}
		h.internalServerError(ctx, err)
	}

	h.ok(ctx, nil)
	return
}

// DeleteTestData godoc
// @Summary Удалить тестовые данные
// @Description Удаление тестового набора данных по ID
// @Tags TestData
// @Accept json
// @Produce json
// @Param id path string true "UUID тестовых данных"
// @Success 200 {object} dto.ErrorResponse "Успешное удаление"
// @Failure 400 {object} dto.ErrorResponse "Некорректный запрос"
// @Failure 404 {object} dto.ErrorResponse "Данные не найдены"
// @Failure 500 {object} dto.ErrorResponse "Внутренняя ошибка"
// @Router /test-data/{id} [delete]
func (h *TestDataHandler) DeleteTestData(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		h.badRequest(ctx, err)
		return
	}

	err = h.DeleteTestDataUseCase.Handle(ctx, id)
	if err != nil {
		if errors.Is(err, ierrors.ErrNotFound) {
			h.notFound(ctx, err)
			return
		}

		h.internalServerError(ctx, err)
		return
	}

	h.ok(ctx, nil)
	return
}
