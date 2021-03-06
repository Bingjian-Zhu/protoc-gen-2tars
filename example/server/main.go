package main

import (
	"github.com/TarsCloud/TarsGo/tars"
	"protoc-gen-2tars/example/pb"
)

type GreeterImp struct {
}

func (imp *GreeterImp) SayHello(input *pb.HelloRequest) (*pb.HelloReply, error) {
	output := new(pb.HelloReply)
	output.Message = "hello" + input.Name
	return output, nil
}

func main() { //Init servant

	imp := new(GreeterImp)                                        //New Imp
	app := new(pb.Greeter)                                        //New init the A JCE
	cfg := tars.GetServerConfig()                                 //Get Config File Object
	app.AddServant(imp, cfg.App+"."+cfg.Server+".GreeterTestObj") //Register Servant
	tars.Run()
}
