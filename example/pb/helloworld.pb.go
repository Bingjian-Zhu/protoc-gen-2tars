// Code generated by protoc-gen-go. DO NOT EDIT.
// source: helloworld.proto

package pb

import (
	context "context"
	fmt "fmt"
	tars "github.com/TarsCloud/TarsGo/tars"
	model "github.com/TarsCloud/TarsGo/tars/model"
	requestf "github.com/TarsCloud/TarsGo/tars/protocol/res/requestf"
	current "github.com/TarsCloud/TarsGo/tars/util/current"
	tools "github.com/TarsCloud/TarsGo/tars/util/tools"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The request message containing the user's name.
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

// The response message containing the greetings
type HelloReply struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloReply) Reset()         { *m = HelloReply{} }
func (m *HelloReply) String() string { return proto.CompactTextString(m) }
func (*HelloReply) ProtoMessage()    {}
func (*HelloReply) Descriptor() ([]byte, []int) {
	return fileDescriptor_17b8c58d586b62f2, []int{1}
}

func (m *HelloReply) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloReply.Unmarshal(m, b)
}
func (m *HelloReply) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloReply.Marshal(b, m, deterministic)
}
func (m *HelloReply) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloReply.Merge(m, src)
}
func (m *HelloReply) XXX_Size() int {
	return xxx_messageInfo_HelloReply.Size(m)
}
func (m *HelloReply) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloReply.DiscardUnknown(m)
}

var xxx_messageInfo_HelloReply proto.InternalMessageInfo

func (m *HelloReply) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*HelloRequest)(nil), "pb.HelloRequest")
	proto.RegisterType((*HelloReply)(nil), "pb.HelloReply")
}

func init() { proto.RegisterFile("helloworld.proto", fileDescriptor_17b8c58d586b62f2) }

var fileDescriptor_17b8c58d586b62f2 = []byte{
	// 142 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xc8, 0x48, 0xcd, 0xc9,
	0xc9, 0x2f, 0xcf, 0x2f, 0xca, 0x49, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48,
	0x52, 0x52, 0xe2, 0xe2, 0xf1, 0x00, 0x89, 0x07, 0xa5, 0x16, 0x96, 0xa6, 0x16, 0x97, 0x08, 0x09,
	0x71, 0xb1, 0xe4, 0x25, 0xe6, 0xa6, 0x4a, 0x30, 0x2a, 0x30, 0x6a, 0x70, 0x06, 0x81, 0xd9, 0x4a,
	0x6a, 0x5c, 0x5c, 0x50, 0x35, 0x05, 0x39, 0x95, 0x42, 0x12, 0x5c, 0xec, 0xb9, 0xa9, 0xc5, 0xc5,
	0x89, 0xe9, 0x30, 0x45, 0x30, 0xae, 0x91, 0x25, 0x17, 0xbb, 0x7b, 0x51, 0x6a, 0x6a, 0x49, 0x6a,
	0x91, 0x90, 0x1e, 0x17, 0x47, 0x70, 0x62, 0x25, 0x58, 0x97, 0x90, 0x80, 0x5e, 0x41, 0x92, 0x1e,
	0xb2, 0x25, 0x52, 0x7c, 0x48, 0x22, 0x05, 0x39, 0x95, 0x4a, 0x0c, 0x49, 0x6c, 0x60, 0x17, 0x19,
	0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xf4, 0xb3, 0x8e, 0xc0, 0xa5, 0x00, 0x00, 0x00,
}

// This following code was generated by tarsrpc
// Gernerated from helloworld.proto
type Greeter struct {
	s model.Servant
}

//SetServant is required by the servant interface.
func (obj *Greeter) SetServant(s model.Servant) {
	obj.s = s
}

//AddServant is required by the servant interface
func (obj *Greeter) AddServant(imp impGreeter, objStr string) {
	tars.AddServant(obj, imp, objStr)
}

////AddServant adds servant  for the service with context
func (obj *Greeter) AddServantWithContext(imp impGreeterWithContext, objStr string) {
	tars.AddServantWithContext(obj, imp, objStr)
}

//TarsSetTimeout is required by the servant interface. t is the timeout in ms.
func (obj *Greeter) TarsSetTimeout(t int) {
	obj.s.TarsSetTimeout(t)
}

//TarsSetProtocol is required by the servant interface. t is the protocol.
func (obj *Greeter) TarsSetProtocol(p model.Protocol) {
	obj.s.TarsSetProtocol(p)
}

type impGreeter interface {
	SayHello(input *HelloRequest) (*HelloReply, error)
}

type impGreeterWithContext interface {
	SayHello(ctx context.Context, input *HelloRequest) (*HelloReply, error)
}

//Dispatch is used to call the user implement of the defined method.
func (obj *Greeter) Dispatch(ctx context.Context, val interface{}, req *requestf.RequestPacket, resp *requestf.ResponsePacket, withContext bool) error {
	input := tools.Int8ToByte(req.SBuffer)
	var (
		output []byte
		err    error
	)
	funcName := req.SFuncName
	switch funcName {

	case "SayHello":
		inputDefine := &HelloRequest{}
		if err = proto.Unmarshal(input, inputDefine); err != nil {
			return err
		}
		var res *HelloReply
		if withContext == false {
			imp := val.(impGreeter)
			res, err = imp.SayHello(inputDefine)
			if err != nil {
				return err
			}
		} else {
			imp := val.(impGreeterWithContext)
			res, err = imp.SayHello(ctx, inputDefine)
			if err != nil {
				return err
			}
		}
		output, err = proto.Marshal(res)
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("func mismatch")
	}
	var _status map[string]string
	s, ok := current.GetResponseStatus(ctx)
	if ok && s != nil {
		_status = s
	}
	var _context map[string]string
	c, ok := current.GetResponseContext(ctx)
	if ok && c != nil {
		_context = c
	}
	*resp = requestf.ResponsePacket{
		IVersion:     1,
		CPacketType:  0,
		IRequestId:   req.IRequestId,
		IMessageType: 0,
		IRet:         0,
		SBuffer:      tools.ByteToInt8(output),
		Status:       _status,
		SResultDesc:  "",
		Context:      _context,
	}
	return nil
}

// SayHello is client rpc method as defined
func (obj *Greeter) SayHello(input *HelloRequest, _opt ...map[string]string) (*HelloReply, error) {
	ctx := context.Background()
	return obj.SayHelloWithContext(ctx, input, _opt...)
}

// SayHelloWithContext is client rpc method as defined
func (obj *Greeter) SayHelloWithContext(ctx context.Context, input *HelloRequest, _opt ...map[string]string) (*HelloReply, error) {
	var inputMarshal []byte
	output := new(HelloReply)
	inputMarshal, err := proto.Marshal(input)
	if err != nil {
		return output, err
	}

	var _status map[string]string
	var _context map[string]string
	if len(_opt) == 1 {
		_context = _opt[0]
	} else if len(_opt) == 2 {
		_context = _opt[0]
		_status = _opt[1]
	}

	resp := new(requestf.ResponsePacket)

	err = obj.s.Tars_invoke(ctx, 0, "SayHello", inputMarshal, _status, _context, resp)
	if err != nil {
		return output, err
	}
	if err = proto.Unmarshal(tools.Int8ToByte(resp.SBuffer), output); err != nil {
		return output, err
	}

	if len(_opt) == 1 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range resp.Context {
			_context[k] = v
		}
	} else if len(_opt) == 2 {
		for k := range _context {
			delete(_context, k)
		}
		for k, v := range resp.Context {
			_context[k] = v
		}
		for k := range _status {
			delete(_status, k)
		}
		for k, v := range resp.Status {
			_status[k] = v
		}
	}

	return output, nil
}
