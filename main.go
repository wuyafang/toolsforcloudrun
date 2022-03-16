package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		fmt.Println(err)
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("helloworld: received a request")
	target := os.Getenv("TARGET")
	if target == "" {
		target = "World"
	}

	duration := 5 * time.Second
	sleepTime := os.Getenv("SLEEP_TIME")
	if sleepTime != "" {
		duration, _ = time.ParseDuration(sleepTime)
	}
	time.Sleep(duration)

	fmt.Fprintf(w, "Hello %s!\n", target)
}

func main() {
	//得到IP地址
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	//开启后台进程
	duration := 5 * time.Second
	sleepTime := os.Getenv("SLEEP_TIME")
	if sleepTime != "" {
		duration, _ = time.ParseDuration(sleepTime)
	}
	startTime := time.Now()
	go func() {
		for {
			println("background thread is running...")
			startDuration := time.Now().Sub(startTime)
			println(fmt.Sprintf("time passed %s since programming start",startDuration.String()))
			time.Sleep(duration)
		}
	}()

	//开启服务器
	fmt.Println("helloworld: starting server...")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(handler),
	}

	// Create channel to listen for signals.
	signalChan := make(chan os.Signal, 1)
	// SIGINT handles Ctrl+C locally.
	// SIGTERM handles Cloud Run termination signal.
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	// Start HTTP server.
	go func() {
		log.Printf("listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	// Receive output from signalChan.
	sig := <-signalChan
	log.Printf("%s signal caught", sig)

	// Timeout if waiting for connections to return idle.
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Add extra handling here to clean up resources, such as flushing logs and
	// closing any database or Redis connections.
	fmt.Println("after 10s the server will be shutdown")
	// Gracefully shutdown the server by waiting on existing requests (except websockets).
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %+v", err)
	}
	log.Print("server exited")
}
