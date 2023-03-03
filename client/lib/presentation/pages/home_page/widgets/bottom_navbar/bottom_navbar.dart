import 'package:flutter/material.dart';

import '../../../../../utils/extensions.dart';
import '../../../../core/constants/colors.dart';
import '../../../../core/constants/dimensions.dart';

class BottomNavBar extends StatelessWidget {
  const BottomNavBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const radius = Radius.circular(10);

    return Positioned(
      bottom: 0,
      child: Container(
        width: context.screenSize.width * 1,
        decoration: const BoxDecoration(
          color: primary,
          borderRadius: BorderRadius.only(
            topLeft: radius,
            topRight: radius,
          ),
        ),
        child: const Padding(
          padding: EdgeInsets.symmetric(
            vertical: verticalPadding,
            horizontal: horizontalPadding,
          ),
          child: Text('Navbar'),
        ),
      ),
    );
  }
}
