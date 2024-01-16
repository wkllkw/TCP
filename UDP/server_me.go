package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	// 解析UDP地址
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("解析地址错误:", err)
		return
	}

	// 创建UDP监听器
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
	// 函数结束时关闭连接
	defer func() { _ = conn.Close() }()

	fmt.Println("UDP服务器监听在 :8080 端口")

	// 创建缓冲区，用于读取数据
	buffer := make([]byte, 1024)

	for {
		// 读取UDP连接上的数据
		n, _, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("读取错误:", err)
			return
		}

		// 将字节转换为字符串，表示接收到的消息
		message := string(buffer[:n])
		message2 := "q!"

		// 判断是否收到关闭连接的消息
		// 如果收到 'q!' 则结束循环，关闭连接
		//if message == "q!"
		if strings.EqualFold(message, message2) {
			fmt.Println("客户端请求关闭连接.")
			fmt.Println("客户端连接关闭")
			break
		} else {
			fmt.Printf("从客户端接收到消息: %s\n", message)
		}
	}
}
