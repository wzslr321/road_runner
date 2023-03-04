import 'package:flutter/material.dart';

import '../../../../../utils/extensions.dart';
import '../../../router/router.dart';
import '../../constants/colors.dart';
import '../../constants/dimensions.dart';
import 'bottom_navbar_element.dart';

class BottomNavBar extends StatelessWidget {
  const BottomNavBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const radius = Radius.circular(30);

    return Positioned(
      bottom: 0,
      child: Container(
        width: context.screenSize.width * 1,
        height: context.screenSize.height * 0.125,
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
            child: Row(
              mainAxisAlignment: MainAxisAlignment.spaceAround,
              children: [
                BottomNavbarElement(
                    icon: Icons.home, redirectionRoute: HomeRoute()),
                BottomNavbarElement(
                  icon: Icons.search,
                  redirectionRoute: SearchRoute(),
                ),
                BottomNavbarElement(
                  icon: Icons.person,
                  redirectionRoute: ProfileRoute(),
                ),
              ],
            )),
      ),
    );
  }
}
