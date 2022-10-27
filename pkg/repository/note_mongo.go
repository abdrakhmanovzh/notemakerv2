package repository

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoteMongo struct {
	db *mongo.Database
}

func NewNoteMongo(db *mongo.Database) *NoteMongo {
	return &NoteMongo{db: db}
}

func (r *NoteMongo) Create(userId int, note model.Note) (int, error) {
	rand.Seed(time.Now().UnixNano())
	newId := rand.Intn(1000)

	_, gg := r.db.Collection(os.Getenv("NOTES")).InsertOne(context.TODO(), bson.D{{Key: "id", Value: newId}, {Key: "title", Value: note.Title}, {Key: "content", Value: note.Content}})
	if gg != nil {
		log.Println("caught exception during insertion#1, aborting.")
		return 0, gg
	}

	_, ok := r.db.Collection(os.Getenv("USERS_NOTES")).InsertOne(context.TODO(), bson.D{{Key: "user_id", Value: userId}, {Key: "note_id", Value: newId}})
	if ok != nil {
		log.Println("caught exception during insertion#2, aborting.")
		return 0, ok
	}
	return newId, nil
}

func (r *NoteMongo) GetAll(userId int) ([]model.Note, error) {
	var notes []model.Note
	var foundNote model.Note

	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}, {Key: "note_id", Value: 1}})

	res, err := r.db.Collection(os.Getenv("USERS_NOTES")).Find(context.TODO(), bson.M{"user_id": userId}, opts)
	if err != nil {
		log.Println("error while trying to find link")
		return nil, err
	}
	var filteredId []map[string]int

	if err = res.All(context.TODO(), &filteredId); err != nil {
		log.Fatal(err)
	}
	for _, i := range filteredId {
		filter := bson.M{"id": i["note_id"]}
		err := r.db.Collection(os.Getenv("NOTES")).FindOne(context.TODO(), filter).Decode(&foundNote)
		if err != nil {
			log.Println("cannot find from notes database")
			return nil, err
		}

		notes = append(notes, foundNote)
	}
	return notes, nil
}

func (r *NoteMongo) GetById(userId, noteId int) (model.Note, error) {
	var note model.Note

	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}, {Key: "note_id", Value: 1}})
	res, err := r.db.Collection(os.Getenv("USERS_NOTES")).Find(context.TODO(), bson.M{"user_id": userId}, opts)
	if err != nil {
		log.Println("error while trying to find link")
	}

	var filteredId []map[string]int

	if err = res.All(context.TODO(), &filteredId); err != nil {
		log.Fatal(err)
	}

	for _, i := range filteredId {
		if i["note_id"] == noteId {
			filter := bson.M{"id": noteId}
			err := r.db.Collection(os.Getenv("NOTES")).FindOne(context.TODO(), filter).Decode(&note)
			if err != nil {
				log.Println("cannot find from notes database")
			}
		}
	}
	return note, nil
}

func (r *NoteMongo) Delete(userId, noteId int) error {
	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}, {Key: "note_id", Value: 1}})
	res, err := r.db.Collection(os.Getenv("USERS_NOTES")).Find(context.TODO(), bson.M{"user_id": userId}, opts)
	if err != nil {
		log.Println("error while trying to find link")
	}

	var filteredId []map[string]int

	if err = res.All(context.TODO(), &filteredId); err != nil {
		log.Fatal(err)
	}
	for _, i := range filteredId {
		if i["note_id"] == noteId {
			filter := bson.D{{Key: "id", Value: noteId}}
			res, err := r.db.Collection(os.Getenv("NOTES")).DeleteOne(context.TODO(), filter)
			if err != nil {
				log.Println("cannot find from notes database")
				return err
			}
			if res.DeletedCount != 1 {
				log.Println("deleted count is wrong")
				return err
			}

			_, gg := r.db.Collection(os.Getenv("USERS_NOTES")).DeleteOne(context.TODO(), bson.D{{Key: "note_id", Value: noteId}})
			if gg != nil {
				log.Println("cannot delete from users_notes")
				return gg
			}
		}
	}
	return nil
}

func (r NoteMongo) Update(userId int, noteId int, input model.UpdateNoteInput) error {
	opts := options.Find().SetProjection(bson.D{{Key: "_id", Value: 0}, {Key: "note_id", Value: 1}})
	res, err := r.db.Collection(os.Getenv("USERS_NOTES")).Find(context.TODO(), bson.M{"user_id": userId}, opts)
	if err != nil {
		log.Println("error while trying to find link")
	}

	var filteredId []map[string]int

	if err = res.All(context.TODO(), &filteredId); err != nil {
		log.Fatal(err)
	}
	for _, i := range filteredId {
		if i["note_id"] == noteId {
			_, err := r.db.Collection(os.Getenv("NOTES")).UpdateOne(
				context.TODO(),
				bson.D{{Key: "id", Value: noteId}},
				bson.D{{Key: "$set", Value: bson.D{{Key: "title", Value: input.Title}, {Key: "content", Value: input.Content}}}},
			)
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
	}
	return nil
}
