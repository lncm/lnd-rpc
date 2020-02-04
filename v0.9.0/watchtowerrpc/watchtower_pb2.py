# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: watchtowerrpc/watchtower.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='watchtowerrpc/watchtower.proto',
  package='watchtowerrpc',
  syntax='proto3',
  serialized_options=_b('Z2github.com/lncm/lnd-rpc/v0.9.0/lnrpc/watchtowerrpc'),
  serialized_pb=_b('\n\x1ewatchtowerrpc/watchtower.proto\x12\rwatchtowerrpc\"\x10\n\x0eGetInfoRequest\"[\n\x0fGetInfoResponse\x12\x16\n\x06pubkey\x18\x01 \x01(\x0cR\x06pubkey\x12\x1c\n\tlisteners\x18\x02 \x03(\tR\tlisteners\x12\x12\n\x04uris\x18\x03 \x03(\tR\x04uris2V\n\nWatchtower\x12H\n\x07GetInfo\x12\x1d.watchtowerrpc.GetInfoRequest\x1a\x1e.watchtowerrpc.GetInfoResponseB4Z2github.com/lncm/lnd-rpc/v0.9.0/lnrpc/watchtowerrpcb\x06proto3')
)




_GETINFOREQUEST = _descriptor.Descriptor(
  name='GetInfoRequest',
  full_name='watchtowerrpc.GetInfoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
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
  serialized_start=49,
  serialized_end=65,
)


_GETINFORESPONSE = _descriptor.Descriptor(
  name='GetInfoResponse',
  full_name='watchtowerrpc.GetInfoResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='watchtowerrpc.GetInfoResponse.pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubkey', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='listeners', full_name='watchtowerrpc.GetInfoResponse.listeners', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='listeners', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='uris', full_name='watchtowerrpc.GetInfoResponse.uris', index=2,
      number=3, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='uris', file=DESCRIPTOR),
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
  serialized_start=67,
  serialized_end=158,
)

DESCRIPTOR.message_types_by_name['GetInfoRequest'] = _GETINFOREQUEST
DESCRIPTOR.message_types_by_name['GetInfoResponse'] = _GETINFORESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

GetInfoRequest = _reflection.GeneratedProtocolMessageType('GetInfoRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETINFOREQUEST,
  __module__ = 'watchtowerrpc.watchtower_pb2'
  # @@protoc_insertion_point(class_scope:watchtowerrpc.GetInfoRequest)
  ))
_sym_db.RegisterMessage(GetInfoRequest)

GetInfoResponse = _reflection.GeneratedProtocolMessageType('GetInfoResponse', (_message.Message,), dict(
  DESCRIPTOR = _GETINFORESPONSE,
  __module__ = 'watchtowerrpc.watchtower_pb2'
  # @@protoc_insertion_point(class_scope:watchtowerrpc.GetInfoResponse)
  ))
_sym_db.RegisterMessage(GetInfoResponse)


DESCRIPTOR._options = None

_WATCHTOWER = _descriptor.ServiceDescriptor(
  name='Watchtower',
  full_name='watchtowerrpc.Watchtower',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=160,
  serialized_end=246,
  methods=[
  _descriptor.MethodDescriptor(
    name='GetInfo',
    full_name='watchtowerrpc.Watchtower.GetInfo',
    index=0,
    containing_service=None,
    input_type=_GETINFOREQUEST,
    output_type=_GETINFORESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_WATCHTOWER)

DESCRIPTOR.services_by_name['Watchtower'] = _WATCHTOWER

# @@protoc_insertion_point(module_scope)
