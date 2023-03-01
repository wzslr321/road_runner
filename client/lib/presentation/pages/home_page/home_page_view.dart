import 'dart:math';

import 'package:flutter/material.dart';

import '../../../utils/extensions.dart';

class HomePageView extends StatelessWidget {
  const HomePageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Text(context.l10n.helloWorld),
    );
  }
}
