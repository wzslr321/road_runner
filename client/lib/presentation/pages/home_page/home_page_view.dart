import 'dart:math';

import 'package:flutter/material.dart';

import '../../../utils/extensions.dart';
import '../../core/widgets/bottom_navbar/bottom_navbar.dart';
import 'widgets/ex_widgets.dart';

class HomePageView extends StatelessWidget {
  const HomePageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Stack(
        children: [
          MapLocation(),
          BottomNavBar(),
        ],
      ),
    );
  }
}
