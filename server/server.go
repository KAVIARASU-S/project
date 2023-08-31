package main

import(
	"banking/controllers"
)


func initDatabase(client *mongo.Client) {
	profileCollection := config.GetCollection(client,"banking","customer")
	controllers.
}