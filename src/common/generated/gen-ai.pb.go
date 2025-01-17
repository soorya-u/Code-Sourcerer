// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v5.29.1
// source: gen-ai.proto

package generated

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

type SourceFileDependencyPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *SourceFileDependencyPayload) Reset() {
	*x = SourceFileDependencyPayload{}
	mi := &file_gen_ai_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceFileDependencyPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceFileDependencyPayload) ProtoMessage() {}

func (x *SourceFileDependencyPayload) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceFileDependencyPayload.ProtoReflect.Descriptor instead.
func (*SourceFileDependencyPayload) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{0}
}

func (x *SourceFileDependencyPayload) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SourceFileDependencyPayload) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type SourceFilePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path         string                         `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content      string                         `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	Dependencies []*SourceFileDependencyPayload `protobuf:"bytes,3,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
}

func (x *SourceFilePayload) Reset() {
	*x = SourceFilePayload{}
	mi := &file_gen_ai_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SourceFilePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SourceFilePayload) ProtoMessage() {}

func (x *SourceFilePayload) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SourceFilePayload.ProtoReflect.Descriptor instead.
func (*SourceFilePayload) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{1}
}

func (x *SourceFilePayload) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *SourceFilePayload) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *SourceFilePayload) GetDependencies() []*SourceFileDependencyPayload {
	if x != nil {
		return x.Dependencies
	}
	return nil
}

type BasicConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestDirectory    string `protobuf:"bytes,1,opt,name=test_directory,json=testDirectory,proto3" json:"test_directory,omitempty"`
	Comments         bool   `protobuf:"varint,2,opt,name=comments,proto3" json:"comments,omitempty"`
	TestingFramework string `protobuf:"bytes,4,opt,name=testing_framework,json=testingFramework,proto3" json:"testing_framework,omitempty"`
	WaterMark        bool   `protobuf:"varint,5,opt,name=water_mark,json=waterMark,proto3" json:"water_mark,omitempty"`
}

func (x *BasicConfig) Reset() {
	*x = BasicConfig{}
	mi := &file_gen_ai_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *BasicConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BasicConfig) ProtoMessage() {}

func (x *BasicConfig) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BasicConfig.ProtoReflect.Descriptor instead.
func (*BasicConfig) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{2}
}

func (x *BasicConfig) GetTestDirectory() string {
	if x != nil {
		return x.TestDirectory
	}
	return ""
}

func (x *BasicConfig) GetComments() bool {
	if x != nil {
		return x.Comments
	}
	return false
}

func (x *BasicConfig) GetTestingFramework() string {
	if x != nil {
		return x.TestingFramework
	}
	return ""
}

func (x *BasicConfig) GetWaterMark() bool {
	if x != nil {
		return x.WaterMark
	}
	return false
}

type Configuration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Configuration *BasicConfig      `protobuf:"bytes,1,opt,name=configuration,proto3" json:"configuration,omitempty"`
	Extras        map[string]string `protobuf:"bytes,2,rep,name=extras,proto3" json:"extras,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Configuration) Reset() {
	*x = Configuration{}
	mi := &file_gen_ai_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Configuration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Configuration) ProtoMessage() {}

func (x *Configuration) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Configuration.ProtoReflect.Descriptor instead.
func (*Configuration) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{3}
}

func (x *Configuration) GetConfiguration() *BasicConfig {
	if x != nil {
		return x.Configuration
	}
	return nil
}

func (x *Configuration) GetExtras() map[string]string {
	if x != nil {
		return x.Extras
	}
	return nil
}

type GithubContextRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MergeId string               `protobuf:"bytes,1,opt,name=merge_id,json=mergeId,proto3" json:"merge_id,omitempty"`
	Context string               `protobuf:"bytes,2,opt,name=context,proto3" json:"context,omitempty"`
	Config  *Configuration       `protobuf:"bytes,3,opt,name=config,proto3" json:"config,omitempty"`
	Files   []*SourceFilePayload `protobuf:"bytes,4,rep,name=files,proto3" json:"files,omitempty"`
}

func (x *GithubContextRequest) Reset() {
	*x = GithubContextRequest{}
	mi := &file_gen_ai_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GithubContextRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GithubContextRequest) ProtoMessage() {}

func (x *GithubContextRequest) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GithubContextRequest.ProtoReflect.Descriptor instead.
func (*GithubContextRequest) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{4}
}

func (x *GithubContextRequest) GetMergeId() string {
	if x != nil {
		return x.MergeId
	}
	return ""
}

func (x *GithubContextRequest) GetContext() string {
	if x != nil {
		return x.Context
	}
	return ""
}

func (x *GithubContextRequest) GetConfig() *Configuration {
	if x != nil {
		return x.Config
	}
	return nil
}

func (x *GithubContextRequest) GetFiles() []*SourceFilePayload {
	if x != nil {
		return x.Files
	}
	return nil
}

type TestFilePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Testname     string `protobuf:"bytes,1,opt,name=testname,proto3" json:"testname,omitempty"`
	Testfilepath string `protobuf:"bytes,2,opt,name=testfilepath,proto3" json:"testfilepath,omitempty"`
	Parentpath   string `protobuf:"bytes,3,opt,name=parentpath,proto3" json:"parentpath,omitempty"`
	Code         string `protobuf:"bytes,4,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *TestFilePayload) Reset() {
	*x = TestFilePayload{}
	mi := &file_gen_ai_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TestFilePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestFilePayload) ProtoMessage() {}

func (x *TestFilePayload) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestFilePayload.ProtoReflect.Descriptor instead.
func (*TestFilePayload) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{5}
}

func (x *TestFilePayload) GetTestname() string {
	if x != nil {
		return x.Testname
	}
	return ""
}

func (x *TestFilePayload) GetTestfilepath() string {
	if x != nil {
		return x.Testfilepath
	}
	return ""
}

func (x *TestFilePayload) GetParentpath() string {
	if x != nil {
		return x.Parentpath
	}
	return ""
}

func (x *TestFilePayload) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

type GeneratedTestsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tests []*TestFilePayload `protobuf:"bytes,1,rep,name=tests,proto3" json:"tests,omitempty"`
}

func (x *GeneratedTestsResponse) Reset() {
	*x = GeneratedTestsResponse{}
	mi := &file_gen_ai_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GeneratedTestsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GeneratedTestsResponse) ProtoMessage() {}

func (x *GeneratedTestsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_gen_ai_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GeneratedTestsResponse.ProtoReflect.Descriptor instead.
func (*GeneratedTestsResponse) Descriptor() ([]byte, []int) {
	return file_gen_ai_proto_rawDescGZIP(), []int{6}
}

func (x *GeneratedTestsResponse) GetTests() []*TestFilePayload {
	if x != nil {
		return x.Tests
	}
	return nil
}

var File_gen_ai_proto protoreflect.FileDescriptor

var file_gen_ai_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x67, 0x65, 0x6e, 0x2d, 0x61, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13,
	0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x67, 0x65,
	0x6e, 0x61, 0x69, 0x22, 0x4b, 0x0a, 0x1b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c,
	0x65, 0x44, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x22, 0x97, 0x01, 0x0a, 0x11, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x12, 0x54, 0x0a, 0x0c, 0x64, 0x65, 0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e,
	0x63, 0x69, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x63, 0x6f, 0x64,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x61, 0x69,
	0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65, 0x44, 0x65, 0x70, 0x65, 0x6e,
	0x64, 0x65, 0x6e, 0x63, 0x79, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x0c, 0x64, 0x65,
	0x70, 0x65, 0x6e, 0x64, 0x65, 0x6e, 0x63, 0x69, 0x65, 0x73, 0x22, 0x9c, 0x01, 0x0a, 0x0b, 0x42,
	0x61, 0x73, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x25, 0x0a, 0x0e, 0x74, 0x65,
	0x73, 0x74, 0x5f, 0x64, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x74, 0x65, 0x73, 0x74, 0x44, 0x69, 0x72, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x08, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12, 0x2b, 0x0a,
	0x11, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x5f, 0x66, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f,
	0x72, 0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x74, 0x65, 0x73, 0x74, 0x69, 0x6e,
	0x67, 0x46, 0x72, 0x61, 0x6d, 0x65, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x61,
	0x74, 0x65, 0x72, 0x5f, 0x6d, 0x61, 0x72, 0x6b, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x09,
	0x77, 0x61, 0x74, 0x65, 0x72, 0x4d, 0x61, 0x72, 0x6b, 0x22, 0xda, 0x01, 0x0a, 0x0d, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x0d, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72,
	0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x42, 0x61, 0x73, 0x69, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x46, 0x0a, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x72, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x45, 0x78, 0x74, 0x72, 0x61, 0x73, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x06, 0x65, 0x78, 0x74, 0x72, 0x61, 0x73, 0x1a, 0x39, 0x0a, 0x0b, 0x45,
	0x78, 0x74, 0x72, 0x61, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0xc5, 0x01, 0x0a, 0x14, 0x47, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x65, 0x72, 0x67, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x78, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x72, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x3c, 0x0a, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x26, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65, 0x72, 0x2e,
	0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x46, 0x69, 0x6c, 0x65,
	0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x22, 0x85,
	0x01, 0x0a, 0x0f, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x74, 0x65, 0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x65, 0x73, 0x74, 0x66, 0x69, 0x6c, 0x65, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x70, 0x61, 0x74, 0x68,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x61, 0x72, 0x65, 0x6e, 0x74, 0x70, 0x61,
	0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22, 0x54, 0x0a, 0x16, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x3a, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x24, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65, 0x72, 0x2e,
	0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65, 0x50, 0x61,
	0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x32, 0x7d, 0x0a, 0x0c,
	0x47, 0x65, 0x6e, 0x41, 0x69, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x6d, 0x0a, 0x11,
	0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x12, 0x29, 0x2e, 0x63, 0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65,
	0x72, 0x2e, 0x67, 0x65, 0x6e, 0x61, 0x69, 0x2e, 0x47, 0x69, 0x74, 0x68, 0x75, 0x62, 0x43, 0x6f,
	0x6e, 0x74, 0x65, 0x78, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e, 0x63,
	0x6f, 0x64, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x72, 0x65, 0x72, 0x2e, 0x67, 0x65, 0x6e,
	0x61, 0x69, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x64, 0x54, 0x65, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x14, 0x5a, 0x12, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65,
	0x64, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_gen_ai_proto_rawDescOnce sync.Once
	file_gen_ai_proto_rawDescData = file_gen_ai_proto_rawDesc
)

func file_gen_ai_proto_rawDescGZIP() []byte {
	file_gen_ai_proto_rawDescOnce.Do(func() {
		file_gen_ai_proto_rawDescData = protoimpl.X.CompressGZIP(file_gen_ai_proto_rawDescData)
	})
	return file_gen_ai_proto_rawDescData
}

var file_gen_ai_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_gen_ai_proto_goTypes = []any{
	(*SourceFileDependencyPayload)(nil), // 0: codesourcerer.genai.SourceFileDependencyPayload
	(*SourceFilePayload)(nil),           // 1: codesourcerer.genai.SourceFilePayload
	(*BasicConfig)(nil),                 // 2: codesourcerer.genai.BasicConfig
	(*Configuration)(nil),               // 3: codesourcerer.genai.Configuration
	(*GithubContextRequest)(nil),        // 4: codesourcerer.genai.GithubContextRequest
	(*TestFilePayload)(nil),             // 5: codesourcerer.genai.TestFilePayload
	(*GeneratedTestsResponse)(nil),      // 6: codesourcerer.genai.GeneratedTestsResponse
	nil,                                 // 7: codesourcerer.genai.Configuration.ExtrasEntry
}
var file_gen_ai_proto_depIdxs = []int32{
	0, // 0: codesourcerer.genai.SourceFilePayload.dependencies:type_name -> codesourcerer.genai.SourceFileDependencyPayload
	2, // 1: codesourcerer.genai.Configuration.configuration:type_name -> codesourcerer.genai.BasicConfig
	7, // 2: codesourcerer.genai.Configuration.extras:type_name -> codesourcerer.genai.Configuration.ExtrasEntry
	3, // 3: codesourcerer.genai.GithubContextRequest.config:type_name -> codesourcerer.genai.Configuration
	1, // 4: codesourcerer.genai.GithubContextRequest.files:type_name -> codesourcerer.genai.SourceFilePayload
	5, // 5: codesourcerer.genai.GeneratedTestsResponse.tests:type_name -> codesourcerer.genai.TestFilePayload
	4, // 6: codesourcerer.genai.GenAiService.GenerateTestFiles:input_type -> codesourcerer.genai.GithubContextRequest
	6, // 7: codesourcerer.genai.GenAiService.GenerateTestFiles:output_type -> codesourcerer.genai.GeneratedTestsResponse
	7, // [7:8] is the sub-list for method output_type
	6, // [6:7] is the sub-list for method input_type
	6, // [6:6] is the sub-list for extension type_name
	6, // [6:6] is the sub-list for extension extendee
	0, // [0:6] is the sub-list for field type_name
}

func init() { file_gen_ai_proto_init() }
func file_gen_ai_proto_init() {
	if File_gen_ai_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_gen_ai_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_gen_ai_proto_goTypes,
		DependencyIndexes: file_gen_ai_proto_depIdxs,
		MessageInfos:      file_gen_ai_proto_msgTypes,
	}.Build()
	File_gen_ai_proto = out.File
	file_gen_ai_proto_rawDesc = nil
	file_gen_ai_proto_goTypes = nil
	file_gen_ai_proto_depIdxs = nil
}
