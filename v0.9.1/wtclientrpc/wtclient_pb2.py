# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: wtclientrpc/wtclient.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='wtclientrpc/wtclient.proto',
  package='wtclientrpc',
  syntax='proto3',
  serialized_options=_b('Z0github.com/lncm/lnd-rpc/v0.9.1/lnrpc/wtclientrpc'),
  serialized_pb=_b('\n\x1awtclientrpc/wtclient.proto\x12\x0bwtclientrpc\"C\n\x0f\x41\x64\x64TowerRequest\x12\x16\n\x06pubkey\x18\x01 \x01(\x0cR\x06pubkey\x12\x18\n\x07\x61\x64\x64ress\x18\x02 \x01(\tR\x07\x61\x64\x64ress\"\x12\n\x10\x41\x64\x64TowerResponse\"F\n\x12RemoveTowerRequest\x12\x16\n\x06pubkey\x18\x01 \x01(\x0cR\x06pubkey\x12\x18\n\x07\x61\x64\x64ress\x18\x02 \x01(\tR\x07\x61\x64\x64ress\"\x15\n\x13RemoveTowerResponse\"Y\n\x13GetTowerInfoRequest\x12\x16\n\x06pubkey\x18\x01 \x01(\x0cR\x06pubkey\x12*\n\x10include_sessions\x18\x02 \x01(\x08R\x10include_sessions\"\xb4\x01\n\x0cTowerSession\x12 \n\x0bnum_backups\x18\x01 \x01(\rR\x0bnum_backups\x12\x30\n\x13num_pending_backups\x18\x02 \x01(\rR\x13num_pending_backups\x12 \n\x0bmax_backups\x18\x03 \x01(\rR\x0bmax_backups\x12.\n\x12sweep_sat_per_byte\x18\x04 \x01(\rR\x12sweep_sat_per_byte\"\xd4\x01\n\x05Tower\x12\x16\n\x06pubkey\x18\x01 \x01(\x0cR\x06pubkey\x12\x1c\n\taddresses\x18\x02 \x03(\tR\taddresses\x12:\n\x18\x61\x63tive_session_candidate\x18\x03 \x01(\x08R\x18\x61\x63tive_session_candidate\x12\"\n\x0cnum_sessions\x18\x04 \x01(\rR\x0cnum_sessions\x12\x35\n\x08sessions\x18\x05 \x03(\x0b\x32\x19.wtclientrpc.TowerSessionR\x08sessions\"?\n\x11ListTowersRequest\x12*\n\x10include_sessions\x18\x01 \x01(\x08R\x10include_sessions\"@\n\x12ListTowersResponse\x12*\n\x06towers\x18\x01 \x03(\x0b\x32\x12.wtclientrpc.TowerR\x06towers\"\x0e\n\x0cStatsRequest\"\x81\x02\n\rStatsResponse\x12 \n\x0bnum_backups\x18\x01 \x01(\rR\x0bnum_backups\x12\x30\n\x13num_pending_backups\x18\x02 \x01(\rR\x13num_pending_backups\x12.\n\x12num_failed_backups\x18\x03 \x01(\rR\x12num_failed_backups\x12\x34\n\x15num_sessions_acquired\x18\x04 \x01(\rR\x15num_sessions_acquired\x12\x36\n\x16num_sessions_exhausted\x18\x05 \x01(\rR\x16num_sessions_exhausted\"\x0f\n\rPolicyRequest\"b\n\x0ePolicyResponse\x12 \n\x0bmax_updates\x18\x01 \x01(\rR\x0bmax_updates\x12.\n\x12sweep_sat_per_byte\x18\x02 \x01(\rR\x12sweep_sat_per_byte2\xc5\x03\n\x10WatchtowerClient\x12G\n\x08\x41\x64\x64Tower\x12\x1c.wtclientrpc.AddTowerRequest\x1a\x1d.wtclientrpc.AddTowerResponse\x12P\n\x0bRemoveTower\x12\x1f.wtclientrpc.RemoveTowerRequest\x1a .wtclientrpc.RemoveTowerResponse\x12M\n\nListTowers\x12\x1e.wtclientrpc.ListTowersRequest\x1a\x1f.wtclientrpc.ListTowersResponse\x12\x44\n\x0cGetTowerInfo\x12 .wtclientrpc.GetTowerInfoRequest\x1a\x12.wtclientrpc.Tower\x12>\n\x05Stats\x12\x19.wtclientrpc.StatsRequest\x1a\x1a.wtclientrpc.StatsResponse\x12\x41\n\x06Policy\x12\x1a.wtclientrpc.PolicyRequest\x1a\x1b.wtclientrpc.PolicyResponseB2Z0github.com/lncm/lnd-rpc/v0.9.1/lnrpc/wtclientrpcb\x06proto3')
)




_ADDTOWERREQUEST = _descriptor.Descriptor(
  name='AddTowerRequest',
  full_name='wtclientrpc.AddTowerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='wtclientrpc.AddTowerRequest.pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubkey', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address', full_name='wtclientrpc.AddTowerRequest.address', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='address', file=DESCRIPTOR),
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
  serialized_start=43,
  serialized_end=110,
)


_ADDTOWERRESPONSE = _descriptor.Descriptor(
  name='AddTowerResponse',
  full_name='wtclientrpc.AddTowerResponse',
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
  serialized_start=112,
  serialized_end=130,
)


_REMOVETOWERREQUEST = _descriptor.Descriptor(
  name='RemoveTowerRequest',
  full_name='wtclientrpc.RemoveTowerRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='wtclientrpc.RemoveTowerRequest.pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubkey', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='address', full_name='wtclientrpc.RemoveTowerRequest.address', index=1,
      number=2, type=9, cpp_type=9, label=1,
      has_default_value=False, default_value=_b("").decode('utf-8'),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='address', file=DESCRIPTOR),
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
  serialized_start=132,
  serialized_end=202,
)


_REMOVETOWERRESPONSE = _descriptor.Descriptor(
  name='RemoveTowerResponse',
  full_name='wtclientrpc.RemoveTowerResponse',
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
  serialized_start=204,
  serialized_end=225,
)


_GETTOWERINFOREQUEST = _descriptor.Descriptor(
  name='GetTowerInfoRequest',
  full_name='wtclientrpc.GetTowerInfoRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='wtclientrpc.GetTowerInfoRequest.pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubkey', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='include_sessions', full_name='wtclientrpc.GetTowerInfoRequest.include_sessions', index=1,
      number=2, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='include_sessions', file=DESCRIPTOR),
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
  serialized_start=227,
  serialized_end=316,
)


_TOWERSESSION = _descriptor.Descriptor(
  name='TowerSession',
  full_name='wtclientrpc.TowerSession',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='num_backups', full_name='wtclientrpc.TowerSession.num_backups', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_pending_backups', full_name='wtclientrpc.TowerSession.num_pending_backups', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_pending_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='max_backups', full_name='wtclientrpc.TowerSession.max_backups', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='max_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sweep_sat_per_byte', full_name='wtclientrpc.TowerSession.sweep_sat_per_byte', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='sweep_sat_per_byte', file=DESCRIPTOR),
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
  serialized_start=319,
  serialized_end=499,
)


_TOWER = _descriptor.Descriptor(
  name='Tower',
  full_name='wtclientrpc.Tower',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='wtclientrpc.Tower.pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pubkey', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='addresses', full_name='wtclientrpc.Tower.addresses', index=1,
      number=2, type=9, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='addresses', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='active_session_candidate', full_name='wtclientrpc.Tower.active_session_candidate', index=2,
      number=3, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='active_session_candidate', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_sessions', full_name='wtclientrpc.Tower.num_sessions', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_sessions', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sessions', full_name='wtclientrpc.Tower.sessions', index=4,
      number=5, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='sessions', file=DESCRIPTOR),
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
  serialized_start=502,
  serialized_end=714,
)


_LISTTOWERSREQUEST = _descriptor.Descriptor(
  name='ListTowersRequest',
  full_name='wtclientrpc.ListTowersRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='include_sessions', full_name='wtclientrpc.ListTowersRequest.include_sessions', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='include_sessions', file=DESCRIPTOR),
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
  serialized_start=716,
  serialized_end=779,
)


_LISTTOWERSRESPONSE = _descriptor.Descriptor(
  name='ListTowersResponse',
  full_name='wtclientrpc.ListTowersResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='towers', full_name='wtclientrpc.ListTowersResponse.towers', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='towers', file=DESCRIPTOR),
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
  serialized_start=781,
  serialized_end=845,
)


_STATSREQUEST = _descriptor.Descriptor(
  name='StatsRequest',
  full_name='wtclientrpc.StatsRequest',
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
  serialized_start=847,
  serialized_end=861,
)


_STATSRESPONSE = _descriptor.Descriptor(
  name='StatsResponse',
  full_name='wtclientrpc.StatsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='num_backups', full_name='wtclientrpc.StatsResponse.num_backups', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_pending_backups', full_name='wtclientrpc.StatsResponse.num_pending_backups', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_pending_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_failed_backups', full_name='wtclientrpc.StatsResponse.num_failed_backups', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_failed_backups', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_sessions_acquired', full_name='wtclientrpc.StatsResponse.num_sessions_acquired', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_sessions_acquired', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='num_sessions_exhausted', full_name='wtclientrpc.StatsResponse.num_sessions_exhausted', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='num_sessions_exhausted', file=DESCRIPTOR),
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
  serialized_start=864,
  serialized_end=1121,
)


_POLICYREQUEST = _descriptor.Descriptor(
  name='PolicyRequest',
  full_name='wtclientrpc.PolicyRequest',
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
  serialized_start=1123,
  serialized_end=1138,
)


_POLICYRESPONSE = _descriptor.Descriptor(
  name='PolicyResponse',
  full_name='wtclientrpc.PolicyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='max_updates', full_name='wtclientrpc.PolicyResponse.max_updates', index=0,
      number=1, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='max_updates', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sweep_sat_per_byte', full_name='wtclientrpc.PolicyResponse.sweep_sat_per_byte', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='sweep_sat_per_byte', file=DESCRIPTOR),
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
  serialized_start=1140,
  serialized_end=1238,
)

_TOWER.fields_by_name['sessions'].message_type = _TOWERSESSION
_LISTTOWERSRESPONSE.fields_by_name['towers'].message_type = _TOWER
DESCRIPTOR.message_types_by_name['AddTowerRequest'] = _ADDTOWERREQUEST
DESCRIPTOR.message_types_by_name['AddTowerResponse'] = _ADDTOWERRESPONSE
DESCRIPTOR.message_types_by_name['RemoveTowerRequest'] = _REMOVETOWERREQUEST
DESCRIPTOR.message_types_by_name['RemoveTowerResponse'] = _REMOVETOWERRESPONSE
DESCRIPTOR.message_types_by_name['GetTowerInfoRequest'] = _GETTOWERINFOREQUEST
DESCRIPTOR.message_types_by_name['TowerSession'] = _TOWERSESSION
DESCRIPTOR.message_types_by_name['Tower'] = _TOWER
DESCRIPTOR.message_types_by_name['ListTowersRequest'] = _LISTTOWERSREQUEST
DESCRIPTOR.message_types_by_name['ListTowersResponse'] = _LISTTOWERSRESPONSE
DESCRIPTOR.message_types_by_name['StatsRequest'] = _STATSREQUEST
DESCRIPTOR.message_types_by_name['StatsResponse'] = _STATSRESPONSE
DESCRIPTOR.message_types_by_name['PolicyRequest'] = _POLICYREQUEST
DESCRIPTOR.message_types_by_name['PolicyResponse'] = _POLICYRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

AddTowerRequest = _reflection.GeneratedProtocolMessageType('AddTowerRequest', (_message.Message,), dict(
  DESCRIPTOR = _ADDTOWERREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.AddTowerRequest)
  ))
_sym_db.RegisterMessage(AddTowerRequest)

AddTowerResponse = _reflection.GeneratedProtocolMessageType('AddTowerResponse', (_message.Message,), dict(
  DESCRIPTOR = _ADDTOWERRESPONSE,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.AddTowerResponse)
  ))
_sym_db.RegisterMessage(AddTowerResponse)

RemoveTowerRequest = _reflection.GeneratedProtocolMessageType('RemoveTowerRequest', (_message.Message,), dict(
  DESCRIPTOR = _REMOVETOWERREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.RemoveTowerRequest)
  ))
_sym_db.RegisterMessage(RemoveTowerRequest)

RemoveTowerResponse = _reflection.GeneratedProtocolMessageType('RemoveTowerResponse', (_message.Message,), dict(
  DESCRIPTOR = _REMOVETOWERRESPONSE,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.RemoveTowerResponse)
  ))
_sym_db.RegisterMessage(RemoveTowerResponse)

GetTowerInfoRequest = _reflection.GeneratedProtocolMessageType('GetTowerInfoRequest', (_message.Message,), dict(
  DESCRIPTOR = _GETTOWERINFOREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.GetTowerInfoRequest)
  ))
_sym_db.RegisterMessage(GetTowerInfoRequest)

TowerSession = _reflection.GeneratedProtocolMessageType('TowerSession', (_message.Message,), dict(
  DESCRIPTOR = _TOWERSESSION,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.TowerSession)
  ))
_sym_db.RegisterMessage(TowerSession)

Tower = _reflection.GeneratedProtocolMessageType('Tower', (_message.Message,), dict(
  DESCRIPTOR = _TOWER,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.Tower)
  ))
_sym_db.RegisterMessage(Tower)

ListTowersRequest = _reflection.GeneratedProtocolMessageType('ListTowersRequest', (_message.Message,), dict(
  DESCRIPTOR = _LISTTOWERSREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.ListTowersRequest)
  ))
_sym_db.RegisterMessage(ListTowersRequest)

ListTowersResponse = _reflection.GeneratedProtocolMessageType('ListTowersResponse', (_message.Message,), dict(
  DESCRIPTOR = _LISTTOWERSRESPONSE,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.ListTowersResponse)
  ))
_sym_db.RegisterMessage(ListTowersResponse)

StatsRequest = _reflection.GeneratedProtocolMessageType('StatsRequest', (_message.Message,), dict(
  DESCRIPTOR = _STATSREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.StatsRequest)
  ))
_sym_db.RegisterMessage(StatsRequest)

StatsResponse = _reflection.GeneratedProtocolMessageType('StatsResponse', (_message.Message,), dict(
  DESCRIPTOR = _STATSRESPONSE,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.StatsResponse)
  ))
_sym_db.RegisterMessage(StatsResponse)

PolicyRequest = _reflection.GeneratedProtocolMessageType('PolicyRequest', (_message.Message,), dict(
  DESCRIPTOR = _POLICYREQUEST,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.PolicyRequest)
  ))
_sym_db.RegisterMessage(PolicyRequest)

PolicyResponse = _reflection.GeneratedProtocolMessageType('PolicyResponse', (_message.Message,), dict(
  DESCRIPTOR = _POLICYRESPONSE,
  __module__ = 'wtclientrpc.wtclient_pb2'
  # @@protoc_insertion_point(class_scope:wtclientrpc.PolicyResponse)
  ))
_sym_db.RegisterMessage(PolicyResponse)


DESCRIPTOR._options = None

_WATCHTOWERCLIENT = _descriptor.ServiceDescriptor(
  name='WatchtowerClient',
  full_name='wtclientrpc.WatchtowerClient',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1241,
  serialized_end=1694,
  methods=[
  _descriptor.MethodDescriptor(
    name='AddTower',
    full_name='wtclientrpc.WatchtowerClient.AddTower',
    index=0,
    containing_service=None,
    input_type=_ADDTOWERREQUEST,
    output_type=_ADDTOWERRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='RemoveTower',
    full_name='wtclientrpc.WatchtowerClient.RemoveTower',
    index=1,
    containing_service=None,
    input_type=_REMOVETOWERREQUEST,
    output_type=_REMOVETOWERRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ListTowers',
    full_name='wtclientrpc.WatchtowerClient.ListTowers',
    index=2,
    containing_service=None,
    input_type=_LISTTOWERSREQUEST,
    output_type=_LISTTOWERSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='GetTowerInfo',
    full_name='wtclientrpc.WatchtowerClient.GetTowerInfo',
    index=3,
    containing_service=None,
    input_type=_GETTOWERINFOREQUEST,
    output_type=_TOWER,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Stats',
    full_name='wtclientrpc.WatchtowerClient.Stats',
    index=4,
    containing_service=None,
    input_type=_STATSREQUEST,
    output_type=_STATSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='Policy',
    full_name='wtclientrpc.WatchtowerClient.Policy',
    index=5,
    containing_service=None,
    input_type=_POLICYREQUEST,
    output_type=_POLICYRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_WATCHTOWERCLIENT)

DESCRIPTOR.services_by_name['WatchtowerClient'] = _WATCHTOWERCLIENT

# @@protoc_insertion_point(module_scope)
