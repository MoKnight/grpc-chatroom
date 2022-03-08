/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-03-04 22:02:45
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-03-07 22:32:13
 */
package main

// 根据proto文件自动生成的代码
import (
	"bufio"
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	proto "chatroom/chatroomservice"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"
)

func main() {
	// 创建连接
	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		fmt.Printf("fail to connect: [%v]\n", err)
		return
	}
	defer conn.Close()

	// 声明客户端
	client := proto.NewChatServiceClient(conn)

	// 声明 context
	ctx := context.Background()

	// 创建双向数据流
	stream, err := client.Talk(ctx, grpc.UseCompressor(gzip.Name))
	if err != nil {
		fmt.Printf("fail to create stream: [%v]\n", err)
	}

	var loginfo proto.Request_Loginfo
	done := make(chan int, 1)

	// 启动一个 goroutine 接收命令行输入的指令
	go func() {
		input := bufio.NewReader(os.Stdin)

		//login step
		for {
			fmt.Println("please enter talkroom account:")
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			name, _ := input.ReadString('\n')

			fmt.Println("then enter talkroom account password:")
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			password, _ := input.ReadString('\n')
			// 向服务端发送 指令
			temp_loginfo := proto.Request_Loginfo{Name: name, Password: password}
			if err := stream.Send(&proto.Request{Type: 1, Loginformation: &temp_loginfo}); err != nil {
				return
			}

			// 接收从 服务端返回的数据流
			output, err := stream.Recv()
			if err == io.EOF {
				fmt.Println("it's time to stop")
				break //如果收到结束信号，则退出“接收循环”，结束客户端程序
			}

			if err != nil {
				// TODO: 处理接收错误
				fmt.Println("fail to get data:", err)
			}

			if output.Output == "you may start" {
				fmt.Println("client start to talk")
				loginfo.Name = temp_loginfo.GetName()
				loginfo.Password = temp_loginfo.GetPassword()
				done <- 1
				break
			}
		}

		//talk step
		for {
			log.Println("please talk:")
			// 获取 命令行输入的字符串， 以回车 \n 作为结束标志
			talkcontent, _ := input.ReadString('\n')
			if err := stream.Send(&proto.Request{TS: &timestamp.Timestamp{Seconds: time.Now().Unix()}, Type: 2, Input: talkcontent, Loginformation: &loginfo}); err != nil {
				return
			}
		}
	}()
	<-done
	for {
		// 接收从 服务端返回的数据流
		answercontent, err := stream.Recv()
		if err == io.EOF || answercontent.Output == "you may stop" {
			fmt.Println("it's time to stop")
			break //如果收到结束信号，则退出“接收循环”，结束客户端程序
		}

		if err != nil {
			// TODO: 处理接收错误
			fmt.Println("fail to get data:", err)
		}

		// 没有错误的情况下，打印来自服务端的消息
		fmt.Printf("%s said: %s", answercontent.Name, answercontent.Output)
	}

}
