///
//  Generated code. Do not modify.
//  source: user.proto
//
// @dart = 2.12
// ignore_for_file: annotate_overrides,camel_case_types,constant_identifier_names,directives_ordering,library_prefixes,non_constant_identifier_names,prefer_final_fields,return_of_invalid_type,unnecessary_const,unnecessary_import,unnecessary_this,unused_import,unused_shown_name

import 'dart:async' as $async;

import 'dart:core' as $core;

import 'package:grpc/service_api.dart' as $grpc;
import 'user.pb.dart' as $0;
export 'user.pb.dart';

class UsersClient extends $grpc.Client {
  static final _$getUser = $grpc.ClientMethod<$0.GetUserRequest, $0.UserData>(
      '/queries.Users/GetUser',
      ($0.GetUserRequest value) => value.writeToBuffer(),
      ($core.List<$core.int> value) => $0.UserData.fromBuffer(value));
  static final _$deleteUser =
      $grpc.ClientMethod<$0.DeleteUserRequest, $0.DeleteUserResponse>(
          '/queries.Users/DeleteUser',
          ($0.DeleteUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DeleteUserResponse.fromBuffer(value));

  UsersClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.UserData> getUser($0.GetUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.DeleteUserResponse> deleteUser(
      $0.DeleteUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteUser, request, options: options);
  }
}

abstract class UsersServiceBase extends $grpc.Service {
  $core.String get $name => 'queries.Users';

  UsersServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.GetUserRequest, $0.UserData>(
        'GetUser',
        getUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetUserRequest.fromBuffer(value),
        ($0.UserData value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteUserRequest, $0.DeleteUserResponse>(
        'DeleteUser',
        deleteUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DeleteUserRequest.fromBuffer(value),
        ($0.DeleteUserResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.UserData> getUser_Pre(
      $grpc.ServiceCall call, $async.Future<$0.GetUserRequest> request) async {
    return getUser(call, await request);
  }

  $async.Future<$0.DeleteUserResponse> deleteUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.DeleteUserRequest> request) async {
    return deleteUser(call, await request);
  }

  $async.Future<$0.UserData> getUser(
      $grpc.ServiceCall call, $0.GetUserRequest request);
  $async.Future<$0.DeleteUserResponse> deleteUser(
      $grpc.ServiceCall call, $0.DeleteUserRequest request);
}
