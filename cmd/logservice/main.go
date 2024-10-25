package main

import (
	"context"
	"fmt"
	"log"
	
	logper "github.com/lyr1cs/fep/log"
	"github.com/lyr1cs/fep/service"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: log服务启动
 * @Date: 2024-10-13 21:48
 */

func main() {  
  logper.Run("./distributed.log")
	host, port := "localhost", "4000"
	serviceName := "Log Service"
	ctx, err := service.Start(
		context.Background(),
		serviceName,
		host,
		port,
		logper.RegisterHandlers,
	)
	if err != nil {
		log.Fatalln("Service start err:" + err.Error())
	}
	<-ctx.Done()
	fmt.Printf("Shutting down: %v\n", serviceName)
}
