package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"main/database"
	"reflect"
)

type TLogger struct {
	logger *log.Logger
}

func Status(logger *log.Logger) *TLogger {
	return &TLogger{logger}
}

func (status *TLogger) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	if request.Method == http.MethodGet {
		getMedicines(response, request)
	}

}

func getMedicines(response http.ResponseWriter, request *http.Request) {
	coll, ctx := database

	cursor, err := coll.Find(
		context.TODO(),
		bson.D{{Key: "NUMERO_REGISTRO_PRODUTO", Value: "154230044"}},
	)

	if err != nil {
		fmt.Println("Error on query", err)
		defer cursor.Close(ctx)
	} else {
		for cursor.Next(ctx) {
			var result bson.M
			err := cursor.Decode(&result)
			if err != nil {
				fmt.Println("Cursos.next() error:", err)
				os.Exit(1)
			} else {
				fmt.Println("\nresult type:", reflect.TypeOf(result))
				fmt.Println("result", result)
			}
		}
	}

	fmt.Fprintf(response, "get")
}
