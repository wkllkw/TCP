package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 解析UDP服务器地址
	serverAddr, err := net.ResolveUDPAddr("udp", "localhost:8080")
	if err != nil {
		fmt.Println("解析服务器地址错误:", err)
		return
	}

	// 创建UDP连接
	conn, err := net.DialUDP("udp", nil, serverAddr)
	if err != nil {
		fmt.Println("连接错误:", err)
		return
	}
	// 函数结束时关闭连接
	defer func() { _ = conn.Close() }()

	// 创建从标准输入读取的 Scanner
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("请输入消息 (输入 'q!' 退出): ")

		// 读取用户输入
		scanner.Scan()
		input := scanner.Text()

		// 检查是否输入 'q!'，如果是则退出循环
		if input == "q!" {
			conn.Write([]byte(input))
			fmt.Println("退出客户端.")
			break
		}

		// 将消息转换为字节并发送到服务器
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println("发送错误:", err)
			return
		}

		fmt.Printf("向UDP服务器发送消息: %s\n", input)
	}
}
