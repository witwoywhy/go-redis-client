package main

import (
	"fmt"
	"gedis"
	"time"
)

type User struct {
	Name       string   `json:"name"`
	Age        int64    `json:"age"`
	Permission []string `json:"permission"`
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	redis, err := gedis.NewGedis(&gedis.Config{
		Host:           "localhost",
		Port:           "6379",
		Password:       "",
		DB:             0,
		ConnectionPool: 10,
	})
	if err != nil {
		panic(err)
	}

	now := time.Now()

	fmt.Println("===== SET =====")
	user := User{
		Name:       "Test Tester",
		Age:        20,
		Permission: []string{"ENABLE"},
	}
	must(redis.Set("TEST:1", user))
	must(redis.Set("TEST:2", 20))
	must(redis.Set("TEST:3", "stringer"))

	fmt.Println("===== GET =====")
	fmt.Println(redis.Get("TEST:1"))
	fmt.Println(redis.Get("TEST:2"))
	fmt.Println(redis.Get("TEST:3"))

	fmt.Println("===== MSET =====")
	user.Name = "Test Test2"
	user.Age = 21
	user.Permission = append(user.Permission, "DISABLE")
	must(
		redis.MSet([]gedis.Multiple{
			{Key: "TEST:4", Value: user},
			{Key: "TEST:5", Value: 21},
			{Key: "TEST:6", Value: 666.777},
		}),
	)
	fmt.Println("===== MGET =====")
	fmt.Println(redis.MGet("TEST:4", "TEST:5", "TEST:6"))

	fmt.Println("===== GETDEL =====")
	fmt.Println(redis.GetDel("TEST:6"))
	fmt.Println(redis.Get("TEST:6"))

	fmt.Println("===== GETSET =====")
	fmt.Println(redis.GetSet("TEST:5", 22))
	fmt.Println(redis.Get("TEST:5"))

	fmt.Println("===== APPEND =====")
	must(redis.Set("TEST:7", "Hello"))
	fmt.Println(redis.Append("TEST:7", " World"))
	fmt.Println(redis.Get("TEST:7"))

	fmt.Println("===== DECR =====")
	fmt.Println(redis.Decr("TEST:8"))

	fmt.Println("===== DECRBY =====")
	fmt.Println(redis.DecrBy("TEST:8", 1))

	fmt.Println("===== GETEX =====")
	must(redis.Set("TEST:9", 9))
	must(redis.Set("TEST:10", 10))
	must(redis.Set("TEST:11", 11))
	must(redis.Set("TEST:12", 12))
	must(redis.Set("TEST:13", 13))
	must(redis.Set("TEST:14", 14))

	tenSecond := now.Add(3 * time.Second)

	fmt.Println(redis.GetEx("TEST:9"))
	fmt.Println(redis.GetEx("TEST:10", gedis.TTL{Option: gedis.EX, Time: "3"}))
	fmt.Println(redis.GetEx("TEST:11", gedis.TTL{Option: gedis.PX, Time: "3000"}))
	fmt.Println(redis.GetEx("TEST:12", gedis.TTL{Option: gedis.EXAT, Time: fmt.Sprintf("%v", tenSecond.Unix())}))
	fmt.Println(redis.GetEx("TEST:13", gedis.TTL{Option: gedis.PXAT, Time: fmt.Sprintf("%v", tenSecond.UnixMilli())}))
	fmt.Println(redis.GetEx("TEST:14", gedis.TTL{Option: gedis.EX, Time: "20"}))
	fmt.Println(redis.GetEx("TEST:14", gedis.TTL{Option: gedis.PERSIST}))

	time.Sleep(3 * time.Second)
	fmt.Println(redis.Get("TEST:9"))
	fmt.Println(redis.Get("TEST:10"))
	fmt.Println(redis.Get("TEST:11"))
	fmt.Println(redis.Get("TEST:12"))
	fmt.Println(redis.Get("TEST:13"))
	fmt.Println(redis.Get("TEST:14"))

	fmt.Println("===== GETRAGE =====")
	fmt.Println(redis.GetRange("TEST:1", 5, 10))

	fmt.Println("===== INCR, INCRBY, INCRBYFLOAT =====")
	fmt.Println(redis.Incr("TEST:15"))
	fmt.Println(redis.IncrBy("TEST:16", 2))
	fmt.Println(redis.IncrByFloat("TEST:17", 3.5))
}
