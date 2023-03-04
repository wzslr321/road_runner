import 'package:flutter/material.dart';

import '../../../../../gen/assets.gen.dart';
import '../../../../../utils/extensions.dart';

class MapLocation extends StatelessWidget {
  const MapLocation({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      height: context.screenSize.height * 1,
      width: context.screenSize.width * 1,
      child: Assets.images.googleMapsMock.image(),
    );
  }
}
