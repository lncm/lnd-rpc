# Generated by the gRPC Python protocol compiler plugin. DO NOT EDIT!
import grpc

from watchtowerrpc import watchtower_pb2 as watchtowerrpc_dot_watchtower__pb2


class WatchtowerStub(object):
  # missing associated documentation comment in .proto file
  pass

  def __init__(self, channel):
    """Constructor.

    Args:
      channel: A grpc.Channel.
    """
    self.GetInfo = channel.unary_unary(
        '/watchtowerrpc.Watchtower/GetInfo',
        request_serializer=watchtowerrpc_dot_watchtower__pb2.GetInfoRequest.SerializeToString,
        response_deserializer=watchtowerrpc_dot_watchtower__pb2.GetInfoResponse.FromString,
        )


class WatchtowerServicer(object):
  # missing associated documentation comment in .proto file
  pass

  def GetInfo(self, request, context):
    """* lncli: tower info
    GetInfo returns general information concerning the companion watchtower
    including it's public key and URIs where the server is currently
    listening for clients.
    """
    context.set_code(grpc.StatusCode.UNIMPLEMENTED)
    context.set_details('Method not implemented!')
    raise NotImplementedError('Method not implemented!')


def add_WatchtowerServicer_to_server(servicer, server):
  rpc_method_handlers = {
      'GetInfo': grpc.unary_unary_rpc_method_handler(
          servicer.GetInfo,
          request_deserializer=watchtowerrpc_dot_watchtower__pb2.GetInfoRequest.FromString,
          response_serializer=watchtowerrpc_dot_watchtower__pb2.GetInfoResponse.SerializeToString,
      ),
  }
  generic_handler = grpc.method_handlers_generic_handler(
      'watchtowerrpc.Watchtower', rpc_method_handlers)
  server.add_generic_rpc_handlers((generic_handler,))