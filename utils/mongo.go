package utils

import (
    "context"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// DB wraps the MongoDB client and collection
type DB struct {
    Client     *mongo.Client
    Collection *mongo.Collection
}

// Connect establishes a connection to MongoDB and selects the database/collection
func Connect(uri, dbName, collectionName string) (*DB, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    clientOptions := options.Client().ApplyURI(uri)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    coll := client.Database(dbName).Collection(collectionName)

    return &DB{
        Client:     client,
        Collection: coll,
    }, nil
}

// InsertOne inserts a single document
func (db *DB) InsertOne(document interface{}) (*mongo.InsertOneResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.Collection.InsertOne(ctx, document)
}

// FindAll returns all documents matching a filter
func (db *DB) FindAll(filter interface{}) ([]bson.M, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    cursor, err := db.Collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(ctx)

    var results []bson.M
    for cursor.Next(ctx) {
        var doc bson.M
        if err := cursor.Decode(&doc); err != nil {
            return nil, err
        }
        results = append(results, doc)
    }
    return results, cursor.Err()
}

// UpdateOne updates the first document matching the filter
func (db *DB) UpdateOne(filter, update interface{}) (*mongo.UpdateResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.Collection.UpdateOne(ctx, filter, update)
}

// DeleteOne deletes the first document matching the filter
func (db *DB) DeleteOne(filter interface{}) (*mongo.DeleteResult, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.Collection.DeleteOne(ctx, filter)
}

// Disconnect closes the connection to MongoDB
func (db *DB) Disconnect() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    return db.Client.Disconnect(ctx)
}
