# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: flyteidl/event/event.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from flyteidl.core import execution_pb2 as flyteidl_dot_core_dot_execution__pb2
from flyteidl.core import identifier_pb2 as flyteidl_dot_core_dot_identifier__pb2
from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2
from google.protobuf import struct_pb2 as google_dot_protobuf_dot_struct__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='flyteidl/event/event.proto',
  package='flyteidl.event',
  syntax='proto3',
  serialized_options=_b('Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/event'),
  serialized_pb=_b('\n\x1a\x66lyteidl/event/event.proto\x12\x0e\x66lyteidl.event\x1a\x1d\x66lyteidl/core/execution.proto\x1a\x1e\x66lyteidl/core/identifier.proto\x1a\x1fgoogle/protobuf/timestamp.proto\x1a\x1cgoogle/protobuf/struct.proto\"\xae\x02\n\x16WorkflowExecutionEvent\x12@\n\x0c\x65xecution_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifier\x12\x13\n\x0bproducer_id\x18\x02 \x01(\t\x12\x35\n\x05phase\x18\x03 \x01(\x0e\x32&.flyteidl.core.WorkflowExecution.Phase\x12/\n\x0boccurred_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x14\n\noutput_uri\x18\x05 \x01(\tH\x00\x12.\n\x05\x65rror\x18\x06 \x01(\x0b\x32\x1d.flyteidl.core.ExecutionErrorH\x00\x42\x0f\n\routput_result\"\xd8\x04\n\x12NodeExecutionEvent\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifier\x12\x13\n\x0bproducer_id\x18\x02 \x01(\t\x12\x31\n\x05phase\x18\x03 \x01(\x0e\x32\".flyteidl.core.NodeExecution.Phase\x12/\n\x0boccurred_at\x18\x04 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tinput_uri\x18\x05 \x01(\t\x12\x14\n\noutput_uri\x18\x06 \x01(\tH\x00\x12.\n\x05\x65rror\x18\x07 \x01(\x0b\x32\x1d.flyteidl.core.ExecutionErrorH\x00\x12\x46\n\x16workflow_node_metadata\x18\x08 \x01(\x0b\x32$.flyteidl.event.WorkflowNodeMetadataH\x01\x12I\n\x14parent_task_metadata\x18\t \x01(\x0b\x32+.flyteidl.event.ParentTaskExecutionMetadata\x12I\n\x14parent_node_metadata\x18\n \x01(\x0b\x32+.flyteidl.event.ParentNodeExecutionMetadata\x12\x10\n\x08group_id\x18\x0b \x01(\t\x12\x15\n\rgraph_node_id\x18\x0c \x01(\t\x12\x11\n\tnode_name\x18\r \x01(\tB\x0f\n\routput_resultB\x11\n\x0ftarget_metadata\"X\n\x14WorkflowNodeMetadata\x12@\n\x0c\x65xecution_id\x18\x01 \x01(\x0b\x32*.flyteidl.core.WorkflowExecutionIdentifier\"Q\n\x1bParentTaskExecutionMetadata\x12\x32\n\x02id\x18\x01 \x01(\x0b\x32&.flyteidl.core.TaskExecutionIdentifier\")\n\x1bParentNodeExecutionMetadata\x12\n\n\x02id\x18\x01 \x01(\t\"\xef\x03\n\x12TaskExecutionEvent\x12*\n\x07task_id\x18\x01 \x01(\x0b\x32\x19.flyteidl.core.Identifier\x12H\n\x18parent_node_execution_id\x18\x02 \x01(\x0b\x32&.flyteidl.core.NodeExecutionIdentifier\x12\x15\n\rretry_attempt\x18\x03 \x01(\r\x12\x31\n\x05phase\x18\x04 \x01(\x0e\x32\".flyteidl.core.TaskExecution.Phase\x12\x13\n\x0bproducer_id\x18\x05 \x01(\t\x12$\n\x04logs\x18\x06 \x03(\x0b\x32\x16.flyteidl.core.TaskLog\x12/\n\x0boccurred_at\x18\x07 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\x11\n\tinput_uri\x18\x08 \x01(\t\x12\x14\n\noutput_uri\x18\t \x01(\tH\x00\x12.\n\x05\x65rror\x18\n \x01(\x0b\x32\x1d.flyteidl.core.ExecutionErrorH\x00\x12,\n\x0b\x63ustom_info\x18\x0b \x01(\x0b\x32\x17.google.protobuf.Struct\x12\x15\n\rphase_version\x18\x0c \x01(\rB\x0f\n\routput_resultB3Z1github.com/lyft/flyteidl/gen/pb-go/flyteidl/eventb\x06proto3')
  ,
  dependencies=[flyteidl_dot_core_dot_execution__pb2.DESCRIPTOR,flyteidl_dot_core_dot_identifier__pb2.DESCRIPTOR,google_dot_protobuf_dot_timestamp__pb2.DESCRIPTOR,google_dot_protobuf_dot_struct__pb2.DESCRIPTOR,])




_WORKFLOWEXECUTIONEVENT = _descriptor.Descriptor(
  name='WorkflowExecutionEvent',
  full_name='flyteidl.event.WorkflowExecutionEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='execution_id', full_name='flyteidl.event.WorkflowExecutionEvent.execution_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='producer_id', full_name='flyteidl.event.WorkflowExecutionEvent.producer_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='flyteidl.event.WorkflowExecutionEvent.phase', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='occurred_at', full_name='flyteidl.event.WorkflowExecutionEvent.occurred_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_uri', full_name='flyteidl.event.WorkflowExecutionEvent.output_uri', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='flyteidl.event.WorkflowExecutionEvent.error', index=5,
      number=6, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='output_result', full_name='flyteidl.event.WorkflowExecutionEvent.output_result',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=173,
  serialized_end=475,
)


_NODEEXECUTIONEVENT = _descriptor.Descriptor(
  name='NodeExecutionEvent',
  full_name='flyteidl.event.NodeExecutionEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.event.NodeExecutionEvent.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='producer_id', full_name='flyteidl.event.NodeExecutionEvent.producer_id', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='flyteidl.event.NodeExecutionEvent.phase', index=2,
      number=3, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='occurred_at', full_name='flyteidl.event.NodeExecutionEvent.occurred_at', index=3,
      number=4, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='input_uri', full_name='flyteidl.event.NodeExecutionEvent.input_uri', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_uri', full_name='flyteidl.event.NodeExecutionEvent.output_uri', index=5,
      number=6, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='flyteidl.event.NodeExecutionEvent.error', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='workflow_node_metadata', full_name='flyteidl.event.NodeExecutionEvent.workflow_node_metadata', index=7,
      number=8, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_task_metadata', full_name='flyteidl.event.NodeExecutionEvent.parent_task_metadata', index=8,
      number=9, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_node_metadata', full_name='flyteidl.event.NodeExecutionEvent.parent_node_metadata', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='group_id', full_name='flyteidl.event.NodeExecutionEvent.group_id', index=10,
      number=11, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='graph_node_id', full_name='flyteidl.event.NodeExecutionEvent.graph_node_id', index=11,
      number=12, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='node_name', full_name='flyteidl.event.NodeExecutionEvent.node_name', index=12,
      number=13, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='output_result', full_name='flyteidl.event.NodeExecutionEvent.output_result',
      index=0, containing_type=None, fields=[]),
    _descriptor.OneofDescriptor(
      name='target_metadata', full_name='flyteidl.event.NodeExecutionEvent.target_metadata',
      index=1, containing_type=None, fields=[]),
  ],
  serialized_start=478,
  serialized_end=1078,
)


_WORKFLOWNODEMETADATA = _descriptor.Descriptor(
  name='WorkflowNodeMetadata',
  full_name='flyteidl.event.WorkflowNodeMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='execution_id', full_name='flyteidl.event.WorkflowNodeMetadata.execution_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1080,
  serialized_end=1168,
)


_PARENTTASKEXECUTIONMETADATA = _descriptor.Descriptor(
  name='ParentTaskExecutionMetadata',
  full_name='flyteidl.event.ParentTaskExecutionMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.event.ParentTaskExecutionMetadata.id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1170,
  serialized_end=1251,
)


_PARENTNODEEXECUTIONMETADATA = _descriptor.Descriptor(
  name='ParentNodeExecutionMetadata',
  full_name='flyteidl.event.ParentNodeExecutionMetadata',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='id', full_name='flyteidl.event.ParentNodeExecutionMetadata.id', index=0,
      number=1, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
  ],
  serialized_start=1253,
  serialized_end=1294,
)


_TASKEXECUTIONEVENT = _descriptor.Descriptor(
  name='TaskExecutionEvent',
  full_name='flyteidl.event.TaskExecutionEvent',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='task_id', full_name='flyteidl.event.TaskExecutionEvent.task_id', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='parent_node_execution_id', full_name='flyteidl.event.TaskExecutionEvent.parent_node_execution_id', index=1,
      number=2, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='retry_attempt', full_name='flyteidl.event.TaskExecutionEvent.retry_attempt', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase', full_name='flyteidl.event.TaskExecutionEvent.phase', index=3,
      number=4, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='producer_id', full_name='flyteidl.event.TaskExecutionEvent.producer_id', index=4,
      number=5, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='logs', full_name='flyteidl.event.TaskExecutionEvent.logs', index=5,
      number=6, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='occurred_at', full_name='flyteidl.event.TaskExecutionEvent.occurred_at', index=6,
      number=7, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='input_uri', full_name='flyteidl.event.TaskExecutionEvent.input_uri', index=7,
      number=8, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output_uri', full_name='flyteidl.event.TaskExecutionEvent.output_uri', index=8,
      number=9, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='error', full_name='flyteidl.event.TaskExecutionEvent.error', index=9,
      number=10, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='custom_info', full_name='flyteidl.event.TaskExecutionEvent.custom_info', index=10,
      number=11, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='phase_version', full_name='flyteidl.event.TaskExecutionEvent.phase_version', index=11,
      number=12, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
  ],
  extensions=[
  ],
  nested_types=[],
  enum_types=[
  ],
  serialized_options=None,
  is_extendable=False,
  syntax='proto3',
  extension_ranges=[],
  oneofs=[
    _descriptor.OneofDescriptor(
      name='output_result', full_name='flyteidl.event.TaskExecutionEvent.output_result',
      index=0, containing_type=None, fields=[]),
  ],
  serialized_start=1297,
  serialized_end=1792,
)

_WORKFLOWEXECUTIONEVENT.fields_by_name['execution_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._WORKFLOWEXECUTIONIDENTIFIER
_WORKFLOWEXECUTIONEVENT.fields_by_name['phase'].enum_type = flyteidl_dot_core_dot_execution__pb2._WORKFLOWEXECUTION_PHASE
_WORKFLOWEXECUTIONEVENT.fields_by_name['occurred_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_WORKFLOWEXECUTIONEVENT.fields_by_name['error'].message_type = flyteidl_dot_core_dot_execution__pb2._EXECUTIONERROR
_WORKFLOWEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _WORKFLOWEXECUTIONEVENT.fields_by_name['output_uri'])
_WORKFLOWEXECUTIONEVENT.fields_by_name['output_uri'].containing_oneof = _WORKFLOWEXECUTIONEVENT.oneofs_by_name['output_result']
_WORKFLOWEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _WORKFLOWEXECUTIONEVENT.fields_by_name['error'])
_WORKFLOWEXECUTIONEVENT.fields_by_name['error'].containing_oneof = _WORKFLOWEXECUTIONEVENT.oneofs_by_name['output_result']
_NODEEXECUTIONEVENT.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
_NODEEXECUTIONEVENT.fields_by_name['phase'].enum_type = flyteidl_dot_core_dot_execution__pb2._NODEEXECUTION_PHASE
_NODEEXECUTIONEVENT.fields_by_name['occurred_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_NODEEXECUTIONEVENT.fields_by_name['error'].message_type = flyteidl_dot_core_dot_execution__pb2._EXECUTIONERROR
_NODEEXECUTIONEVENT.fields_by_name['workflow_node_metadata'].message_type = _WORKFLOWNODEMETADATA
_NODEEXECUTIONEVENT.fields_by_name['parent_task_metadata'].message_type = _PARENTTASKEXECUTIONMETADATA
_NODEEXECUTIONEVENT.fields_by_name['parent_node_metadata'].message_type = _PARENTNODEEXECUTIONMETADATA
_NODEEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _NODEEXECUTIONEVENT.fields_by_name['output_uri'])
_NODEEXECUTIONEVENT.fields_by_name['output_uri'].containing_oneof = _NODEEXECUTIONEVENT.oneofs_by_name['output_result']
_NODEEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _NODEEXECUTIONEVENT.fields_by_name['error'])
_NODEEXECUTIONEVENT.fields_by_name['error'].containing_oneof = _NODEEXECUTIONEVENT.oneofs_by_name['output_result']
_NODEEXECUTIONEVENT.oneofs_by_name['target_metadata'].fields.append(
  _NODEEXECUTIONEVENT.fields_by_name['workflow_node_metadata'])
_NODEEXECUTIONEVENT.fields_by_name['workflow_node_metadata'].containing_oneof = _NODEEXECUTIONEVENT.oneofs_by_name['target_metadata']
_WORKFLOWNODEMETADATA.fields_by_name['execution_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._WORKFLOWEXECUTIONIDENTIFIER
_PARENTTASKEXECUTIONMETADATA.fields_by_name['id'].message_type = flyteidl_dot_core_dot_identifier__pb2._TASKEXECUTIONIDENTIFIER
_TASKEXECUTIONEVENT.fields_by_name['task_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._IDENTIFIER
_TASKEXECUTIONEVENT.fields_by_name['parent_node_execution_id'].message_type = flyteidl_dot_core_dot_identifier__pb2._NODEEXECUTIONIDENTIFIER
_TASKEXECUTIONEVENT.fields_by_name['phase'].enum_type = flyteidl_dot_core_dot_execution__pb2._TASKEXECUTION_PHASE
_TASKEXECUTIONEVENT.fields_by_name['logs'].message_type = flyteidl_dot_core_dot_execution__pb2._TASKLOG
_TASKEXECUTIONEVENT.fields_by_name['occurred_at'].message_type = google_dot_protobuf_dot_timestamp__pb2._TIMESTAMP
_TASKEXECUTIONEVENT.fields_by_name['error'].message_type = flyteidl_dot_core_dot_execution__pb2._EXECUTIONERROR
_TASKEXECUTIONEVENT.fields_by_name['custom_info'].message_type = google_dot_protobuf_dot_struct__pb2._STRUCT
_TASKEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _TASKEXECUTIONEVENT.fields_by_name['output_uri'])
_TASKEXECUTIONEVENT.fields_by_name['output_uri'].containing_oneof = _TASKEXECUTIONEVENT.oneofs_by_name['output_result']
_TASKEXECUTIONEVENT.oneofs_by_name['output_result'].fields.append(
  _TASKEXECUTIONEVENT.fields_by_name['error'])
_TASKEXECUTIONEVENT.fields_by_name['error'].containing_oneof = _TASKEXECUTIONEVENT.oneofs_by_name['output_result']
DESCRIPTOR.message_types_by_name['WorkflowExecutionEvent'] = _WORKFLOWEXECUTIONEVENT
DESCRIPTOR.message_types_by_name['NodeExecutionEvent'] = _NODEEXECUTIONEVENT
DESCRIPTOR.message_types_by_name['WorkflowNodeMetadata'] = _WORKFLOWNODEMETADATA
DESCRIPTOR.message_types_by_name['ParentTaskExecutionMetadata'] = _PARENTTASKEXECUTIONMETADATA
DESCRIPTOR.message_types_by_name['ParentNodeExecutionMetadata'] = _PARENTNODEEXECUTIONMETADATA
DESCRIPTOR.message_types_by_name['TaskExecutionEvent'] = _TASKEXECUTIONEVENT
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

WorkflowExecutionEvent = _reflection.GeneratedProtocolMessageType('WorkflowExecutionEvent', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWEXECUTIONEVENT,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.WorkflowExecutionEvent)
  ))
_sym_db.RegisterMessage(WorkflowExecutionEvent)

NodeExecutionEvent = _reflection.GeneratedProtocolMessageType('NodeExecutionEvent', (_message.Message,), dict(
  DESCRIPTOR = _NODEEXECUTIONEVENT,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.NodeExecutionEvent)
  ))
_sym_db.RegisterMessage(NodeExecutionEvent)

WorkflowNodeMetadata = _reflection.GeneratedProtocolMessageType('WorkflowNodeMetadata', (_message.Message,), dict(
  DESCRIPTOR = _WORKFLOWNODEMETADATA,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.WorkflowNodeMetadata)
  ))
_sym_db.RegisterMessage(WorkflowNodeMetadata)

ParentTaskExecutionMetadata = _reflection.GeneratedProtocolMessageType('ParentTaskExecutionMetadata', (_message.Message,), dict(
  DESCRIPTOR = _PARENTTASKEXECUTIONMETADATA,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.ParentTaskExecutionMetadata)
  ))
_sym_db.RegisterMessage(ParentTaskExecutionMetadata)

ParentNodeExecutionMetadata = _reflection.GeneratedProtocolMessageType('ParentNodeExecutionMetadata', (_message.Message,), dict(
  DESCRIPTOR = _PARENTNODEEXECUTIONMETADATA,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.ParentNodeExecutionMetadata)
  ))
_sym_db.RegisterMessage(ParentNodeExecutionMetadata)

TaskExecutionEvent = _reflection.GeneratedProtocolMessageType('TaskExecutionEvent', (_message.Message,), dict(
  DESCRIPTOR = _TASKEXECUTIONEVENT,
  __module__ = 'flyteidl.event.event_pb2'
  # @@protoc_insertion_point(class_scope:flyteidl.event.TaskExecutionEvent)
  ))
_sym_db.RegisterMessage(TaskExecutionEvent)


DESCRIPTOR._options = None
# @@protoc_insertion_point(module_scope)
