import 'package:fpdart/fpdart.dart';

import '../../gen/proto/users.pb.dart';
import '../users/call_failure.dart';

typedef CallResult = Either<CallFailure, GetUserResponse>;
