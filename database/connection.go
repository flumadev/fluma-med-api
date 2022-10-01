package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() (*mongo.Collection, context.Context) {
	uri := "mongodb+srv://fluma:9TfCTLnvVOqU1Azx@medicinecluster.pohrbuy.mongodb.net/?retryWrites=true&w=majority"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	coll := client.Database("MedicineCluster").Collection("medicines")

	return coll, ctx
}
