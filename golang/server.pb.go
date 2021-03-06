// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.4
// source: server.proto

package golang

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

type EchoRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EchoRequest) Reset() {
	*x = EchoRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoRequest) ProtoMessage() {}

func (x *EchoRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoRequest.ProtoReflect.Descriptor instead.
func (*EchoRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{0}
}

func (x *EchoRequest) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type EchoResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Msg string `protobuf:"bytes,1,opt,name=msg,proto3" json:"msg,omitempty"`
}

func (x *EchoResponse) Reset() {
	*x = EchoResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EchoResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EchoResponse) ProtoMessage() {}

func (x *EchoResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EchoResponse.ProtoReflect.Descriptor instead.
func (*EchoResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{1}
}

func (x *EchoResponse) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type Value struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to ValueType:
	//	*Value_F64
	//	*Value_Str
	ValueType isValue_ValueType `protobuf_oneof:"ValueType"`
}

func (x *Value) Reset() {
	*x = Value{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Value) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Value) ProtoMessage() {}

func (x *Value) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Value.ProtoReflect.Descriptor instead.
func (*Value) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{2}
}

func (m *Value) GetValueType() isValue_ValueType {
	if m != nil {
		return m.ValueType
	}
	return nil
}

func (x *Value) GetF64() float64 {
	if x, ok := x.GetValueType().(*Value_F64); ok {
		return x.F64
	}
	return 0
}

func (x *Value) GetStr() string {
	if x, ok := x.GetValueType().(*Value_Str); ok {
		return x.Str
	}
	return ""
}

type isValue_ValueType interface {
	isValue_ValueType()
}

type Value_F64 struct {
	F64 float64 `protobuf:"fixed64,2,opt,name=f64,proto3,oneof"`
}

type Value_Str struct {
	Str string `protobuf:"bytes,3,opt,name=str,proto3,oneof"`
}

func (*Value_F64) isValue_ValueType() {}

func (*Value_Str) isValue_ValueType() {}

type Sample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        uint64   `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Timestamp uint64   `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	Values    []*Value `protobuf:"bytes,3,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *Sample) Reset() {
	*x = Sample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{3}
}

func (x *Sample) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Sample) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *Sample) GetValues() []*Value {
	if x != nil {
		return x.Values
	}
	return nil
}

type Samples struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags    map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Samples []*Sample         `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *Samples) Reset() {
	*x = Samples{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Samples) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Samples) ProtoMessage() {}

func (x *Samples) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Samples.ProtoReflect.Descriptor instead.
func (*Samples) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{4}
}

func (x *Samples) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *Samples) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type PushRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Samples []*Samples `protobuf:"bytes,1,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *PushRequest) Reset() {
	*x = PushRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushRequest) ProtoMessage() {}

func (x *PushRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushRequest.ProtoReflect.Descriptor instead.
func (*PushRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{5}
}

func (x *PushRequest) GetSamples() []*Samples {
	if x != nil {
		return x.Samples
	}
	return nil
}

type SampleResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Result bool   `protobuf:"varint,2,opt,name=result,proto3" json:"result,omitempty"`
}

func (x *SampleResult) Reset() {
	*x = SampleResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SampleResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SampleResult) ProtoMessage() {}

func (x *SampleResult) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SampleResult.ProtoReflect.Descriptor instead.
func (*SampleResult) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{6}
}

func (x *SampleResult) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SampleResult) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

type PushResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Responses []*SampleResult `protobuf:"bytes,1,rep,name=responses,proto3" json:"responses,omitempty"`
}

func (x *PushResponse) Reset() {
	*x = PushResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PushResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PushResponse) ProtoMessage() {}

func (x *PushResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PushResponse.ProtoReflect.Descriptor instead.
func (*PushResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{7}
}

func (x *PushResponse) GetResponses() []*SampleResult {
	if x != nil {
		return x.Responses
	}
	return nil
}

type ReadRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Regex string `protobuf:"bytes,1,opt,name=regex,proto3" json:"regex,omitempty"`
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{8}
}

func (x *ReadRequest) GetRegex() string {
	if x != nil {
		return x.Regex
	}
	return ""
}

type KafkaConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Bootstrap string `protobuf:"bytes,1,opt,name=bootstrap,proto3" json:"bootstrap,omitempty"`
	Topic     string `protobuf:"bytes,2,opt,name=topic,proto3" json:"topic,omitempty"`
}

func (x *KafkaConfig) Reset() {
	*x = KafkaConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *KafkaConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KafkaConfig) ProtoMessage() {}

func (x *KafkaConfig) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KafkaConfig.ProtoReflect.Descriptor instead.
func (*KafkaConfig) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{9}
}

func (x *KafkaConfig) GetBootstrap() string {
	if x != nil {
		return x.Bootstrap
	}
	return ""
}

func (x *KafkaConfig) GetTopic() string {
	if x != nil {
		return x.Topic
	}
	return ""
}

type FileConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dir  string `protobuf:"bytes,1,opt,name=dir,proto3" json:"dir,omitempty"`
	File string `protobuf:"bytes,2,opt,name=file,proto3" json:"file,omitempty"`
}

func (x *FileConfig) Reset() {
	*x = FileConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FileConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FileConfig) ProtoMessage() {}

func (x *FileConfig) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FileConfig.ProtoReflect.Descriptor instead.
func (*FileConfig) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{10}
}

func (x *FileConfig) GetDir() string {
	if x != nil {
		return x.Dir
	}
	return ""
}

func (x *FileConfig) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

type QueueConfig struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Configs:
	//	*QueueConfig_Kafka
	//	*QueueConfig_File
	Configs isQueueConfig_Configs `protobuf_oneof:"configs"`
}

func (x *QueueConfig) Reset() {
	*x = QueueConfig{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *QueueConfig) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueueConfig) ProtoMessage() {}

func (x *QueueConfig) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueueConfig.ProtoReflect.Descriptor instead.
func (*QueueConfig) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{11}
}

func (m *QueueConfig) GetConfigs() isQueueConfig_Configs {
	if m != nil {
		return m.Configs
	}
	return nil
}

func (x *QueueConfig) GetKafka() *KafkaConfig {
	if x, ok := x.GetConfigs().(*QueueConfig_Kafka); ok {
		return x.Kafka
	}
	return nil
}

func (x *QueueConfig) GetFile() *FileConfig {
	if x, ok := x.GetConfigs().(*QueueConfig_File); ok {
		return x.File
	}
	return nil
}

type isQueueConfig_Configs interface {
	isQueueConfig_Configs()
}

type QueueConfig_Kafka struct {
	Kafka *KafkaConfig `protobuf:"bytes,1,opt,name=kafka,proto3,oneof"`
}

type QueueConfig_File struct {
	File *FileConfig `protobuf:"bytes,2,opt,name=file,proto3,oneof"`
}

func (*QueueConfig_Kafka) isQueueConfig_Configs() {}

func (*QueueConfig_File) isQueueConfig_Configs() {}

type SeriesSnapshot struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tags          map[string]string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ResponseQueue *QueueConfig      `protobuf:"bytes,2,opt,name=response_queue,json=responseQueue,proto3" json:"response_queue,omitempty"`
	DataQueue     *QueueConfig      `protobuf:"bytes,3,opt,name=data_queue,json=dataQueue,proto3" json:"data_queue,omitempty"`
	Offset        uint64            `protobuf:"varint,4,opt,name=offset,proto3" json:"offset,omitempty"`
}

func (x *SeriesSnapshot) Reset() {
	*x = SeriesSnapshot{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriesSnapshot) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriesSnapshot) ProtoMessage() {}

func (x *SeriesSnapshot) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriesSnapshot.ProtoReflect.Descriptor instead.
func (*SeriesSnapshot) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{12}
}

func (x *SeriesSnapshot) GetTags() map[string]string {
	if x != nil {
		return x.Tags
	}
	return nil
}

func (x *SeriesSnapshot) GetResponseQueue() *QueueConfig {
	if x != nil {
		return x.ResponseQueue
	}
	return nil
}

func (x *SeriesSnapshot) GetDataQueue() *QueueConfig {
	if x != nil {
		return x.DataQueue
	}
	return nil
}

func (x *SeriesSnapshot) GetOffset() uint64 {
	if x != nil {
		return x.Offset
	}
	return 0
}

type ReadResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Snapshots []*SeriesSnapshot `protobuf:"bytes,1,rep,name=snapshots,proto3" json:"snapshots,omitempty"`
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_server_proto_msgTypes[13]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_server_proto_msgTypes[13]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_server_proto_rawDescGZIP(), []int{13}
}

func (x *ReadResponse) GetSnapshots() []*SeriesSnapshot {
	if x != nil {
		return x.Snapshots
	}
	return nil
}

var File_server_proto protoreflect.FileDescriptor

var file_server_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x22, 0x1f, 0x0a, 0x0b, 0x45, 0x63, 0x68, 0x6f,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x20, 0x0a, 0x0c, 0x45, 0x63, 0x68,
	0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6d, 0x73, 0x67, 0x22, 0x3c, 0x0a, 0x05, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x12, 0x12, 0x0a, 0x03, 0x66, 0x36, 0x34, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x01, 0x48, 0x00, 0x52, 0x03, 0x66, 0x36, 0x34, 0x12, 0x12, 0x0a, 0x03, 0x73, 0x74, 0x72, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x03, 0x73, 0x74, 0x72, 0x42, 0x0b, 0x0a, 0x09,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x54, 0x79, 0x70, 0x65, 0x22, 0x5f, 0x0a, 0x06, 0x53, 0x61, 0x6d,
	0x70, 0x6c, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x12, 0x27, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x0f, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x22, 0x9f, 0x01, 0x0a, 0x07, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x12, 0x2f, 0x0a, 0x04, 0x74, 0x61, 0x67, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72,
	0x79, 0x52, 0x04, 0x74, 0x61, 0x67, 0x73, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07, 0x73, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x1a, 0x37, 0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b,
	0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x3a, 0x0a, 0x0b,
	0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x73,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x6d,
	0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x52,
	0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x36, 0x0a, 0x0c, 0x53, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x72, 0x65, 0x73, 0x75,
	0x6c, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x22, 0x44, 0x0a, 0x0c, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x34, 0x0a, 0x09, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x09, 0x72, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x73, 0x22, 0x23, 0x0a, 0x0b, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x72, 0x65, 0x67, 0x65, 0x78, 0x22, 0x41, 0x0a, 0x0b, 0x4b,
	0x61, 0x66, 0x6b, 0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x1c, 0x0a, 0x09, 0x62, 0x6f,
	0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x62,
	0x6f, 0x6f, 0x74, 0x73, 0x74, 0x72, 0x61, 0x70, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x70, 0x69,
	0x63, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x70, 0x69, 0x63, 0x22, 0x32,
	0x0a, 0x0a, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x10, 0x0a, 0x03,
	0x64, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x69, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x69,
	0x6c, 0x65, 0x22, 0x73, 0x0a, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x12, 0x2d, 0x0a, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x4b, 0x61, 0x66, 0x6b,
	0x61, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x05, 0x6b, 0x61, 0x66, 0x6b, 0x61,
	0x12, 0x2a, 0x0a, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x48, 0x00, 0x52, 0x04, 0x66, 0x69, 0x6c, 0x65, 0x42, 0x09, 0x0a, 0x07,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x22, 0x8d, 0x02, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x69,
	0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x12, 0x36, 0x0a, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f,
	0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70, 0x73, 0x68,
	0x6f, 0x74, 0x2e, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x74, 0x61,
	0x67, 0x73, 0x12, 0x3c, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x5f, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x09, 0x64, 0x61, 0x74,
	0x61, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x1a, 0x37,
	0x0a, 0x09, 0x54, 0x61, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x46, 0x0a, 0x0c, 0x52, 0x65, 0x61, 0x64, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73,
	0x68, 0x6f, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6d, 0x61, 0x63,
	0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x53, 0x6e, 0x61, 0x70,
	0x73, 0x68, 0x6f, 0x74, 0x52, 0x09, 0x73, 0x6e, 0x61, 0x70, 0x73, 0x68, 0x6f, 0x74, 0x73, 0x32,
	0x85, 0x02, 0x0a, 0x0b, 0x54, 0x73, 0x64, 0x62, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x37, 0x0a, 0x04, 0x45, 0x63, 0x68, 0x6f, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72,
	0x70, 0x63, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16,
	0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x41, 0x0a, 0x0a, 0x45, 0x63, 0x68, 0x6f,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x45, 0x63, 0x68, 0x6f, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x41, 0x0a, 0x0a, 0x50,
	0x75, 0x73, 0x68, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68,
	0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x16, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x50, 0x75, 0x73, 0x68,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x12, 0x37,
	0x0a, 0x04, 0x52, 0x65, 0x61, 0x64, 0x12, 0x15, 0x2e, 0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70,
	0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e,
	0x6d, 0x61, 0x63, 0x68, 0x5f, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x27, 0x5a, 0x25, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x73, 0x6f, 0x6c, 0x6c, 0x65, 0x7a, 0x61, 0x2f, 0x6d,
	0x61, 0x63, 0x68, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_server_proto_rawDescOnce sync.Once
	file_server_proto_rawDescData = file_server_proto_rawDesc
)

func file_server_proto_rawDescGZIP() []byte {
	file_server_proto_rawDescOnce.Do(func() {
		file_server_proto_rawDescData = protoimpl.X.CompressGZIP(file_server_proto_rawDescData)
	})
	return file_server_proto_rawDescData
}

var file_server_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_server_proto_goTypes = []interface{}{
	(*EchoRequest)(nil),    // 0: mach_rpc.EchoRequest
	(*EchoResponse)(nil),   // 1: mach_rpc.EchoResponse
	(*Value)(nil),          // 2: mach_rpc.Value
	(*Sample)(nil),         // 3: mach_rpc.Sample
	(*Samples)(nil),        // 4: mach_rpc.Samples
	(*PushRequest)(nil),    // 5: mach_rpc.PushRequest
	(*SampleResult)(nil),   // 6: mach_rpc.SampleResult
	(*PushResponse)(nil),   // 7: mach_rpc.PushResponse
	(*ReadRequest)(nil),    // 8: mach_rpc.ReadRequest
	(*KafkaConfig)(nil),    // 9: mach_rpc.KafkaConfig
	(*FileConfig)(nil),     // 10: mach_rpc.FileConfig
	(*QueueConfig)(nil),    // 11: mach_rpc.QueueConfig
	(*SeriesSnapshot)(nil), // 12: mach_rpc.SeriesSnapshot
	(*ReadResponse)(nil),   // 13: mach_rpc.ReadResponse
	nil,                    // 14: mach_rpc.Samples.TagsEntry
	nil,                    // 15: mach_rpc.SeriesSnapshot.TagsEntry
}
var file_server_proto_depIdxs = []int32{
	2,  // 0: mach_rpc.Sample.values:type_name -> mach_rpc.Value
	14, // 1: mach_rpc.Samples.tags:type_name -> mach_rpc.Samples.TagsEntry
	3,  // 2: mach_rpc.Samples.samples:type_name -> mach_rpc.Sample
	4,  // 3: mach_rpc.PushRequest.samples:type_name -> mach_rpc.Samples
	6,  // 4: mach_rpc.PushResponse.responses:type_name -> mach_rpc.SampleResult
	9,  // 5: mach_rpc.QueueConfig.kafka:type_name -> mach_rpc.KafkaConfig
	10, // 6: mach_rpc.QueueConfig.file:type_name -> mach_rpc.FileConfig
	15, // 7: mach_rpc.SeriesSnapshot.tags:type_name -> mach_rpc.SeriesSnapshot.TagsEntry
	11, // 8: mach_rpc.SeriesSnapshot.response_queue:type_name -> mach_rpc.QueueConfig
	11, // 9: mach_rpc.SeriesSnapshot.data_queue:type_name -> mach_rpc.QueueConfig
	12, // 10: mach_rpc.ReadResponse.snapshots:type_name -> mach_rpc.SeriesSnapshot
	0,  // 11: mach_rpc.TsdbService.Echo:input_type -> mach_rpc.EchoRequest
	0,  // 12: mach_rpc.TsdbService.EchoStream:input_type -> mach_rpc.EchoRequest
	5,  // 13: mach_rpc.TsdbService.PushStream:input_type -> mach_rpc.PushRequest
	8,  // 14: mach_rpc.TsdbService.Read:input_type -> mach_rpc.ReadRequest
	1,  // 15: mach_rpc.TsdbService.Echo:output_type -> mach_rpc.EchoResponse
	1,  // 16: mach_rpc.TsdbService.EchoStream:output_type -> mach_rpc.EchoResponse
	7,  // 17: mach_rpc.TsdbService.PushStream:output_type -> mach_rpc.PushResponse
	13, // 18: mach_rpc.TsdbService.Read:output_type -> mach_rpc.ReadResponse
	15, // [15:19] is the sub-list for method output_type
	11, // [11:15] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_server_proto_init() }
func file_server_proto_init() {
	if File_server_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_server_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoRequest); i {
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
		file_server_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EchoResponse); i {
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
		file_server_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Value); i {
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
		file_server_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Sample); i {
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
		file_server_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Samples); i {
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
		file_server_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushRequest); i {
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
		file_server_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SampleResult); i {
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
		file_server_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PushResponse); i {
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
		file_server_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadRequest); i {
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
		file_server_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*KafkaConfig); i {
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
		file_server_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FileConfig); i {
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
		file_server_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*QueueConfig); i {
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
		file_server_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriesSnapshot); i {
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
		file_server_proto_msgTypes[13].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadResponse); i {
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
	file_server_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*Value_F64)(nil),
		(*Value_Str)(nil),
	}
	file_server_proto_msgTypes[11].OneofWrappers = []interface{}{
		(*QueueConfig_Kafka)(nil),
		(*QueueConfig_File)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_server_proto_goTypes,
		DependencyIndexes: file_server_proto_depIdxs,
		MessageInfos:      file_server_proto_msgTypes,
	}.Build()
	File_server_proto = out.File
	file_server_proto_rawDesc = nil
	file_server_proto_goTypes = nil
	file_server_proto_depIdxs = nil
}
