package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"strconv"

	proto "github.com/shivarajshanthaiah-grpc-sample/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	proto.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()
	proto.RegisterExampleServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

//Unray Operation
// func (s *server) ServerReply(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
// 	fmt.Println("Recieve Request from client", req.SomeString)
// 	fmt.Println("hello from server")
// 	return &proto.HelloResponse{}, errors.New("")
// }

//Client streaming
// func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
// 	total := 0
// 	for {
// 		request, err := stream.Recv()
// 		if err == io.EOF {
// 			return stream.SendAndClose(&proto.HelloResponse{
// 				Reply: strconv.Itoa(total),
// 			})
// 		}
// 		if err != nil {
// 			return err
// 		}
// 		total++
// 		fmt.Println(request)
// 	}
// }

// Server Streaming
// func (s *server) ServerReply(req *proto.HelloRequest, stream proto.Example_ServerReplyServer) error {
// 	fmt.Println(req.SomeString)
// 	time.Sleep(5 * time.Second)
// 	serverReplay := []*proto.HelloResponse{
// 		{Reply: "Reply 1"},
// 		{Reply: "Reply 2"},
// 		{Reply: "Reply 3"},
// 		{Reply: "Reply 4"},
// 	}

// 	for _, msg := range serverReplay {
// 		err := stream.Send(msg)
// 		if err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// Bi-directional streaming
func (s *server) ServerReply(stream proto.Example_ServerReplyServer) error {
	for i := 0; i < 5; i++ {
		err := stream.Send(&proto.HelloResponse{Reply: "message" + strconv.Itoa(i) + "from server"})
		if err != nil {
			return errors.New("Unable to send data from server")
		}
	}
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println(req.SomeString)
	}
	return nil
}
