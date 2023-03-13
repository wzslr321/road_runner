import 'package:fpdart/fpdart.dart';

import '../../gen/proto/user.pb.dart';
import '../users/call_failure.dart';

typedef CallResult = Either<CallFailure, UserData>;
