package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/knight"
	"saint-seiya-back/internal/application/knight/dto"
	knightDomain "saint-seiya-back/internal/domain/knight"
	"saint-seiya-back/internal/responses"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type KnightController struct {
	createKnightUseCase *knight.CreateKnightUseCase
}

func NewKnightController(createKnightUseCase *knight.CreateKnightUseCase) *KnightController {
	return &KnightController{createKnightUseCase: createKnightUseCase}
}

func (kc *KnightController) CreateKnight(c *gin.Context) {
	var req dto.CreateKnightRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	knightDomain := &knightDomain.KnightDomain{
		Name:            req.Name,
		Rank:            req.Rank,
		Pv:              req.Pv,
		AtkC:            req.AtkC,
		DefC:            req.DefC,
		DefF:            req.DefF,
		AtkF:            req.AtkF,
		Speed:           req.Speed,
		StatusHit:       req.StatusHit,
		StatusResist:    req.StatusResist,
		CritDamageC:     req.CritDamageC,
		ResistDamageC:   req.ResistDamageC,
		PerfurationDefC: req.PerfurationDefC,
		ReflectDamage:   req.ReflectDamage,
		Heal:            req.Heal,
		CritLevelF:      req.CritLevelF,
		CritEffectF:     req.CritEffectF,
		CritResistF:     req.ResistCritF,
		ResistDamageF:   req.ResistDamageF,
		PerfurationDefF: req.PerfurationDefF,
		LifeTheft:       req.LifeTheft,
		CritBasicF:      req.CritBasicF,
		ImageURL:        req.ImageURL,
	}

	result, err := kc.createKnightUseCase.Execute(knightDomain)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to create knight", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Knight created successfully", result)
}
