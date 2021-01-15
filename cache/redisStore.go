package cache

import (
	"context"
	"fmt"
	"github.com/liuhongdi/digv19/global"
	"log"
	"time"
)

var ctx = context.Background()

const CAPTCHA = "captcha:"

type RedisStore struct {
}

//set a capt
func (r RedisStore) Set(id string, value string) {
	key := CAPTCHA + id
	err := global.RedisDb.Set(ctx,key, value, time.Minute*2).Err()
	if err != nil {
		log.Println(err.Error())
	}
}
//get a capt
func (r RedisStore) Get(id string, clear bool) string {
	key := CAPTCHA + id
	val, err := global.RedisDb.Get(ctx, key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	if clear {
		err := global.RedisDb.Del(ctx,key).Err()
		if err != nil {
			fmt.Println(err)
		    return ""
	    }
	}
	return val
}
//verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := RedisStore{}.Get(id, clear)
	//fmt.Println("key:"+id+";value:"+v+";answer:"+answer)
	return v == answer
}

