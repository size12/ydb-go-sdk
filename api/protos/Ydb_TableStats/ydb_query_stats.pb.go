// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_query_stats.proto

package Ydb_TableStats

import (
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

// Describes select, update (insert, upsert, replace) and delete operations
type OperationStats struct {
	Rows                 uint64   `protobuf:"varint,1,opt,name=rows,proto3" json:"rows,omitempty"`
	Bytes                uint64   `protobuf:"varint,2,opt,name=bytes,proto3" json:"bytes,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationStats) Reset()         { *m = OperationStats{} }
func (m *OperationStats) String() string { return proto.CompactTextString(m) }
func (*OperationStats) ProtoMessage()    {}
func (*OperationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{0}
}

func (m *OperationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationStats.Unmarshal(m, b)
}
func (m *OperationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationStats.Marshal(b, m, deterministic)
}
func (m *OperationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationStats.Merge(m, src)
}
func (m *OperationStats) XXX_Size() int {
	return xxx_messageInfo_OperationStats.Size(m)
}
func (m *OperationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationStats.DiscardUnknown(m)
}

var xxx_messageInfo_OperationStats proto.InternalMessageInfo

func (m *OperationStats) GetRows() uint64 {
	if m != nil {
		return m.Rows
	}
	return 0
}

func (m *OperationStats) GetBytes() uint64 {
	if m != nil {
		return m.Bytes
	}
	return 0
}

// Describes all operations on a table
type TableAccessStats struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Reads                *OperationStats `protobuf:"bytes,3,opt,name=reads,proto3" json:"reads,omitempty"`
	Updates              *OperationStats `protobuf:"bytes,4,opt,name=updates,proto3" json:"updates,omitempty"`
	Deletes              *OperationStats `protobuf:"bytes,5,opt,name=deletes,proto3" json:"deletes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *TableAccessStats) Reset()         { *m = TableAccessStats{} }
func (m *TableAccessStats) String() string { return proto.CompactTextString(m) }
func (*TableAccessStats) ProtoMessage()    {}
func (*TableAccessStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{1}
}

func (m *TableAccessStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableAccessStats.Unmarshal(m, b)
}
func (m *TableAccessStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableAccessStats.Marshal(b, m, deterministic)
}
func (m *TableAccessStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableAccessStats.Merge(m, src)
}
func (m *TableAccessStats) XXX_Size() int {
	return xxx_messageInfo_TableAccessStats.Size(m)
}
func (m *TableAccessStats) XXX_DiscardUnknown() {
	xxx_messageInfo_TableAccessStats.DiscardUnknown(m)
}

var xxx_messageInfo_TableAccessStats proto.InternalMessageInfo

func (m *TableAccessStats) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TableAccessStats) GetReads() *OperationStats {
	if m != nil {
		return m.Reads
	}
	return nil
}

func (m *TableAccessStats) GetUpdates() *OperationStats {
	if m != nil {
		return m.Updates
	}
	return nil
}

func (m *TableAccessStats) GetDeletes() *OperationStats {
	if m != nil {
		return m.Deletes
	}
	return nil
}

type QueryPhaseStats struct {
	DurationUs           uint64              `protobuf:"varint,1,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	TableAccess          []*TableAccessStats `protobuf:"bytes,2,rep,name=table_access,json=tableAccess,proto3" json:"table_access,omitempty"`
	CpuTimeUs            uint64              `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *QueryPhaseStats) Reset()         { *m = QueryPhaseStats{} }
func (m *QueryPhaseStats) String() string { return proto.CompactTextString(m) }
func (*QueryPhaseStats) ProtoMessage()    {}
func (*QueryPhaseStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{2}
}

func (m *QueryPhaseStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryPhaseStats.Unmarshal(m, b)
}
func (m *QueryPhaseStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryPhaseStats.Marshal(b, m, deterministic)
}
func (m *QueryPhaseStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryPhaseStats.Merge(m, src)
}
func (m *QueryPhaseStats) XXX_Size() int {
	return xxx_messageInfo_QueryPhaseStats.Size(m)
}
func (m *QueryPhaseStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryPhaseStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryPhaseStats proto.InternalMessageInfo

func (m *QueryPhaseStats) GetDurationUs() uint64 {
	if m != nil {
		return m.DurationUs
	}
	return 0
}

func (m *QueryPhaseStats) GetTableAccess() []*TableAccessStats {
	if m != nil {
		return m.TableAccess
	}
	return nil
}

func (m *QueryPhaseStats) GetCpuTimeUs() uint64 {
	if m != nil {
		return m.CpuTimeUs
	}
	return 0
}

type CompilationStats struct {
	FromCache            bool     `protobuf:"varint,1,opt,name=from_cache,json=fromCache,proto3" json:"from_cache,omitempty"`
	DurationUs           uint64   `protobuf:"varint,2,opt,name=duration_us,json=durationUs,proto3" json:"duration_us,omitempty"`
	CpuTimeUs            uint64   `protobuf:"varint,3,opt,name=cpu_time_us,json=cpuTimeUs,proto3" json:"cpu_time_us,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CompilationStats) Reset()         { *m = CompilationStats{} }
func (m *CompilationStats) String() string { return proto.CompactTextString(m) }
func (*CompilationStats) ProtoMessage()    {}
func (*CompilationStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{3}
}

func (m *CompilationStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CompilationStats.Unmarshal(m, b)
}
func (m *CompilationStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CompilationStats.Marshal(b, m, deterministic)
}
func (m *CompilationStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CompilationStats.Merge(m, src)
}
func (m *CompilationStats) XXX_Size() int {
	return xxx_messageInfo_CompilationStats.Size(m)
}
func (m *CompilationStats) XXX_DiscardUnknown() {
	xxx_messageInfo_CompilationStats.DiscardUnknown(m)
}

var xxx_messageInfo_CompilationStats proto.InternalMessageInfo

func (m *CompilationStats) GetFromCache() bool {
	if m != nil {
		return m.FromCache
	}
	return false
}

func (m *CompilationStats) GetDurationUs() uint64 {
	if m != nil {
		return m.DurationUs
	}
	return 0
}

func (m *CompilationStats) GetCpuTimeUs() uint64 {
	if m != nil {
		return m.CpuTimeUs
	}
	return 0
}

type QueryStats struct {
	// A query might have one or more execution phases
	QueryPhases          []*QueryPhaseStats `protobuf:"bytes,1,rep,name=query_phases,json=queryPhases,proto3" json:"query_phases,omitempty"`
	Compilation          *CompilationStats  `protobuf:"bytes,2,opt,name=compilation,proto3" json:"compilation,omitempty"`
	ProcessCpuTimeUs     uint64             `protobuf:"varint,3,opt,name=process_cpu_time_us,json=processCpuTimeUs,proto3" json:"process_cpu_time_us,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *QueryStats) Reset()         { *m = QueryStats{} }
func (m *QueryStats) String() string { return proto.CompactTextString(m) }
func (*QueryStats) ProtoMessage()    {}
func (*QueryStats) Descriptor() ([]byte, []int) {
	return fileDescriptor_bd6647573551cb14, []int{4}
}

func (m *QueryStats) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryStats.Unmarshal(m, b)
}
func (m *QueryStats) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryStats.Marshal(b, m, deterministic)
}
func (m *QueryStats) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryStats.Merge(m, src)
}
func (m *QueryStats) XXX_Size() int {
	return xxx_messageInfo_QueryStats.Size(m)
}
func (m *QueryStats) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryStats.DiscardUnknown(m)
}

var xxx_messageInfo_QueryStats proto.InternalMessageInfo

func (m *QueryStats) GetQueryPhases() []*QueryPhaseStats {
	if m != nil {
		return m.QueryPhases
	}
	return nil
}

func (m *QueryStats) GetCompilation() *CompilationStats {
	if m != nil {
		return m.Compilation
	}
	return nil
}

func (m *QueryStats) GetProcessCpuTimeUs() uint64 {
	if m != nil {
		return m.ProcessCpuTimeUs
	}
	return 0
}

func init() {
	proto.RegisterType((*OperationStats)(nil), "Ydb.TableStats.OperationStats")
	proto.RegisterType((*TableAccessStats)(nil), "Ydb.TableStats.TableAccessStats")
	proto.RegisterType((*QueryPhaseStats)(nil), "Ydb.TableStats.QueryPhaseStats")
	proto.RegisterType((*CompilationStats)(nil), "Ydb.TableStats.CompilationStats")
	proto.RegisterType((*QueryStats)(nil), "Ydb.TableStats.QueryStats")
}

func init() { proto.RegisterFile("ydb_query_stats.proto", fileDescriptor_bd6647573551cb14) }

var fileDescriptor_bd6647573551cb14 = []byte{
	// 398 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x93, 0xbf, 0xae, 0xd3, 0x30,
	0x18, 0xc5, 0xe5, 0x24, 0x85, 0xdb, 0x2f, 0x57, 0x25, 0xf2, 0x05, 0x29, 0x0b, 0xf7, 0x46, 0x99,
	0xba, 0x90, 0xe1, 0xc2, 0x80, 0xd8, 0x68, 0x36, 0x16, 0x20, 0xb4, 0x03, 0x53, 0xe4, 0xd8, 0x46,
	0x8d, 0xd4, 0xd4, 0xae, 0xed, 0x08, 0xf2, 0x24, 0x3c, 0x10, 0xef, 0xc0, 0xb3, 0x30, 0x22, 0x3b,
	0x4d, 0xff, 0xa4, 0x48, 0x74, 0x73, 0x4f, 0xbf, 0xe3, 0x73, 0xbe, 0x5f, 0x12, 0x78, 0xd1, 0xb1,
	0xaa, 0xdc, 0xb5, 0x5c, 0x75, 0xa5, 0x36, 0xc4, 0xe8, 0x4c, 0x2a, 0x61, 0x04, 0x9e, 0x7d, 0x65,
	0x55, 0xb6, 0x24, 0xd5, 0x86, 0x7f, 0xb1, 0x6a, 0xfa, 0x0e, 0x66, 0x1f, 0x25, 0x57, 0xc4, 0xd4,
	0x62, 0xeb, 0x14, 0x8c, 0x21, 0x50, 0xe2, 0xbb, 0x8e, 0x51, 0x82, 0xe6, 0x41, 0xe1, 0xce, 0xf8,
	0x39, 0x4c, 0xaa, 0xce, 0x70, 0x1d, 0x7b, 0x4e, 0xec, 0x7f, 0xa4, 0xbf, 0x11, 0x44, 0xee, 0xaa,
	0xf7, 0x94, 0x72, 0xad, 0x0f, 0xf6, 0x2d, 0x69, 0xb8, 0xb3, 0x4f, 0x0b, 0x77, 0xc6, 0x6f, 0x60,
	0xa2, 0x38, 0x61, 0x3a, 0xf6, 0x13, 0x34, 0x0f, 0x1f, 0xef, 0xb3, 0xf3, 0x12, 0xd9, 0x79, 0x83,
	0xa2, 0x1f, 0xc6, 0x6f, 0xe1, 0x69, 0x2b, 0x19, 0xb1, 0xb1, 0xc1, 0x55, 0xbe, 0x61, 0xdc, 0x3a,
	0x19, 0xdf, 0x70, 0xeb, 0x9c, 0x5c, 0xe7, 0xdc, 0x8f, 0x7f, 0x08, 0x6e, 0xbc, 0xc8, 0x4f, 0x7f,
	0x22, 0x78, 0xf6, 0xd9, 0xa2, 0xfb, 0xb4, 0x26, 0xba, 0x77, 0xe0, 0x07, 0x08, 0x59, 0xdb, 0x7b,
	0xca, 0x76, 0xa0, 0x03, 0x83, 0xb4, 0xd2, 0x38, 0x87, 0x5b, 0x63, 0x03, 0x4a, 0xe2, 0x68, 0xc4,
	0x5e, 0xe2, 0xcf, 0xc3, 0xc7, 0x64, 0x9c, 0x3c, 0x06, 0x56, 0x84, 0xe6, 0xa8, 0xe0, 0x7b, 0x08,
	0xa9, 0x6c, 0x4b, 0x53, 0x37, 0xdc, 0xa6, 0xf8, 0x2e, 0x65, 0x4a, 0x65, 0xbb, 0xac, 0x1b, 0xbe,
	0xd2, 0xa9, 0x82, 0x28, 0x17, 0x8d, 0xac, 0x37, 0x27, 0x0f, 0xec, 0x25, 0xc0, 0x37, 0x25, 0x9a,
	0x92, 0x12, 0xba, 0xee, 0xb9, 0xdf, 0x14, 0x53, 0xab, 0xe4, 0x56, 0x18, 0x17, 0xf7, 0x2e, 0x8a,
	0xff, 0x2f, 0xf3, 0x17, 0x02, 0x70, 0x34, 0xfa, 0xb8, 0x05, 0xdc, 0xf6, 0xaf, 0x95, 0xb4, 0x70,
	0x2c, 0x09, 0xbb, 0xe7, 0xc3, 0x78, 0xcf, 0x11, 0xbf, 0x22, 0xdc, 0x1d, 0x04, 0x7b, 0x47, 0x48,
	0x8f, 0x6b, 0xb8, 0x4e, 0xff, 0x40, 0x35, 0xde, 0xb4, 0x38, 0x35, 0xe1, 0x57, 0x70, 0x27, 0x95,
	0xb0, 0xd4, 0xca, 0xcb, 0xfa, 0xd1, 0xfe, 0xaf, 0x7c, 0xd8, 0x62, 0x71, 0x07, 0x33, 0x2a, 0x9a,
	0xac, 0x23, 0x5b, 0xc6, 0x7f, 0x64, 0x1d, 0xab, 0xfe, 0x20, 0x54, 0x3d, 0x71, 0x1f, 0xc5, 0xeb,
	0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7c, 0xed, 0x33, 0xba, 0x2d, 0x03, 0x00, 0x00,
}

const ()
