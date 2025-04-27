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
