import 'package:auto_route/auto_route.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

class BottomNavbarElement extends StatelessWidget {
  const BottomNavbarElement({
    Key? key,
    required this.icon,
    required this.redirectionRoute,
  }) : super(key: key);

  final IconData icon;
  final PageRouteInfo redirectionRoute;

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () => context.router.push(redirectionRoute),
      child: Icon(
        icon,
        color: Colors.white,
        size: 35,
      ),
    );
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(DiagnosticsProperty<IconData>('icon', icon))
      ..properties.add(DiagnosticsProperty<PageRouteInfo>(
          'redirectionRoute', redirectionRoute));
  }
}
