import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'settings_element_core.dart';

class SettingsElement extends StatelessWidget {
  const SettingsElement({
    Key? key,
    required this.title,
    required this.onTap,
  }) : super(key: key);

  final String title;
  final Function() onTap;

  @override
  Widget build(BuildContext context) {
    return SettingsElementCore(title: title, onTap: onTap);
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties
      ..add(StringProperty('title', title))
      ..add(ObjectFlagProperty<Function()>.has('onTap', onTap));
  }
}
