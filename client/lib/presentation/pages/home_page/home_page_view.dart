import 'dart:math';

import 'package:flutter/material.dart';

import '../../../utils/extensions.dart';
import 'widgets/bottom_navbar/bottom_navbar.dart';

class HomePageView extends StatelessWidget {
  const HomePageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return const Scaffold(
      body: Stack(
        children: [
          BottomNavBar(),
        ],
      ),
    );
  }
}
