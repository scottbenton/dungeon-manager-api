package campaigns

import (
	"DungeonManagerAPI/utils"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

func GetCampaignRouter() *chi.Mux {
	campaignRouter := chi.NewRouter()
	campaignRouter.Use(utils.VerifySession)
	campaignRouter.Use(middleware.RequestID)
	campaignRouter.Use(middleware.Logger)
	campaignRouter.Use(middleware.Recoverer)
	campaignRouter.Use(middleware.URLFormat)
	campaignRouter.Use(render.SetContentType(render.ContentTypeJSON))

	campaignRouter.Post("/", createCampaignAPI)
	campaignRouter.Get("/", getUsersCampaignsAPI)
	campaignRouter.Get("/{campaignId}", getCampaignAPI)

	return campaignRouter
}

func createCampaignAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
	log.Println("User ID: ", uid)

	var createCampaignInput = new(createCampaignDTO);
	if err := render.DecodeJSON(r.Body, createCampaignInput); err != nil {
		render.Render(w, r, ErrInvalidRequest(err))
		log.Println("Failed to decode create campaign request", err)
		return
	}

	log.Println("Creating campaign: ", createCampaignInput.Name)


	campaign, err := createCampaign(uid, createCampaignInput.Name)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		log.Println("Failed to create campaign", err)
		return
	}

	render.Status(r, http.StatusCreated)
	render.JSON(w, r, campaign)
}

func getUsersCampaignsAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
	log.Println("User ID: ", uid)

	campaigns, err := getUsersCampaigns(uid)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		log.Println("Failed to get campaigns", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaigns)
}

func getCampaignAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
		campaignId := chi.URLParam(r, "campaignId")

	log.Println("User ID: ", uid)
	log.Println("Campaign ID: ", campaignId)

	campaign, err := getCampaign(uid, campaignId)
	if err != nil {
		render.Render(w, r, ErrInternalServer(err))
		log.Println("Failed to get campaign", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaign)
}

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrInvalidRequest(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		StatusText:     "Invalid request.",
		ErrorText:      err.Error(),
	}
}

func ErrInternalServer(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		StatusText:     "Server Error",
		ErrorText:      err.Error(),
	}
}
