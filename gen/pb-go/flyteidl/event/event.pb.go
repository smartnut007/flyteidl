// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/event/event.proto

package event

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	core "github.com/lyft/flyteidl/gen/pb-go/flyteidl/core"
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

type WorkflowExecutionEvent struct {
	// Workflow execution id
	ExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                       `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.WorkflowExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.WorkflowExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the workflow.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*WorkflowExecutionEvent_OutputUri
	//	*WorkflowExecutionEvent_Error
	OutputResult         isWorkflowExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *WorkflowExecutionEvent) Reset()         { *m = WorkflowExecutionEvent{} }
func (m *WorkflowExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*WorkflowExecutionEvent) ProtoMessage()    {}
func (*WorkflowExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{0}
}

func (m *WorkflowExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowExecutionEvent.Unmarshal(m, b)
}
func (m *WorkflowExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowExecutionEvent.Marshal(b, m, deterministic)
}
func (m *WorkflowExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowExecutionEvent.Merge(m, src)
}
func (m *WorkflowExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_WorkflowExecutionEvent.Size(m)
}
func (m *WorkflowExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowExecutionEvent proto.InternalMessageInfo

func (m *WorkflowExecutionEvent) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetPhase() core.WorkflowExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.WorkflowExecution_UNDEFINED
}

func (m *WorkflowExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

type isWorkflowExecutionEvent_OutputResult interface {
	isWorkflowExecutionEvent_OutputResult()
}

type WorkflowExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,5,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type WorkflowExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,6,opt,name=error,proto3,oneof"`
}

func (*WorkflowExecutionEvent_OutputUri) isWorkflowExecutionEvent_OutputResult() {}

func (*WorkflowExecutionEvent_Error) isWorkflowExecutionEvent_OutputResult() {}

func (m *WorkflowExecutionEvent) GetOutputResult() isWorkflowExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *WorkflowExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *WorkflowExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*WorkflowExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*WorkflowExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*WorkflowExecutionEvent_OutputUri)(nil),
		(*WorkflowExecutionEvent_Error)(nil),
	}
}

type NodeExecutionEvent struct {
	// Unique identifier for this node execution
	Id *core.NodeExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the id of the originator (Propeller) of the event
	ProducerId string                   `protobuf:"bytes,2,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	Phase      core.NodeExecution_Phase `protobuf:"varint,3,opt,name=phase,proto3,enum=flyteidl.core.NodeExecution_Phase" json:"phase,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the node.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,4,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	InputUri   string               `protobuf:"bytes,5,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*NodeExecutionEvent_OutputUri
	//	*NodeExecutionEvent_Error
	OutputResult isNodeExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Additional metadata to do with this event's node target based
	// on the node type
	//
	// Types that are valid to be assigned to TargetMetadata:
	//	*NodeExecutionEvent_WorkflowNodeMetadata
	TargetMetadata isNodeExecutionEvent_TargetMetadata `protobuf_oneof:"target_metadata"`
	// [To be deprecated] Specifies which task (if any) launched this node.
	ParentTaskMetadata *ParentTaskExecutionMetadata `protobuf:"bytes,9,opt,name=parent_task_metadata,json=parentTaskMetadata,proto3" json:"parent_task_metadata,omitempty"`
	// Specifies the parent node of the current node execution. Node executions at level zero will not have a parent node.
	ParentNodeMetadata *ParentNodeExecutionMetadata `protobuf:"bytes,10,opt,name=parent_node_metadata,json=parentNodeMetadata,proto3" json:"parent_node_metadata,omitempty"`
	// Group identifier to indicate arbitrary grouping like retries, sub workflow
	GroupId string `protobuf:"bytes,11,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	// Identifier of the node in the original workflow/graph
	GraphNodeId string `protobuf:"bytes,12,opt,name=graph_node_id,json=graphNodeId,proto3" json:"graph_node_id,omitempty"`
	// Friendly readable name for the node
	NodeName             string   `protobuf:"bytes,13,opt,name=node_name,json=nodeName,proto3" json:"node_name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NodeExecutionEvent) Reset()         { *m = NodeExecutionEvent{} }
func (m *NodeExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*NodeExecutionEvent) ProtoMessage()    {}
func (*NodeExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{1}
}

func (m *NodeExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NodeExecutionEvent.Unmarshal(m, b)
}
func (m *NodeExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NodeExecutionEvent.Marshal(b, m, deterministic)
}
func (m *NodeExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NodeExecutionEvent.Merge(m, src)
}
func (m *NodeExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_NodeExecutionEvent.Size(m)
}
func (m *NodeExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_NodeExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_NodeExecutionEvent proto.InternalMessageInfo

func (m *NodeExecutionEvent) GetId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *NodeExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *NodeExecutionEvent) GetPhase() core.NodeExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.NodeExecution_UNDEFINED
}

func (m *NodeExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *NodeExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isNodeExecutionEvent_OutputResult interface {
	isNodeExecutionEvent_OutputResult()
}

type NodeExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,6,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type NodeExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,7,opt,name=error,proto3,oneof"`
}

func (*NodeExecutionEvent_OutputUri) isNodeExecutionEvent_OutputResult() {}

func (*NodeExecutionEvent_Error) isNodeExecutionEvent_OutputResult() {}

func (m *NodeExecutionEvent) GetOutputResult() isNodeExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *NodeExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *NodeExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*NodeExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

type isNodeExecutionEvent_TargetMetadata interface {
	isNodeExecutionEvent_TargetMetadata()
}

type NodeExecutionEvent_WorkflowNodeMetadata struct {
	WorkflowNodeMetadata *WorkflowNodeMetadata `protobuf:"bytes,8,opt,name=workflow_node_metadata,json=workflowNodeMetadata,proto3,oneof"`
}

func (*NodeExecutionEvent_WorkflowNodeMetadata) isNodeExecutionEvent_TargetMetadata() {}

func (m *NodeExecutionEvent) GetTargetMetadata() isNodeExecutionEvent_TargetMetadata {
	if m != nil {
		return m.TargetMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetWorkflowNodeMetadata() *WorkflowNodeMetadata {
	if x, ok := m.GetTargetMetadata().(*NodeExecutionEvent_WorkflowNodeMetadata); ok {
		return x.WorkflowNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentTaskMetadata() *ParentTaskExecutionMetadata {
	if m != nil {
		return m.ParentTaskMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetParentNodeMetadata() *ParentNodeExecutionMetadata {
	if m != nil {
		return m.ParentNodeMetadata
	}
	return nil
}

func (m *NodeExecutionEvent) GetGroupId() string {
	if m != nil {
		return m.GroupId
	}
	return ""
}

func (m *NodeExecutionEvent) GetGraphNodeId() string {
	if m != nil {
		return m.GraphNodeId
	}
	return ""
}

func (m *NodeExecutionEvent) GetNodeName() string {
	if m != nil {
		return m.NodeName
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*NodeExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*NodeExecutionEvent_OutputUri)(nil),
		(*NodeExecutionEvent_Error)(nil),
		(*NodeExecutionEvent_WorkflowNodeMetadata)(nil),
	}
}

// For Workflow Nodes we need to send information about the workflow that's launched
type WorkflowNodeMetadata struct {
	ExecutionId          *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=execution_id,json=executionId,proto3" json:"execution_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                          `json:"-"`
	XXX_unrecognized     []byte                            `json:"-"`
	XXX_sizecache        int32                             `json:"-"`
}

func (m *WorkflowNodeMetadata) Reset()         { *m = WorkflowNodeMetadata{} }
func (m *WorkflowNodeMetadata) String() string { return proto.CompactTextString(m) }
func (*WorkflowNodeMetadata) ProtoMessage()    {}
func (*WorkflowNodeMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{2}
}

func (m *WorkflowNodeMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_WorkflowNodeMetadata.Unmarshal(m, b)
}
func (m *WorkflowNodeMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_WorkflowNodeMetadata.Marshal(b, m, deterministic)
}
func (m *WorkflowNodeMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowNodeMetadata.Merge(m, src)
}
func (m *WorkflowNodeMetadata) XXX_Size() int {
	return xxx_messageInfo_WorkflowNodeMetadata.Size(m)
}
func (m *WorkflowNodeMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowNodeMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowNodeMetadata proto.InternalMessageInfo

func (m *WorkflowNodeMetadata) GetExecutionId() *core.WorkflowExecutionIdentifier {
	if m != nil {
		return m.ExecutionId
	}
	return nil
}

type ParentTaskExecutionMetadata struct {
	Id                   *core.TaskExecutionIdentifier `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                      `json:"-"`
	XXX_unrecognized     []byte                        `json:"-"`
	XXX_sizecache        int32                         `json:"-"`
}

func (m *ParentTaskExecutionMetadata) Reset()         { *m = ParentTaskExecutionMetadata{} }
func (m *ParentTaskExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentTaskExecutionMetadata) ProtoMessage()    {}
func (*ParentTaskExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{3}
}

func (m *ParentTaskExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentTaskExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *ParentTaskExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentTaskExecutionMetadata.Merge(m, src)
}
func (m *ParentTaskExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentTaskExecutionMetadata.Size(m)
}
func (m *ParentTaskExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentTaskExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentTaskExecutionMetadata proto.InternalMessageInfo

func (m *ParentTaskExecutionMetadata) GetId() *core.TaskExecutionIdentifier {
	if m != nil {
		return m.Id
	}
	return nil
}

type ParentNodeExecutionMetadata struct {
	// Unique identifier of the parent node id within the execution
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParentNodeExecutionMetadata) Reset()         { *m = ParentNodeExecutionMetadata{} }
func (m *ParentNodeExecutionMetadata) String() string { return proto.CompactTextString(m) }
func (*ParentNodeExecutionMetadata) ProtoMessage()    {}
func (*ParentNodeExecutionMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{4}
}

func (m *ParentNodeExecutionMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Unmarshal(m, b)
}
func (m *ParentNodeExecutionMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Marshal(b, m, deterministic)
}
func (m *ParentNodeExecutionMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParentNodeExecutionMetadata.Merge(m, src)
}
func (m *ParentNodeExecutionMetadata) XXX_Size() int {
	return xxx_messageInfo_ParentNodeExecutionMetadata.Size(m)
}
func (m *ParentNodeExecutionMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_ParentNodeExecutionMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_ParentNodeExecutionMetadata proto.InternalMessageInfo

func (m *ParentNodeExecutionMetadata) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// Plugin specific execution event information. For tasks like Python, Hive, Spark, DynamicJob.
type TaskExecutionEvent struct {
	// ID of the task. In combination with the retryAttempt this will indicate
	// the task execution uniquely for a given parent node execution.
	TaskId *core.Identifier `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	// A task execution is always kicked off by a node execution, the event consumer
	// will use the parent_id to relate the task to it's parent node execution
	ParentNodeExecutionId *core.NodeExecutionIdentifier `protobuf:"bytes,2,opt,name=parent_node_execution_id,json=parentNodeExecutionId,proto3" json:"parent_node_execution_id,omitempty"`
	// retry attempt number for this task, ie., 2 for the second attempt
	RetryAttempt uint32 `protobuf:"varint,3,opt,name=retry_attempt,json=retryAttempt,proto3" json:"retry_attempt,omitempty"`
	// Phase associated with the event
	Phase core.TaskExecution_Phase `protobuf:"varint,4,opt,name=phase,proto3,enum=flyteidl.core.TaskExecution_Phase" json:"phase,omitempty"`
	// id of the process that sent this event, mainly for trace debugging
	ProducerId string `protobuf:"bytes,5,opt,name=producer_id,json=producerId,proto3" json:"producer_id,omitempty"`
	// log information for the task execution
	Logs []*core.TaskLog `protobuf:"bytes,6,rep,name=logs,proto3" json:"logs,omitempty"`
	// This timestamp represents when the original event occurred, it is generated
	// by the executor of the task.
	OccurredAt *timestamp.Timestamp `protobuf:"bytes,7,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	// URI of the input file, it encodes all the information
	// including Cloud source provider. ie., s3://...
	InputUri string `protobuf:"bytes,8,opt,name=input_uri,json=inputUri,proto3" json:"input_uri,omitempty"`
	// Types that are valid to be assigned to OutputResult:
	//	*TaskExecutionEvent_OutputUri
	//	*TaskExecutionEvent_Error
	OutputResult isTaskExecutionEvent_OutputResult `protobuf_oneof:"output_result"`
	// Custom data that the task plugin sends back. This is extensible to allow various plugins in the system.
	CustomInfo *_struct.Struct `protobuf:"bytes,11,opt,name=custom_info,json=customInfo,proto3" json:"custom_info,omitempty"`
	// Some phases, like RUNNING, can send multiple events with changed metadata (new logs, additional custom_info, etc)
	// that should be recorded regardless of the lack of phase change.
	// The version field should be incremented when metadata changes across the duration of an individual phase.
	PhaseVersion         uint32   `protobuf:"varint,12,opt,name=phase_version,json=phaseVersion,proto3" json:"phase_version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TaskExecutionEvent) Reset()         { *m = TaskExecutionEvent{} }
func (m *TaskExecutionEvent) String() string { return proto.CompactTextString(m) }
func (*TaskExecutionEvent) ProtoMessage()    {}
func (*TaskExecutionEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_4b035d24120b1b12, []int{5}
}

func (m *TaskExecutionEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TaskExecutionEvent.Unmarshal(m, b)
}
func (m *TaskExecutionEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TaskExecutionEvent.Marshal(b, m, deterministic)
}
func (m *TaskExecutionEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TaskExecutionEvent.Merge(m, src)
}
func (m *TaskExecutionEvent) XXX_Size() int {
	return xxx_messageInfo_TaskExecutionEvent.Size(m)
}
func (m *TaskExecutionEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_TaskExecutionEvent.DiscardUnknown(m)
}

var xxx_messageInfo_TaskExecutionEvent proto.InternalMessageInfo

func (m *TaskExecutionEvent) GetTaskId() *core.Identifier {
	if m != nil {
		return m.TaskId
	}
	return nil
}

func (m *TaskExecutionEvent) GetParentNodeExecutionId() *core.NodeExecutionIdentifier {
	if m != nil {
		return m.ParentNodeExecutionId
	}
	return nil
}

func (m *TaskExecutionEvent) GetRetryAttempt() uint32 {
	if m != nil {
		return m.RetryAttempt
	}
	return 0
}

func (m *TaskExecutionEvent) GetPhase() core.TaskExecution_Phase {
	if m != nil {
		return m.Phase
	}
	return core.TaskExecution_UNDEFINED
}

func (m *TaskExecutionEvent) GetProducerId() string {
	if m != nil {
		return m.ProducerId
	}
	return ""
}

func (m *TaskExecutionEvent) GetLogs() []*core.TaskLog {
	if m != nil {
		return m.Logs
	}
	return nil
}

func (m *TaskExecutionEvent) GetOccurredAt() *timestamp.Timestamp {
	if m != nil {
		return m.OccurredAt
	}
	return nil
}

func (m *TaskExecutionEvent) GetInputUri() string {
	if m != nil {
		return m.InputUri
	}
	return ""
}

type isTaskExecutionEvent_OutputResult interface {
	isTaskExecutionEvent_OutputResult()
}

type TaskExecutionEvent_OutputUri struct {
	OutputUri string `protobuf:"bytes,9,opt,name=output_uri,json=outputUri,proto3,oneof"`
}

type TaskExecutionEvent_Error struct {
	Error *core.ExecutionError `protobuf:"bytes,10,opt,name=error,proto3,oneof"`
}

func (*TaskExecutionEvent_OutputUri) isTaskExecutionEvent_OutputResult() {}

func (*TaskExecutionEvent_Error) isTaskExecutionEvent_OutputResult() {}

func (m *TaskExecutionEvent) GetOutputResult() isTaskExecutionEvent_OutputResult {
	if m != nil {
		return m.OutputResult
	}
	return nil
}

func (m *TaskExecutionEvent) GetOutputUri() string {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_OutputUri); ok {
		return x.OutputUri
	}
	return ""
}

func (m *TaskExecutionEvent) GetError() *core.ExecutionError {
	if x, ok := m.GetOutputResult().(*TaskExecutionEvent_Error); ok {
		return x.Error
	}
	return nil
}

func (m *TaskExecutionEvent) GetCustomInfo() *_struct.Struct {
	if m != nil {
		return m.CustomInfo
	}
	return nil
}

func (m *TaskExecutionEvent) GetPhaseVersion() uint32 {
	if m != nil {
		return m.PhaseVersion
	}
	return 0
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TaskExecutionEvent) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TaskExecutionEvent_OutputUri)(nil),
		(*TaskExecutionEvent_Error)(nil),
	}
}

func init() {
	proto.RegisterType((*WorkflowExecutionEvent)(nil), "flyteidl.event.WorkflowExecutionEvent")
	proto.RegisterType((*NodeExecutionEvent)(nil), "flyteidl.event.NodeExecutionEvent")
	proto.RegisterType((*WorkflowNodeMetadata)(nil), "flyteidl.event.WorkflowNodeMetadata")
	proto.RegisterType((*ParentTaskExecutionMetadata)(nil), "flyteidl.event.ParentTaskExecutionMetadata")
	proto.RegisterType((*ParentNodeExecutionMetadata)(nil), "flyteidl.event.ParentNodeExecutionMetadata")
	proto.RegisterType((*TaskExecutionEvent)(nil), "flyteidl.event.TaskExecutionEvent")
}

func init() { proto.RegisterFile("flyteidl/event/event.proto", fileDescriptor_4b035d24120b1b12) }

var fileDescriptor_4b035d24120b1b12 = []byte{
	// 764 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x56, 0xdf, 0x6b, 0xdb, 0x3a,
	0x18, 0x6d, 0xd2, 0xfc, 0xfc, 0x9c, 0xb4, 0x5c, 0xd1, 0xdb, 0xeb, 0xa6, 0xb7, 0xb7, 0x21, 0x77,
	0x8c, 0xd0, 0x51, 0x9b, 0xa5, 0x6c, 0x14, 0xb6, 0x97, 0x16, 0x0a, 0x0d, 0xac, 0xa5, 0x78, 0xed,
	0x06, 0x63, 0xc3, 0x38, 0xb6, 0xe2, 0x8a, 0xc6, 0x96, 0x91, 0xe5, 0x76, 0xfd, 0x9f, 0xf7, 0x3a,
	0xf6, 0x3a, 0x24, 0xc7, 0x4e, 0xe5, 0x84, 0x6c, 0x1d, 0x7d, 0x09, 0xf8, 0xd3, 0xf1, 0xd1, 0xd1,
	0xd1, 0xf7, 0x9d, 0x18, 0x3a, 0xe3, 0xc9, 0x3d, 0xc7, 0xc4, 0x9b, 0x98, 0xf8, 0x16, 0x87, 0x3c,
	0xfd, 0x35, 0x22, 0x46, 0x39, 0x45, 0x6b, 0xd9, 0x9a, 0x21, 0xab, 0x9d, 0x9d, 0x1c, 0xeb, 0x52,
	0x86, 0x4d, 0xfc, 0x15, 0xbb, 0x09, 0x27, 0x34, 0x4c, 0xe1, 0x9d, 0xff, 0xd4, 0x65, 0xe2, 0xe1,
	0x90, 0x93, 0x31, 0xc1, 0x6c, 0xba, 0xbe, 0xeb, 0x53, 0xea, 0x4f, 0xb0, 0x29, 0x9f, 0x46, 0xc9,
	0xd8, 0xe4, 0x24, 0xc0, 0x31, 0x77, 0x82, 0x68, 0x0a, 0xf8, 0xb7, 0x08, 0x88, 0x39, 0x4b, 0xdc,
	0xa9, 0x9a, 0xde, 0xb7, 0x32, 0x6c, 0x7e, 0xa4, 0xec, 0x66, 0x3c, 0xa1, 0x77, 0x27, 0xd9, 0xd6,
	0x27, 0x42, 0x18, 0x3a, 0x83, 0x56, 0x2e, 0xc6, 0x26, 0x9e, 0x5e, 0xea, 0x96, 0xfa, 0xda, 0x60,
	0xcf, 0xc8, 0xf5, 0x0b, 0x41, 0xc6, 0xdc, 0xcb, 0xc3, 0x5c, 0xa1, 0xa5, 0xe1, 0x59, 0x11, 0xed,
	0x82, 0x16, 0x31, 0xea, 0x25, 0x2e, 0x66, 0x82, 0xad, 0xdc, 0x2d, 0xf5, 0x9b, 0x16, 0x64, 0xa5,
	0xa1, 0x87, 0xde, 0x42, 0x35, 0xba, 0x76, 0x62, 0xac, 0xaf, 0x76, 0x4b, 0xfd, 0xb5, 0xc1, 0xf3,
	0x5f, 0x6d, 0x64, 0x5c, 0x08, 0xb4, 0x95, 0xbe, 0x84, 0xde, 0x80, 0x46, 0x5d, 0x37, 0x61, 0x0c,
	0x7b, 0xb6, 0xc3, 0xf5, 0x8a, 0x14, 0xdb, 0x31, 0xd2, 0xc3, 0x1b, 0xd9, 0xe1, 0x8d, 0xcb, 0xcc,
	0x1d, 0x0b, 0x32, 0xf8, 0x11, 0x47, 0xbb, 0x00, 0x34, 0xe1, 0x51, 0xc2, 0xed, 0x84, 0x11, 0xbd,
	0x2a, 0xa4, 0x9d, 0xae, 0x58, 0xcd, 0xb4, 0x76, 0xc5, 0x08, 0x7a, 0x05, 0x55, 0xcc, 0x18, 0x65,
	0x7a, 0x4d, 0xf2, 0xee, 0x14, 0xb4, 0xcd, 0x9c, 0x13, 0xa0, 0xd3, 0x15, 0x2b, 0x45, 0x1f, 0xaf,
	0x43, 0x7b, 0xca, 0xcb, 0x70, 0x9c, 0x4c, 0x78, 0xef, 0x7b, 0x15, 0xd0, 0x39, 0xf5, 0x70, 0xc1,
	0xea, 0xd7, 0x50, 0xce, 0x0d, 0x2e, 0x9e, 0x5b, 0x81, 0x3f, 0x30, 0xb7, 0x4c, 0x7e, 0xc3, 0xd3,
	0x43, 0xd5, 0xd3, 0xde, 0x32, 0xee, 0x27, 0xf4, 0x73, 0x1b, 0x9a, 0x24, 0x54, 0xec, 0xb4, 0x1a,
	0xb2, 0x20, 0xbc, 0x54, 0xcd, 0xae, 0x2d, 0x31, 0xbb, 0xfe, 0x18, 0xb3, 0xd1, 0x67, 0xd8, 0xbc,
	0x9b, 0xf6, 0x88, 0x1d, 0x52, 0x0f, 0xdb, 0x01, 0xe6, 0x8e, 0xe7, 0x70, 0x47, 0x6f, 0x48, 0x9e,
	0x67, 0x86, 0x3a, 0x79, 0x79, 0x47, 0x09, 0x17, 0xce, 0xa6, 0xd8, 0xd3, 0x92, 0xb5, 0x71, 0xb7,
	0xa0, 0x8e, 0xbe, 0xc0, 0x46, 0xe4, 0x30, 0x1c, 0x72, 0x9b, 0x3b, 0xf1, 0xcd, 0x8c, 0xbb, 0x29,
	0xb9, 0x5f, 0x14, 0xb9, 0x2f, 0x24, 0xf6, 0xd2, 0x89, 0x6f, 0x72, 0xb9, 0x19, 0x95, 0x85, 0xa2,
	0x7c, 0x71, 0x01, 0xbd, 0x2a, 0x1d, 0x96, 0xd1, 0x2b, 0xd7, 0x57, 0xa4, 0x57, 0xd4, 0x6f, 0x41,
	0xc3, 0x67, 0x34, 0x89, 0x44, 0x97, 0x68, 0xf2, 0x3e, 0xea, 0xf2, 0x79, 0xe8, 0xa1, 0x1e, 0xb4,
	0x7d, 0xe6, 0x44, 0xd7, 0xe9, 0xc6, 0xc4, 0xd3, 0x5b, 0x72, 0x5d, 0x93, 0x45, 0x41, 0x32, 0xf4,
	0xc4, 0x7d, 0xca, 0xd5, 0xd0, 0x09, 0xb0, 0xde, 0x4e, 0xef, 0x53, 0x14, 0xce, 0x9d, 0x00, 0xcf,
	0x35, 0xf9, 0xf1, 0x5f, 0xb0, 0xce, 0x1d, 0xe6, 0x63, 0x9e, 0x1f, 0xa3, 0x87, 0x61, 0x63, 0x91,
	0xdb, 0x4f, 0x9c, 0x31, 0xbd, 0x2b, 0xd8, 0x5e, 0x62, 0xfc, 0xd2, 0x31, 0x53, 0xde, 0x50, 0xc7,
	0xac, 0xb7, 0x9f, 0xd1, 0x2e, 0x34, 0x1c, 0xad, 0xe5, 0xb4, 0x4d, 0x09, 0xff, 0x51, 0x01, 0xa4,
	0xd0, 0xa5, 0x43, 0x3e, 0x80, 0xba, 0x6c, 0x9d, 0x5c, 0xc2, 0x56, 0x41, 0xc2, 0x83, 0x5d, 0x6b,
	0x02, 0x39, 0xf4, 0x90, 0x0d, 0xfa, 0xc3, 0xb6, 0x50, 0xbc, 0x2a, 0x3f, 0x2a, 0x2e, 0xfe, 0x8e,
	0xe6, 0x4f, 0x30, 0xf4, 0xd0, 0xff, 0xd0, 0x66, 0x98, 0xb3, 0x7b, 0xdb, 0xe1, 0x1c, 0x07, 0x11,
	0x97, 0x41, 0xd1, 0xb6, 0x5a, 0xb2, 0x78, 0x94, 0xd6, 0x66, 0x29, 0x52, 0x59, 0x98, 0x22, 0xca,
	0x59, 0xd5, 0x14, 0x29, 0x04, 0x54, 0x75, 0x2e, 0xa0, 0xf6, 0xa0, 0x32, 0xa1, 0x7e, 0xac, 0xd7,
	0xba, 0xab, 0x7d, 0x6d, 0xb0, 0xb9, 0x80, 0xf9, 0x1d, 0xf5, 0x2d, 0x89, 0x29, 0x46, 0x52, 0xfd,
	0xcf, 0x23, 0xa9, 0xb1, 0x34, 0x92, 0x9a, 0x4b, 0x22, 0x09, 0x1e, 0x15, 0x49, 0x87, 0xa0, 0xb9,
	0x49, 0xcc, 0x69, 0x60, 0x93, 0x70, 0x4c, 0xe5, 0xe4, 0x69, 0x83, 0x7f, 0xe6, 0x14, 0xbf, 0x97,
	0xff, 0xc8, 0x16, 0xa4, 0xd8, 0x61, 0x38, 0xa6, 0xe2, 0x5e, 0xa4, 0x83, 0xf6, 0x2d, 0x66, 0x31,
	0xa1, 0xa1, 0x9c, 0xca, 0xb6, 0xd5, 0x92, 0xc5, 0x0f, 0x69, 0x6d, 0x7e, 0xf2, 0x0e, 0x3e, 0xbd,
	0xf4, 0x09, 0xbf, 0x4e, 0x46, 0x86, 0x4b, 0x03, 0x73, 0x72, 0x3f, 0xe6, 0x66, 0xfe, 0xf9, 0xe0,
	0xe3, 0xd0, 0x8c, 0x46, 0xfb, 0x3e, 0x35, 0xd5, 0x8f, 0x93, 0x51, 0x4d, 0xea, 0x38, 0xf8, 0x19,
	0x00, 0x00, 0xff, 0xff, 0xe6, 0x2b, 0x04, 0xdb, 0xb5, 0x08, 0x00, 0x00,
}
