// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.1
// 	protoc        v5.27.0
// source: bookstore.proto

package gen

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Author struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	DateOfBirth int64  `protobuf:"varint,3,opt,name=date_of_birth,json=dateOfBirth,proto3" json:"date_of_birth,omitempty"`
}

func (x *Author) Reset() {
	*x = Author{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Author) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Author) ProtoMessage() {}

func (x *Author) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Author.ProtoReflect.Descriptor instead.
func (*Author) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{0}
}

func (x *Author) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Author) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Author) GetDateOfBirth() int64 {
	if x != nil {
		return x.DateOfBirth
	}
	return 0
}

type GetAuthorsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAuthorsRequest) Reset() {
	*x = GetAuthorsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsRequest) ProtoMessage() {}

func (x *GetAuthorsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorsRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{1}
}

type GetAuthorsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Authors []*Author `protobuf:"bytes,1,rep,name=authors,proto3" json:"authors,omitempty"`
}

func (x *GetAuthorsResponse) Reset() {
	*x = GetAuthorsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorsResponse) ProtoMessage() {}

func (x *GetAuthorsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorsResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorsResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{2}
}

func (x *GetAuthorsResponse) GetAuthors() []*Author {
	if x != nil {
		return x.Authors
	}
	return nil
}

type GetAuthorRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetAuthorRequest) Reset() {
	*x = GetAuthorRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorRequest) ProtoMessage() {}

func (x *GetAuthorRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorRequest.ProtoReflect.Descriptor instead.
func (*GetAuthorRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{3}
}

func (x *GetAuthorRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetAuthorResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Author *Author `protobuf:"bytes,1,opt,name=author,proto3" json:"author,omitempty"`
}

func (x *GetAuthorResponse) Reset() {
	*x = GetAuthorResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAuthorResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAuthorResponse) ProtoMessage() {}

func (x *GetAuthorResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAuthorResponse.ProtoReflect.Descriptor instead.
func (*GetAuthorResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{4}
}

func (x *GetAuthorResponse) GetAuthor() *Author {
	if x != nil {
		return x.Author
	}
	return nil
}

type Book struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Title    string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	AuthorId string `protobuf:"bytes,3,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Synopsis string `protobuf:"bytes,4,opt,name=synopsis,proto3" json:"synopsis,omitempty"`
	ImageUrl string `protobuf:"bytes,5,opt,name=image_url,json=imageUrl,proto3" json:"image_url,omitempty"`
	Genre    string `protobuf:"bytes,6,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *Book) Reset() {
	*x = Book{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Book) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Book) ProtoMessage() {}

func (x *Book) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
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
	return file_bookstore_proto_rawDescGZIP(), []int{5}
}

func (x *Book) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Book) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Book) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *Book) GetSynopsis() string {
	if x != nil {
		return x.Synopsis
	}
	return ""
}

func (x *Book) GetImageUrl() string {
	if x != nil {
		return x.ImageUrl
	}
	return ""
}

func (x *Book) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type GetBooksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title    string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	AuthorId string `protobuf:"bytes,2,opt,name=author_id,json=authorId,proto3" json:"author_id,omitempty"`
	Genre    string `protobuf:"bytes,3,opt,name=genre,proto3" json:"genre,omitempty"`
}

func (x *GetBooksRequest) Reset() {
	*x = GetBooksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksRequest) ProtoMessage() {}

func (x *GetBooksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksRequest.ProtoReflect.Descriptor instead.
func (*GetBooksRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{6}
}

func (x *GetBooksRequest) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *GetBooksRequest) GetAuthorId() string {
	if x != nil {
		return x.AuthorId
	}
	return ""
}

func (x *GetBooksRequest) GetGenre() string {
	if x != nil {
		return x.Genre
	}
	return ""
}

type GetBooksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Books []*Book `protobuf:"bytes,1,rep,name=books,proto3" json:"books,omitempty"`
}

func (x *GetBooksResponse) Reset() {
	*x = GetBooksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBooksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBooksResponse) ProtoMessage() {}

func (x *GetBooksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBooksResponse.ProtoReflect.Descriptor instead.
func (*GetBooksResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{7}
}

func (x *GetBooksResponse) GetBooks() []*Book {
	if x != nil {
		return x.Books
	}
	return nil
}

type GetBookRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetBookRequest) Reset() {
	*x = GetBookRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookRequest) ProtoMessage() {}

func (x *GetBookRequest) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookRequest.ProtoReflect.Descriptor instead.
func (*GetBookRequest) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{8}
}

func (x *GetBookRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetBookResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Book *Book `protobuf:"bytes,1,opt,name=book,proto3" json:"book,omitempty"`
}

func (x *GetBookResponse) Reset() {
	*x = GetBookResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_bookstore_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetBookResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetBookResponse) ProtoMessage() {}

func (x *GetBookResponse) ProtoReflect() protoreflect.Message {
	mi := &file_bookstore_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetBookResponse.ProtoReflect.Descriptor instead.
func (*GetBookResponse) Descriptor() ([]byte, []int) {
	return file_bookstore_proto_rawDescGZIP(), []int{9}
}

func (x *GetBookResponse) GetBook() *Book {
	if x != nil {
		return x.Book
	}
	return nil
}

var File_bookstore_proto protoreflect.FileDescriptor

var file_bookstore_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x50, 0x0a, 0x06, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x22, 0x0a, 0x0d, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x6f, 0x66, 0x5f, 0x62, 0x69, 0x72, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x66, 0x42, 0x69,
	0x72, 0x74, 0x68, 0x22, 0x13, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x07, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x07, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72,
	0x73, 0x22, 0x22, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x34, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x06, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x41, 0x75, 0x74,
	0x68, 0x6f, 0x72, 0x52, 0x06, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x22, 0x98, 0x01, 0x0a, 0x04,
	0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x61, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x61,
	0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70,
	0x73, 0x69, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x79, 0x6e, 0x6f, 0x70,
	0x73, 0x69, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x5f, 0x75, 0x72, 0x6c,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x55, 0x72, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x67, 0x65, 0x6e, 0x72, 0x65, 0x22, 0x5a, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f,
	0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1b, 0x0a, 0x09, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05,
	0x67, 0x65, 0x6e, 0x72, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x67, 0x65, 0x6e,
	0x72, 0x65, 0x22, 0x2f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b, 0x0a, 0x05, 0x62, 0x6f, 0x6f, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x05, 0x62, 0x6f,
	0x6f, 0x6b, 0x73, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x2c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x04, 0x62, 0x6f, 0x6f, 0x6b,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x05, 0x2e, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x04, 0x62,
	0x6f, 0x6f, 0x6b, 0x32, 0x7a, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x35, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f,
	0x72, 0x73, 0x12, 0x12, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75, 0x74, 0x68,
	0x6f, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x32, 0x0a, 0x09, 0x47,
	0x65, 0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x12, 0x11, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x75,
	0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x47, 0x65,
	0x74, 0x41, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32,
	0x6c, 0x0a, 0x0b, 0x42, 0x6f, 0x6f, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x2f,
	0x0a, 0x08, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x12, 0x10, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x47,
	0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2c, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x12, 0x0f, 0x2e, 0x47, 0x65, 0x74,
	0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x10, 0x2e, 0x47, 0x65,
	0x74, 0x42, 0x6f, 0x6f, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a,
	0x04, 0x2f, 0x67, 0x65, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_bookstore_proto_rawDescOnce sync.Once
	file_bookstore_proto_rawDescData = file_bookstore_proto_rawDesc
)

func file_bookstore_proto_rawDescGZIP() []byte {
	file_bookstore_proto_rawDescOnce.Do(func() {
		file_bookstore_proto_rawDescData = protoimpl.X.CompressGZIP(file_bookstore_proto_rawDescData)
	})
	return file_bookstore_proto_rawDescData
}

var file_bookstore_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_bookstore_proto_goTypes = []interface{}{
	(*Author)(nil),             // 0: Author
	(*GetAuthorsRequest)(nil),  // 1: GetAuthorsRequest
	(*GetAuthorsResponse)(nil), // 2: GetAuthorsResponse
	(*GetAuthorRequest)(nil),   // 3: GetAuthorRequest
	(*GetAuthorResponse)(nil),  // 4: GetAuthorResponse
	(*Book)(nil),               // 5: Book
	(*GetBooksRequest)(nil),    // 6: GetBooksRequest
	(*GetBooksResponse)(nil),   // 7: GetBooksResponse
	(*GetBookRequest)(nil),     // 8: GetBookRequest
	(*GetBookResponse)(nil),    // 9: GetBookResponse
}
var file_bookstore_proto_depIdxs = []int32{
	0, // 0: GetAuthorsResponse.authors:type_name -> Author
	0, // 1: GetAuthorResponse.author:type_name -> Author
	5, // 2: GetBooksResponse.books:type_name -> Book
	5, // 3: GetBookResponse.book:type_name -> Book
	1, // 4: AuthorService.GetAuthors:input_type -> GetAuthorsRequest
	3, // 5: AuthorService.GetAuthor:input_type -> GetAuthorRequest
	6, // 6: BookService.GetBooks:input_type -> GetBooksRequest
	8, // 7: BookService.GetBook:input_type -> GetBookRequest
	2, // 8: AuthorService.GetAuthors:output_type -> GetAuthorsResponse
	4, // 9: AuthorService.GetAuthor:output_type -> GetAuthorResponse
	7, // 10: BookService.GetBooks:output_type -> GetBooksResponse
	9, // 11: BookService.GetBook:output_type -> GetBookResponse
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_bookstore_proto_init() }
func file_bookstore_proto_init() {
	if File_bookstore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_bookstore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Author); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAuthorResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Book); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBooksResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_bookstore_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetBookResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_bookstore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_bookstore_proto_goTypes,
		DependencyIndexes: file_bookstore_proto_depIdxs,
		MessageInfos:      file_bookstore_proto_msgTypes,
	}.Build()
	File_bookstore_proto = out.File
	file_bookstore_proto_rawDesc = nil
	file_bookstore_proto_goTypes = nil
	file_bookstore_proto_depIdxs = nil
}
