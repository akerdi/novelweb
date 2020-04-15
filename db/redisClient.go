package db

import (
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	jsoniter "github.com/json-iterator/go"
	"novelweb/config"
	"novelweb/db/schema"
	"time"
)

type RedisClient struct {
	client *redis.Client
}

var RedisConnector *RedisClient

func InitRedisClient() {
	addr := fmt.Sprintf("%s:%d", config.GetConfig().Redis.Host, config.GetConfig().Redis.Port)
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: config.GetConfig().Redis.Pass,
		DB:       config.GetConfig().Redis.DB,
	})
	RedisConnector = &RedisClient{
		client: client,
	}
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("pongpongpong: ", pong)
}

func (connector *RedisClient) SaveChapterByNovelmd5(novelId string, chapter *schema.NovelChapter) (err error) {
	if chapter == nil {
		err = errors.New(fmt.Sprintf("novel:%s save with empty chapter!", novelId))
		return
	}
	chapterJson, err := jsoniter.ConfigFastest.MarshalToString(chapter)
	if err != nil {
		return
	}
	err = connector.client.Set(novelId, chapterJson, 24*2.5*time.Hour).Err()
	if err != nil {
		return
	}
	return
}

func (connector *RedisClient) ChapterWithNovelmd5(novelId string) (*schema.NovelChapter, error) {
	var chapters *schema.NovelChapter
	chapterJson, err := connector.client.Get(novelId).Result()
	if err != redis.Nil || chapterJson == "" {
		return chapters, errors.New(fmt.Sprintf("Redis Error: %s not exits??", novelId))
	}
	err = jsoniter.ConfigFastest.UnmarshalFromString(chapterJson, &chapters)
	if err != nil {
		return nil, err
	}
	return chapters, nil
}
