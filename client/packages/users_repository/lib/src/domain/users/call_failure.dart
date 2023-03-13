import 'package:freezed_annotation/freezed_annotation.dart';

part 'call_failure.freezed.dart';

@freezed
class CallFailure with _$CallFailure {
  const factory CallFailure.serverError() = ServerErrror;
}
