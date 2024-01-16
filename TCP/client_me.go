package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// 创建TCP连接
	conn, err := net.Dial("tcp", "localhost:8080")
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

		// 将消息转换为字节并发送到服务器
		_, err := conn.Write([]byte(input + "\n"))
		if err != nil {
			fmt.Println("发送错误:", err)
			return
		}

		// 检查是否输入 'q!'，如果是则退出循环
		if input == "q!" {
			fmt.Println("退出客户端.")
			break
		}

		fmt.Printf("向TCP服务器发送消息: %s\n", input)

		// 读取服务器的响应
		_, err = bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			fmt.Println("读取响应错误:", err)
			return
		}
	}
}
