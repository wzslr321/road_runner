import 'package:flutter/foundation.dart';
import 'package:fpdart/fpdart.dart';
import 'package:grpc/grpc.dart';

import '../domain/core/typedefs.dart';
import '../domain/users/call_failure.dart';
import '../domain/users/users_repository_i.dart';
import '../gen/proto/user.pbgrpc.dart';

class UsersRepository implements UsersRepositoryI {
  @override
  Future<CallResult> getUser(String username) async {
    final channel = ClientChannel(
      '192.168.0.100',
      port: 50051,
      options: ChannelOptions(
        credentials: const ChannelCredentials.insecure(),
        codecRegistry:
            CodecRegistry(codecs: const [GzipCodec(), IdentityCodec()]),
      ),
    );

    final client = UsersClient(channel);
    try {
      final response = await client.getUser(
        GetUserRequest()..username = username,
      );

      await channel.shutdown();
      return right(response);
    } on GrpcError catch (err) {
      if (kDebugMode) {
        debugPrint(err.toString());
      }
      await channel.shutdown();
      return left(const CallFailure.serverError());
    }
  }
}
