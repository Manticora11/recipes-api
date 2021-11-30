// Recipes API
//
// This is a sample recipes API. You can find out more about the API at https://github.com/Manticora11/recipes-api
//
//  Schemes: http
//  Host: localhost:8080
//  BasePath: /
//  Version: 1.0.0
//  Contact: Jose A. Devia <totes111167@gmail.com> https://twitter.com/Manticora1127
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta
package main

import (
	"context"
	"log"
	"os"

	handlers "github.com/Manticora11/recipes-api/handlers"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var recipesHandler *handlers.RecipesHandler
var ctx context.Context
var err error
var client *mongo.Client

func init() {
	/*
		recipes = make([]Recipe, 0)
		file, _ := ioutil.ReadFile("recipes.json")
		_ = json.Unmarshal([]byte(file), &recipes)
	*/
	ctx = context.Background()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err = client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB")
	collection := client.Database(os.Getenv("MONGO_DATABASE")).Collection("recipes")
	recipesHandler = handlers.NewRecipeHandler(ctx, collection)
	/*
		var listOfRecipes []interface{}
			for _, recipe := range recipes {
				listOfRecipes = append(listOfRecipes, recipe)
			}
			insertManyResult, err := collection.InsertMany(ctx, listOfRecipes)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("Inserted recipes: ", len(insertManyResult.InsertedIDs))
	*/
}

func main() {
	router := gin.Default()
	router.POST("/recipes", recipesHandler.NewRecipeHandler)
	router.GET("/recipes", recipesHandler.ListRecipesHandler)
	router.PUT("/recipes/:id", recipesHandler.UpdateRecipeHandler)
	router.DELETE("/recipes/:id", recipesHandler.DeleteRecipeHandler)
	//router.GET("/recipes/search", recipesHandler.SearchRecipesHandler)
	router.Run()
}
