import 'package:flutter/material.dart';

import '../../../utils/extensions.dart';
import '../../core/constants/dimensions.dart';
import '../../core/widgets/bottom_navbar/bottom_navbar.dart';
import 'widgets/ex_widgets.dart';
import 'widgets/greeting.dart';
import 'widgets/settings_element/settings_element_log_in.dart';

class ProfilePageView extends StatelessWidget {
  const ProfilePageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          Padding(
            padding: const EdgeInsets.symmetric(
              horizontal: horizontalPadding,
              vertical: verticalPadding,
            ),
            child: Column(
              children: [
                const Greeting(),
                SingleChildScrollView(
                  child: Column(
                    children: [
                      SettingsElementLogIn(
                        title: context.l10n.settingsLogIn,
                      ),
                      SettingsElement(
                        title: context.l10n.settingsCommunityGuidelines,
                        onTap: () => {},
                      ),
                      SettingsElement(
                        title: context.l10n.settingsPrivacyPolicy,
                        onTap: () => {},
                      ),
                      SettingsElement(
                        title: context.l10n.settingsTermsOfService,
                        onTap: () => {},
                      ),
                    ],
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
