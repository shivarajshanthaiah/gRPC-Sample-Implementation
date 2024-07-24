package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	proto "github.com/shivarajshanthaiah-grpc-sample/protoc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	connection, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client = proto.NewExampleClient(connection)
	//implementing REST API
	r := gin.Default()
	//Unary operation
	// r.GET("/sent-message-to-server/:message", ClientConnectionServer)

	// Client streaming
	r.GET("/sent", ClientConnectionServer)
	r.Run(":8080")
}

//Unary Operation
// func ClientConnectionServer(c *gin.Context) {
// 	msg := c.Param("message")

// 	req := &proto.HelloRequest{SomeString: msg}
// 	client.ServerReply(context.TODO(), req)
// 	c.JSON(http.StatusOK, gin.H{
// 		"message": "message sent succefully to server" + msg,
// 	})
// }

//Client Streaming
// func ClientConnectionServer(c *gin.Context) {
// 	req := []*proto.HelloRequest{
// 		{SomeString: "Request 1"},
// 		{SomeString: "Request 2"},
// 		{SomeString: "Request 3"},
// 		{SomeString: "Request 4"},
// 	}

// 	stream, err := client.ServerReply(context.TODO())
// 	if err != nil {
// 		fmt.Println("Some error occured")
// 		return
// 	}
// 	for _, re := range req {
// 		err = stream.Send(re)
// 		if err != nil {
// 			fmt.Println("Request has not been fullfilled")
// 			return
// 		}
// 	}

// 	response, err := stream.CloseAndRecv()
// 	if err != nil {
// 		fmt.Println("Some error occured", err)
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"Message count": response,
// 	})
// }

// Server Streaming
// func ClientConnectionServer(c *gin.Context) {
// 	stream, err := client.ServerReply(context.TODO(), &proto.HelloRequest{SomeString: "This is Client"})
// 	if err != nil {
// 		fmt.Println("An Error Occured")
// 		return
// 	}

// 	count := 0
// 	for {
// 		message, err := stream.Recv()
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Println("Server message", message)
// 		time.Sleep(1 * time.Second)
// 		count++
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"Message Count": count,
// 	})
// }

// Bi-directional streaming
func ClientConnectionServer(c *gin.Context) {
	stream, err := client.ServerReply(context.TODO())
	if err != nil {
		fmt.Println("Some error occured")
		return
	}

	send, recieve := 0, 0
	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.HelloRequest{SomeString: "message" + strconv.Itoa(i) + "from client"})
		if err != nil {
			fmt.Println("Unable to send data")
			return
		}
		send++
	}
	if err := stream.CloseSend(); err != nil {
		log.Println(err)
	}

	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server message:- ", message)
		recieve++
	}
	c.JSON(http.StatusOK, gin.H{
		"Messages sent":     send,
		"Messages ricieved": recieve,
	})
}
