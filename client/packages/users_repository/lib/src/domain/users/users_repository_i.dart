import '../core/typedefs.dart';

abstract class UsersRepositoryI{
  Future<CallResult> getUser(String username);
  Future<CallResult> updateUser(String username);
}
