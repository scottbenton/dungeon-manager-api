package campaigns

type createCampaignDAO struct {
	Name string
	Owners []string
}

type campaignDAO struct {
	Id string `bson:"_id"`
	Name string `bson:"name"`
	Owners []string `bson:"owners"`
}

func (dao campaignDAO) toDTO() campaignDTO {
	return campaignDTO{
		Id: dao.Id,
		Name: dao.Name,
		Owners: dao.Owners,
	}
}
