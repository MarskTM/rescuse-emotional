package controller

import (
	"rescues/model"
	"rescues/service"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type AnswerController interface {
	GetById(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FilterByQuestionId(w http.ResponseWriter, r *http.Request)
}

type answerController struct {
	answerService service.AnswerService
}

// Get answer by id godoc
// @tags answer-manager-apis
// @Summary Get answer by id
// @Description input: answer's id => output: struct answer
// @Accept json
// @Produce json
// @Param id path integer true "answer's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/{id} [get]
func (c *answerController) GetById(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "get answer failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.answerService.GetById(id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get answer successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Get all answers godoc
// @tags answer-manager-apis
// @Summary Get answers
// @Description output: struct answers
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/all [get]
func (c *answerController) GetAll(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	tmp, err := c.answerService.GetAll()
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get answers successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Create answer godoc
// @tags answer-manager-apis
// @Summary Create answers
// @Description input: answer model.Answer => output: status
// @Accept json
// @Produce json
// @param answer body model.Answer true "fill answer"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/create [post]
func (c *answerController) Create(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var answer model.Answer

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "create answer failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}
	tmp, err := c.answerService.Create(&answer)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "create answers successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Update answer godoc
// @tags answer-manager-apis
// @Summary Update answer's
// @Description input: answer model.Answer => output: status
// @Accept json
// @Produce json
// @param answer body model.Answer true "change answer"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/update [put]
func (c *answerController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var answer *model.Answer

	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "update answer failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.answerService.Update(*answer)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update answer successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Delete answer by id godoc
// @tags answer-manager-apis
// @Summary Delete answer by id
// @Description input: answer's id => output: status
// @Accept json
// @Produce json
// @Param id path integer true "answer's id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/delete/{id} [delete]
func (c *answerController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		res = &model.Response{
			Data:    nil,
			Message: "delete answer failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	if err := c.answerService.Delete(id); err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete answer successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

// Filter amswer by questionId godoc
// @tags amswer-manager-apis
// @Summary filter amswer by question_id
// @Description input: amswer's question_id => output:  amswers
// @Accept json
// @Produce json
// @Param question_id path integer true "question_id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /answer/filter/{question_id} [get]
func (c *answerController) FilterByQuestionId(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "question_id")
	question_id, err := strconv.Atoi(idStr)

	answers, err := c.answerService.FilterByQuestionId(question_id)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    answers,
			Message: "filter question successfully.",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewAnswerController() AnswerController {
	answerService := service.NewAnswerService()
	return &answerController{
		answerService: answerService,
	}
}