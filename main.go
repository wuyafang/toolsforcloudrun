package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
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
			println(fmt.Sprintf("time passed %d seconds since programming start",time.Now().Second()-startTime.Second()))
			time.Sleep(duration)
		}
	}()

	//开启服务器
	fmt.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("helloworld: listening on port %s", port)
	fmt.Print(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
