package storage

import (
	"fmt"
	"time"

	"github.com/go-redis/redis"
)


var RedisDB redis.Client

func InitNewClient(addr string) (error) {
	RedisDB = *redis.NewClient(&redis.Options{
		Addr: addr,
	})
	err := RedisDB.Ping().Err()
	if err != nil {
		fmt.Printf("Ошибка соединения с Redis: %s\n", err)
		return err
	}
	fmt.Println("Соединение с Redis установлено")
	return nil
}

func SetToRedis(key string, value []byte) {
	err := RedisDB.Set(key, value, 5*time.Second).Err()
	if err != nil {
		fmt.Printf("Ошибка при вставке JSON в Redis: %v", err)
	}
	// fmt.Printf("JSON успешно вставлен в Redis под ключом: %s\n", key)
}


