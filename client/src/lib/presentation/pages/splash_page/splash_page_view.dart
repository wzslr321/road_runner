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
        backgroundColor: const Color(0xffC7AEAE),
        body: InkWell(
          onTap: () {
            context.router.popForced();
            context.router.push(const HomeRoute());
          },
          child: Padding(
            padding: const EdgeInsets.all(10),
            child: Assets.images.splashScreen.image(
              fit: BoxFit.fill,
              height: double.infinity,
              width: double.infinity,
            ),
          ),
        ),
      ),
    );
  }
}
