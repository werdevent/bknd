package db

import (
	"context"
	"errors"
	"time"

	"github.com/GeorgeHN666/werdevent-backend/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DB) UpdateConsumerDetails(uuid string, CD map[string]interface{}) error {

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()

	db := s.Conn.Database(s.DB).Collection(constants.USER_COLLECTION)

	id, err := primitive.ObjectIDFromHex(uuid)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": CD,
	}

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	res, err := db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if res.ModifiedCount < 1 {
		return errors.New("no documents were modified")
	}

	return nil
}
