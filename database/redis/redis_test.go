package redis

// go test redis.go redis_test.go -v
// 指定函数 go test redis.go redis_test.go -run='TestGet' -v
import (
	"testing"
)

func TestSet(t *testing.T) {
	initClient()
	redisExampleSet()
}

func TestGet(t *testing.T) {
	initClient()
	redisExampleGet()
}
