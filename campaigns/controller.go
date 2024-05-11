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
	campaignRouter.Put("/{campaignId}", updateCampaignAPI)
	campaignRouter.Delete("/{campaignId}", deleteCampaignAPI)

	return campaignRouter
}

func createCampaignAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
	log.Println("User ID: ", uid)

	var createCampaignInput = new(createCampaignDTO)
	if err := render.DecodeJSON(r.Body, createCampaignInput); err != nil {
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
		log.Println("Failed to decode create campaign request", err)
		return
	}

	log.Println("Creating campaign: ", createCampaignInput.Name)

	campaign, err := createCampaign(uid, createCampaignInput.Name)
	if err != nil {
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
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
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaigns"))
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
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
		log.Println("Failed to get campaign", err)
		return
	}

	render.Status(r, http.StatusOK)
	render.JSON(w, r, campaign)
}

func updateCampaignAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
	campaignId := chi.URLParam(r, "campaignId")

	log.Println("User ID: ", uid)
	log.Println("Campaign ID: ", campaignId)

	var updateCampaignInput = new(updateCampaignDTO)
	if err := render.DecodeJSON(r.Body, updateCampaignInput); err != nil {
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
		log.Println("Failed to decode update campaign request", err)
		return
	}

	log.Println("Updating campaign: ", updateCampaignInput)

	err := updateCampaign(uid, campaignId, *updateCampaignInput)
	if err != nil {
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
		log.Println("Failed to update campaign", err)
		return
	}

	render.Status(r, http.StatusOK)
}

func deleteCampaignAPI(w http.ResponseWriter, r *http.Request) {
	// Get the uid from the request context
	uid := utils.GetUidFromContext(r.Context())
	campaignId := chi.URLParam(r, "campaignId")

	log.Println("User ID: ", uid)
	log.Println("Campaign ID: ", campaignId)

	err := deleteCampaign(uid, campaignId)
	if err != nil {
		render.Render(w, r, utils.GetHttpErrorFromError(err, "campaign"))
		log.Println("Failed to delete campaign", err)
		return
	}

	render.Status(r, http.StatusOK)
}
