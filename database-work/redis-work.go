package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"go-person-server/config"
)

func main() {
	conn := config.RedisConnPool().Get()
	defer conn.Close()
	//write
	//_, err := conn.Do("set", "test1", "huzonghong")
	//if err != nil {
	//	fmt.Println("redis set failed", err)
	//	return
	//}

	//r, err := redis.String(conn.Do("get", "test"))
	//if err != nil {
	//	fmt.Println("redis get failed", err)
	//	return
	//}

	//批量操作
	//_, err := conn.Do("MSet", "hgg", 100, "hgg1", 101, "hgg2", 102)
	//if err != nil {
	//	fmt.Println("redis MSet failed", err)
	//	return
	//}
	//r, err := redis.Ints(conn.Do("MGet", "hgg*", "hgg1", "hgg2"))
	//if err != nil {
	//	fmt.Println("redis MGet failed", err)
	//	return
	//}
	//for _, v := range r {
	//	fmt.Println("result:", v)
	//}

	//设置过期时间
	//_, err := conn.Do("expire", "hgg", 10)
	//if err != nil {
	//	fmt.Println("redis expire failed", err)
	//	return
	//}
	//v, _ := conn.Do("ttl", "hgg")
	//fmt.Println("expire:", v)

	//List  队列操作
	//_, err := conn.Do("lpush", "book_list_03", 100, 200, 300, 400, 500)
	//if err != nil {
	//	fmt.Println("redis lpush failed", err)
	//	return
	//}
	//v, _ := redis.Ints(conn.Do("lrange", "book_list_03", 0, -1))
	//
	//for _, s := range v {
	//	fmt.Println("lpop:", s)
	//}

	// ddl   删除操作
	//_, err := conn.Do("del", "test1", "test2", "hgg")
	//if err != nil {
	//	fmt.Println("redis del   failed", err)
	//	return
	//}
	//v, _ := redis.String(conn.Do("get", "test1"))
	//fmt.Println("删除以后key是否存在：", v)

	//  hash  操作

	_, err := conn.Do("HSet", "hash-test", "hash-test04", "1004", "hash-test05", "1005")
	if err != nil {
		fmt.Println("redis hash failed ", err)
		return
	}
	v, _ := redis.String(conn.Do("HGet", "hash-test", "hash-test05"))
	fmt.Println("HGet value :", v)

}
