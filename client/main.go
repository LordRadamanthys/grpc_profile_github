package main

import (
	"context"
	"fmt"
	"log"

	user "github.com/LordRadamanthys/grpc_profile_github/pb"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8200", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	defer conn.Close()

	usr := user.NewUserServiceClient(conn)

	fmt.Print("Enter the username: ")
	var inputText string
	fmt.Scan(&inputText)

	response, err := usr.GetUser(context.Background(), &user.UserRequest{
		Username: inputText,
	})

	if err != nil {
		log.Fatal("Error calling GetUser")
	}

	log.Println(response)
}
