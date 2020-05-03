# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: signrpc/signer.proto

import sys
_b=sys.version_info[0]<3 and (lambda x:x) or (lambda x:x.encode('latin1'))
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from google.protobuf import reflection as _reflection
from google.protobuf import symbol_database as _symbol_database
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()




DESCRIPTOR = _descriptor.FileDescriptor(
  name='signrpc/signer.proto',
  package='signrpc',
  syntax='proto3',
  serialized_options=_b('Z,github.com/lncm/lnd-rpc/v0.9.2/lnrpc/signrpc'),
  serialized_pb=_b('\n\x14signrpc/signer.proto\x12\x07signrpc\"3\n\nKeyLocator\x12\x12\n\nkey_family\x18\x01 \x01(\x05\x12\x11\n\tkey_index\x18\x02 \x01(\x05\"L\n\rKeyDescriptor\x12\x15\n\rraw_key_bytes\x18\x01 \x01(\x0c\x12$\n\x07key_loc\x18\x02 \x01(\x0b\x32\x13.signrpc.KeyLocator\")\n\x05TxOut\x12\r\n\x05value\x18\x01 \x01(\x03\x12\x11\n\tpk_script\x18\x02 \x01(\x0c\"\xc4\x01\n\x0eSignDescriptor\x12(\n\x08key_desc\x18\x01 \x01(\x0b\x32\x16.signrpc.KeyDescriptor\x12\x14\n\x0csingle_tweak\x18\x02 \x01(\x0c\x12\x14\n\x0c\x64ouble_tweak\x18\x03 \x01(\x0c\x12\x16\n\x0ewitness_script\x18\x04 \x01(\x0c\x12\x1e\n\x06output\x18\x05 \x01(\x0b\x32\x0e.signrpc.TxOut\x12\x0f\n\x07sighash\x18\x07 \x01(\r\x12\x13\n\x0binput_index\x18\x08 \x01(\x05\"L\n\x07SignReq\x12\x14\n\x0craw_tx_bytes\x18\x01 \x01(\x0c\x12+\n\nsign_descs\x18\x02 \x03(\x0b\x32\x17.signrpc.SignDescriptor\"\x1c\n\x08SignResp\x12\x10\n\x08raw_sigs\x18\x01 \x03(\x0c\"2\n\x0bInputScript\x12\x0f\n\x07witness\x18\x01 \x03(\x0c\x12\x12\n\nsig_script\x18\x02 \x01(\x0c\">\n\x0fInputScriptResp\x12+\n\rinput_scripts\x18\x01 \x03(\x0b\x32\x14.signrpc.InputScript\"C\n\x0eSignMessageReq\x12\x0b\n\x03msg\x18\x01 \x01(\x0c\x12$\n\x07key_loc\x18\x02 \x01(\x0b\x32\x13.signrpc.KeyLocator\"$\n\x0fSignMessageResp\x12\x11\n\tsignature\x18\x01 \x01(\x0c\"B\n\x10VerifyMessageReq\x12\x0b\n\x03msg\x18\x01 \x01(\x0c\x12\x11\n\tsignature\x18\x02 \x01(\x0c\x12\x0e\n\x06pubkey\x18\x03 \x01(\x0c\"\"\n\x11VerifyMessageResp\x12\r\n\x05valid\x18\x01 \x01(\x08\"R\n\x10SharedKeyRequest\x12\x18\n\x10\x65phemeral_pubkey\x18\x01 \x01(\x0c\x12$\n\x07key_loc\x18\x02 \x01(\x0b\x32\x13.signrpc.KeyLocator\"\'\n\x11SharedKeyResponse\x12\x12\n\nshared_key\x18\x01 \x01(\x0c\x32\xd4\x02\n\x06Signer\x12\x34\n\rSignOutputRaw\x12\x10.signrpc.SignReq\x1a\x11.signrpc.SignResp\x12@\n\x12\x43omputeInputScript\x12\x10.signrpc.SignReq\x1a\x18.signrpc.InputScriptResp\x12@\n\x0bSignMessage\x12\x17.signrpc.SignMessageReq\x1a\x18.signrpc.SignMessageResp\x12\x46\n\rVerifyMessage\x12\x19.signrpc.VerifyMessageReq\x1a\x1a.signrpc.VerifyMessageResp\x12H\n\x0f\x44\x65riveSharedKey\x12\x19.signrpc.SharedKeyRequest\x1a\x1a.signrpc.SharedKeyResponseB.Z,github.com/lncm/lnd-rpc/v0.9.2/lnrpc/signrpcb\x06proto3')
)




_KEYLOCATOR = _descriptor.Descriptor(
  name='KeyLocator',
  full_name='signrpc.KeyLocator',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key_family', full_name='signrpc.KeyLocator.key_family', index=0,
      number=1, type=5, cpp_type=1, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_index', full_name='signrpc.KeyLocator.key_index', index=1,
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
  serialized_start=33,
  serialized_end=84,
)


_KEYDESCRIPTOR = _descriptor.Descriptor(
  name='KeyDescriptor',
  full_name='signrpc.KeyDescriptor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='raw_key_bytes', full_name='signrpc.KeyDescriptor.raw_key_bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_loc', full_name='signrpc.KeyDescriptor.key_loc', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=86,
  serialized_end=162,
)


_TXOUT = _descriptor.Descriptor(
  name='TxOut',
  full_name='signrpc.TxOut',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='value', full_name='signrpc.TxOut.value', index=0,
      number=1, type=3, cpp_type=2, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pk_script', full_name='signrpc.TxOut.pk_script', index=1,
      number=2, type=12, cpp_type=9, label=1,
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
  serialized_start=164,
  serialized_end=205,
)


_SIGNDESCRIPTOR = _descriptor.Descriptor(
  name='SignDescriptor',
  full_name='signrpc.SignDescriptor',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='key_desc', full_name='signrpc.SignDescriptor.key_desc', index=0,
      number=1, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='single_tweak', full_name='signrpc.SignDescriptor.single_tweak', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='double_tweak', full_name='signrpc.SignDescriptor.double_tweak', index=2,
      number=3, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='witness_script', full_name='signrpc.SignDescriptor.witness_script', index=3,
      number=4, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='output', full_name='signrpc.SignDescriptor.output', index=4,
      number=5, type=11, cpp_type=10, label=1,
      has_default_value=False, default_value=None,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sighash', full_name='signrpc.SignDescriptor.sighash', index=5,
      number=7, type=13, cpp_type=3, label=1,
      has_default_value=False, default_value=0,
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='input_index', full_name='signrpc.SignDescriptor.input_index', index=6,
      number=8, type=5, cpp_type=1, label=1,
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
  serialized_start=208,
  serialized_end=404,
)


_SIGNREQ = _descriptor.Descriptor(
  name='SignReq',
  full_name='signrpc.SignReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='raw_tx_bytes', full_name='signrpc.SignReq.raw_tx_bytes', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sign_descs', full_name='signrpc.SignReq.sign_descs', index=1,
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
  serialized_start=406,
  serialized_end=482,
)


_SIGNRESP = _descriptor.Descriptor(
  name='SignResp',
  full_name='signrpc.SignResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='raw_sigs', full_name='signrpc.SignResp.raw_sigs', index=0,
      number=1, type=12, cpp_type=9, label=3,
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
  serialized_start=484,
  serialized_end=512,
)


_INPUTSCRIPT = _descriptor.Descriptor(
  name='InputScript',
  full_name='signrpc.InputScript',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='witness', full_name='signrpc.InputScript.witness', index=0,
      number=1, type=12, cpp_type=9, label=3,
      has_default_value=False, default_value=[],
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='sig_script', full_name='signrpc.InputScript.sig_script', index=1,
      number=2, type=12, cpp_type=9, label=1,
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
  serialized_start=514,
  serialized_end=564,
)


_INPUTSCRIPTRESP = _descriptor.Descriptor(
  name='InputScriptResp',
  full_name='signrpc.InputScriptResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='input_scripts', full_name='signrpc.InputScriptResp.input_scripts', index=0,
      number=1, type=11, cpp_type=10, label=3,
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
  serialized_start=566,
  serialized_end=628,
)


_SIGNMESSAGEREQ = _descriptor.Descriptor(
  name='SignMessageReq',
  full_name='signrpc.SignMessageReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='signrpc.SignMessageReq.msg', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_loc', full_name='signrpc.SignMessageReq.key_loc', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=630,
  serialized_end=697,
)


_SIGNMESSAGERESP = _descriptor.Descriptor(
  name='SignMessageResp',
  full_name='signrpc.SignMessageResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='signature', full_name='signrpc.SignMessageResp.signature', index=0,
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
  serialized_start=699,
  serialized_end=735,
)


_VERIFYMESSAGEREQ = _descriptor.Descriptor(
  name='VerifyMessageReq',
  full_name='signrpc.VerifyMessageReq',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='msg', full_name='signrpc.VerifyMessageReq.msg', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='signature', full_name='signrpc.VerifyMessageReq.signature', index=1,
      number=2, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='pubkey', full_name='signrpc.VerifyMessageReq.pubkey', index=2,
      number=3, type=12, cpp_type=9, label=1,
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
  serialized_start=737,
  serialized_end=803,
)


_VERIFYMESSAGERESP = _descriptor.Descriptor(
  name='VerifyMessageResp',
  full_name='signrpc.VerifyMessageResp',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='valid', full_name='signrpc.VerifyMessageResp.valid', index=0,
      number=1, type=8, cpp_type=7, label=1,
      has_default_value=False, default_value=False,
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
  serialized_start=805,
  serialized_end=839,
)


_SHAREDKEYREQUEST = _descriptor.Descriptor(
  name='SharedKeyRequest',
  full_name='signrpc.SharedKeyRequest',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='ephemeral_pubkey', full_name='signrpc.SharedKeyRequest.ephemeral_pubkey', index=0,
      number=1, type=12, cpp_type=9, label=1,
      has_default_value=False, default_value=_b(""),
      message_type=None, enum_type=None, containing_type=None,
      is_extension=False, extension_scope=None,
      serialized_options=None, file=DESCRIPTOR),
    _descriptor.FieldDescriptor(
      name='key_loc', full_name='signrpc.SharedKeyRequest.key_loc', index=1,
      number=2, type=11, cpp_type=10, label=1,
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
  serialized_start=841,
  serialized_end=923,
)


_SHAREDKEYRESPONSE = _descriptor.Descriptor(
  name='SharedKeyResponse',
  full_name='signrpc.SharedKeyResponse',
  filename=None,
  file=DESCRIPTOR,
  containing_type=None,
  fields=[
    _descriptor.FieldDescriptor(
      name='shared_key', full_name='signrpc.SharedKeyResponse.shared_key', index=0,
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
  serialized_start=925,
  serialized_end=964,
)

_KEYDESCRIPTOR.fields_by_name['key_loc'].message_type = _KEYLOCATOR
_SIGNDESCRIPTOR.fields_by_name['key_desc'].message_type = _KEYDESCRIPTOR
_SIGNDESCRIPTOR.fields_by_name['output'].message_type = _TXOUT
_SIGNREQ.fields_by_name['sign_descs'].message_type = _SIGNDESCRIPTOR
_INPUTSCRIPTRESP.fields_by_name['input_scripts'].message_type = _INPUTSCRIPT
_SIGNMESSAGEREQ.fields_by_name['key_loc'].message_type = _KEYLOCATOR
_SHAREDKEYREQUEST.fields_by_name['key_loc'].message_type = _KEYLOCATOR
DESCRIPTOR.message_types_by_name['KeyLocator'] = _KEYLOCATOR
DESCRIPTOR.message_types_by_name['KeyDescriptor'] = _KEYDESCRIPTOR
DESCRIPTOR.message_types_by_name['TxOut'] = _TXOUT
DESCRIPTOR.message_types_by_name['SignDescriptor'] = _SIGNDESCRIPTOR
DESCRIPTOR.message_types_by_name['SignReq'] = _SIGNREQ
DESCRIPTOR.message_types_by_name['SignResp'] = _SIGNRESP
DESCRIPTOR.message_types_by_name['InputScript'] = _INPUTSCRIPT
DESCRIPTOR.message_types_by_name['InputScriptResp'] = _INPUTSCRIPTRESP
DESCRIPTOR.message_types_by_name['SignMessageReq'] = _SIGNMESSAGEREQ
DESCRIPTOR.message_types_by_name['SignMessageResp'] = _SIGNMESSAGERESP
DESCRIPTOR.message_types_by_name['VerifyMessageReq'] = _VERIFYMESSAGEREQ
DESCRIPTOR.message_types_by_name['VerifyMessageResp'] = _VERIFYMESSAGERESP
DESCRIPTOR.message_types_by_name['SharedKeyRequest'] = _SHAREDKEYREQUEST
DESCRIPTOR.message_types_by_name['SharedKeyResponse'] = _SHAREDKEYRESPONSE
_sym_db.RegisterFileDescriptor(DESCRIPTOR)

KeyLocator = _reflection.GeneratedProtocolMessageType('KeyLocator', (_message.Message,), dict(
  DESCRIPTOR = _KEYLOCATOR,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.KeyLocator)
  ))
_sym_db.RegisterMessage(KeyLocator)

KeyDescriptor = _reflection.GeneratedProtocolMessageType('KeyDescriptor', (_message.Message,), dict(
  DESCRIPTOR = _KEYDESCRIPTOR,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.KeyDescriptor)
  ))
_sym_db.RegisterMessage(KeyDescriptor)

TxOut = _reflection.GeneratedProtocolMessageType('TxOut', (_message.Message,), dict(
  DESCRIPTOR = _TXOUT,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.TxOut)
  ))
_sym_db.RegisterMessage(TxOut)

SignDescriptor = _reflection.GeneratedProtocolMessageType('SignDescriptor', (_message.Message,), dict(
  DESCRIPTOR = _SIGNDESCRIPTOR,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SignDescriptor)
  ))
_sym_db.RegisterMessage(SignDescriptor)

SignReq = _reflection.GeneratedProtocolMessageType('SignReq', (_message.Message,), dict(
  DESCRIPTOR = _SIGNREQ,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SignReq)
  ))
_sym_db.RegisterMessage(SignReq)

SignResp = _reflection.GeneratedProtocolMessageType('SignResp', (_message.Message,), dict(
  DESCRIPTOR = _SIGNRESP,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SignResp)
  ))
_sym_db.RegisterMessage(SignResp)

InputScript = _reflection.GeneratedProtocolMessageType('InputScript', (_message.Message,), dict(
  DESCRIPTOR = _INPUTSCRIPT,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.InputScript)
  ))
_sym_db.RegisterMessage(InputScript)

InputScriptResp = _reflection.GeneratedProtocolMessageType('InputScriptResp', (_message.Message,), dict(
  DESCRIPTOR = _INPUTSCRIPTRESP,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.InputScriptResp)
  ))
_sym_db.RegisterMessage(InputScriptResp)

SignMessageReq = _reflection.GeneratedProtocolMessageType('SignMessageReq', (_message.Message,), dict(
  DESCRIPTOR = _SIGNMESSAGEREQ,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SignMessageReq)
  ))
_sym_db.RegisterMessage(SignMessageReq)

SignMessageResp = _reflection.GeneratedProtocolMessageType('SignMessageResp', (_message.Message,), dict(
  DESCRIPTOR = _SIGNMESSAGERESP,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SignMessageResp)
  ))
_sym_db.RegisterMessage(SignMessageResp)

VerifyMessageReq = _reflection.GeneratedProtocolMessageType('VerifyMessageReq', (_message.Message,), dict(
  DESCRIPTOR = _VERIFYMESSAGEREQ,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.VerifyMessageReq)
  ))
_sym_db.RegisterMessage(VerifyMessageReq)

VerifyMessageResp = _reflection.GeneratedProtocolMessageType('VerifyMessageResp', (_message.Message,), dict(
  DESCRIPTOR = _VERIFYMESSAGERESP,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.VerifyMessageResp)
  ))
_sym_db.RegisterMessage(VerifyMessageResp)

SharedKeyRequest = _reflection.GeneratedProtocolMessageType('SharedKeyRequest', (_message.Message,), dict(
  DESCRIPTOR = _SHAREDKEYREQUEST,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SharedKeyRequest)
  ))
_sym_db.RegisterMessage(SharedKeyRequest)

SharedKeyResponse = _reflection.GeneratedProtocolMessageType('SharedKeyResponse', (_message.Message,), dict(
  DESCRIPTOR = _SHAREDKEYRESPONSE,
  __module__ = 'signrpc.signer_pb2'
  # @@protoc_insertion_point(class_scope:signrpc.SharedKeyResponse)
  ))
_sym_db.RegisterMessage(SharedKeyResponse)


DESCRIPTOR._options = None

_SIGNER = _descriptor.ServiceDescriptor(
  name='Signer',
  full_name='signrpc.Signer',
  file=DESCRIPTOR,
  index=0,
  serialized_options=None,
  serialized_start=967,
  serialized_end=1307,
  methods=[
  _descriptor.MethodDescriptor(
    name='SignOutputRaw',
    full_name='signrpc.Signer.SignOutputRaw',
    index=0,
    containing_service=None,
    input_type=_SIGNREQ,
    output_type=_SIGNRESP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='ComputeInputScript',
    full_name='signrpc.Signer.ComputeInputScript',
    index=1,
    containing_service=None,
    input_type=_SIGNREQ,
    output_type=_INPUTSCRIPTRESP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='SignMessage',
    full_name='signrpc.Signer.SignMessage',
    index=2,
    containing_service=None,
    input_type=_SIGNMESSAGEREQ,
    output_type=_SIGNMESSAGERESP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='VerifyMessage',
    full_name='signrpc.Signer.VerifyMessage',
    index=3,
    containing_service=None,
    input_type=_VERIFYMESSAGEREQ,
    output_type=_VERIFYMESSAGERESP,
    serialized_options=None,
  ),
  _descriptor.MethodDescriptor(
    name='DeriveSharedKey',
    full_name='signrpc.Signer.DeriveSharedKey',
    index=4,
    containing_service=None,
    input_type=_SHAREDKEYREQUEST,
    output_type=_SHAREDKEYRESPONSE,
    serialized_options=None,
  ),
])
_sym_db.RegisterServiceDescriptor(_SIGNER)

DESCRIPTOR.services_by_name['Signer'] = _SIGNER

# @@protoc_insertion_point(module_scope)
