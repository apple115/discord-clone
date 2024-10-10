package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ChannelData struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name        string             `bson:"name" json:"name"`
	Description string             `bson:"description" json:"description"`
	UserID      string             `bson:"user_id" json:"user_id"`
	CreatedAt   time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time          `bson:"updated_at" json:"updated_at"`
}

func GetChannel() ([]ChannelData, error) {
	var channelList []ChannelData
	collection := MongoDB.Collection("channelList")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var channel ChannelData
		err := cursor.Decode(&channel)
		if err != nil {
			return nil, err
		}
		channelList = append(channelList, channel)
	}
	return channelList, nil
}

func GetChannelByID(channelID string) (ChannelData, error) {
	var channel ChannelData
	channelObjID, err := primitive.ObjectIDFromHex(channelID)
	if err != nil {
		return channel, err
	}

	collection := MongoDB.Collection("channelList")
	filter := bson.M{"_id": channelObjID}
	err = collection.FindOne(context.TODO(), filter).Decode(&channel)
	return channel, err
}

func AddChannel(data map[string]interface{}) error {
	channel := ChannelData{
		Name:        data["name"].(string),
		Description: data["description"].(string),
		UserID:      data["userID"].(string),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	collection := MongoDB.Collection("channelList")
	_, err := collection.InsertOne(context.TODO(), channel)
	return err
}

func ExitChannel(channelID string) (bool, error) {
	channelObjID, err := primitive.ObjectIDFromHex(channelID) // string to objectID
	if err != nil {
		return false, nil
	}
	collection := MongoDB.Collection("channelList")
	filter := bson.M{"_id": channelObjID}
	update := bson.M{"$set": bson.M{"updated_at": time.Now()}}

	result, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return false, err
	}
	return result.ModifiedCount > 0, nil
}

func EditChannel(channelID string, data map[string]interface{}) error {
	channelObjID, err := primitive.ObjectIDFromHex(channelID)
	if err != nil {
		return err
	}

	collection := MongoDB.Collection("channelList")
	filter := bson.M{"_id": channelObjID}
	update := bson.M{"$set": bson.M{
		"name":        data["name"],
		"description": data["description"],
		"updated_at":  time.Now(),
	}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}

func DeleteChannel(channelID string) error {
	channelObjID, err := primitive.ObjectIDFromHex(channelID)
	if err != nil {
		return err
	}

	collection := MongoDB.Collection("channelList")
	filter := bson.M{"_id": channelObjID}

	_, err = collection.DeleteOne(context.TODO(), filter)
	return err
}
