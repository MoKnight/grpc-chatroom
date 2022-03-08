/*
 * @Description:
 * @version:
 * @Author: MoonKnight
 * @Date: 2022-03-04 22:02:54
 * @LastEditors: MoonKnight
 * @LastEditTime: 2022-03-07 22:31:24
 */

package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"

	proto "chatroom/chatroomservice"

	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

//服务端
type Service struct {
	proto.UnimplementedChatServiceServer
}

type ConnectPool struct {
	sync.Map
}

var connect_pool *ConnectPool

func (p *ConnectPool) Get(name string) proto.ChatService_TalkServer {
	if stream, ok := p.Load(name); ok {

		return stream.(proto.ChatService_TalkServer)
	} else {
		return nil
	}
}

func (p *ConnectPool) Add(name string, stream proto.ChatService_TalkServer) {
	p.Store(name, stream)
}

func (p *ConnectPool) Del(name string) {
	p.Delete(name)
}

func (p *ConnectPool) BroadCast(from, message string) {
	log.Printf("BroadCast from: %s, message: %s\n", from, message)
	p.Range(func(username_i, stream_i interface{}) bool {
		username := username_i.(string)
		stream := stream_i.(proto.ChatService_TalkServer)
		if username == from { //不发给自身
			return true
		} else {
			stream.Send(&proto.Response{
				TS:     &timestamp.Timestamp{Seconds: time.Now().Unix()},
				Name:   username,
				Output: message,
			})
		}
		return true
	})
}

func (s *Service) Talk(stream proto.ChatService_TalkServer) error {
	peer, _ := peer.FromContext(stream.Context())
	fmt.Printf("Received new connection.  %s\n", peer.Addr.String())
	var username string

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		if req.Type == 1 { //请求登录
			username = req.GetLoginformation().GetName()
			if connect_pool.Get(username) != nil {
				stream.Send(&proto.Response{
					TS:     &timestamp.Timestamp{Seconds: time.Now().Unix()},
					Name:   username,
					Output: fmt.Sprintf("username %s already exists!", username),
				})
				return nil
			} else { // 连接成功
				connect_pool.Store(username, stream)
				stream.Send(&proto.Response{
					TS:     &timestamp.Timestamp{Seconds: time.Now().Unix()},
					Name:   username,
					Output: "you may start",
				})
				break
			}
		}
	}

	go func() {
		<-stream.Context().Done()
		connect_pool.Del(username)
		connect_pool.BroadCast(username, fmt.Sprintf("%s leave room", username))
	}()
	connect_pool.BroadCast(username, fmt.Sprintf("Welcome %s!", username))

	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		if req.Type == 2 {
			connect_pool.BroadCast(username, fmt.Sprintf("%s: %s", username, req.Input))
		} else if req.Type == 3 {
			stream.Send(&proto.Response{
				TS:     &timestamp.Timestamp{Seconds: time.Now().Unix()},
				Name:   username,
				Output: "it's time to stop",
			})
		}

	}
	return nil
}

func main() {
	//gin part
	//r := gin.Default()
	connect_pool = &ConnectPool{}

	//go r.Run(":8000")

	//grpc part
	server := grpc.NewServer(
		grpc.RPCCompressor(grpc.NewGZIPCompressor()),
		grpc.RPCDecompressor(grpc.NewGZIPDecompressor()),
	)
	// 注册 ChatServer
	proto.RegisterChatServiceServer(server, &Service{})
	address, err := net.Listen("tcp", ":3000")
	fmt.Println("启动服务端...")
	if err != nil {
		panic(err)
	}
	if err := server.Serve(address); err != nil {
		panic(err)
	}
}
