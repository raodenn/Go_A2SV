package repositories

import (
	"context"
	"errors"
	"fmt"
	domain "task_manager/domain"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type UserRepo struct {
	Collection *mongo.Collection
}

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Username     string             `bson:"username,omitempty"`
	Password     *string            `bson:"password,omitempty"`
	UserType     string             `bson:"user_type,omitempty"`
	Token        *string            `bson:"token,omitempty"`
	RefreshToken *string            `bson:"refresh_token,omitempty"`
	CreatedAt    time.Time          `bson:"created_at,omitempty"`
	UpdatedAt    time.Time          `bson:"updated_at,omitempty"`
}

func initdb() *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("MongoDB connection failed: %v", err))
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("MongoDB ping failed: %v", err))
	}
	return client.Database("task_manager")
}

func NewUserRepo() domain.UserRepository {
	col := initdb().Collection("users")
	return &UserRepo{
		Collection: col,
	}
}

func FromDomain(u *domain.User) *User {
	return &User{
		ID:        primitive.NewObjectID(),
		Username:  u.Username,
		Password:  u.Password,
		UserType:  u.UserType,
		Token:     u.Token,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func ToDomain(u *User) *domain.User {
	return &domain.User{
		ID:        u.ID.Hex(),
		Username:  u.Username,
		Password:  u.Password,
		UserType:  u.UserType,
		Token:     u.Token,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (r *UserRepo) CreateUser(ctx context.Context, user *domain.User) error {
	err := r.Collection.FindOne(ctx, bson.M{"username": user.Username}).Err()
	if err == nil {
		return errors.New("username already exists")
	} else if err != mongo.ErrNoDocuments {
		return err
	}

	doc := FromDomain(user)
	_, insertErr := r.Collection.InsertOne(ctx, doc)
	return insertErr
}

func (r *UserRepo) GetUser(ctx context.Context, username string) (*domain.User, error) {
	var user User
	err := r.Collection.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return ToDomain(&user), nil
}
