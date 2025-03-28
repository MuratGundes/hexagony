package controller

import (
	"encoding/json"
	"hexagony/internal/app/domain"
	"hexagony/internal/app/pkg/rest"
	"hexagony/internal/app/pkg/validation"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type AlbumHandler struct {
	albumUseCase domain.AlbumUseCase
}

func NewAlbumHandler(c *chi.Mux, as domain.AlbumUseCase) {
	handler := AlbumHandler{albumUseCase: as}

	c.Route("/album", func(r chi.Router) {
		r.Get("/", handler.FindAll)
		r.Get("/{uuid}", handler.FindByID)
		r.Post("/", handler.Add)
		r.Put("/{uuid}", handler.Update)
		r.Delete("/{uuid}", handler.Delete)
	})
}

// FindAll godoc
// @Summary      List of albums
// @Description  lists all albums
// @Tags         album
// @Accept       json
// @Produce      json
// @Success      200  {object}  []domain.Album
// @Failure      422   {object}  rest.APIMessage
// @Router       /album [get]
func (a *AlbumHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	albums, err := a.albumUseCase.FindAll(r.Context())
	if err != nil {
		rest.DecodeError(w, r, domain.ErrFindAll, http.StatusUnprocessableEntity)
		return
	}

	rest.EncodeJSON(w, http.StatusOK, &albums)
}

// FindByID godoc
// @Summary      List an album
// @Description  lists an album by uuid
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        uuid  path      string  true  "album uuid"
// @Success      200   {object}  domain.Album
// @Failure      422      {object}  rest.APIMessage
// @Router       /album/{uuid} [get]
func (a *AlbumHandler) FindByID(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(chi.URLParam(r, "uuid"))
	if err != nil {
		rest.DecodeError(w, r, domain.ErrUUIDParse, http.StatusUnprocessableEntity)
		return
	}

	album, err := a.albumUseCase.FindByID(r.Context(), uuid)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrFindByID, http.StatusUnprocessableEntity)
		return
	}

	rest.EncodeJSON(w, http.StatusOK, album)
}

// Add godoc
// @Summary      Add an album
// @Description  add a new album
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        payload  body      domain.Album  true  "add a new album"
// @Success      201      {object}  rest.APIMessage
// @Failure      422      {object}  rest.APIMessage
// @Router       /album [post]
func (a *AlbumHandler) Add(w http.ResponseWriter, r *http.Request) {
	var album domain.Album

	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrAdd, http.StatusUnprocessableEntity)
		return
	}

	if err := validator.New().StructCtx(r.Context(), album); err != nil {
		validation.Message(w, err)
		return
	}

	album.UUID = uuid.New()
	album.CreatedAt = time.Now()
	album.UpdatedAt = time.Now()

	err = a.albumUseCase.Add(r.Context(), &album)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrAdd, http.StatusUnprocessableEntity)
		return
	}

	rest.EncodeJSON(w, http.StatusCreated, &rest.APIMessage{Message: "Created"})
}

// Update godoc
// @Summary      Update an album
// @Description  update an album by uuid
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        uuid     path      string        true  "album uuid"
// @Param        payload  body      domain.Album  true  "update an album by uuid"
// @Success      200      {object}  rest.APIMessage
// @Failure      422   {object}  rest.APIMessage
// @Router       /album/{uuid} [put]
func (a *AlbumHandler) Update(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(chi.URLParam(r, "uuid"))
	if err != nil {
		rest.DecodeError(w, r, domain.ErrUUIDParse, http.StatusUnprocessableEntity)
		return
	}

	var album domain.Album

	err = json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrUpdate, http.StatusUnprocessableEntity)
		return
	}

	if err := validator.New().StructCtx(r.Context(), album); err != nil {
		validation.Message(w, err)
		return
	}

	album.UpdatedAt = time.Now()

	err = a.albumUseCase.Update(r.Context(), uuid, &album)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrUpdate, http.StatusUnprocessableEntity)
		return
	}

	rest.EncodeJSON(w, http.StatusOK, &rest.APIMessage{Message: "Updated"})
}

// Update godoc
// @Summary      Delete an album
// @Description  delete an album by uuid
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        uuid  path      string  true  "album uuid"
// @Success      200   {object}  rest.APIMessage
// @Failure      422  {object}  rest.APIMessage
// @Router       /album/{uuid} [delete]
func (a *AlbumHandler) Delete(w http.ResponseWriter, r *http.Request) {
	uuid, err := uuid.Parse(chi.URLParam(r, "uuid"))
	if err != nil {
		rest.DecodeError(w, r, domain.ErrDelete, http.StatusUnprocessableEntity)
		return
	}

	err = a.albumUseCase.Delete(r.Context(), uuid)
	if err != nil {
		rest.DecodeError(w, r, domain.ErrDelete, http.StatusUnprocessableEntity)
		return
	}

	rest.EncodeJSON(w, http.StatusOK, &rest.APIMessage{Message: "Deleted"})
}
