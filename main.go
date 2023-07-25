package main

import (
	"GoORM-Mongo/mongorm"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	mongorm.Model
	FirstName string `bson:"fist_name"`
	LastName  string `bson:"last_name"`
	Email     string `bson:"email"`
}

func main() {
	fmt.Println("App started!")
	client, err := mongorm.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}

	db := client.Database("test_db")

	// Create new user
	user := User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "j.d@example.com",
	}
	err = user.Create(context.Background(), db, "users", &user)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User created: %v\n", user)

	// Read a user bu ID
	var readUser User
	err = readUser.Read(context.Background(), db, "users", bson.M{"_id": user.ID}, &readUser)
	if err != nil {
		panic(err)
	}
	fmt.Printf("User read: %v\n", readUser)

  // Update a uers's email
  update := bson.M{"$set": bson.M{"email": "j.d.updated@examgle.com"}}
  err = user.Update(context.Background(), db, "users", bson.M{"_id": user.ID}, update)
  if err != nil {
    panic(err)
  }
  fmt.Printf("User updated: %v\n", user)
  
  // Delete a user by _id
  err = user.Delete(context.Background(), db, "users", bson.M{"_id": user.ID})
  if err != nil {
    panic(err)
  }
  fmt.Printf("User deleted")
}
