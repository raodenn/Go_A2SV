package data

import (
	"Task_management_api/middleware"
	"Task_management_api/models"
	"context"
	"errors"
	"log"
	"time"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

// var UserCollection *mongo.Collection = MongoClient.Database("taskdb").Collection("users")

func HashPassword(password string) *string {
	pass, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic(err)

	}
	hashed := string(pass)
	return &hashed
}

func VerifyPassword(userPass string, foundPass string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(foundPass), []byte(userPass))
	if err != nil {
		return false, err
	}
	return true, nil
}

func CreateUser(user *models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := validate.Struct(user)
	if err != nil {
		return err
	}
	var existingUser models.User
	err = UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		//username already exists

		return errors.New("username already exists")
	}

	user.Password = HashPassword(*user.Password)

	user.CreatedAt = time.Now().Truncate(time.Second)
	user.UpdatedAt = time.Now().Truncate(time.Second)
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()
	token, refreshToken, _ := middleware.GenerateToken(*user.Username, *user.UserType, user.UserId)
	user.Token = &token
	user.RefreshToken = &refreshToken

	_, insertErr := UserCollection.InsertOne(ctx, user)
	if insertErr != nil {
		return insertErr
	}
	return err
}

func Login(user models.User) (models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var foundUser models.User
	err := UserCollection.FindOne(ctx, bson.M{"username": user.Username}).Decode(&foundUser)
	defer cancel()
	if err != nil {
		return models.User{}, err
	}
	valid, err := VerifyPassword(*user.Password, *foundUser.Password)
	if err != nil {
		return models.User{}, err
	}
	if !valid {
		log.Panic("Wrong password")
	}
	defer cancel()
	token, refreshtoken, _ := middleware.GenerateToken(*foundUser.Username, *foundUser.UserType, foundUser.UserId)
	update := bson.M{"$set": bson.M{
		"token":         token,
		"refresh_token": refreshtoken,
		"updated_at":    time.Now(),
	}}
	_, err = UserCollection.UpdateOne(ctx, bson.M{"user_id": foundUser.UserId}, update)
	if err != nil {
		return models.User{}, err
	}
	foundUser.Token = &token
	foundUser.RefreshToken = &refreshtoken
	return foundUser, nil
}
