import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:road_runner/presentation/core/my_app_widget.dart';
import 'package:road_runner/presentation/pages/ex_pages.dart';
import 'package:road_runner/presentation/router/router.dart';

void main() {
  group('MyApp', () {
    testWidgets('renders SplashPage', (tester) async {
      await tester.pumpWidget(
        MyAppWidget(),
      );
      expect(find.byType(MaterialApp), findsOneWidget);
    });
  });
}
