package database

import (
	"context"
	"errors"
	"time"

	"github.com/GeorgeHN666/werdevent-backend/app/models"
	"github.com/GeorgeHN666/werdevent-backend/app/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// InsertUser will insert a new user document to the database
func (s *DB) InsertUser(user *models.User) (*models.User, error) {

	ctx, closeDB := context.WithTimeout(context.Background(), 1*time.Minute)
	defer closeDB()

	db := s.Client.Database(s.Database).Collection(models.USERS_DB)

	user.ID = primitive.NewObjectID()
	user.CreatedAt = time.Now().Local().String()
	user.UpdatedAt = time.Now().Local().String()

	_, err := db.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUser will get a specified user by id from the database
func (s *DB) GetUser(userID string) (*models.User, error) {

	ctx, closeDB := context.WithTimeout(context.Background(), 1*time.Minute)
	defer closeDB()

	db := s.Client.Database(s.Database).Collection(models.USERS_DB)

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	filter := bson.M{
		"_id": id,
	}

	var result *models.User

	err = db.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no document were found")
		} else {
			return nil, err
		}
	}

	return result, nil
}

// UpdateUserDetails will filter and update the given field to a specific user in the database
func (s *DB) UpdateUserDetails(UserID string, UpdateModel *models.User) error {

	ctx, closeDB := context.WithTimeout(context.Background(), 1*time.Minute)
	defer closeDB()

	db := s.Client.Database(s.Database).Collection(models.USERS_DB)

	updateDoc := utils.FilterEmptyStructs(UpdateModel)

	updatedUser, err := db.UpdateByID(ctx, UserID, updateDoc)
	if err != nil {
		return err
	}
	if updatedUser.ModifiedCount == 0 {
		return errors.New("no users were founded")
	}

	return nil
}

// DeleteUser will delete the document for a specified user in the database
func (s *DB) DeleteUser(userID string) error {

	ctx, closeDB := context.WithTimeout(context.Background(), 1*time.Minute)
	defer closeDB()

	db := s.Client.Database(s.Database).Collection(models.USERS_DB)

	id, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": id,
	}

	_, err = db.DeleteOne(ctx, filter)
	if err != nil {
		return err
	}
	return nil
}
