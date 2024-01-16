package main

import (
	"bufio"
	"fmt"
	"net"
)

// 处理单个客户端连接的函数
func handleConnection(conn net.Conn) {
	// 在函数结束时关闭连接
	defer func() {
		_ = conn.Close()
		fmt.Println("客户端连接关闭.")
	}()

	// 创建从连接中读取的 Scanner
	scanner := bufio.NewScanner(conn)

	for {
		// 读取客户端发送的数据
		scanned := scanner.Scan()
		if !scanned {
			break
		}

		// 获取客户端发送的消息
		message := scanner.Text()

		// 如果收到 'q!' 则结束循环，关闭连接
		if message == "q!" {
			fmt.Println("客户端请求关闭连接.")
			break
		}

		//输出接收到的消息
		fmt.Printf("从客户端接收到消息: %s\n", message)

		// 向客户端发送响应数据
		response := "收到你的消息: " + message
		_, err := conn.Write([]byte(response + "\n"))
		if err != nil {
			fmt.Println("发送响应错误:", err)
			break
		}
	}
}

func main() {
	// 监听本地8080端口
	listener, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("监听错误:", err)
		return
	}
	// 函数结束时关闭监听器
	defer func() { _ = listener.Close() }()

	fmt.Println("TCP服务器监听在 :8080 端口")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("接受连接错误:", err)
			break
		}

		fmt.Println("新的客户端连接.")

		// 启动新的 goroutine 处理连接
		go handleConnection(conn)
	}
}
