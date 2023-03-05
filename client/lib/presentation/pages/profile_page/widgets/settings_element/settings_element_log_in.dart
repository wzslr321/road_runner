import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';

import 'settings_element_core.dart';

class SettingsElementLogIn extends StatelessWidget {
  const SettingsElementLogIn({
    Key? key,
    required this.title,
  }) : super(key: key);

  final String title;

  @override
  Widget build(BuildContext context) {
    return SettingsElementCore(title: title, onTap: () => {});
  }

  @override
  void debugFillProperties(DiagnosticPropertiesBuilder properties) {
    super.debugFillProperties(properties);
    properties.add(StringProperty('title', title));
  }
}
