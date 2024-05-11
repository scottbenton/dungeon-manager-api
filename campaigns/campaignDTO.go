package campaigns

type campaignDTO struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Owners []string `json:"owners"`
}

type createCampaignDTO struct {
	Name string `json:"name"`
}

// All fields are optional since we merge it with the existing object
type updateCampaignDTO struct {
	Name *string `json:"name"`
	Owners *[]string `json:"owners"`
}