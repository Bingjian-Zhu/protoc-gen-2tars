package main

import (
	"fmt"
	"github.com/TarsCloud/TarsGo/tars"
	"protoc-gen-2tars/example/pb"
)

func main() {
	comm := tars.NewCommunicator()
	obj := fmt.Sprintf("StressTest.HelloPbServer.GreeterTestObj@tcp -h 127.0.0.1  -p 10015  -t 60000")
	app := new(pb.Greeter)
	comm.StringToProxy(obj, app)
	input := &pb.HelloRequest{Name: "sandyskies"}
	output, err := app.SayHello(input)
	if err != nil {
		fmt.Println("err: ", err)
	}
	fmt.Println("result is:", output.Message)
}
