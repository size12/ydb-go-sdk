// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_experimental.proto

package Ydb_Experimental

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb"
	Ydb_Issue "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Issue"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/api/protos/Ydb_Operations"
	fmt "fmt"
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

type ExecuteStreamQueryRequest_ProfileMode int32

const (
	ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED ExecuteStreamQueryRequest_ProfileMode = 0
	ExecuteStreamQueryRequest_NONE                     ExecuteStreamQueryRequest_ProfileMode = 1
	ExecuteStreamQueryRequest_BASIC                    ExecuteStreamQueryRequest_ProfileMode = 2
	ExecuteStreamQueryRequest_FULL                     ExecuteStreamQueryRequest_ProfileMode = 3
)

var ExecuteStreamQueryRequest_ProfileMode_name = map[int32]string{
	0: "PROFILE_MODE_UNSPECIFIED",
	1: "NONE",
	2: "BASIC",
	3: "FULL",
}

var ExecuteStreamQueryRequest_ProfileMode_value = map[string]int32{
	"PROFILE_MODE_UNSPECIFIED": 0,
	"NONE":                     1,
	"BASIC":                    2,
	"FULL":                     3,
}

func (x ExecuteStreamQueryRequest_ProfileMode) String() string {
	return proto.EnumName(ExecuteStreamQueryRequest_ProfileMode_name, int32(x))
}

func (ExecuteStreamQueryRequest_ProfileMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3, 0}
}

type UploadRowsRequest struct {
	Table                string                          `protobuf:"bytes,1,opt,name=table,proto3" json:"table,omitempty"`
	Rows                 *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=rows,proto3" json:"rows,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,3,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *UploadRowsRequest) Reset()         { *m = UploadRowsRequest{} }
func (m *UploadRowsRequest) String() string { return proto.CompactTextString(m) }
func (*UploadRowsRequest) ProtoMessage()    {}
func (*UploadRowsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{0}
}

func (m *UploadRowsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsRequest.Unmarshal(m, b)
}
func (m *UploadRowsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsRequest.Marshal(b, m, deterministic)
}
func (m *UploadRowsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsRequest.Merge(m, src)
}
func (m *UploadRowsRequest) XXX_Size() int {
	return xxx_messageInfo_UploadRowsRequest.Size(m)
}
func (m *UploadRowsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsRequest proto.InternalMessageInfo

func (m *UploadRowsRequest) GetTable() string {
	if m != nil {
		return m.Table
	}
	return ""
}

func (m *UploadRowsRequest) GetRows() *Ydb.TypedValue {
	if m != nil {
		return m.Rows
	}
	return nil
}

func (m *UploadRowsRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type UploadRowsResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *UploadRowsResponse) Reset()         { *m = UploadRowsResponse{} }
func (m *UploadRowsResponse) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResponse) ProtoMessage()    {}
func (*UploadRowsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{1}
}

func (m *UploadRowsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResponse.Unmarshal(m, b)
}
func (m *UploadRowsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResponse.Marshal(b, m, deterministic)
}
func (m *UploadRowsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResponse.Merge(m, src)
}
func (m *UploadRowsResponse) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResponse.Size(m)
}
func (m *UploadRowsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResponse proto.InternalMessageInfo

func (m *UploadRowsResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type UploadRowsResult struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *UploadRowsResult) Reset()         { *m = UploadRowsResult{} }
func (m *UploadRowsResult) String() string { return proto.CompactTextString(m) }
func (*UploadRowsResult) ProtoMessage()    {}
func (*UploadRowsResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{2}
}

func (m *UploadRowsResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UploadRowsResult.Unmarshal(m, b)
}
func (m *UploadRowsResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UploadRowsResult.Marshal(b, m, deterministic)
}
func (m *UploadRowsResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UploadRowsResult.Merge(m, src)
}
func (m *UploadRowsResult) XXX_Size() int {
	return xxx_messageInfo_UploadRowsResult.Size(m)
}
func (m *UploadRowsResult) XXX_DiscardUnknown() {
	xxx_messageInfo_UploadRowsResult.DiscardUnknown(m)
}

var xxx_messageInfo_UploadRowsResult proto.InternalMessageInfo

type ExecuteStreamQueryRequest struct {
	YqlText              string                                `protobuf:"bytes,1,opt,name=yql_text,json=yqlText,proto3" json:"yql_text,omitempty"`
	Parameters           map[string]*Ydb.TypedValue            `protobuf:"bytes,2,rep,name=parameters,proto3" json:"parameters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ProfileMode          ExecuteStreamQueryRequest_ProfileMode `protobuf:"varint,3,opt,name=profile_mode,json=profileMode,proto3,enum=Ydb.Experimental.ExecuteStreamQueryRequest_ProfileMode" json:"profile_mode,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *ExecuteStreamQueryRequest) Reset()         { *m = ExecuteStreamQueryRequest{} }
func (m *ExecuteStreamQueryRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryRequest) ProtoMessage()    {}
func (*ExecuteStreamQueryRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{3}
}

func (m *ExecuteStreamQueryRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryRequest.Merge(m, src)
}
func (m *ExecuteStreamQueryRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryRequest.Size(m)
}
func (m *ExecuteStreamQueryRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryRequest proto.InternalMessageInfo

func (m *ExecuteStreamQueryRequest) GetYqlText() string {
	if m != nil {
		return m.YqlText
	}
	return ""
}

func (m *ExecuteStreamQueryRequest) GetParameters() map[string]*Ydb.TypedValue {
	if m != nil {
		return m.Parameters
	}
	return nil
}

func (m *ExecuteStreamQueryRequest) GetProfileMode() ExecuteStreamQueryRequest_ProfileMode {
	if m != nil {
		return m.ProfileMode
	}
	return ExecuteStreamQueryRequest_PROFILE_MODE_UNSPECIFIED
}

type ExecuteStreamQueryResponse struct {
	Status               Ydb.StatusIds_StatusCode  `protobuf:"varint,1,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues               []*Ydb_Issue.IssueMessage `protobuf:"bytes,2,rep,name=issues,proto3" json:"issues,omitempty"`
	Result               *ExecuteStreamQueryResult `protobuf:"bytes,3,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *ExecuteStreamQueryResponse) Reset()         { *m = ExecuteStreamQueryResponse{} }
func (m *ExecuteStreamQueryResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResponse) ProtoMessage()    {}
func (*ExecuteStreamQueryResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{4}
}

func (m *ExecuteStreamQueryResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResponse.Merge(m, src)
}
func (m *ExecuteStreamQueryResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResponse.Size(m)
}
func (m *ExecuteStreamQueryResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResponse proto.InternalMessageInfo

func (m *ExecuteStreamQueryResponse) GetStatus() Ydb.StatusIds_StatusCode {
	if m != nil {
		return m.Status
	}
	return Ydb.StatusIds_STATUS_CODE_UNSPECIFIED
}

func (m *ExecuteStreamQueryResponse) GetIssues() []*Ydb_Issue.IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *ExecuteStreamQueryResponse) GetResult() *ExecuteStreamQueryResult {
	if m != nil {
		return m.Result
	}
	return nil
}

type StreamQueryProgress struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamQueryProgress) Reset()         { *m = StreamQueryProgress{} }
func (m *StreamQueryProgress) String() string { return proto.CompactTextString(m) }
func (*StreamQueryProgress) ProtoMessage()    {}
func (*StreamQueryProgress) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{5}
}

func (m *StreamQueryProgress) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamQueryProgress.Unmarshal(m, b)
}
func (m *StreamQueryProgress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamQueryProgress.Marshal(b, m, deterministic)
}
func (m *StreamQueryProgress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamQueryProgress.Merge(m, src)
}
func (m *StreamQueryProgress) XXX_Size() int {
	return xxx_messageInfo_StreamQueryProgress.Size(m)
}
func (m *StreamQueryProgress) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamQueryProgress.DiscardUnknown(m)
}

var xxx_messageInfo_StreamQueryProgress proto.InternalMessageInfo

type ExecuteStreamQueryResult struct {
	// Types that are valid to be assigned to Result:
	//	*ExecuteStreamQueryResult_ResultSet
	//	*ExecuteStreamQueryResult_Profile
	//	*ExecuteStreamQueryResult_Progress
	Result               isExecuteStreamQueryResult_Result `protobuf_oneof:"result"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ExecuteStreamQueryResult) Reset()         { *m = ExecuteStreamQueryResult{} }
func (m *ExecuteStreamQueryResult) String() string { return proto.CompactTextString(m) }
func (*ExecuteStreamQueryResult) ProtoMessage()    {}
func (*ExecuteStreamQueryResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{6}
}

func (m *ExecuteStreamQueryResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteStreamQueryResult.Unmarshal(m, b)
}
func (m *ExecuteStreamQueryResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteStreamQueryResult.Marshal(b, m, deterministic)
}
func (m *ExecuteStreamQueryResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteStreamQueryResult.Merge(m, src)
}
func (m *ExecuteStreamQueryResult) XXX_Size() int {
	return xxx_messageInfo_ExecuteStreamQueryResult.Size(m)
}
func (m *ExecuteStreamQueryResult) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteStreamQueryResult.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteStreamQueryResult proto.InternalMessageInfo

type isExecuteStreamQueryResult_Result interface {
	isExecuteStreamQueryResult_Result()
}

type ExecuteStreamQueryResult_ResultSet struct {
	ResultSet *Ydb.ResultSet `protobuf:"bytes,1,opt,name=result_set,json=resultSet,proto3,oneof"`
}

type ExecuteStreamQueryResult_Profile struct {
	Profile string `protobuf:"bytes,2,opt,name=profile,proto3,oneof"`
}

type ExecuteStreamQueryResult_Progress struct {
	Progress *StreamQueryProgress `protobuf:"bytes,3,opt,name=progress,proto3,oneof"`
}

func (*ExecuteStreamQueryResult_ResultSet) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Profile) isExecuteStreamQueryResult_Result() {}

func (*ExecuteStreamQueryResult_Progress) isExecuteStreamQueryResult_Result() {}

func (m *ExecuteStreamQueryResult) GetResult() isExecuteStreamQueryResult_Result {
	if m != nil {
		return m.Result
	}
	return nil
}

func (m *ExecuteStreamQueryResult) GetResultSet() *Ydb.ResultSet {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_ResultSet); ok {
		return x.ResultSet
	}
	return nil
}

func (m *ExecuteStreamQueryResult) GetProfile() string {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_Profile); ok {
		return x.Profile
	}
	return ""
}

func (m *ExecuteStreamQueryResult) GetProgress() *StreamQueryProgress {
	if x, ok := m.GetResult().(*ExecuteStreamQueryResult_Progress); ok {
		return x.Progress
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*ExecuteStreamQueryResult) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*ExecuteStreamQueryResult_ResultSet)(nil),
		(*ExecuteStreamQueryResult_Profile)(nil),
		(*ExecuteStreamQueryResult_Progress)(nil),
	}
}

type GetDiskSpaceUsageRequest struct {
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,1,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	Database             string                          `protobuf:"bytes,2,opt,name=database,proto3" json:"database,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *GetDiskSpaceUsageRequest) Reset()         { *m = GetDiskSpaceUsageRequest{} }
func (m *GetDiskSpaceUsageRequest) String() string { return proto.CompactTextString(m) }
func (*GetDiskSpaceUsageRequest) ProtoMessage()    {}
func (*GetDiskSpaceUsageRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{7}
}

func (m *GetDiskSpaceUsageRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiskSpaceUsageRequest.Unmarshal(m, b)
}
func (m *GetDiskSpaceUsageRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiskSpaceUsageRequest.Marshal(b, m, deterministic)
}
func (m *GetDiskSpaceUsageRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiskSpaceUsageRequest.Merge(m, src)
}
func (m *GetDiskSpaceUsageRequest) XXX_Size() int {
	return xxx_messageInfo_GetDiskSpaceUsageRequest.Size(m)
}
func (m *GetDiskSpaceUsageRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiskSpaceUsageRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiskSpaceUsageRequest proto.InternalMessageInfo

func (m *GetDiskSpaceUsageRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

func (m *GetDiskSpaceUsageRequest) GetDatabase() string {
	if m != nil {
		return m.Database
	}
	return ""
}

type GetDiskSpaceUsageResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *GetDiskSpaceUsageResponse) Reset()         { *m = GetDiskSpaceUsageResponse{} }
func (m *GetDiskSpaceUsageResponse) String() string { return proto.CompactTextString(m) }
func (*GetDiskSpaceUsageResponse) ProtoMessage()    {}
func (*GetDiskSpaceUsageResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{8}
}

func (m *GetDiskSpaceUsageResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiskSpaceUsageResponse.Unmarshal(m, b)
}
func (m *GetDiskSpaceUsageResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiskSpaceUsageResponse.Marshal(b, m, deterministic)
}
func (m *GetDiskSpaceUsageResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiskSpaceUsageResponse.Merge(m, src)
}
func (m *GetDiskSpaceUsageResponse) XXX_Size() int {
	return xxx_messageInfo_GetDiskSpaceUsageResponse.Size(m)
}
func (m *GetDiskSpaceUsageResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiskSpaceUsageResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiskSpaceUsageResponse proto.InternalMessageInfo

func (m *GetDiskSpaceUsageResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type GetDiskSpaceUsageResult struct {
	CloudId    string `protobuf:"bytes,1,opt,name=cloud_id,json=cloudId,proto3" json:"cloud_id,omitempty"`
	FolderId   string `protobuf:"bytes,2,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	DatabaseId string `protobuf:"bytes,3,opt,name=database_id,json=databaseId,proto3" json:"database_id,omitempty"`
	// in bytes
	TotalSize            uint64   `protobuf:"varint,4,opt,name=total_size,json=totalSize,proto3" json:"total_size,omitempty"`
	DataSize             uint64   `protobuf:"varint,5,opt,name=data_size,json=dataSize,proto3" json:"data_size,omitempty"`
	IndexSize            uint64   `protobuf:"varint,6,opt,name=index_size,json=indexSize,proto3" json:"index_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetDiskSpaceUsageResult) Reset()         { *m = GetDiskSpaceUsageResult{} }
func (m *GetDiskSpaceUsageResult) String() string { return proto.CompactTextString(m) }
func (*GetDiskSpaceUsageResult) ProtoMessage()    {}
func (*GetDiskSpaceUsageResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac21a693e2c386a5, []int{9}
}

func (m *GetDiskSpaceUsageResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetDiskSpaceUsageResult.Unmarshal(m, b)
}
func (m *GetDiskSpaceUsageResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetDiskSpaceUsageResult.Marshal(b, m, deterministic)
}
func (m *GetDiskSpaceUsageResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetDiskSpaceUsageResult.Merge(m, src)
}
func (m *GetDiskSpaceUsageResult) XXX_Size() int {
	return xxx_messageInfo_GetDiskSpaceUsageResult.Size(m)
}
func (m *GetDiskSpaceUsageResult) XXX_DiscardUnknown() {
	xxx_messageInfo_GetDiskSpaceUsageResult.DiscardUnknown(m)
}

var xxx_messageInfo_GetDiskSpaceUsageResult proto.InternalMessageInfo

func (m *GetDiskSpaceUsageResult) GetCloudId() string {
	if m != nil {
		return m.CloudId
	}
	return ""
}

func (m *GetDiskSpaceUsageResult) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *GetDiskSpaceUsageResult) GetDatabaseId() string {
	if m != nil {
		return m.DatabaseId
	}
	return ""
}

func (m *GetDiskSpaceUsageResult) GetTotalSize() uint64 {
	if m != nil {
		return m.TotalSize
	}
	return 0
}

func (m *GetDiskSpaceUsageResult) GetDataSize() uint64 {
	if m != nil {
		return m.DataSize
	}
	return 0
}

func (m *GetDiskSpaceUsageResult) GetIndexSize() uint64 {
	if m != nil {
		return m.IndexSize
	}
	return 0
}

func init() {
	proto.RegisterEnum("Ydb.Experimental.ExecuteStreamQueryRequest_ProfileMode", ExecuteStreamQueryRequest_ProfileMode_name, ExecuteStreamQueryRequest_ProfileMode_value)
	proto.RegisterType((*UploadRowsRequest)(nil), "Ydb.Experimental.UploadRowsRequest")
	proto.RegisterType((*UploadRowsResponse)(nil), "Ydb.Experimental.UploadRowsResponse")
	proto.RegisterType((*UploadRowsResult)(nil), "Ydb.Experimental.UploadRowsResult")
	proto.RegisterType((*ExecuteStreamQueryRequest)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest")
	proto.RegisterMapType((map[string]*Ydb.TypedValue)(nil), "Ydb.Experimental.ExecuteStreamQueryRequest.ParametersEntry")
	proto.RegisterType((*ExecuteStreamQueryResponse)(nil), "Ydb.Experimental.ExecuteStreamQueryResponse")
	proto.RegisterType((*StreamQueryProgress)(nil), "Ydb.Experimental.StreamQueryProgress")
	proto.RegisterType((*ExecuteStreamQueryResult)(nil), "Ydb.Experimental.ExecuteStreamQueryResult")
	proto.RegisterType((*GetDiskSpaceUsageRequest)(nil), "Ydb.Experimental.GetDiskSpaceUsageRequest")
	proto.RegisterType((*GetDiskSpaceUsageResponse)(nil), "Ydb.Experimental.GetDiskSpaceUsageResponse")
	proto.RegisterType((*GetDiskSpaceUsageResult)(nil), "Ydb.Experimental.GetDiskSpaceUsageResult")
}

func init() { proto.RegisterFile("ydb_experimental.proto", fileDescriptor_ac21a693e2c386a5) }

var fileDescriptor_ac21a693e2c386a5 = []byte{
	// 793 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x55, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0xae, 0x9b, 0x36, 0x9b, 0x9c, 0xa0, 0x36, 0x0c, 0x3f, 0x9b, 0x64, 0x59, 0x6d, 0x65, 0xb4,
	0x52, 0x84, 0x90, 0x03, 0x05, 0x69, 0x11, 0x5c, 0x91, 0xd6, 0xa5, 0x5e, 0xf5, 0x27, 0x4c, 0x5a,
	0x24, 0xe0, 0xc2, 0x9a, 0x64, 0xce, 0xae, 0xac, 0x3a, 0x19, 0x77, 0x66, 0xcc, 0xc6, 0x7b, 0xc9,
	0x2d, 0x6f, 0xc0, 0x83, 0xf0, 0x04, 0xdc, 0xf2, 0x3e, 0x5c, 0xa2, 0x19, 0x8f, 0xd3, 0x6c, 0x49,
	0x2b, 0x7e, 0x6e, 0x22, 0xcf, 0xf7, 0x9d, 0xef, 0xcc, 0xf9, 0xcd, 0xc0, 0xfb, 0x05, 0x9f, 0xc4,
	0xb8, 0xc8, 0x50, 0x26, 0x33, 0x9c, 0x6b, 0x96, 0x06, 0x99, 0x14, 0x5a, 0x90, 0xf6, 0xf7, 0x7c,
	0x12, 0x84, 0x2b, 0x78, 0xef, 0x93, 0xab, 0xe4, 0x2a, 0x99, 0xc9, 0x41, 0x96, 0x4f, 0xd2, 0x64,
	0x3a, 0x60, 0x59, 0x32, 0xb0, 0xa6, 0x6a, 0x60, 0x5c, 0x24, 0x4a, 0xe5, 0x18, 0xcf, 0x50, 0x29,
	0xf6, 0x12, 0x4b, 0x1f, 0xbd, 0x8f, 0xef, 0x55, 0x88, 0x0c, 0x25, 0xd3, 0x89, 0x98, 0x3b, 0xeb,
	0xc1, 0xbd, 0xd6, 0x4a, 0x33, 0x9d, 0xab, 0x78, 0x2a, 0x38, 0x2a, 0x27, 0xe8, 0xdf, 0x2b, 0xf8,
	0x89, 0xa5, 0xb9, 0x0b, 0xc4, 0xff, 0xd5, 0x83, 0xb7, 0x2f, 0xb3, 0x54, 0x30, 0x4e, 0xc5, 0x2b,
	0x45, 0xf1, 0x3a, 0x47, 0xa5, 0xc9, 0xbb, 0xb0, 0xad, 0xd9, 0x24, 0xc5, 0x8e, 0xb7, 0xe7, 0xf5,
	0x9b, 0xb4, 0x3c, 0x90, 0x0f, 0x61, 0x4b, 0x8a, 0x57, 0xaa, 0xb3, 0xb9, 0xe7, 0xf5, 0x5b, 0xfb,
	0xbb, 0x81, 0xa9, 0xc3, 0x45, 0x91, 0x21, 0xff, 0xce, 0x38, 0xa4, 0x96, 0x24, 0xcf, 0xa1, 0xbd,
	0x0c, 0x3f, 0xce, 0x98, 0x64, 0x33, 0xd5, 0xa9, 0x59, 0xc1, 0x13, 0x2b, 0x38, 0xaf, 0x48, 0x75,
	0xf3, 0x39, 0xb2, 0x66, 0x74, 0x57, 0xbc, 0x09, 0xf8, 0xa7, 0x40, 0x56, 0x63, 0x53, 0x99, 0x98,
	0x2b, 0x24, 0xcf, 0xa0, 0xb9, 0x34, 0xb4, 0x01, 0xb6, 0xf6, 0xbb, 0x77, 0xba, 0xa6, 0x37, 0xb6,
	0x3e, 0x81, 0xf6, 0x1b, 0xee, 0xf2, 0x54, 0xfb, 0xbf, 0xd4, 0xa0, 0x1b, 0x2e, 0x70, 0x9a, 0x6b,
	0x1c, 0x6b, 0x89, 0x6c, 0xf6, 0x6d, 0x8e, 0xb2, 0xa8, 0xea, 0xd0, 0x85, 0x46, 0x71, 0x9d, 0xc6,
	0x1a, 0x17, 0xda, 0x95, 0xe2, 0x41, 0x71, 0x9d, 0x5e, 0xe0, 0x42, 0x93, 0x1f, 0x01, 0x6c, 0x76,
	0xa8, 0x51, 0x9a, 0x92, 0xd4, 0xfa, 0xad, 0xfd, 0xaf, 0x82, 0xdb, 0xa3, 0x11, 0xdc, 0xe9, 0x3b,
	0x18, 0x2d, 0xd5, 0xe1, 0x5c, 0xcb, 0x82, 0xae, 0xb8, 0x23, 0x3f, 0xc0, 0x5b, 0x99, 0x14, 0x2f,
	0x92, 0x14, 0xe3, 0x99, 0xe0, 0x68, 0x0b, 0xb8, 0xb3, 0xff, 0xec, 0x5f, 0xb9, 0x2f, 0xf5, 0xa7,
	0x82, 0x23, 0x6d, 0x65, 0x37, 0x87, 0xde, 0x19, 0xec, 0xde, 0xba, 0x9a, 0xb4, 0xa1, 0x76, 0x85,
	0x85, 0xcb, 0xd0, 0x7c, 0x92, 0xa7, 0xb0, 0x6d, 0xa7, 0xe4, 0xae, 0x5e, 0x97, 0xec, 0x97, 0x9b,
	0x5f, 0x78, 0xfe, 0x73, 0x68, 0xad, 0xdc, 0x45, 0x3e, 0x80, 0xce, 0x88, 0x9e, 0x1f, 0x45, 0x27,
	0x61, 0x7c, 0x7a, 0x7e, 0x18, 0xc6, 0x97, 0x67, 0xe3, 0x51, 0x78, 0x10, 0x1d, 0x45, 0xe1, 0x61,
	0x7b, 0x83, 0x34, 0x60, 0xeb, 0xec, 0xfc, 0x2c, 0x6c, 0x7b, 0xa4, 0x09, 0xdb, 0xc3, 0xaf, 0xc7,
	0xd1, 0x41, 0x7b, 0xd3, 0x80, 0x47, 0x97, 0x27, 0x27, 0xed, 0x9a, 0xff, 0xbb, 0x07, 0xbd, 0x75,
	0x29, 0xb9, 0xce, 0x7f, 0x0a, 0xf5, 0x72, 0xd8, 0x6d, 0xa8, 0x3b, 0xae, 0xed, 0x63, 0x0b, 0x45,
	0x5c, 0xb9, 0xaf, 0x03, 0x93, 0xb2, 0x33, 0x24, 0x03, 0xa8, 0xdb, 0xfd, 0xab, 0x5a, 0xf4, 0xd0,
	0x4a, 0x22, 0x03, 0x95, 0xbf, 0xa7, 0xe5, 0x5e, 0x52, 0x67, 0x46, 0x86, 0x50, 0x97, 0x76, 0x34,
	0xdc, 0xd4, 0x7e, 0xf4, 0xcf, 0x8a, 0x6e, 0x14, 0xd4, 0x29, 0xfd, 0xf7, 0xe0, 0x9d, 0x15, 0x72,
	0x24, 0xc5, 0x4b, 0x89, 0x4a, 0xf9, 0xbf, 0x79, 0xd0, 0xb9, 0x4b, 0x4b, 0x06, 0x00, 0xa5, 0x3a,
	0x56, 0xa8, 0xdd, 0x58, 0xef, 0xd8, 0xbb, 0x4b, 0x83, 0x31, 0xea, 0xe3, 0x0d, 0xda, 0x94, 0xd5,
	0x81, 0xf4, 0xe0, 0x81, 0x6b, 0xab, 0x6d, 0x52, 0xf3, 0x78, 0x83, 0x56, 0x00, 0x39, 0x80, 0x46,
	0xe6, 0x6e, 0x75, 0x69, 0x3c, 0xfd, 0x7b, 0x1a, 0x6b, 0x42, 0x3c, 0xde, 0xa0, 0x4b, 0xe1, 0xb0,
	0x51, 0x55, 0xc2, 0xff, 0xd9, 0x83, 0xce, 0x37, 0xa8, 0x0f, 0x13, 0x75, 0x35, 0xce, 0xd8, 0x14,
	0x2f, 0x6d, 0xc5, 0xdc, 0x8e, 0xac, 0x5b, 0x78, 0xef, 0xbf, 0x2d, 0x3c, 0xe9, 0x41, 0x83, 0x33,
	0xcd, 0x26, 0x4c, 0xb9, 0xa4, 0xe8, 0xf2, 0xec, 0x5f, 0x40, 0x77, 0x4d, 0x0c, 0xff, 0xf7, 0x3f,
	0xe1, 0x0f, 0x0f, 0x1e, 0xae, 0x73, 0x6b, 0x5a, 0xd2, 0x85, 0xc6, 0x34, 0x15, 0x39, 0x8f, 0x13,
	0x5e, 0x6d, 0xbf, 0x3d, 0x47, 0x9c, 0x3c, 0x82, 0xe6, 0x0b, 0x91, 0x72, 0x94, 0x86, 0x73, 0x91,
	0x96, 0x40, 0xc4, 0xc9, 0x13, 0x68, 0x55, 0x51, 0x1b, 0xba, 0x66, 0x69, 0xa8, 0xa0, 0x88, 0x93,
	0xc7, 0x00, 0x5a, 0x68, 0x96, 0xc6, 0x2a, 0x79, 0x8d, 0x9d, 0xad, 0x3d, 0xaf, 0xbf, 0x45, 0x9b,
	0x16, 0x19, 0x27, 0xaf, 0xd1, 0x38, 0x37, 0xc6, 0x25, 0xbb, 0x6d, 0x59, 0x5b, 0x06, 0x4b, 0x3e,
	0x06, 0x48, 0xe6, 0x1c, 0x17, 0x25, 0x5b, 0x2f, 0xb5, 0x16, 0x31, 0xf4, 0xf0, 0x73, 0x78, 0x34,
	0x15, 0xb3, 0xa0, 0x60, 0x06, 0x09, 0x0a, 0x3e, 0x09, 0x56, 0x5f, 0xb0, 0x21, 0x59, 0x9d, 0x80,
	0x91, 0x7d, 0x12, 0xfe, 0xf4, 0xbc, 0x49, 0xdd, 0x3e, 0x06, 0x9f, 0xfd, 0x15, 0x00, 0x00, 0xff,
	0xff, 0xa6, 0x7b, 0xc0, 0xb2, 0xf3, 0x06, 0x00, 0x00,
}

const ()

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *UploadRowsRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}

// SetOperationParams implements ydb generic interface for setting
// operation parameters inside driver implementation.
func (m *GetDiskSpaceUsageRequest) SetOperationParams(v *Ydb_Operations.OperationParams) {
	m.OperationParams = v
}
