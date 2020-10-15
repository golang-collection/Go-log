package main

import (
	"Go-log/global"
	"Go-log/logger"
	"context"
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

/**
* @Author: super
* @Date: 2020-10-15 15:14
* @Description:
**/

func setupLogger() error {
	//通过配置文件进行配置
	fileName := "runtime/app.log"
	fmt.Println("log file name ", fileName)
	global.Logger = logger.NewLogger(&lumberjack.Logger{
		Filename:  fileName,
		MaxSize:   500,
		MaxAge:    10,
		LocalTime: true,
	}, "TRACE: ", log.LstdFlags).WithCaller(2)

	return nil
}

func init() {
	//初始化日志
	err := setupLogger()
	if err != nil {
		log.Printf("init setupLogger err: %v\n", err)
	}
}

func main() {
	global.Logger.Info(context.Background(), "log info")
	global.Logger.Warn(context.Background(), "log warn")
	global.Logger.Debug(context.Background(), "log debug")
	global.Logger.Error(context.Background(), "log error")
	global.Logger.Panic(context.Background(), "log panic")
}
