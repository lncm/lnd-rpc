# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: walletrpc/walletkit.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf.internal import enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from lnrpc import rpc_pb2 as lnrpc_dot_rpc__pb2
from signrpc import signer_pb2 as signrpc_dot_signer__pb2


DESCRIPTOR = _descriptor.FileDescriptor(
  name='walletrpc/walletkit.proto',
  package='walletrpc',
  syntax='proto3',
  serialized_options=_b('Z.github.com/lncm/lnd-rpc/v0.8.1/lnrpc/walletrpc'),
  serialized_pb=_b('\n\x19walletrpc/walletkit.proto\x12\twalletrpc\x1a\x0flnrpc/rpc.proto\x1a\x14signrpc/signer.proto\"6\n\x06KeyReq\x12\x18\n\x10key_finger_print\x18\x01 \x01(\x05\x12\x12\n\nkey_family\x18\x02 \x01(\x05\"\r\n\x0b\x41\x64\x64rRequest\"\x1c\n\x0c\x41\x64\x64rResponse\x12\x0c\n\x04\x61\x64\x64r\x18\x01 \x01(\t\"\x1d\n\x0bTransaction\x12\x0e\n\x06tx_hex\x18\x01 \x01(\x0c\"(\n\x0fPublishResponse\x12\x15\n\rpublish_error\x18\x01 \x01(\t\"I\n\x12SendOutputsRequest\x12\x12\n\nsat_per_kw\x18\x01 \x01(\x03\x12\x1f\n\x07outputs\x18\x02 \x03(\x0b\x32\x0e.signrpc.TxOut\"%\n\x13SendOutputsResponse\x12\x0e\n\x06raw_tx\x18\x01 \x01(\x0c\")\n\x12\x45stimateFeeRequest\x12\x13\n\x0b\x63onf_target\x18\x01 \x01(\x05\")\n\x13\x45stimateFeeResponse\x12\x12\n\nsat_per_kw\x18\x01 \x01(\x03\"\xa1\x02\n\x0cPendingSweep\x12+\n\x08outpoint\x18\x01 \x01(\x0b\x32\x0f.lnrpc.OutPointR\x08outpoint\x12:\n\x0cwitness_type\x18\x02 \x01(\x0e\x32\x16.walletrpc.WitnessTypeR\x0cwitness_type\x12\x1e\n\namount_sat\x18\x03 \x01(\rR\namount_sat\x12\"\n\x0csat_per_byte\x18\x04 \x01(\rR\x0csat_per_byte\x12.\n\x12\x62roadcast_attempts\x18\x05 \x01(\rR\x12\x62roadcast_attempts\x12\x34\n\x15next_broadcast_height\x18\x06 \x01(\rR\x15next_broadcast_height\"\x16\n\x14PendingSweepsRequest\"X\n\x15PendingSweepsResponse\x12?\n\x0epending_sweeps\x18\x01 \x03(\x0b\x32\x17.walletrpc.PendingSweepR\x0epending_sweeps\"\x83\x01\n\x0e\x42umpFeeRequest\x12+\n\x08outpoint\x18\x01 \x01(\x0b\x32\x0f.lnrpc.OutPointR\x08outpoint\x12 \n\x0btarget_conf\x18\x02 \x01(\rR\x0btarget_conf\x12\"\n\x0csat_per_byte\x18\x03 \x01(\rR\x0csat_per_byte\"\x11\n\x0f\x42umpFeeResponse*\x82\x03\n\x0bWitnessType\x12\x13\n\x0fUNKNOWN_WITNESS\x10\x00\x12\x18\n\x14\x43OMMITMENT_TIME_LOCK\x10\x01\x12\x17\n\x13\x43OMMITMENT_NO_DELAY\x10\x02\x12\x15\n\x11\x43OMMITMENT_REVOKE\x10\x03\x12\x17\n\x13HTLC_OFFERED_REVOKE\x10\x04\x12\x18\n\x14HTLC_ACCEPTED_REVOKE\x10\x05\x12%\n!HTLC_OFFERED_TIMEOUT_SECOND_LEVEL\x10\x06\x12&\n\"HTLC_ACCEPTED_SUCCESS_SECOND_LEVEL\x10\x07\x12\x1f\n\x1bHTLC_OFFERED_REMOTE_TIMEOUT\x10\x08\x12 \n\x1cHTLC_ACCEPTED_REMOTE_SUCCESS\x10\t\x12\x1c\n\x18HTLC_SECOND_LEVEL_REVOKE\x10\n\x12\x14\n\x10WITNESS_KEY_HASH\x10\x0b\x12\x1b\n\x17NESTED_WITNESS_KEY_HASH\x10\x0c\x32\xba\x04\n\tWalletKit\x12:\n\rDeriveNextKey\x12\x11.walletrpc.KeyReq\x1a\x16.signrpc.KeyDescriptor\x12\x38\n\tDeriveKey\x12\x13.signrpc.KeyLocator\x1a\x16.signrpc.KeyDescriptor\x12;\n\x08NextAddr\x12\x16.walletrpc.AddrRequest\x1a\x17.walletrpc.AddrResponse\x12H\n\x12PublishTransaction\x12\x16.walletrpc.Transaction\x1a\x1a.walletrpc.PublishResponse\x12L\n\x0bSendOutputs\x12\x1d.walletrpc.SendOutputsRequest\x1a\x1e.walletrpc.SendOutputsResponse\x12L\n\x0b\x45stimateFee\x12\x1d.walletrpc.EstimateFeeRequest\x1a\x1e.walletrpc.EstimateFeeResponse\x12R\n\rPendingSweeps\x12\x1f.walletrpc.PendingSweepsRequest\x1a .walletrpc.PendingSweepsResponse\x12@\n\x07\x42umpFee\x12\x19.walletrpc.BumpFeeRequest\x1a\x1a.walletrpc.BumpFeeResponseB0Z.github.com/lncm/lnd-rpc/v0.8.1/lnrpc/walletrpcb\x06proto3')
  ,
  dependencies=[lnrpc_dot_rpc__pb2.DESCRIPTOR,signrpc_dot_signer__pb2.DESCRIPTOR,])

_WITNESSTYPE = _descriptor.EnumDescriptor(
  name='WitnessType',
  full_name='walletrpc.WitnessType',
  filename=None,
  file=DESCRIPTOR,
  values=[
    _descriptor.EnumValueDescriptor(
      name='UNKNOWN_WITNESS', index=0, number=0,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMMITMENT_TIME_LOCK', index=1, number=1,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMMITMENT_NO_DELAY', index=2, number=2,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='COMMITMENT_REVOKE', index=3, number=3,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_OFFERED_REVOKE', index=4, number=4,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_ACCEPTED_REVOKE', index=5, number=5,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_OFFERED_TIMEOUT_SECOND_LEVEL', index=6, number=6,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_ACCEPTED_SUCCESS_SECOND_LEVEL', index=7, number=7,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_OFFERED_REMOTE_TIMEOUT', index=8, number=8,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_ACCEPTED_REMOTE_SUCCESS', index=9, number=9,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='HTLC_SECOND_LEVEL_REVOKE', index=10, number=10,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='WITNESS_KEY_HASH', index=11, number=11,
      serialized_options=None,
      type=None),
    _descriptor.EnumValueDescriptor(
      name='NESTED_WITNESS_KEY_HASH', index=12, number=12,
      serialized_options=None,
      type=None),
  ],
  containing_type=None,
  serialized_options=None,
  serialized_start=1013,
  serialized_end=1399,
)
_sym_db.RegisterEnumDescriptor(_WITNESSTYPE)

WitnessType = enum_type_wrapper.EnumTypeWrapper(_WITNESSTYPE)
UNKNOWN_WITNESS = 0
COMMITMENT_TIME_LOCK = 1
COMMITMENT_NO_DELAY = 2
COMMITMENT_REVOKE = 3
HTLC_OFFERED_REVOKE = 4
HTLC_ACCEPTED_REVOKE = 5
HTLC_OFFERED_TIMEOUT_SECOND_LEVEL = 6
HTLC_ACCEPTED_SUCCESS_SECOND_LEVEL = 7
HTLC_OFFERED_REMOTE_TIMEOUT = 8
HTLC_ACCEPTED_REMOTE_SUCCESS = 9
HTLC_SECOND_LEVEL_REVOKE = 10
WITNESS_KEY_HASH = 11
NESTED_WITNESS_KEY_HASH = 12



_KEYREQ = _descriptor.Descriptor(
  name='KeyReq',
  full_name='walletrpc.KeyReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key_finger_print', full_name='walletrpc.KeyReq.key_finger_print', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_family', full_name='walletrpc.KeyReq.key_family', index=1,
      number=2, type=5, cpp_type=1, label=1,
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
  ],
  serialized_start=79,
  serialized_end=133,
)


_ADDRREQUEST = _descriptor.Descriptor(
  name='AddrRequest',
  full_name='walletrpc.AddrRequest',
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
  serialized_start=135,
  serialized_end=148,
)


_ADDRRESPONSE = _descriptor.Descriptor(
  name='AddrResponse',
  full_name='walletrpc.AddrResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='addr', full_name='walletrpc.AddrResponse.addr', index=0,
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
  serialized_start=150,
  serialized_end=178,
)


_TRANSACTION = _descriptor.Descriptor(
  name='Transaction',
  full_name='walletrpc.Transaction',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='tx_hex', full_name='walletrpc.Transaction.tx_hex', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=180,
  serialized_end=209,
)


_PUBLISHRESPONSE = _descriptor.Descriptor(
  name='PublishResponse',
  full_name='walletrpc.PublishResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='publish_error', full_name='walletrpc.PublishResponse.publish_error', index=0,
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
  serialized_start=211,
  serialized_end=251,
)


_SENDOUTPUTSREQUEST = _descriptor.Descriptor(
  name='SendOutputsRequest',
  full_name='walletrpc.SendOutputsRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sat_per_kw', full_name='walletrpc.SendOutputsRequest.sat_per_kw', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='outputs', full_name='walletrpc.SendOutputsRequest.outputs', index=1,
      number=2, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
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
  serialized_start=253,
  serialized_end=326,
)


_SENDOUTPUTSRESPONSE = _descriptor.Descriptor(
  name='SendOutputsResponse',
  full_name='walletrpc.SendOutputsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='raw_tx', full_name='walletrpc.SendOutputsResponse.raw_tx', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
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
  serialized_start=328,
  serialized_end=365,
)


_ESTIMATEFEEREQUEST = _descriptor.Descriptor(
  name='EstimateFeeRequest',
  full_name='walletrpc.EstimateFeeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='conf_target', full_name='walletrpc.EstimateFeeRequest.conf_target', index=0,
      number=1, type=5, cpp_type=1, label=1,
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
  ],
  serialized_start=367,
  serialized_end=408,
)


_ESTIMATEFEERESPONSE = _descriptor.Descriptor(
  name='EstimateFeeResponse',
  full_name='walletrpc.EstimateFeeResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='sat_per_kw', full_name='walletrpc.EstimateFeeResponse.sat_per_kw', index=0,
      number=1, type=3, cpp_type=2, label=1,
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
  ],
  serialized_start=410,
  serialized_end=451,
)


_PENDINGSWEEP = _descriptor.Descriptor(
  name='PendingSweep',
  full_name='walletrpc.PendingSweep',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='outpoint', full_name='walletrpc.PendingSweep.outpoint', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='outpoint', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='witness_type', full_name='walletrpc.PendingSweep.witness_type', index=1,
      number=2, type=14, cpp_type=8, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='witness_type', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='amount_sat', full_name='walletrpc.PendingSweep.amount_sat', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='amount_sat', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sat_per_byte', full_name='walletrpc.PendingSweep.sat_per_byte', index=3,
      number=4, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='sat_per_byte', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='broadcast_attempts', full_name='walletrpc.PendingSweep.broadcast_attempts', index=4,
      number=5, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='broadcast_attempts', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='next_broadcast_height', full_name='walletrpc.PendingSweep.next_broadcast_height', index=5,
      number=6, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='next_broadcast_height', file=DESCRIPTOR),
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
  serialized_start=454,
  serialized_end=743,
)


_PENDINGSWEEPSREQUEST = _descriptor.Descriptor(
  name='PendingSweepsRequest',
  full_name='walletrpc.PendingSweepsRequest',
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
  serialized_start=745,
  serialized_end=767,
)


_PENDINGSWEEPSRESPONSE = _descriptor.Descriptor(
  name='PendingSweepsResponse',
  full_name='walletrpc.PendingSweepsResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='pending_sweeps', full_name='walletrpc.PendingSweepsResponse.pending_sweeps', index=0,
      number=1, type=11, cpp_type=10, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='pending_sweeps', file=DESCRIPTOR),
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
  serialized_start=769,
  serialized_end=857,
)


_BUMPFEEREQUEST = _descriptor.Descriptor(
  name='BumpFeeRequest',
  full_name='walletrpc.BumpFeeRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='outpoint', full_name='walletrpc.BumpFeeRequest.outpoint', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='outpoint', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='target_conf', full_name='walletrpc.BumpFeeRequest.target_conf', index=1,
      number=2, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='target_conf', file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sat_per_byte', full_name='walletrpc.BumpFeeRequest.sat_per_byte', index=2,
      number=3, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, json_name='sat_per_byte', file=DESCRIPTOR),
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
  serialized_start=860,
  serialized_end=991,
)


_BUMPFEERESPONSE = _descriptor.Descriptor(
  name='BumpFeeResponse',
  full_name='walletrpc.BumpFeeResponse',
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
  serialized_start=993,
  serialized_end=1010,
)

_SENDOUTPUTSREQUEST.fields_by_name['outputs'].message_type = signrpc_dot_signer__pb2._TXOUT
_PENDINGSWEEP.fields_by_name['outpoint'].message_type = lnrpc_dot_rpc__pb2._OUTPOINT
_PENDINGSWEEP.fields_by_name['witness_type'].enum_type = _WITNESSTYPE
_PENDINGSWEEPSRESPONSE.fields_by_name['pending_sweeps'].message_type = _PENDINGSWEEP
_BUMPFEEREQUEST.fields_by_name['outpoint'].message_type = lnrpc_dot_rpc__pb2._OUTPOINT
DESCRIPTOR.message_types_by_name['KeyReq'] = _KEYREQ
DESCRIPTOR.message_types_by_name['AddrRequest'] = _ADDRREQUEST
DESCRIPTOR.message_types_by_name['AddrResponse'] = _ADDRRESPONSE
DESCRIPTOR.message_types_by_name['Transaction'] = _TRANSACTION
DESCRIPTOR.message_types_by_name['PublishResponse'] = _PUBLISHRESPONSE
DESCRIPTOR.message_types_by_name['SendOutputsRequest'] = _SENDOUTPUTSREQUEST
DESCRIPTOR.message_types_by_name['SendOutputsResponse'] = _SENDOUTPUTSRESPONSE
DESCRIPTOR.message_types_by_name['EstimateFeeRequest'] = _ESTIMATEFEEREQUEST
DESCRIPTOR.message_types_by_name['EstimateFeeResponse'] = _ESTIMATEFEERESPONSE
DESCRIPTOR.message_types_by_name['PendingSweep'] = _PENDINGSWEEP
DESCRIPTOR.message_types_by_name['PendingSweepsRequest'] = _PENDINGSWEEPSREQUEST
DESCRIPTOR.message_types_by_name['PendingSweepsResponse'] = _PENDINGSWEEPSRESPONSE
DESCRIPTOR.message_types_by_name['BumpFeeRequest'] = _BUMPFEEREQUEST
DESCRIPTOR.message_types_by_name['BumpFeeResponse'] = _BUMPFEERESPONSE
DESCRIPTOR.enum_types_by_name['WitnessType'] = _WITNESSTYPE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

KeyReq = _reflection.GeneratedProtocolMessageType('KeyReq', (_message.Message,), dict(
  DESCRIPTOR = _KEYREQ,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.KeyReq)
  ))
_sym_db.RegisterMessage(KeyReq)

AddrRequest = _reflection.GeneratedProtocolMessageType('AddrRequest', (_message.Message,), dict(
  DESCRIPTOR = _ADDRREQUEST,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.AddrRequest)
  ))
_sym_db.RegisterMessage(AddrRequest)

AddrResponse = _reflection.GeneratedProtocolMessageType('AddrResponse', (_message.Message,), dict(
  DESCRIPTOR = _ADDRRESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.AddrResponse)
  ))
_sym_db.RegisterMessage(AddrResponse)

Transaction = _reflection.GeneratedProtocolMessageType('Transaction', (_message.Message,), dict(
  DESCRIPTOR = _TRANSACTION,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.Transaction)
  ))
_sym_db.RegisterMessage(Transaction)

PublishResponse = _reflection.GeneratedProtocolMessageType('PublishResponse', (_message.Message,), dict(
  DESCRIPTOR = _PUBLISHRESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.PublishResponse)
  ))
_sym_db.RegisterMessage(PublishResponse)

SendOutputsRequest = _reflection.GeneratedProtocolMessageType('SendOutputsRequest', (_message.Message,), dict(
  DESCRIPTOR = _SENDOUTPUTSREQUEST,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.SendOutputsRequest)
  ))
_sym_db.RegisterMessage(SendOutputsRequest)

SendOutputsResponse = _reflection.GeneratedProtocolMessageType('SendOutputsResponse', (_message.Message,), dict(
  DESCRIPTOR = _SENDOUTPUTSRESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.SendOutputsResponse)
  ))
_sym_db.RegisterMessage(SendOutputsResponse)

EstimateFeeRequest = _reflection.GeneratedProtocolMessageType('EstimateFeeRequest', (_message.Message,), dict(
  DESCRIPTOR = _ESTIMATEFEEREQUEST,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.EstimateFeeRequest)
  ))
_sym_db.RegisterMessage(EstimateFeeRequest)

EstimateFeeResponse = _reflection.GeneratedProtocolMessageType('EstimateFeeResponse', (_message.Message,), dict(
  DESCRIPTOR = _ESTIMATEFEERESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.EstimateFeeResponse)
  ))
_sym_db.RegisterMessage(EstimateFeeResponse)

PendingSweep = _reflection.GeneratedProtocolMessageType('PendingSweep', (_message.Message,), dict(
  DESCRIPTOR = _PENDINGSWEEP,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.PendingSweep)
  ))
_sym_db.RegisterMessage(PendingSweep)

PendingSweepsRequest = _reflection.GeneratedProtocolMessageType('PendingSweepsRequest', (_message.Message,), dict(
  DESCRIPTOR = _PENDINGSWEEPSREQUEST,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.PendingSweepsRequest)
  ))
_sym_db.RegisterMessage(PendingSweepsRequest)

PendingSweepsResponse = _reflection.GeneratedProtocolMessageType('PendingSweepsResponse', (_message.Message,), dict(
  DESCRIPTOR = _PENDINGSWEEPSRESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.PendingSweepsResponse)
  ))
_sym_db.RegisterMessage(PendingSweepsResponse)

BumpFeeRequest = _reflection.GeneratedProtocolMessageType('BumpFeeRequest', (_message.Message,), dict(
  DESCRIPTOR = _BUMPFEEREQUEST,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.BumpFeeRequest)
  ))
_sym_db.RegisterMessage(BumpFeeRequest)

BumpFeeResponse = _reflection.GeneratedProtocolMessageType('BumpFeeResponse', (_message.Message,), dict(
  DESCRIPTOR = _BUMPFEERESPONSE,
  __module__ = 'walletrpc.walletkit_pb2'
  # @@protoc_insertion_point(class_scope:walletrpc.BumpFeeResponse)
  ))
_sym_db.RegisterMessage(BumpFeeResponse)


DESCRIPTOR._options = None

_WALLETKIT = _descriptor.ServiceDescriptor(
  name='WalletKit',
  full_name='walletrpc.WalletKit',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=1402,
  serialized_end=1972,
  methods=[
  _descriptor.MethodDescriptor(
    name='DeriveNextKey',
    full_name='walletrpc.WalletKit.DeriveNextKey',
    index=0,
    containing_service=None,
    input_type=_KEYREQ,
    output_type=signrpc_dot_signer__pb2._KEYDESCRIPTOR,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeriveKey',
    full_name='walletrpc.WalletKit.DeriveKey',
    index=1,
    containing_service=None,
    input_type=signrpc_dot_signer__pb2._KEYLOCATOR,
    output_type=signrpc_dot_signer__pb2._KEYDESCRIPTOR,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='NextAddr',
    full_name='walletrpc.WalletKit.NextAddr',
    index=2,
    containing_service=None,
    input_type=_ADDRREQUEST,
    output_type=_ADDRRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PublishTransaction',
    full_name='walletrpc.WalletKit.PublishTransaction',
    index=3,
    containing_service=None,
    input_type=_TRANSACTION,
    output_type=_PUBLISHRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SendOutputs',
    full_name='walletrpc.WalletKit.SendOutputs',
    index=4,
    containing_service=None,
    input_type=_SENDOUTPUTSREQUEST,
    output_type=_SENDOUTPUTSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='EstimateFee',
    full_name='walletrpc.WalletKit.EstimateFee',
    index=5,
    containing_service=None,
    input_type=_ESTIMATEFEEREQUEST,
    output_type=_ESTIMATEFEERESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='PendingSweeps',
    full_name='walletrpc.WalletKit.PendingSweeps',
    index=6,
    containing_service=None,
    input_type=_PENDINGSWEEPSREQUEST,
    output_type=_PENDINGSWEEPSRESPONSE,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='BumpFee',
    full_name='walletrpc.WalletKit.BumpFee',
    index=7,
    containing_service=None,
    input_type=_BUMPFEEREQUEST,
    output_type=_BUMPFEERESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_WALLETKIT)

DESCRIPTOR.services_by_name['WalletKit'] = _WALLETKIT

# @@protoc_insertion_point(module_scope)
