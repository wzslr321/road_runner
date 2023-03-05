import 'app_localizations.dart';

/// The translations for English (`en`).
class AppLocalizationsEn extends AppLocalizations {
  AppLocalizationsEn([String locale = 'en']) : super(locale);

  @override
  String get unauthorizedGreeting => 'Hello There!';

  @override
  String get unauthorizedFeaturesInfo => 'Access more features by logging in.';

  @override
  String get settingsLogIn => 'Log in';

  @override
  String get settingsTermsOfService => 'Terms of service';

  @override
  String get settingsCommunityGuidelines => 'Community Guidelines';

  @override
  String get settingsPrivacyPolicy => 'Privacy Policy';

  @override
  String get searchInputHolder => 'Where do you want to go?';
}
