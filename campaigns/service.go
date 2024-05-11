package campaigns

import (
	"DungeonManagerAPI/utils"
	"slices"
)

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

func getCampaign(userId string, campaignId string) (campaignDTO, error) {
	return getCampaignDB(userId, campaignId);
}

func updateCampaign(userId string, campaignId string, update updateCampaignDTO) error {
	err := makeSureUserIsOwner(userId, campaignId);
	if(err != nil) {
		return err;
	}

	return updateCampaignDB(campaignId, update);
} 

func deleteCampaign(userId string, campaignId string) error {
	err := makeSureUserIsOwner(userId, campaignId);
	if(err != nil) {
		return err;
	}

	return deleteCampaignDB(campaignId);
}

func makeSureUserIsOwner(userId string, campaignId string) error {
	campaign, err := getCampaign(userId, campaignId);
	if(err != nil) {
		return err;
	}

	if(!slices.Contains(campaign.Owners, userId)) {
		return utils.ErrNotAuthorized;
	}

	return nil;
}