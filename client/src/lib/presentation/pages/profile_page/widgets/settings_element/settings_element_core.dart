import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import '../../../../../utils/extensions.dart';
import '../../../../core/constants/colors.dart';

class SettingsElementCore extends StatelessWidget {
  const SettingsElementCore({
    Key? key,
    required this.title,
    required this.onTap,
  }) : super(key: key);

  final String title;
  final Function() onTap;

  @override
  Widget build(BuildContext context) {
    final width = context.screenSize.width;
    final height = context.screenSize.height;

    return Padding(
      padding: EdgeInsets.only(bottom: height * 0.01),
      child: Container(
        padding: EdgeInsets.symmetric(
          horizontal: width * 0.025,
          vertical: height * 0.0225,
        ),
        decoration: BoxDecoration(
          color: primary,
          borderRadius: BorderRadius.circular(15),
        ),
        child: Row(
          children: [
            Text(
              title,
              style: TextStyle(
                color: Colors.white.withOpacity(0.9),
                fontSize: 16,
                letterSpacing: 0.5,
              ),
            ),
            Expanded(
              child: Align(
                alignment: Alignment.centerRight,
                child: InkWell(
                  onTap: onTap,
                  child: Icon(
                    Icons.arrow_forward_ios,
                    color: Colors.white.withOpacity(0.9),
                    size: 15,
                  ),
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(StringProperty('title', title))
      ..add(ObjectFlagProperty<Function()>.has('onTap', onTap));
  }
}
