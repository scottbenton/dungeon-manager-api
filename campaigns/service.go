package campaigns

func createCampaign(userId string, campaignName string) (campaignDTO, error) {
	campaign := createCampaignInput{
		name: campaignName,
		owners: []string{userId},
	}
	
	campaignId, err := createCampaignDB(campaign);

	if(err != nil) {
		return campaignDTO{}, err;
	}

	return campaignDTO {
		Id: campaignId,
		Name: campaignName,
		Owners: []string{userId},
	}, nil
}

func getUsersCampaigns(userId string) ([]campaignDTO, error) {
	return getUsersCampaignsDB(userId);
}