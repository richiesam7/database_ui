package chat

import (
	"log"
	"golang.org/x/net/context"
	"encoding/json"
	"github.com/go-redis/redis"
	"fmt"
)

var client *redis.Client 

type Author struct {
	Name string `json:"name"`
	Age int `json:"age"`
}

type Server struct {
}

func init() {
	connectToRedis()
	pingRedis()
	set("id1234", "Value1234")
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	log.Printf("Receive message body from client: %s", in.Body)
	return &Message{Body: "Hello From the Server!"}, nil
}

func (s *Server) Get(ctx context.Context, in *RedisMessage) (*RedisMessage, error) {
	value := get(in.Key)
	log.Printf("Message from Client: Key- %s Value- %s", in.Key, in.Value)
	return &RedisMessage{Key: in.Key, Value: value}, nil
}

func (s *Server) Set(ctx context.Context, in *RedisMessage) (*RedisMessage, error) {
	log.Printf("Message from Client: Key- %s Value- %s", in.Key, in.Value)
	set(in.Key, in.Value)
	return &RedisMessage{Key: in.Key, Value: in.Value}, nil
}

func connectToRedis() {
	if client == nil {
		client = redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "",
			DB: 0,
		})	
	}
}

func pingRedis() bool {
	pong, err := client.Ping().Result()
	if (pong != "" && err == nil) {
		return true
	} else {
		return false
	}
}

func get(key string) string {
    val, err := client.Get(key).Result()
    if err != nil {
        fmt.Println(err)
    }
    return val
}

func set(key string, value string) {
    err := client.Set(key, value, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
}

func setTest(key string, value Author) {
	json, err := json.Marshal(value)
    if err != nil {
        fmt.Println(err)
    }

    err = client.Set(key, json, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
}