package repository

import (
	"context"
	"errors"
	"math/rand"
	"os"
	"time"

	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthMongo struct {
	db *mongo.Database
}

func NewAuthMongo(db *mongo.Database) *AuthMongo {
	return &AuthMongo{db: db}
}

func (r *AuthMongo) CreateUser(user model.User) (string, error) {
	user.ID = primitive.NewObjectID()
	rand.Seed(time.Now().UnixNano())
	user.Id = rand.Intn(10000)

	collection := r.db.Collection(os.Getenv("USERS"))

	res, insertErr := collection.InsertOne(context.TODO(), user)
	if insertErr != nil {
		return "", errors.New("user is already registered in the system")
	}

	id := res.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

func (r *AuthMongo) GetUser(username string) (model.User, error) {
	var user model.User
	collection := r.db.Collection(os.Getenv("USERS"))

	if err := collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user); err != nil {
		//log.Fatal(err)
		return user, err
	}

	return user, nil
}
