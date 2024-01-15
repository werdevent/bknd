package db

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"github.com/GeorgeHN666/werdevent-backend/constants"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *DB) InsertUser(u *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	db := s.Conn.Database(s.DB).Collection(constants.USER_COLLECTION)

	u.ID = primitive.NewObjectID()
	u.Created_At = fmt.Sprint(time.Now().Unix())
	u.Updated_At = fmt.Sprint(time.Now().Unix())
	u.Last_Updated = "Default"

	_, err := db.InsertOne(ctx, u)
	if err != nil {
		return err
	}
	return nil
}

func (s *DB) GetUser(e string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	db := s.Conn.Database(s.DB).Collection(constants.USER_COLLECTION)

	var res models.User

	err := db.FindOne(ctx, bson.M{"email": e}).Decode(&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (s *DB) UpdateUserDetails(UID string, payloadmodel *models.User) error {

	ctx, cancel := context.WithTimeout(context.Background(), 45*time.Second)
	defer cancel()
	db := s.Conn.Database(s.DB).Collection(constants.USER_COLLECTION)

	UpdatedModel := utils.FilterEmptyStructs(payloadmodel)
	UpdatedModel["valid_code"] = payloadmodel.Valid_Code

	update := bson.M{
		"$set": UpdatedModel,
	}

	id, _ := primitive.ObjectIDFromHex(UID)

	filter := bson.M{
		"_id": bson.M{"$eq": id},
	}

	UpdatedUser, err := db.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	if UpdatedUser.ModifiedCount == 0 {
		return errors.New("no changes were done")
	}

	return nil
}

// Delete function
