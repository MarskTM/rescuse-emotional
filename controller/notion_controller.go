package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	
	"rescues/model"
	"rescues/service"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type notionController struct {
	notionService service.NotionService
}

type NotionController interface {
	GetByProfileId(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// --------------------Notion module------------------------

//Get notion's godoc
// @tags notion-manager-apis
// @Summary Get By ProfileId notion
// @Description input: ProfileID => output: notion's list
// @Accept json
// @Produce json
// @param profile_id path integer true "profile_id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /notion/profile/{profile_id} [get]
func (c *notionController) GetByProfileId(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	profileIdStr := chi.URLParam(r, "profile_id")
	profileId, err := strconv.Atoi(profileIdStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)
		
		res = &model.Response{
			Data:    nil,
			Message: "get notion failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.notionService.GetByProfileId(profileId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)

		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "get notion successfully",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

//Create notion godoc
// @tags notion-manager-apis
// @Summary Create notion
// @Description input: notion's data => output: notion's data created
// @Accept json
// @Produce json
// @Param notion body model.Notion true "notion"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /notion/create [post]
func (c *notionController) Create(w http.ResponseWriter, r *http.Request) {
	var res model.Response
	var notion model.Notion

	err := json.NewDecoder(r.Body).Decode(&notion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)

		res = model.Response{
			Data:    nil,
			Message: "create notion failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.notionService.Create(&notion)
	if err != nil {
		res = model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = model.Response{
			Data:    tmp,
			Message: "create notion successfully",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

//Update notion godoc
// @tags notion-manager-apis
// @Summary Update notion
// @Description input: notion's data => output: notion's data updated
// @Accept json
// @Produce json
// @Param notion body model.Notion true "notion"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /notion/update [put]
func (c *notionController) Update(w http.ResponseWriter, r *http.Request) {
	var res *model.Response
	var notion model.Notion

	err := json.NewDecoder(r.Body).Decode(&notion)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)

		res = &model.Response{
			Data:    nil,
			Message: "update notion failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	tmp, err := c.notionService.Update(&notion)
	if err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    tmp,
			Message: "update notion successfully",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

//Delete notion godoc
// @tags notion-manager-apis
// @Summary Delete notion
// @Description input: notion's id => output: notion's data deleted
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Security ApiKeyAuth
// @Success 200 {object} model.Response
// @Router /notion/delete/{id} [delete]
func (c *notionController) Delete(w http.ResponseWriter, r *http.Request) {
	var res *model.Response

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		http.Error(w, http.StatusText(400), 400)

		res = &model.Response{
			Data:    nil,
			Message: "delete notion failed: " + err.Error(),
			Success: false,
		}
		render.JSON(w, r, res)
		return
	}

	
	if err := c.notionService.Delete(id);err != nil {
		res = &model.Response{
			Data:    nil,
			Message: err.Error(),
			Success: false,
		}
	} else {
		res = &model.Response{
			Data:    nil,
			Message: "delete notion successfully",
			Success: true,
		}
	}
	render.JSON(w, r, res)
}

func NewNotionController() NotionController {
	notionService := service.NewNotionService()
	return &notionController{
		notionService: notionService,
	}
}