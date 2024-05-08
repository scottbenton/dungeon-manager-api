package campaigns

type campaignDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Owners []string `json:"owners"`
}

type createCampaignDTO struct {
	Name string `json:"name"`
}