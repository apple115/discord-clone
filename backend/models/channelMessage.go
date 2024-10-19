package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ChannelMessage struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	ChannelID primitive.ObjectID `bson:"channel_id" json:"channel_id"`
	Text      string             `bson:"text" json:"text"`
	UserID    string             `bson:"user_id" json:"user_id"`
	CreatedAt time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at" json:"updated_at"`
	Timestamp string             `bson:"timestamp" json:"timestamp"`
}

// 得到全部的消息，每次选择最后的50条，然后根据时间排序，最新的在最前面
func GetChannelMessages(channelID string) ([]ChannelMessage, error) {
	var messages []ChannelMessage
	channelObjID, err := primitive.ObjectIDFromHex(channelID) // string to objectID
	if err != nil {
		return nil, err
	}
	collection := MongoDB.Collection("channelMessage")

	// 查询条件
	filter := bson.M{"channel_id": channelObjID}

	// 设置排序和分页选项
	findOptions := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}). // 按 created_at 降序排序
		SetLimit(50)                                     // 限制返回的文档数量为50

	// 查询
	cursor, err := collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	if err = cursor.All(context.TODO(), &messages); err != nil {
		return nil, err
	}
	return messages, nil
}

func AddChannelMessage(data map[string]interface{}) error {
	message := ChannelMessage{
		ChannelID: data["channel_id"].(primitive.ObjectID),
		Text:      data["text"].(string),
		UserID:    data["user_id"].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	collection := MongoDB.Collection("channelMessage")
	_, err := collection.InsertOne(context.TODO(), message)
	return err
}

func EditChannelMessage(id string, data map[string]interface{}) error {
	// 将 id 转换为 ObjectID
	messageObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := MongoDB.Collection("channelMessage")
	filter := bson.M{"_id": messageObjID}
	update := bson.M{"$set": bson.M{
		"text":       data["text"],
		"updated_at": time.Now(),
	}}

	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}

// 删除消息，只是将内容变成“消息已删除”
func DeleteChannelMessage(id string) error {
	// 将 id 转换为 ObjectID
	messageObjID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	collection := MongoDB.Collection("channelMessage")
	filter := bson.M{"_id": messageObjID}
	update := bson.M{"$set": bson.M{
		"text":       "消息已删除",
		"updated_at": time.Now(),
	}}
	_, err = collection.UpdateOne(context.TODO(), filter, update)
	return err
}
