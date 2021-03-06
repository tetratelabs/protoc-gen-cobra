// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: timer.proto

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type TickRequest struct {
	Interval             int32    `protobuf:"varint,1,opt,name=interval,proto3" json:"interval,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickRequest) Reset()         { *m = TickRequest{} }
func (m *TickRequest) String() string { return proto.CompactTextString(m) }
func (*TickRequest) ProtoMessage()    {}
func (*TickRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_timer_468c66ad9c568581, []int{0}
}
func (m *TickRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TickRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickRequest.Merge(dst, src)
}
func (m *TickRequest) XXX_Size() int {
	return m.Size()
}
func (m *TickRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TickRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TickRequest proto.InternalMessageInfo

func (m *TickRequest) GetInterval() int32 {
	if m != nil {
		return m.Interval
	}
	return 0
}

type TickResponse struct {
	Time                 string   `protobuf:"bytes,1,opt,name=time,proto3" json:"time,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TickResponse) Reset()         { *m = TickResponse{} }
func (m *TickResponse) String() string { return proto.CompactTextString(m) }
func (*TickResponse) ProtoMessage()    {}
func (*TickResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_timer_468c66ad9c568581, []int{1}
}
func (m *TickResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TickResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TickResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *TickResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TickResponse.Merge(dst, src)
}
func (m *TickResponse) XXX_Size() int {
	return m.Size()
}
func (m *TickResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TickResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TickResponse proto.InternalMessageInfo

func (m *TickResponse) GetTime() string {
	if m != nil {
		return m.Time
	}
	return ""
}

func init() {
	proto.RegisterType((*TickRequest)(nil), "pb.TickRequest")
	proto.RegisterType((*TickResponse)(nil), "pb.TickResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TimerClient is the client API for Timer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TimerClient interface {
	Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (Timer_TickClient, error)
}

type timerClient struct {
	cc *grpc.ClientConn
}

func NewTimerClient(cc *grpc.ClientConn) TimerClient {
	return &timerClient{cc}
}

func (c *timerClient) Tick(ctx context.Context, in *TickRequest, opts ...grpc.CallOption) (Timer_TickClient, error) {
	stream, err := c.cc.NewStream(ctx, &_Timer_serviceDesc.Streams[0], "/pb.Timer/Tick", opts...)
	if err != nil {
		return nil, err
	}
	x := &timerTickClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Timer_TickClient interface {
	Recv() (*TickResponse, error)
	grpc.ClientStream
}

type timerTickClient struct {
	grpc.ClientStream
}

func (x *timerTickClient) Recv() (*TickResponse, error) {
	m := new(TickResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TimerServer is the server API for Timer service.
type TimerServer interface {
	Tick(*TickRequest, Timer_TickServer) error
}

func RegisterTimerServer(s *grpc.Server, srv TimerServer) {
	s.RegisterService(&_Timer_serviceDesc, srv)
}

func _Timer_Tick_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TickRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(TimerServer).Tick(m, &timerTickServer{stream})
}

type Timer_TickServer interface {
	Send(*TickResponse) error
	grpc.ServerStream
}

type timerTickServer struct {
	grpc.ServerStream
}

func (x *timerTickServer) Send(m *TickResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _Timer_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Timer",
	HandlerType: (*TimerServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Tick",
			Handler:       _Timer_Tick_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "timer.proto",
}

func (m *TickRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Interval != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintTimer(dAtA, i, uint64(m.Interval))
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *TickResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TickResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Time) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintTimer(dAtA, i, uint64(len(m.Time)))
		i += copy(dAtA[i:], m.Time)
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintTimer(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TickRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Interval != 0 {
		n += 1 + sovTimer(uint64(m.Interval))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *TickResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Time)
	if l > 0 {
		n += 1 + l + sovTimer(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovTimer(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozTimer(x uint64) (n int) {
	return sovTimer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TickRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Interval", wireType)
			}
			m.Interval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Interval |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *TickResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: TickResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TickResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Time", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTimer
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Time = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTimer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTimer
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipTimer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTimer
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTimer
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthTimer
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowTimer
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipTimer(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthTimer = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTimer   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("timer.proto", fileDescriptor_timer_468c66ad9c568581) }

var fileDescriptor_timer_468c66ad9c568581 = []byte{
	// 154 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0xc9, 0xcc, 0x4d,
	0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0xd2, 0xe4, 0xe2, 0x0e,
	0xc9, 0x4c, 0xce, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x92, 0xe2, 0xe2, 0xc8, 0xcc,
	0x2b, 0x49, 0x2d, 0x2a, 0x4b, 0xcc, 0x91, 0x60, 0x54, 0x60, 0xd4, 0x60, 0x0d, 0x82, 0xf3, 0x95,
	0x94, 0xb8, 0x78, 0x20, 0x4a, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0x85, 0x84, 0xb8, 0x58, 0x40,
	0xa6, 0x81, 0xd5, 0x71, 0x06, 0x81, 0xd9, 0x46, 0x26, 0x5c, 0xac, 0x21, 0x20, 0x1b, 0x84, 0xb4,
	0xb9, 0x58, 0x40, 0x8a, 0x85, 0xf8, 0xf5, 0x0a, 0x92, 0xf4, 0x90, 0x6c, 0x90, 0x12, 0x40, 0x08,
	0x40, 0xcc, 0x31, 0x60, 0x74, 0x12, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07,
	0x8f, 0xe4, 0x18, 0x67, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0xbb, 0xd0, 0x18, 0x10, 0x00, 0x00,
	0xff, 0xff, 0xa0, 0xa1, 0x58, 0x36, 0xb0, 0x00, 0x00, 0x00,
}
