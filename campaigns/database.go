package campaigns

import (
	"DungeonManagerAPI/config"
	"DungeonManagerAPI/utils"
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func getCampaignCollection() *mongo.Collection {
	db := config.GetMongoDatabase();
	return db.Collection("campaigns");
}

func getUsersCampaignsDB(userId string) ([]campaignDTO, error) {
	campaignsCollection := getCampaignCollection();

	// Filter by campaigns where the user is an owner
	filter := bson.D{{Key: "owners", Value: userId}};

	// Get campaigns
	result, err := campaignsCollection.Find(context.TODO(), filter, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))

	if(err != nil) {
		log.Println("Failed to get campaigns", err);
		return []campaignDTO{}, errors.New("Failed to get campaigns");
	}

	// Extract campaigns from results
	var campaigns []campaignDTO;
	for result.Next(context.TODO()) {
		var campaign campaignDAO;
		err := result.Decode(&campaign);
		if(err != nil) {
			log.Println("Failed to decode campaign", err);
			return []campaignDTO{}, errors.New("Failed to decode campaign");
		}
		campaigns = append(campaigns, campaign.toDTO());
	}

	log.Println(campaigns);

	return campaigns, nil;
}

func getCampaignDB(userId string, campaignId string) (campaignDTO, error) {
	campaignsCollection := getCampaignCollection();

	id := utils.ConvertStringToObjectId(campaignId);
	// Filter by campaigns where the user is an owner
	filter := bson.D{{Key: "owners", Value: userId}, {Key: "_id", Value: id}};

	// Get campaigns
	result := campaignsCollection.FindOne(context.TODO(), filter, options.FindOne())

	// Extract campaigns from results
	var campaign campaignDAO;
	err := result.Decode(&campaign);
	if(err != nil) {
		log.Println("Failed to decode campaign", err);
		return campaignDTO{}, errors.New("Failed to decode campaign");
	}
	return campaign.toDTO(), nil;
}

type createCampaignInput struct {
	name string
	owners []string
}

func createCampaignDB(campaignInput createCampaignInput) (string, error) {
	campaignsCollection := getCampaignCollection();

	campaign := createCampaignDAO {
		Name: campaignInput.name,
		Owners: campaignInput.owners,
	}

	log.Println("Inserting campaign", campaign);

	result, err := campaignsCollection.InsertOne(context.TODO(), campaign)

	if(result.InsertedID == nil || err != nil) {
		log.Println("Failed to insert campaign", err);
		return "", errors.New("Failed to insert campaign");
	}

	return utils.ConvertObjectIdToString(result.InsertedID), nil;
}

func updateCampaignDB(campaignId string, updateInput updateCampaignDTO) error {
	campaignsCollection := getCampaignCollection();

	id := utils.ConvertStringToObjectId(campaignId);
	filter := bson.D{{Key: "_id", Value: id}};

	update := bson.D{};

	if(updateInput.Name != nil) {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "name", Value: *updateInput.Name}}});
	}

	if(updateInput.Owners != nil) {
		update = append(update, bson.E{Key: "$set", Value: bson.D{{Key: "owners", Value: *updateInput.Owners}}});
	}

	log.Println("Updating campaign", update);

	_, err := campaignsCollection.UpdateOne(context.TODO(), filter, update);

	if(err != nil) {
		log.Println("Failed to update campaign", err);
		return errors.New("Failed to update campaign");
	}

	return nil;
}

func deleteCampaignDB(campaignId string) error {
	campaignsCollection := getCampaignCollection();

	id := utils.ConvertStringToObjectId(campaignId);
	filter := bson.D{{Key: "_id", Value: id}};
	_, err := campaignsCollection.DeleteOne(context.TODO(), filter);

	if(err != nil) {
		log.Println("Failed to remove campaign", err);
		return errors.New("Failed to remove campaign");
	}

	return nil;
}