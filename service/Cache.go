package service

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
	"time"
)

var System_cache cache.Cache

func CacheInit(){
	var err error
	System_cache,err=cache.NewCache("redis", `{"key":"lbccsu","conn":":6379","dbNum":"0"}`)
	if err!=nil{
		beego.Error("redis cache init err:",err)
	}

}


func Cache_Get(key string) string{
	if System_cache.Get(key) !=nil{
		return string(System_cache.Get(key).([]byte)[:])
	}else{
		return ""
	}
}

func CachePut(key,value string,time time.Duration) error{
	return System_cache.Put(key,[]byte(value),time)
}