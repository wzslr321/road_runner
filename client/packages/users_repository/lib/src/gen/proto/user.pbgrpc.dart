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
  static final _$createUser =
      $grpc.ClientMethod<$0.CreateUserRequest, $0.CreateUserResponse>(
          '/queries.Users/CreateUser',
          ($0.CreateUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.CreateUserResponse.fromBuffer(value));
  static final _$getUser =
      $grpc.ClientMethod<$0.GetUserRequest, $0.GetUserResponse>(
          '/queries.Users/GetUser',
          ($0.GetUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.GetUserResponse.fromBuffer(value));
  static final _$updateUser =
      $grpc.ClientMethod<$0.UpdateUserRequest, $0.UpdateUserResponse>(
          '/queries.Users/UpdateUser',
          ($0.UpdateUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.UpdateUserResponse.fromBuffer(value));
  static final _$deleteUser =
      $grpc.ClientMethod<$0.DeleteUserRequest, $0.DeleteUserResponse>(
          '/queries.Users/DeleteUser',
          ($0.DeleteUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.DeleteUserResponse.fromBuffer(value));
  static final _$loginUser =
      $grpc.ClientMethod<$0.LoginUserRequest, $0.LoginUserResponse>(
          '/queries.Users/LoginUser',
          ($0.LoginUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.LoginUserResponse.fromBuffer(value));
  static final _$logoutUser =
      $grpc.ClientMethod<$0.LogoutUserRequest, $0.LogoutUserResponse>(
          '/queries.Users/LogoutUser',
          ($0.LogoutUserRequest value) => value.writeToBuffer(),
          ($core.List<$core.int> value) =>
              $0.LogoutUserResponse.fromBuffer(value));

  UsersClient($grpc.ClientChannel channel,
      {$grpc.CallOptions? options,
      $core.Iterable<$grpc.ClientInterceptor>? interceptors})
      : super(channel, options: options, interceptors: interceptors);

  $grpc.ResponseFuture<$0.CreateUserResponse> createUser(
      $0.CreateUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$createUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.GetUserResponse> getUser($0.GetUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$getUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.UpdateUserResponse> updateUser(
      $0.UpdateUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$updateUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.DeleteUserResponse> deleteUser(
      $0.DeleteUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$deleteUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.LoginUserResponse> loginUser(
      $0.LoginUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$loginUser, request, options: options);
  }

  $grpc.ResponseFuture<$0.LogoutUserResponse> logoutUser(
      $0.LogoutUserRequest request,
      {$grpc.CallOptions? options}) {
    return $createUnaryCall(_$logoutUser, request, options: options);
  }
}

abstract class UsersServiceBase extends $grpc.Service {
  $core.String get $name => 'queries.Users';

  UsersServiceBase() {
    $addMethod($grpc.ServiceMethod<$0.CreateUserRequest, $0.CreateUserResponse>(
        'CreateUser',
        createUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.CreateUserRequest.fromBuffer(value),
        ($0.CreateUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.GetUserRequest, $0.GetUserResponse>(
        'GetUser',
        getUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.GetUserRequest.fromBuffer(value),
        ($0.GetUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.UpdateUserRequest, $0.UpdateUserResponse>(
        'UpdateUser',
        updateUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.UpdateUserRequest.fromBuffer(value),
        ($0.UpdateUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.DeleteUserRequest, $0.DeleteUserResponse>(
        'DeleteUser',
        deleteUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.DeleteUserRequest.fromBuffer(value),
        ($0.DeleteUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LoginUserRequest, $0.LoginUserResponse>(
        'LoginUser',
        loginUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LoginUserRequest.fromBuffer(value),
        ($0.LoginUserResponse value) => value.writeToBuffer()));
    $addMethod($grpc.ServiceMethod<$0.LogoutUserRequest, $0.LogoutUserResponse>(
        'LogoutUser',
        logoutUser_Pre,
        false,
        false,
        ($core.List<$core.int> value) => $0.LogoutUserRequest.fromBuffer(value),
        ($0.LogoutUserResponse value) => value.writeToBuffer()));
  }

  $async.Future<$0.CreateUserResponse> createUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.CreateUserRequest> request) async {
    return createUser(call, await request);
  }

  $async.Future<$0.GetUserResponse> getUser_Pre(
      $grpc.ServiceCall call, $async.Future<$0.GetUserRequest> request) async {
    return getUser(call, await request);
  }

  $async.Future<$0.UpdateUserResponse> updateUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.UpdateUserRequest> request) async {
    return updateUser(call, await request);
  }

  $async.Future<$0.DeleteUserResponse> deleteUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.DeleteUserRequest> request) async {
    return deleteUser(call, await request);
  }

  $async.Future<$0.LoginUserResponse> loginUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.LoginUserRequest> request) async {
    return loginUser(call, await request);
  }

  $async.Future<$0.LogoutUserResponse> logoutUser_Pre($grpc.ServiceCall call,
      $async.Future<$0.LogoutUserRequest> request) async {
    return logoutUser(call, await request);
  }

  $async.Future<$0.CreateUserResponse> createUser(
      $grpc.ServiceCall call, $0.CreateUserRequest request);
  $async.Future<$0.GetUserResponse> getUser(
      $grpc.ServiceCall call, $0.GetUserRequest request);
  $async.Future<$0.UpdateUserResponse> updateUser(
      $grpc.ServiceCall call, $0.UpdateUserRequest request);
  $async.Future<$0.DeleteUserResponse> deleteUser(
      $grpc.ServiceCall call, $0.DeleteUserRequest request);
  $async.Future<$0.LoginUserResponse> loginUser(
      $grpc.ServiceCall call, $0.LoginUserRequest request);
  $async.Future<$0.LogoutUserResponse> logoutUser(
      $grpc.ServiceCall call, $0.LogoutUserRequest request);
}
