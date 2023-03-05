import 'package:flutter/material.dart';

import '../../../../utils/extensions.dart';

class Greeting extends StatelessWidget {
  const Greeting({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Padding(
      padding: EdgeInsets.only(
        top: context.screenSize.height * 0.1,
        bottom: context.screenSize.height * 0.05,
      ),
      child: Column(
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(
            context.l10n.unauthorizedGreeting,
            style: const TextStyle(
              fontSize: 24,
              fontWeight: FontWeight.w600,
            ),
          ),
          const Divider(
            height: 10,
          ),
          Text(
            context.l10n.unauthorizedFeaturesInfo,
            style: TextStyle(
              fontSize: 16,
              color: Colors.grey.withOpacity(0.9),
            ),
          ),
        ],
      ),
    );
  }
}
