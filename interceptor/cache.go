package interceptor

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"helloworld/protobuf"
	"log"
	"strings"
	"time"
)

var redisClient *redis.Client

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
}

// UnaryCacheServerInterceptor ...
func UnaryCacheServerInterceptor(ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response interface{}, err error) {
	var key string
	var result interface{}

	switch info.FullMethod {
	case "/protobuf.Greeter/SayHello":
		r := request.(*protobuf.HelloRequest)
		key = fmt.Sprintf("%s:%s", info.FullMethod, strings.ReplaceAll(strings.ToLower(r.Name), " ", ""))
		result = &protobuf.HelloReply{}
	}

	if cacheExists(ctx, key) {
		err = getFromCache(ctx, key, &result)
		if err == nil {
			return result, nil
		}
		log.Println(err)
	}

	response, err = handler(ctx, request)
	if len(key) > 0 {
		setCache(ctx, key, response)
	}
	return
}

func getFromCache(ctx context.Context, key string, result interface{}) (err error) {
	log.Printf("get from cache: %s", key)
	cachedResponse := getCache(ctx, key)
	return json.Unmarshal([]byte(cachedResponse), result)
}

func cacheExists(ctx context.Context, key string) bool {
	if len(key) > 0 {
		r := redisClient.Exists(ctx, key)
		if r != nil {
			return r.Val() > 0
		}
	}
	return false
}

func getCache(ctx context.Context, key string) string {
	r := redisClient.Get(ctx, key)
	if r == nil {
		return ""
	}
	return r.Val()
}

func setCache(ctx context.Context, key string, response interface{}) {
	if len(key) > 0 {
		value, err := json.Marshal(response)
		if err == nil {
			redisClient.Set(ctx, key, value, 5*time.Minute)
		} else {
			log.Println(err)
		}
	}
}
