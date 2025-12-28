package controllers

import (
	"net/http"
	"saint-seiya-back/internal/application/knight"
	"saint-seiya-back/internal/application/knight/dto"
	knightDomain "saint-seiya-back/internal/domain/knight"
	"saint-seiya-back/internal/responses"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type KnightController struct {
	createKnightUseCase      *knight.CreateKnightUseCase
	getKnightsUseCase        *knight.GetKnightsUseCase
	getKnightByIdUseCase     *knight.GetKnightByIdUseCase
	createKnightSkillUseCase *knight.CreateKnightSkillUseCase
}

func NewKnightController(
	createKnightUseCase *knight.CreateKnightUseCase,
	getKnightsUseCase *knight.GetKnightsUseCase,
	getKnightByIdUseCase *knight.GetKnightByIdUseCase,
	createKnightSkillUseCase *knight.CreateKnightSkillUseCase,
) *KnightController {
	return &KnightController{
		createKnightUseCase:      createKnightUseCase,
		getKnightsUseCase:        getKnightsUseCase,
		getKnightByIdUseCase:     getKnightByIdUseCase,
		createKnightSkillUseCase: createKnightSkillUseCase,
	}
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

func (kc *KnightController) GetKnights(c *gin.Context) {
	page := 1
	limit := 20
	rank := c.Query("rank")
	name := c.Query("name")

	if p := c.Query("page"); p != "" {
		if parsed, err := strconv.Atoi(p); err == nil && parsed > 0 {
			page = parsed
		}
	}

	if l := c.Query("limit"); l != "" {
		if parsed, err := strconv.Atoi(l); err == nil && parsed > 0 {
			limit = parsed
		}
	}

	input := knight.GetKnightsInput{
		Page:  page,
		Limit: limit,
		Rank:  rank,
		Name:  name,
	}

	result, err := kc.getKnightsUseCase.Execute(input)
	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Error internal to get", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "Knights retrieved successfully", result)
}

func (kc *KnightController) GetKnightByID(c *gin.Context) {
	idString := c.Param("id")
	id, err := strconv.ParseUint(idString, 10, 64)

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "ID invalid", "Id must be a number")
		return
	}

	result, err := kc.getKnightByIdUseCase.Execute(uint(id))
	if err != nil {
		responses.Error(c, http.StatusNotFound, "knight not found", err.Error())
		return
	}

	responses.Success(c, http.StatusOK, "knight found", result)
}

func (kc *KnightController) CreateKnightSkill(c *gin.Context) {
	knightIDString := c.Param("id")
	knightID, err := strconv.ParseUint(knightIDString, 10, 64)

	if err != nil {
		responses.Error(c, http.StatusBadRequest, "Id invalid", "Id must be a number")
		return
	}

	var req dto.CreateKnightSkillRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		if _, ok := err.(validator.ValidationErrors); ok {
			responses.ValidationError(c, http.StatusBadRequest, err)
			return
		}
		responses.Error(c, http.StatusBadRequest, "Invalid JSON sent", err.Error())
		return
	}

	result, err := kc.createKnightSkillUseCase.Execute(knight.CreateKnightSkillInput{
		KnightID:    uint(knightID),
		Name:        req.Name,
		Type:        req.Type,
		ImageURL:    req.ImageURL,
		Description: req.Description,
		Levels:      req.Levels,
	})

	if err != nil {
		responses.Error(c, http.StatusInternalServerError, "Failed to create knight skill", err.Error())
		return
	}

	responses.Success(c, http.StatusCreated, "Knight skill created successfully", result)
}
