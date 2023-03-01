import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import '../../../gen/assets.gen.dart';

import '../../router/router.dart';

class SplashPageView extends StatelessWidget {
  const SplashPageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SafeArea(
      child: Scaffold(
        body: InkWell(
          onTap: () => context.router.push(const HomeRoute()),
          child: Assets.images.splashPage.image(),
        ),
      ),
    );
  }
}
