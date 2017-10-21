package jz

import (
	"log"
	redis "github.com/jonnywang/go-kits/redis"
)

func init()  {
	redis.Logger.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
}

var JzLogger = redis.Logger