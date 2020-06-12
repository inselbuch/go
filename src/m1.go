package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// You will be using this Trainer type later in the program
type Trainer struct {
    Name string
    Age  int
    City string
}

func main() {

   // Set client options
   clientOptions := options.Client().ApplyURI("mongodb://192.168.1.189:27017")

   // Connect to MongoDB
   ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
   client, err := mongo.Connect(context.TODO(), clientOptions)

   if err != nil {
      log.Fatal(err)
   }

   // Check the connection
   err = client.Ping(context.TODO(), nil)

   if err != nil {
      log.Fatal(err)
   }

   fmt.Println("Connected to MongoDB!")

   databases, err := client.ListDatabaseNames(ctx, bson.M{})
   if err != nil {
      log.Fatal(err)
   }
   fmt.Println(databases)

// mongoDB primitives

// primitive.ObjectId [objectid hash]
// bson.D [document]
// bson.M [map]
// bson.A [array]

   quickstartDatabase := client.Database("quickstart")
   podcastsCollection := quickstartDatabase.Collection("podcasts")
   //episodesCollection := quickstartDatabase.Collection("episodes")

   // create (CRUD)
   document := bson.D {
      {"title", "The Polyglot Developer Podcast"},
      {"author", "Nic Raboy"},
      {"tags", bson.A {"development","programming","coding"}},
   }

   podcastResult, err := podcastsCollection.InsertOne(ctx,document)

   if err != nil {
      log.Fatal(err)
   }

   fmt.Printf("Inserted %v into collection.\n",podcastResult.InsertedID)
}

