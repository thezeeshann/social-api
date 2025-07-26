package main

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/thezeeshann/social/internal/store"
)

type createPostPayload struct {
	Title   string   `json:"title" validate:"required,max=100"`
	Content string   `json:"content" validate:"required,max=1000"`
	Tags    []string `json:"tags"`
}

func (app *application) createPostsHandler(w http.ResponseWriter, r *http.Request) {
	var payload createPostPayload
	if err := readJson(w, r, &payload); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := Validate.Struct(payload); err != nil {
		writeJsonError(w, http.StatusBadRequest, err.Error())
		return
	}

	post := &store.Post{
		Title:   payload.Title,
		Content: payload.Content,
		Tags:    payload.Tags,
		// TODO: change the userid
		UserId: 1,
	}

	ctx := r.Context()

	if err := app.store.Posts.Create(ctx, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}

	if err := writeJson(w, http.StatusCreated, post); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
	}
}

func (app *application) getPostHandler(w http.ResponseWriter, r *http.Request) {
	idPrams := chi.URLParam(r, "postID")
	id, err := strconv.ParseInt(idPrams, 10, 64)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
		return
	}
	ctx := r.Context()
	posts, err := app.store.Posts.GetById(ctx, id)
	if err != nil {
		writeJsonError(w, http.StatusNotFound, err.Error())
		return
	}
	if err := writeJson(w, http.StatusCreated, posts); err != nil {
		writeJsonError(w, http.StatusInternalServerError, err.Error())
	}
}
