import 'package:flutter/material.dart';

import '../../../utils/extensions.dart';
import '../../core/constants/dimensions.dart';
import '../../core/widgets/bottom_navbar/bottom_navbar.dart';

class SearchPageView extends StatelessWidget {
  const SearchPageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Padding(
            padding: const EdgeInsets.only(
              left: horizontalPadding,
              bottom: horizontalPadding,
              top: 75,
            ),
            child: Column(
              children: [
                TextFormField(
                  decoration: InputDecoration(
                    hintText: context.l10n.searchInputHolder,
                    prefixIcon: const Icon(Icons.search),
                  ),
                )
              ],
            ),
          ),
          const BottomNavBar(),
        ],
      ),
    );
  }
}
