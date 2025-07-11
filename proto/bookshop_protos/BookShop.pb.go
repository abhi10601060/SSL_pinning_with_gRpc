// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: BookShop.proto

package bookshop_protos

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Void struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Void) Reset() {
	*x = Void{}
	mi := &file_BookShop_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Void) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Void) ProtoMessage() {}

func (x *Void) ProtoReflect() protoreflect.Message {
	mi := &file_BookShop_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Void.ProtoReflect.Descriptor instead.
func (*Void) Descriptor() ([]byte, []int) {
	return file_BookShop_proto_rawDescGZIP(), []int{0}
}

type Book struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Book) Reset() {
	*x = Book{}
	mi := &file_BookShop_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_BookShop_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Book.ProtoReflect.Descriptor instead.
func (*Book) Descriptor() ([]byte, []int) {
	return file_BookShop_proto_rawDescGZIP(), []int{1}
}

func (x *Book) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Book) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ListAllResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Books         []*Book                `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
	Count         int32                  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ListAllResponse) Reset() {
	*x = ListAllResponse{}
	mi := &file_BookShop_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ListAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAllResponse) ProtoMessage() {}

func (x *ListAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BookShop_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAllResponse.ProtoReflect.Descriptor instead.
func (*ListAllResponse) Descriptor() ([]byte, []int) {
	return file_BookShop_proto_rawDescGZIP(), []int{2}
}

func (x *ListAllResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

func (x *ListAllResponse) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type AddBookResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	IsSuccessful  bool                   `protobuf:"varint,1,opt,name=isSuccessful,proto3" json:"isSuccessful,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddBookResponse) Reset() {
	*x = AddBookResponse{}
	mi := &file_BookShop_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddBookResponse) ProtoMessage() {}

func (x *AddBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_BookShop_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddBookResponse.ProtoReflect.Descriptor instead.
func (*AddBookResponse) Descriptor() ([]byte, []int) {
	return file_BookShop_proto_rawDescGZIP(), []int{3}
}

func (x *AddBookResponse) GetIsSuccessful() bool {
	if x != nil {
		return x.IsSuccessful
	}
	return false
}

func (x *AddBookResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_BookShop_proto protoreflect.FileDescriptor

const file_BookShop_proto_rawDesc = "" +
	"\n" +
	"\x0eBookShop.proto\x12\bbookshop\"\x06\n" +
	"\x04Void\"*\n" +
	"\x04Book\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x05R\x02id\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\"M\n" +
	"\x0fListAllResponse\x12$\n" +
	"\x05books\x18\x01 \x03(\v2\x0e.bookshop.BookR\x05books\x12\x14\n" +
	"\x05count\x18\x02 \x01(\x05R\x05count\"O\n" +
	"\x0fAddBookResponse\x12\"\n" +
	"\fisSuccessful\x18\x01 \x01(\bR\fisSuccessful\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage2}\n" +
	"\x0fBookShopService\x124\n" +
	"\aListAll\x12\x0e.bookshop.Void\x1a\x19.bookshop.ListAllResponse\x124\n" +
	"\aAddBook\x12\x0e.bookshop.Book\x1a\x19.bookshop.AddBookResponseBCZAgithub.com/abhi1060/proto_example/bookshop_protos;bookshop_protosb\x06proto3"

var (
	file_BookShop_proto_rawDescOnce sync.Once
	file_BookShop_proto_rawDescData []byte
)

func file_BookShop_proto_rawDescGZIP() []byte {
	file_BookShop_proto_rawDescOnce.Do(func() {
		file_BookShop_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_BookShop_proto_rawDesc), len(file_BookShop_proto_rawDesc)))
	})
	return file_BookShop_proto_rawDescData
}

var file_BookShop_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_BookShop_proto_goTypes = []any{
	(*Void)(nil),            // 0: bookshop.Void
	(*Book)(nil),            // 1: bookshop.Book
	(*ListAllResponse)(nil), // 2: bookshop.ListAllResponse
	(*AddBookResponse)(nil), // 3: bookshop.AddBookResponse
}
var file_BookShop_proto_depIdxs = []int32{
	1, // 0: bookshop.ListAllResponse.books:type_name -> bookshop.Book
	0, // 1: bookshop.BookShopService.ListAll:input_type -> bookshop.Void
	1, // 2: bookshop.BookShopService.AddBook:input_type -> bookshop.Book
	2, // 3: bookshop.BookShopService.ListAll:output_type -> bookshop.ListAllResponse
	3, // 4: bookshop.BookShopService.AddBook:output_type -> bookshop.AddBookResponse
	3, // [3:5] is the sub-list for method output_type
	1, // [1:3] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_BookShop_proto_init() }
func file_BookShop_proto_init() {
	if File_BookShop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_BookShop_proto_rawDesc), len(file_BookShop_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_BookShop_proto_goTypes,
		DependencyIndexes: file_BookShop_proto_depIdxs,
		MessageInfos:      file_BookShop_proto_msgTypes,
	}.Build()
	File_BookShop_proto = out.File
	file_BookShop_proto_goTypes = nil
	file_BookShop_proto_depIdxs = nil
}
