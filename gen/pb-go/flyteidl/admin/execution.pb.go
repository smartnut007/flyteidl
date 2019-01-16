// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/admin/execution.proto

package admin // import "github.com/lyft/flyteidl/gen/pb-go/flyteidl/admin"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import duration "github.com/golang/protobuf/ptypes/duration"
import timestamp "github.com/golang/protobuf/ptypes/timestamp"
import core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExecutionMetadata_ExecutionMode int32

const (
	ExecutionMetadata_MANUAL    ExecutionMetadata_ExecutionMode = 0
	ExecutionMetadata_SCHEDULED ExecutionMetadata_ExecutionMode = 1
	ExecutionMetadata_SYSTEM    ExecutionMetadata_ExecutionMode = 2
	ExecutionMetadata_RELAUNCH  ExecutionMetadata_ExecutionMode = 3
)

var ExecutionMetadata_ExecutionMode_name = map[int32]string{
	0: "MANUAL",
	1: "SCHEDULED",
	2: "SYSTEM",
	3: "RELAUNCH",
}
var ExecutionMetadata_ExecutionMode_value = map[string]int32{
	"MANUAL":    0,
	"SCHEDULED": 1,
	"SYSTEM":    2,
	"RELAUNCH":  3,
}

func (x ExecutionMetadata_ExecutionMode) String() string {
	return proto.EnumName(ExecutionMetadata_ExecutionMode_name, int32(x))
}
func (ExecutionMetadata_ExecutionMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{8, 0}
}

type ExecutionCreateRequest struct {
	// Name of the project the execution belongs to.
	Project string `protobuf:"bytes,1,opt,name=project,proto3" json:"project,omitempty"`
	// Name of the domain the execution belongs to.
	// A domain can be considered as a subset within a specific project.
	Domain string `protobuf:"bytes,2,opt,name=domain,proto3" json:"domain,omitempty"`
	// User provided value for the resource.
	// If none is provided the system will generate a unique string.
	// +optional
	Name                 string         `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Spec                 *ExecutionSpec `protobuf:"bytes,4,opt,name=spec,proto3" json:"spec,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecutionCreateRequest) Reset()         { *m = ExecutionCreateRequest{} }
func (m *ExecutionCreateRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionCreateRequest) ProtoMessage()    {}
func (*ExecutionCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{0}
}
func (m *ExecutionCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionCreateRequest.Unmarshal(m, b)
}
func (m *ExecutionCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionCreateRequest.Marshal(b, m, deterministic)
}
func (dst *ExecutionCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionCreateRequest.Merge(dst, src)
}
func (m *ExecutionCreateRequest) XXX_Size() int {
	return xxx_messageInfo_ExecutionCreateRequest.Size(m)
}
func (m *ExecutionCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionCreateRequest proto.InternalMessageInfo

func (m *ExecutionCreateRequest) GetProject() string {
	if m != nil {
		return m.Project
	}
	return ""
}

func (m *ExecutionCreateRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *ExecutionCreateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExecutionCreateRequest) GetSpec() *ExecutionSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

type ExecutionRelaunchRequest struct {
	// Identifier of the workflow execution to relaunch.
	Id *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// User provided value for the relaunched execution.
	// If none is provided the system will generate a unique string.
	// +optional
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionRelaunchRequest) Reset()         { *m = ExecutionRelaunchRequest{} }
func (m *ExecutionRelaunchRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionRelaunchRequest) ProtoMessage()    {}
func (*ExecutionRelaunchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{1}
}
func (m *ExecutionRelaunchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionRelaunchRequest.Unmarshal(m, b)
}
func (m *ExecutionRelaunchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionRelaunchRequest.Marshal(b, m, deterministic)
}
func (dst *ExecutionRelaunchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionRelaunchRequest.Merge(dst, src)
}
func (m *ExecutionRelaunchRequest) XXX_Size() int {
	return xxx_messageInfo_ExecutionRelaunchRequest.Size(m)
}
func (m *ExecutionRelaunchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionRelaunchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionRelaunchRequest proto.InternalMessageInfo

func (m *ExecutionRelaunchRequest) GetId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ExecutionRelaunchRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type ExecutionCreateResponse struct {
	Id                   *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *ExecutionCreateResponse) Reset()         { *m = ExecutionCreateResponse{} }
func (m *ExecutionCreateResponse) String() string { return proto.CompactTextString(m) }
func (*ExecutionCreateResponse) ProtoMessage()    {}
func (*ExecutionCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{2}
}
func (m *ExecutionCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionCreateResponse.Unmarshal(m, b)
}
func (m *ExecutionCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionCreateResponse.Marshal(b, m, deterministic)
}
func (dst *ExecutionCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionCreateResponse.Merge(dst, src)
}
func (m *ExecutionCreateResponse) XXX_Size() int {
	return xxx_messageInfo_ExecutionCreateResponse.Size(m)
}
func (m *ExecutionCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionCreateResponse proto.InternalMessageInfo

func (m *ExecutionCreateResponse) GetId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

// A message used to fetch a single workflow execution entity.
type WorkflowExecutionGetRequest struct {
	// Uniquely identifies an individual workflow execution.
	Id                   *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowExecutionGetRequest) Reset()         { *m = WorkflowExecutionGetRequest{} }
func (m *WorkflowExecutionGetRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionGetRequest) ProtoMessage()    {}
func (*WorkflowExecutionGetRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{3}
}
func (m *WorkflowExecutionGetRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionGetRequest.Unmarshal(m, b)
}
func (m *WorkflowExecutionGetRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionGetRequest.Marshal(b, m, deterministic)
}
func (dst *WorkflowExecutionGetRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionGetRequest.Merge(dst, src)
}
func (m *WorkflowExecutionGetRequest) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionGetRequest.Size(m)
}
func (m *WorkflowExecutionGetRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionGetRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionGetRequest proto.InternalMessageInfo

func (m *WorkflowExecutionGetRequest) GetId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

type Execution struct {
	Id                   *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Spec                 *ExecutionSpec                    `protobuf:"bytes,2,opt,name=spec,proto3" json:"spec,omitempty"`
	Closure              *ExecutionClosure                 `protobuf:"bytes,3,opt,name=closure,proto3" json:"closure,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *Execution) Reset()         { *m = Execution{} }
func (m *Execution) String() string { return proto.CompactTextString(m) }
func (*Execution) ProtoMessage()    {}
func (*Execution) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{4}
}
func (m *Execution) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Execution.Unmarshal(m, b)
}
func (m *Execution) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Execution.Marshal(b, m, deterministic)
}
func (dst *Execution) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Execution.Merge(dst, src)
}
func (m *Execution) XXX_Size() int {
	return xxx_messageInfo_Execution.Size(m)
}
func (m *Execution) XXX_DiscardUnknown() {
	xxx_messageInfo_Execution.DiscardUnknown(m)
}

var xxx_messageInfo_Execution proto.InternalMessageInfo

func (m *Execution) GetId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *Execution) GetSpec() *ExecutionSpec {
	if m != nil {
		return m.Spec
	}
	return nil
}

func (m *Execution) GetClosure() *ExecutionClosure {
	if m != nil {
		return m.Closure
	}
	return nil
}

type ExecutionList struct {
	Executions []*Execution `protobuf:"bytes,1,rep,name=executions,proto3" json:"executions,omitempty"`
	// In the case of multiple pages of results, the server-provided token can be used to fetch the next page
	// in a query. If there are no more results, this value will be empty.
	Token                string   `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionList) Reset()         { *m = ExecutionList{} }
func (m *ExecutionList) String() string { return proto.CompactTextString(m) }
func (*ExecutionList) ProtoMessage()    {}
func (*ExecutionList) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{5}
}
func (m *ExecutionList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionList.Unmarshal(m, b)
}
func (m *ExecutionList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionList.Marshal(b, m, deterministic)
}
func (dst *ExecutionList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionList.Merge(dst, src)
}
func (m *ExecutionList) XXX_Size() int {
	return xxx_messageInfo_ExecutionList.Size(m)
}
func (m *ExecutionList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionList.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionList proto.InternalMessageInfo

func (m *ExecutionList) GetExecutions() []*Execution {
	if m != nil {
		return m.Executions
	}
	return nil
}

func (m *ExecutionList) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

// Input/output data can represented by actual values or a link to where it is stored
type LiteralMapBlob struct {
	// Types that are valid to be assigned to Data:
	//	*LiteralMapBlob_Values
	//	*LiteralMapBlob_Uri
	Data                 isLiteralMapBlob_Data `protobuf_oneof:"data"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *LiteralMapBlob) Reset()         { *m = LiteralMapBlob{} }
func (m *LiteralMapBlob) String() string { return proto.CompactTextString(m) }
func (*LiteralMapBlob) ProtoMessage()    {}
func (*LiteralMapBlob) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{6}
}
func (m *LiteralMapBlob) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiteralMapBlob.Unmarshal(m, b)
}
func (m *LiteralMapBlob) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiteralMapBlob.Marshal(b, m, deterministic)
}
func (dst *LiteralMapBlob) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiteralMapBlob.Merge(dst, src)
}
func (m *LiteralMapBlob) XXX_Size() int {
	return xxx_messageInfo_LiteralMapBlob.Size(m)
}
func (m *LiteralMapBlob) XXX_DiscardUnknown() {
	xxx_messageInfo_LiteralMapBlob.DiscardUnknown(m)
}

var xxx_messageInfo_LiteralMapBlob proto.InternalMessageInfo

type isLiteralMapBlob_Data interface {
	isLiteralMapBlob_Data()
}

type LiteralMapBlob_Values struct {
	Values *core.LiteralMap `protobuf:"bytes,1,opt,name=values,proto3,oneof"`
}

type LiteralMapBlob_Uri struct {
	Uri string `protobuf:"bytes,2,opt,name=uri,proto3,oneof"`
}

func (*LiteralMapBlob_Values) isLiteralMapBlob_Data() {}

func (*LiteralMapBlob_Uri) isLiteralMapBlob_Data() {}

func (m *LiteralMapBlob) GetData() isLiteralMapBlob_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *LiteralMapBlob) GetValues() *core.LiteralMap {
	if x, ok := m.GetData().(*LiteralMapBlob_Values); ok {
		return x.Values
	}
	return nil
}

func (m *LiteralMapBlob) GetUri() string {
	if x, ok := m.GetData().(*LiteralMapBlob_Uri); ok {
		return x.Uri
	}
	return ""
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*LiteralMapBlob) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _LiteralMapBlob_OneofMarshaler, _LiteralMapBlob_OneofUnmarshaler, _LiteralMapBlob_OneofSizer, []interface{}{
		(*LiteralMapBlob_Values)(nil),
		(*LiteralMapBlob_Uri)(nil),
	}
}

func _LiteralMapBlob_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*LiteralMapBlob)
	// data
	switch x := m.Data.(type) {
	case *LiteralMapBlob_Values:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Values); err != nil {
			return err
		}
	case *LiteralMapBlob_Uri:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.Uri)
	case nil:
	default:
		return fmt.Errorf("LiteralMapBlob.Data has unexpected type %T", x)
	}
	return nil
}

func _LiteralMapBlob_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*LiteralMapBlob)
	switch tag {
	case 1: // data.values
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.LiteralMap)
		err := b.DecodeMessage(msg)
		m.Data = &LiteralMapBlob_Values{msg}
		return true, err
	case 2: // data.uri
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.Data = &LiteralMapBlob_Uri{x}
		return true, err
	default:
		return false, nil
	}
}

func _LiteralMapBlob_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*LiteralMapBlob)
	// data
	switch x := m.Data.(type) {
	case *LiteralMapBlob_Values:
		s := proto.Size(x.Values)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *LiteralMapBlob_Uri:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.Uri)))
		n += len(x.Uri)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

// Encapsulates the results of the Execution
type ExecutionClosure struct {
	// Types that are valid to be assigned to OutputResult:
	//	*ExecutionClosure_Outputs
	//	*ExecutionClosure_Error
	//	*ExecutionClosure_AbortCause
	OutputResult isExecutionClosure_OutputResult `protobuf_oneof:"output_result"`
	// Inputs computed and passed for execution.
	// computed_inputs depends on inputs in ExecutionSpec, fixed and default inputs in launch plan
	ComputedInputs *core.LiteralMap `protobuf:"bytes,3,opt,name=computed_inputs,json=computedInputs,proto3" json:"computed_inputs,omitempty"`
	// Phase of the executions
	Phase core.WorkflowExecutionPhase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecutionPhase" json:"phase,omitempty"`
	// Time at which the execution began running.
	StartedAt *timestamp.Timestamp `protobuf:"bytes,5,opt,name=started_at,json=startedAt,proto3" json:"started_at,omitempty"`
	// The amount of time the execution spent running.
	Duration *duration.Duration `protobuf:"bytes,6,opt,name=duration,proto3" json:"duration,omitempty"`
	// Time at which the execution was created.
	CreatedAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Time at which the execution was last updated.
	UpdatedAt *timestamp.Timestamp `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	// The notification settings to use after merging the CreateExecutionRequest and the launch plan
	// notification settings.
	Notifications []*Notification `protobuf:"bytes,9,rep,name=notifications,proto3" json:"notifications,omitempty"`
	// Identifies the workflow definition for this execution.
	WorkflowId           *core.Identifier `protobuf:"bytes,11,opt,name=workflow_id,json=workflowId,proto3" json:"workflow_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *ExecutionClosure) Reset()         { *m = ExecutionClosure{} }
func (m *ExecutionClosure) String() string { return proto.CompactTextString(m) }
func (*ExecutionClosure) ProtoMessage()    {}
func (*ExecutionClosure) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{7}
}
func (m *ExecutionClosure) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionClosure.Unmarshal(m, b)
}
func (m *ExecutionClosure) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionClosure.Marshal(b, m, deterministic)
}
func (dst *ExecutionClosure) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionClosure.Merge(dst, src)
}
func (m *ExecutionClosure) XXX_Size() int {
	return xxx_messageInfo_ExecutionClosure.Size(m)
}
func (m *ExecutionClosure) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionClosure.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionClosure proto.InternalMessageInfo

type isExecutionClosure_OutputResult interface {
	isExecutionClosure_OutputResult()
}

type ExecutionClosure_Outputs struct {
	Outputs *LiteralMapBlob `protobuf:"bytes,1,opt,name=outputs,proto3,oneof"`
}

type ExecutionClosure_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,2,opt,name=error,proto3,oneof"`
}

type ExecutionClosure_AbortCause struct {
	AbortCause string `protobuf:"bytes,10,opt,name=abort_cause,json=abortCause,proto3,oneof"`
}

func (*ExecutionClosure_Outputs) isExecutionClosure_OutputResult() {}

func (*ExecutionClosure_Error) isExecutionClosure_OutputResult() {}

func (*ExecutionClosure_AbortCause) isExecutionClosure_OutputResult() {}

func (m *ExecutionClosure) GetOutputResult() isExecutionClosure_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *ExecutionClosure) GetOutputs() *LiteralMapBlob {
	if x, ok := m.GetOutputResult().(*ExecutionClosure_Outputs); ok {
		return x.Outputs
	}
	return nil
}

func (m *ExecutionClosure) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*ExecutionClosure_Error); ok {
		return x.Error
	}
	return nil
}

func (m *ExecutionClosure) GetAbortCause() string {
	if x, ok := m.GetOutputResult().(*ExecutionClosure_AbortCause); ok {
		return x.AbortCause
	}
	return ""
}

func (m *ExecutionClosure) GetComputedInputs() *core.LiteralMap {
	if m != nil {
		return m.ComputedInputs
	}
	return nil
}

func (m *ExecutionClosure) GetPhase() core.WorkflowExecutionPhase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecutionPhase_WORKFLOW_PHASE_UNDEFINED
}

func (m *ExecutionClosure) GetStartedAt() *timestamp.Timestamp {
	if m != nil {
		return m.StartedAt
	}
	return nil
}

func (m *ExecutionClosure) GetDuration() *duration.Duration {
	if m != nil {
		return m.Duration
	}
	return nil
}

func (m *ExecutionClosure) GetCreatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.CreatedAt
	}
	return nil
}

func (m *ExecutionClosure) GetUpdatedAt() *timestamp.Timestamp {
	if m != nil {
		return m.UpdatedAt
	}
	return nil
}

func (m *ExecutionClosure) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

func (m *ExecutionClosure) GetWorkflowId() *core.Identifier {
	if m != nil {
		return m.WorkflowId
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*ExecutionClosure) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _ExecutionClosure_OneofMarshaler, _ExecutionClosure_OneofUnmarshaler, _ExecutionClosure_OneofSizer, []interface{}{
		(*ExecutionClosure_Outputs)(nil),
		(*ExecutionClosure_Error)(nil),
		(*ExecutionClosure_AbortCause)(nil),
	}
}

func _ExecutionClosure_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*ExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *ExecutionClosure_Outputs:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Outputs); err != nil {
			return err
		}
	case *ExecutionClosure_Error:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Error); err != nil {
			return err
		}
	case *ExecutionClosure_AbortCause:
		b.EncodeVarint(10<<3 | proto.WireBytes)
		b.EncodeStringBytes(x.AbortCause)
	case nil:
	default:
		return fmt.Errorf("ExecutionClosure.OutputResult has unexpected type %T", x)
	}
	return nil
}

func _ExecutionClosure_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*ExecutionClosure)
	switch tag {
	case 1: // output_result.outputs
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(LiteralMapBlob)
		err := b.DecodeMessage(msg)
		m.OutputResult = &ExecutionClosure_Outputs{msg}
		return true, err
	case 2: // output_result.error
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(core.ExecutionError)
		err := b.DecodeMessage(msg)
		m.OutputResult = &ExecutionClosure_Error{msg}
		return true, err
	case 10: // output_result.abort_cause
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		x, err := b.DecodeStringBytes()
		m.OutputResult = &ExecutionClosure_AbortCause{x}
		return true, err
	default:
		return false, nil
	}
}

func _ExecutionClosure_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*ExecutionClosure)
	// output_result
	switch x := m.OutputResult.(type) {
	case *ExecutionClosure_Outputs:
		s := proto.Size(x.Outputs)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExecutionClosure_Error:
		s := proto.Size(x.Error)
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(s))
		n += s
	case *ExecutionClosure_AbortCause:
		n += 1 // tag and wire
		n += proto.SizeVarint(uint64(len(x.AbortCause)))
		n += len(x.AbortCause)
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type ExecutionMetadata struct {
	Mode ExecutionMetadata_ExecutionMode `protobuf:"varint,1,opt,name=mode,proto3,enum=flyteidl.admin.ExecutionMetadata_ExecutionMode" json:"mode,omitempty"`
	// Identifier of the entity that triggered this execution.
	Principal string `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	// Indicates the "nestedness" of this execution.
	// If a user launches a workflow execution, the default nesting is 0.
	// If this execution further launches a workflow (child workflow), the nesting level is incremented by 0 => 1
	// Generally, if workflow at nesting level k launches a workflow then the child workflow will have
	// nesting = k + 1.
	Nesting              int32    `protobuf:"varint,3,opt,name=nesting,proto3" json:"nesting,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionMetadata) Reset()         { *m = ExecutionMetadata{} }
func (m *ExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ExecutionMetadata) ProtoMessage()    {}
func (*ExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{8}
}
func (m *ExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionMetadata.Unmarshal(m, b)
}
func (m *ExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionMetadata.Marshal(b, m, deterministic)
}
func (dst *ExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionMetadata.Merge(dst, src)
}
func (m *ExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ExecutionMetadata.Size(m)
}
func (m *ExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionMetadata proto.InternalMessageInfo

func (m *ExecutionMetadata) GetMode() ExecutionMetadata_ExecutionMode {
	if m != nil {
		return m.Mode
	}
	return ExecutionMetadata_MANUAL
}

func (m *ExecutionMetadata) GetPrincipal() string {
	if m != nil {
		return m.Principal
	}
	return ""
}

func (m *ExecutionMetadata) GetNesting() int32 {
	if m != nil {
		return m.Nesting
	}
	return 0
}

type ExecutionSpec struct {
	// Launch plan to be executed
	LaunchPlan *core.Identifier `protobuf:"bytes,1,opt,name=launch_plan,json=launchPlan,proto3" json:"launch_plan,omitempty"`
	// Input values to be passed for the execution
	Inputs *core.LiteralMap `protobuf:"bytes,2,opt,name=inputs,proto3" json:"inputs,omitempty"`
	// Metadata for the execution
	Metadata *ExecutionMetadata `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// List of notifications based on Execution status transitions
	Notifications        []*Notification `protobuf:"bytes,4,rep,name=notifications,proto3" json:"notifications,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *ExecutionSpec) Reset()         { *m = ExecutionSpec{} }
func (m *ExecutionSpec) String() string { return proto.CompactTextString(m) }
func (*ExecutionSpec) ProtoMessage()    {}
func (*ExecutionSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{9}
}
func (m *ExecutionSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionSpec.Unmarshal(m, b)
}
func (m *ExecutionSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionSpec.Marshal(b, m, deterministic)
}
func (dst *ExecutionSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionSpec.Merge(dst, src)
}
func (m *ExecutionSpec) XXX_Size() int {
	return xxx_messageInfo_ExecutionSpec.Size(m)
}
func (m *ExecutionSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionSpec proto.InternalMessageInfo

func (m *ExecutionSpec) GetLaunchPlan() *core.Identifier {
	if m != nil {
		return m.LaunchPlan
	}
	return nil
}

func (m *ExecutionSpec) GetInputs() *core.LiteralMap {
	if m != nil {
		return m.Inputs
	}
	return nil
}

func (m *ExecutionSpec) GetMetadata() *ExecutionMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *ExecutionSpec) GetNotifications() []*Notification {
	if m != nil {
		return m.Notifications
	}
	return nil
}

type ExecutionTerminateRequest struct {
	// Uniquely identifies the individual workflow execution to be terminated.
	Id *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Optional reason for aborting.
	Cause                string   `protobuf:"bytes,2,opt,name=cause,proto3" json:"cause,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionTerminateRequest) Reset()         { *m = ExecutionTerminateRequest{} }
func (m *ExecutionTerminateRequest) String() string { return proto.CompactTextString(m) }
func (*ExecutionTerminateRequest) ProtoMessage()    {}
func (*ExecutionTerminateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{10}
}
func (m *ExecutionTerminateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionTerminateRequest.Unmarshal(m, b)
}
func (m *ExecutionTerminateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionTerminateRequest.Marshal(b, m, deterministic)
}
func (dst *ExecutionTerminateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionTerminateRequest.Merge(dst, src)
}
func (m *ExecutionTerminateRequest) XXX_Size() int {
	return xxx_messageInfo_ExecutionTerminateRequest.Size(m)
}
func (m *ExecutionTerminateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionTerminateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionTerminateRequest proto.InternalMessageInfo

func (m *ExecutionTerminateRequest) GetId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *ExecutionTerminateRequest) GetCause() string {
	if m != nil {
		return m.Cause
	}
	return ""
}

type ExecutionTerminateResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecutionTerminateResponse) Reset()         { *m = ExecutionTerminateResponse{} }
func (m *ExecutionTerminateResponse) String() string { return proto.CompactTextString(m) }
func (*ExecutionTerminateResponse) ProtoMessage()    {}
func (*ExecutionTerminateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_execution_65768c2d6295fb4f, []int{11}
}
func (m *ExecutionTerminateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecutionTerminateResponse.Unmarshal(m, b)
}
func (m *ExecutionTerminateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecutionTerminateResponse.Marshal(b, m, deterministic)
}
func (dst *ExecutionTerminateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecutionTerminateResponse.Merge(dst, src)
}
func (m *ExecutionTerminateResponse) XXX_Size() int {
	return xxx_messageInfo_ExecutionTerminateResponse.Size(m)
}
func (m *ExecutionTerminateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecutionTerminateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecutionTerminateResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ExecutionCreateRequest)(nil), "flyteidl.admin.ExecutionCreateRequest")
	proto.RegisterType((*ExecutionRelaunchRequest)(nil), "flyteidl.admin.ExecutionRelaunchRequest")
	proto.RegisterType((*ExecutionCreateResponse)(nil), "flyteidl.admin.ExecutionCreateResponse")
	proto.RegisterType((*WorkflowExecutionGetRequest)(nil), "flyteidl.admin.WorkflowExecutionGetRequest")
	proto.RegisterType((*Execution)(nil), "flyteidl.admin.Execution")
	proto.RegisterType((*ExecutionList)(nil), "flyteidl.admin.ExecutionList")
	proto.RegisterType((*LiteralMapBlob)(nil), "flyteidl.admin.LiteralMapBlob")
	proto.RegisterType((*ExecutionClosure)(nil), "flyteidl.admin.ExecutionClosure")
	proto.RegisterType((*ExecutionMetadata)(nil), "flyteidl.admin.ExecutionMetadata")
	proto.RegisterType((*ExecutionSpec)(nil), "flyteidl.admin.ExecutionSpec")
	proto.RegisterType((*ExecutionTerminateRequest)(nil), "flyteidl.admin.ExecutionTerminateRequest")
	proto.RegisterType((*ExecutionTerminateResponse)(nil), "flyteidl.admin.ExecutionTerminateResponse")
	proto.RegisterEnum("flyteidl.admin.ExecutionMetadata_ExecutionMode", ExecutionMetadata_ExecutionMode_name, ExecutionMetadata_ExecutionMode_value)
}

func init() {
	proto.RegisterFile("flyteidl/admin/execution.proto", fileDescriptor_execution_65768c2d6295fb4f)
}

var fileDescriptor_execution_65768c2d6295fb4f = []byte{
	// 911 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdd, 0x6e, 0xdb, 0x36,
	0x14, 0xf6, 0x7f, 0xe2, 0xe3, 0xc5, 0xf5, 0x88, 0xa2, 0x53, 0xdc, 0x34, 0x4b, 0x05, 0x0c, 0x08,
	0x06, 0xcc, 0x42, 0x1c, 0xf4, 0x62, 0x19, 0x76, 0x61, 0x3b, 0xc6, 0x1c, 0xc0, 0x09, 0x0a, 0x3a,
	0xc1, 0xd0, 0xdd, 0x78, 0xb4, 0x44, 0x3b, 0x6c, 0x25, 0x51, 0xa3, 0xa8, 0x75, 0xbd, 0xdc, 0x23,
	0xec, 0x15, 0xf6, 0x10, 0x7b, 0x94, 0x3d, 0xcf, 0x20, 0x8a, 0x54, 0x64, 0x27, 0x81, 0x37, 0x64,
	0x77, 0x3a, 0x3c, 0xdf, 0xc7, 0x73, 0x78, 0x78, 0xbe, 0x43, 0xc1, 0xe1, 0xd2, 0xff, 0x24, 0x29,
	0xf3, 0x7c, 0x87, 0x78, 0x01, 0x0b, 0x1d, 0xfa, 0x1b, 0x75, 0x13, 0xc9, 0x78, 0xd8, 0x8b, 0x04,
	0x97, 0x1c, 0xb5, 0x8d, 0xbf, 0xa7, 0xfc, 0xdd, 0x97, 0x1b, 0x78, 0x97, 0x07, 0x81, 0x01, 0x77,
	0x0f, 0x72, 0xa7, 0xcb, 0x05, 0x75, 0x7c, 0x26, 0xa9, 0x20, 0x7e, 0xac, 0xbd, 0xaf, 0xd6, 0xbd,
	0x1b, 0x91, 0xba, 0x87, 0xeb, 0x6e, 0xe6, 0xd1, 0x50, 0xb2, 0x25, 0xa3, 0xc2, 0xf8, 0x57, 0x9c,
	0xaf, 0x7c, 0xea, 0x28, 0x6b, 0x91, 0x2c, 0x1d, 0x2f, 0x11, 0xa4, 0xc0, 0xff, 0x72, 0xd3, 0x2f,
	0x59, 0x40, 0x63, 0x49, 0x82, 0x28, 0x03, 0xd8, 0x7f, 0x94, 0xe1, 0xc5, 0xd8, 0x04, 0x1d, 0x09,
	0x4a, 0x24, 0xc5, 0xf4, 0x97, 0x84, 0xc6, 0x12, 0x59, 0xb0, 0x13, 0x09, 0xfe, 0x9e, 0xba, 0xd2,
	0x2a, 0x1f, 0x95, 0x8f, 0x9b, 0xd8, 0x98, 0xe8, 0x05, 0x34, 0x3c, 0x1e, 0x10, 0x16, 0x5a, 0x15,
	0xe5, 0xd0, 0x16, 0x42, 0x50, 0x0b, 0x49, 0x40, 0xad, 0xaa, 0x5a, 0x55, 0xdf, 0xe8, 0x04, 0x6a,
	0x71, 0x44, 0x5d, 0xab, 0x76, 0x54, 0x3e, 0x6e, 0xf5, 0x5f, 0xf5, 0xd6, 0x4b, 0xd7, 0xcb, 0x63,
	0xcf, 0x22, 0xea, 0x62, 0x05, 0xb5, 0xdf, 0x83, 0x95, 0x2f, 0x63, 0xea, 0x93, 0x24, 0x74, 0x6f,
	0x4d, 0x52, 0x67, 0x50, 0x61, 0x9e, 0xca, 0xa7, 0xd5, 0xff, 0xfa, 0x6e, 0xb3, 0xb4, 0x3a, 0xbd,
	0x1f, 0xb9, 0xf8, 0xb0, 0xf4, 0xf9, 0xc7, 0x9c, 0x7c, 0x91, 0x97, 0x0b, 0x57, 0x98, 0xf7, 0x50,
	0x7a, 0xf6, 0x0d, 0x7c, 0x71, 0xef, 0xf8, 0x71, 0xc4, 0xc3, 0x98, 0x3e, 0x25, 0x94, 0xfd, 0x0e,
	0x5e, 0xde, 0x83, 0xfc, 0x40, 0xe5, 0xff, 0x70, 0x0a, 0xfb, 0xaf, 0x32, 0x34, 0x73, 0xdf, 0x93,
	0xea, 0x61, 0xae, 0xa6, 0xf2, 0xaf, 0xaf, 0x06, 0x9d, 0xc1, 0x8e, 0xeb, 0xf3, 0x38, 0x11, 0x59,
	0x15, 0x5b, 0xfd, 0xa3, 0x47, 0x59, 0xa3, 0x0c, 0x87, 0x0d, 0xc1, 0xfe, 0x19, 0xf6, 0x72, 0xe7,
	0x94, 0xc5, 0x12, 0x7d, 0x0b, 0x90, 0xf7, 0x7b, 0x6c, 0x95, 0x8f, 0xaa, 0xc7, 0xad, 0xfe, 0xfe,
	0xa3, 0xfb, 0xe1, 0x02, 0x18, 0x3d, 0x87, 0xba, 0xe4, 0x1f, 0xa8, 0x69, 0xc0, 0xcc, 0xb0, 0x09,
	0xb4, 0xa7, 0x99, 0xbc, 0x2e, 0x49, 0x34, 0xf4, 0xf9, 0x02, 0x9d, 0x42, 0xe3, 0x57, 0xe2, 0x27,
	0x34, 0xd6, 0x25, 0xda, 0xdf, 0x28, 0xd1, 0x1d, 0x7c, 0x52, 0xc2, 0x1a, 0x8a, 0x10, 0x54, 0x13,
	0xc1, 0xb2, 0xad, 0x27, 0x25, 0x9c, 0x1a, 0xc3, 0x06, 0xd4, 0x3c, 0x22, 0x89, 0xfd, 0x67, 0x1d,
	0x3a, 0x9b, 0x47, 0x4c, 0xab, 0xc2, 0x13, 0x19, 0x25, 0xd2, 0x84, 0x39, 0xdc, 0x3c, 0xc5, 0x7a,
	0x5a, 0x93, 0x12, 0x36, 0x04, 0xf4, 0x06, 0xea, 0x54, 0x08, 0x2e, 0xee, 0xdf, 0x82, 0x4a, 0x30,
	0x8f, 0x35, 0x4e, 0x41, 0x93, 0x12, 0xce, 0xd0, 0xe8, 0x35, 0xb4, 0xc8, 0x82, 0x0b, 0x39, 0x77,
	0x49, 0x12, 0x53, 0x0b, 0x74, 0xae, 0xa0, 0x16, 0x47, 0xe9, 0x1a, 0x1a, 0xc2, 0x33, 0x97, 0x07,
	0x51, 0x22, 0xa9, 0x37, 0x67, 0xa1, 0xca, 0xae, 0xba, 0xa5, 0x08, 0xb8, 0x6d, 0x18, 0x17, 0x8a,
	0x80, 0xbe, 0x83, 0x7a, 0x74, 0x4b, 0x62, 0xaa, 0xe4, 0xdb, 0xee, 0x7f, 0xb5, 0xad, 0xc3, 0xde,
	0xa6, 0x60, 0x9c, 0x71, 0xd2, 0xfb, 0x8d, 0x25, 0x11, 0x69, 0x7c, 0x22, 0xad, 0xba, 0x8a, 0xdd,
	0xed, 0x65, 0x13, 0xa9, 0x67, 0x26, 0x52, 0xef, 0xda, 0x4c, 0x24, 0xdc, 0xd4, 0xe8, 0x81, 0x44,
	0x6f, 0x60, 0xd7, 0x4c, 0x32, 0xab, 0xa1, 0x93, 0xde, 0x24, 0x9e, 0x6b, 0x00, 0xce, 0xa1, 0x69,
	0x44, 0x57, 0x89, 0x58, 0x45, 0xdc, 0xd9, 0x1e, 0x51, 0xa3, 0x07, 0xaa, 0x19, 0x93, 0xc8, 0x33,
	0xd4, 0xdd, 0xed, 0x54, 0x8d, 0x1e, 0x48, 0x34, 0x84, 0xbd, 0x90, 0xa7, 0xba, 0x72, 0x49, 0xd6,
	0xca, 0x4d, 0xd5, 0xca, 0x07, 0x9b, 0x4d, 0x70, 0x55, 0x00, 0xe1, 0x75, 0x0a, 0x3a, 0x83, 0xd6,
	0x47, 0x5d, 0xcc, 0x39, 0xf3, 0xac, 0xd6, 0x83, 0x17, 0x55, 0xd0, 0x2f, 0x18, 0xf4, 0x85, 0x37,
	0x7c, 0x06, 0x7b, 0x59, 0x37, 0xcd, 0x05, 0x8d, 0x13, 0x5f, 0xda, 0x7f, 0x97, 0xe1, 0xf3, 0xfc,
	0x4a, 0x2e, 0xa9, 0x24, 0x69, 0xeb, 0xa2, 0x11, 0xd4, 0x02, 0xee, 0x51, 0xd5, 0xa2, 0xed, 0xbe,
	0xf3, 0xa8, 0xd0, 0x0c, 0xa1, 0xb0, 0xc2, 0x3d, 0x8a, 0x15, 0x19, 0x1d, 0x40, 0x33, 0x12, 0x2c,
	0x74, 0x59, 0x44, 0x7c, 0x2d, 0xbe, 0xbb, 0x85, 0xf4, 0xc9, 0x08, 0x69, 0x2c, 0x59, 0xb8, 0x52,
	0xad, 0x56, 0xc7, 0xc6, 0xb4, 0xcf, 0x0b, 0xe2, 0x4f, 0xb7, 0x43, 0x00, 0x8d, 0xcb, 0xc1, 0xd5,
	0xcd, 0x60, 0xda, 0x29, 0xa1, 0x3d, 0x68, 0xce, 0x46, 0x93, 0xf1, 0xf9, 0xcd, 0x74, 0x7c, 0xde,
	0x29, 0xa7, 0xae, 0xd9, 0xbb, 0xd9, 0xf5, 0xf8, 0xb2, 0x53, 0x41, 0x9f, 0xc1, 0x2e, 0x1e, 0x4f,
	0x07, 0x37, 0x57, 0xa3, 0x49, 0xa7, 0x6a, 0xff, 0x5e, 0x29, 0x6c, 0x33, 0xcb, 0x06, 0x52, 0x2b,
	0x7b, 0x20, 0xe6, 0x91, 0x4f, 0xc2, 0x47, 0x54, 0x5e, 0xac, 0x5b, 0x86, 0x7e, 0xeb, 0x93, 0x10,
	0x9d, 0x40, 0x43, 0xeb, 0xa2, 0xb2, 0x4d, 0x17, 0x1a, 0x88, 0xbe, 0x87, 0xdd, 0x40, 0x97, 0x47,
	0x8b, 0xe9, 0xf5, 0xd6, 0x3a, 0xe2, 0x9c, 0x72, 0xbf, 0x53, 0x6a, 0xff, 0xb9, 0x53, 0xec, 0x00,
	0xf6, 0xf3, 0x10, 0xd7, 0x54, 0x04, 0x2c, 0x2c, 0xbc, 0xd9, 0x4f, 0x79, 0x0e, 0x9e, 0x43, 0x3d,
	0x1b, 0x26, 0x7a, 0xa6, 0x2a, 0xc3, 0x3e, 0x80, 0xee, 0x43, 0xe1, 0xb2, 0x37, 0x72, 0x78, 0xfa,
	0xd3, 0xc9, 0x8a, 0xc9, 0xdb, 0x64, 0xd1, 0x73, 0x79, 0xe0, 0xf8, 0x9f, 0x96, 0xd2, 0xc9, 0xff,
	0x58, 0x56, 0x34, 0x74, 0xa2, 0xc5, 0x37, 0x2b, 0xee, 0xac, 0xff, 0x1e, 0x2d, 0x1a, 0x4a, 0x4e,
	0xa7, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xe7, 0x44, 0xf2, 0x2f, 0x67, 0x09, 0x00, 0x00,
}
