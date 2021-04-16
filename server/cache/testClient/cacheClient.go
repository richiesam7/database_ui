package main

import (
	"log"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"github.com/tutorialedge/go-grpc-beginners-tutorial/chat"
)

func cacheSet(key string, value string) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	response, err := c.Set(context.Background(), &chat.RedisMessage{Key: key, Value: value})
	if err != nil {
		log.Fatalf("Error when calling Set: %s", err)
	}
	log.Printf("Message from Server: Key- %s Value- %s", response.Key, response.Value)
	response, err = c.Set(context.Background(), &chat.RedisMessage{Key: "id1234", Value: "NewValue1234"})
	if err != nil {
		log.Fatalf("Error when calling Set: %s", err)
	}
	log.Printf("Message from Server: Key- %s Value- %s", response.Key, response.Value)
}

func cacheGet() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)
	response, err := c.Get(context.Background(), &chat.RedisMessage{Key: "id1234", Value: ""})
	if err != nil {
		log.Fatalf("Error when calling Set: %s", err)
	}
	log.Printf("Message from Server: Key- %s Value- %s", response.Key, response.Value)
	response, err = c.Set(context.Background(), &chat.RedisMessage{Key: "id1234", Value: "NewValue1234"})
	if err != nil {
		log.Fatalf("Error when calling Set: %s", err)
	}
	log.Printf("Message from Server: Key- %s Value- %s", response.Key, response.Value)
}
