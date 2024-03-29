// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: book.proto

package book

import (
	context "context"
	fmt "fmt"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	io "io"
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
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// 请求书详情的参数结构  book_id 32位整形
type BookInfoParams struct {
	BookId int32 `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
}

func (m *BookInfoParams) Reset()         { *m = BookInfoParams{} }
func (m *BookInfoParams) String() string { return proto.CompactTextString(m) }
func (*BookInfoParams) ProtoMessage()    {}
func (*BookInfoParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{0}
}
func (m *BookInfoParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookInfoParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookInfoParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookInfoParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfoParams.Merge(m, src)
}
func (m *BookInfoParams) XXX_Size() int {
	return m.Size()
}
func (m *BookInfoParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfoParams.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfoParams proto.InternalMessageInfo

func (m *BookInfoParams) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

// 书详情信息的结构   book_name字符串类型
type BookInfo struct {
	BookId   int32  `protobuf:"varint,1,opt,name=book_id,json=bookId,proto3" json:"book_id,omitempty"`
	BookName string `protobuf:"bytes,2,opt,name=book_name,json=bookName,proto3" json:"book_name,omitempty"`
}

func (m *BookInfo) Reset()         { *m = BookInfo{} }
func (m *BookInfo) String() string { return proto.CompactTextString(m) }
func (*BookInfo) ProtoMessage()    {}
func (*BookInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{1}
}
func (m *BookInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookInfo.Merge(m, src)
}
func (m *BookInfo) XXX_Size() int {
	return m.Size()
}
func (m *BookInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BookInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BookInfo proto.InternalMessageInfo

func (m *BookInfo) GetBookId() int32 {
	if m != nil {
		return m.BookId
	}
	return 0
}

func (m *BookInfo) GetBookName() string {
	if m != nil {
		return m.BookName
	}
	return ""
}

// 请求书列表的参数结构  page、limit   32位整形
type BookListParams struct {
	Page  int32 `protobuf:"varint,1,opt,name=page,proto3" json:"page,omitempty"`
	Limit int32 `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
}

func (m *BookListParams) Reset()         { *m = BookListParams{} }
func (m *BookListParams) String() string { return proto.CompactTextString(m) }
func (*BookListParams) ProtoMessage()    {}
func (*BookListParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{2}
}
func (m *BookListParams) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookListParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookListParams.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookListParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookListParams.Merge(m, src)
}
func (m *BookListParams) XXX_Size() int {
	return m.Size()
}
func (m *BookListParams) XXX_DiscardUnknown() {
	xxx_messageInfo_BookListParams.DiscardUnknown(m)
}

var xxx_messageInfo_BookListParams proto.InternalMessageInfo

func (m *BookListParams) GetPage() int32 {
	if m != nil {
		return m.Page
	}
	return 0
}

func (m *BookListParams) GetLimit() int32 {
	if m != nil {
		return m.Limit
	}
	return 0
}

// 书列表的结构    BookInfo结构数组
type BookList struct {
	BookList []*BookInfo `protobuf:"bytes,1,rep,name=book_list,json=bookList,proto3" json:"book_list,omitempty"`
}

func (m *BookList) Reset()         { *m = BookList{} }
func (m *BookList) String() string { return proto.CompactTextString(m) }
func (*BookList) ProtoMessage()    {}
func (*BookList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1e89d0eaa98dc5d8, []int{3}
}
func (m *BookList) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BookList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BookList.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BookList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BookList.Merge(m, src)
}
func (m *BookList) XXX_Size() int {
	return m.Size()
}
func (m *BookList) XXX_DiscardUnknown() {
	xxx_messageInfo_BookList.DiscardUnknown(m)
}

var xxx_messageInfo_BookList proto.InternalMessageInfo

func (m *BookList) GetBookList() []*BookInfo {
	if m != nil {
		return m.BookList
	}
	return nil
}

func init() {
	proto.RegisterType((*BookInfoParams)(nil), "book.BookInfoParams")
	proto.RegisterType((*BookInfo)(nil), "book.BookInfo")
	proto.RegisterType((*BookListParams)(nil), "book.BookListParams")
	proto.RegisterType((*BookList)(nil), "book.BookList")
}

func init() { proto.RegisterFile("book.proto", fileDescriptor_1e89d0eaa98dc5d8) }

var fileDescriptor_1e89d0eaa98dc5d8 = []byte{
	// 251 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4a, 0xca, 0xcf, 0xcf,
	0xd6, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x01, 0xb1, 0x95, 0x34, 0xb9, 0xf8, 0x9c, 0xf2,
	0xf3, 0xb3, 0x3d, 0xf3, 0xd2, 0xf2, 0x03, 0x12, 0x8b, 0x12, 0x73, 0x8b, 0x85, 0xc4, 0xb9, 0xd8,
	0x41, 0x32, 0xf1, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41, 0x6c, 0x20, 0xae, 0x67,
	0x8a, 0x92, 0x03, 0x17, 0x07, 0x4c, 0x29, 0x4e, 0x45, 0x42, 0xd2, 0x5c, 0x9c, 0x60, 0x89, 0xbc,
	0xc4, 0xdc, 0x54, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x0e, 0x90, 0x80, 0x5f, 0x62, 0x6e,
	0xaa, 0x92, 0x15, 0xc4, 0x32, 0x9f, 0xcc, 0xe2, 0x12, 0xa8, 0x65, 0x42, 0x5c, 0x2c, 0x05, 0x89,
	0xe9, 0xa9, 0x50, 0x43, 0xc0, 0x6c, 0x21, 0x11, 0x2e, 0xd6, 0x9c, 0xcc, 0xdc, 0xcc, 0x12, 0xb0,
	0x76, 0xd6, 0x20, 0x08, 0x47, 0xc9, 0x1c, 0x62, 0x3b, 0x48, 0xaf, 0x90, 0x36, 0xd4, 0x92, 0x9c,
	0xcc, 0xe2, 0x12, 0x09, 0x46, 0x05, 0x66, 0x0d, 0x6e, 0x23, 0x3e, 0x3d, 0xb0, 0xd7, 0x60, 0x0e,
	0x84, 0x58, 0x0a, 0x52, 0x6c, 0x54, 0xc2, 0xc5, 0x02, 0x12, 0x15, 0x32, 0xe5, 0xe2, 0x76, 0x4f,
	0x2d, 0x81, 0xfb, 0x40, 0x04, 0x55, 0x03, 0xc4, 0x3d, 0x52, 0x68, 0xc6, 0x28, 0x31, 0x20, 0x69,
	0x03, 0x5b, 0x8d, 0xa4, 0x0d, 0xe1, 0x0d, 0x64, 0x6d, 0x20, 0x51, 0x25, 0x06, 0x27, 0x89, 0x13,
	0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48, 0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86,
	0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x48, 0x62, 0x03, 0x07, 0xbf, 0x31, 0x20, 0x00,
	0x00, 0xff, 0xff, 0xd1, 0xec, 0xae, 0xac, 0x8c, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BookClient is the client API for Book service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BookClient interface {
	GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error)
	GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error)
}

type bookClient struct {
	cc *grpc.ClientConn
}

func NewBookClient(cc *grpc.ClientConn) BookClient {
	return &bookClient{cc}
}

func (c *bookClient) GetBookInfo(ctx context.Context, in *BookInfoParams, opts ...grpc.CallOption) (*BookInfo, error) {
	out := new(BookInfo)
	err := c.cc.Invoke(ctx, "/book.Book/GetBookInfo", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *bookClient) GetBookList(ctx context.Context, in *BookListParams, opts ...grpc.CallOption) (*BookList, error) {
	out := new(BookList)
	err := c.cc.Invoke(ctx, "/book.Book/GetBookList", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookServer is the server API for Book service.
type BookServer interface {
	GetBookInfo(context.Context, *BookInfoParams) (*BookInfo, error)
	GetBookList(context.Context, *BookListParams) (*BookList, error)
}

func RegisterBookServer(s *grpc.Server, srv BookServer) {
	s.RegisterService(&_Book_serviceDesc, srv)
}

func _Book_GetBookInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookInfoParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.Book/GetBookInfo",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookInfo(ctx, req.(*BookInfoParams))
	}
	return interceptor(ctx, in, info, handler)
}

func _Book_GetBookList_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BookListParams)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookServer).GetBookList(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.Book/GetBookList",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookServer).GetBookList(ctx, req.(*BookListParams))
	}
	return interceptor(ctx, in, info, handler)
}

var _Book_serviceDesc = grpc.ServiceDesc{
	ServiceName: "book.Book",
	HandlerType: (*BookServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetBookInfo",
			Handler:    _Book_GetBookInfo_Handler,
		},
		{
			MethodName: "GetBookList",
			Handler:    _Book_GetBookList_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "book.proto",
}

func (m *BookInfoParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookInfoParams) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BookId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBook(dAtA, i, uint64(m.BookId))
	}
	return i, nil
}

func (m *BookInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookInfo) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.BookId != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBook(dAtA, i, uint64(m.BookId))
	}
	if len(m.BookName) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintBook(dAtA, i, uint64(len(m.BookName)))
		i += copy(dAtA[i:], m.BookName)
	}
	return i, nil
}

func (m *BookListParams) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookListParams) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Page != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintBook(dAtA, i, uint64(m.Page))
	}
	if m.Limit != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintBook(dAtA, i, uint64(m.Limit))
	}
	return i, nil
}

func (m *BookList) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BookList) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.BookList) > 0 {
		for _, msg := range m.BookList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintBook(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintBook(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *BookInfoParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BookId != 0 {
		n += 1 + sovBook(uint64(m.BookId))
	}
	return n
}

func (m *BookInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.BookId != 0 {
		n += 1 + sovBook(uint64(m.BookId))
	}
	l = len(m.BookName)
	if l > 0 {
		n += 1 + l + sovBook(uint64(l))
	}
	return n
}

func (m *BookListParams) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Page != 0 {
		n += 1 + sovBook(uint64(m.Page))
	}
	if m.Limit != 0 {
		n += 1 + sovBook(uint64(m.Limit))
	}
	return n
}

func (m *BookList) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.BookList) > 0 {
		for _, e := range m.BookList {
			l = e.Size()
			n += 1 + l + sovBook(uint64(l))
		}
	}
	return n
}

func sovBook(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozBook(x uint64) (n int) {
	return sovBook(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BookInfoParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BookInfoParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookInfoParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookId", wireType)
			}
			m.BookId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BookId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BookInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BookInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookId", wireType)
			}
			m.BookId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.BookId |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthBook
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BookListParams) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BookListParams: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookListParams: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Page", wireType)
			}
			m.Page = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Page |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Limit", wireType)
			}
			m.Limit = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Limit |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *BookList) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBook
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BookList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BookList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field BookList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBook
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthBook
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthBook
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.BookList = append(m.BookList, &BookInfo{})
			if err := m.BookList[len(m.BookList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBook(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthBook
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipBook(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBook
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
					return 0, ErrIntOverflowBook
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
					return 0, ErrIntOverflowBook
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
			if length < 0 {
				return 0, ErrInvalidLengthBook
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthBook
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowBook
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
				next, err := skipBook(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthBook
				}
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
	ErrInvalidLengthBook = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBook   = fmt.Errorf("proto: integer overflow")
)


