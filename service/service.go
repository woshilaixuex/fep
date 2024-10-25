package service

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

/*
 * @Author: lyr1cs
 * @Email: linyugang7295@gmail.com
 * @Description: 服务启动
 * @Date: 2024-10-11 22:47
 */

/*
 * 这段代码的关键在于registerHandlersFunc是如何绑定到srv上的
 *
 */
func Start(ctx context.Context, serviceName, host, port string,
	registerHandlersFunc func()) (context.Context, error) {
	registerHandlersFunc()
	ctx = startService(ctx, serviceName, host, port)
	return ctx, nil
}
func startService(ctx context.Context, serviceName, host, port string) context.Context {
  ctx, cancel := context.WithCancel(ctx)
	var srv http.Server
	srv.Addr = host + ":" + port
	go func() {
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Printf("HTTP server ListenAndServe: %v", err)
		}
	}()
	go func() {
		fmt.Printf("%v started Press any key top stop\n", serviceName)
		var s string
		fmt.Scanln(&s)
		srv.Shutdown(ctx)
		cancel()
	}()
	return ctx
}
