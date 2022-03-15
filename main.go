package main

import (
	"fmt"
	"net"
	"strings"
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
	go func() {
		for {
			println("background thread is running...")
			time.Sleep(duration)
		}
	}
	
	//开启服务器
	log.Print("helloworld: starting server...")

	http.HandleFunc("/", handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("helloworld: listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
