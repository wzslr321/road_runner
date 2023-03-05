import 'app_localizations.dart';

/// The translations for Polish (`pl`).
class AppLocalizationsPl extends AppLocalizations {
  AppLocalizationsPl([String locale = 'pl']) : super(locale);

  @override
  String get unauthorizedGreeting => 'Cześć użytkowniku!';

  @override
  String get unauthorizedFeaturesInfo => 'Zyskaj dostęp do większej ilości funkcji, logując się!';

  @override
  String get settingsLogIn => 'Zaloguj się';

  @override
  String get settingsTermsOfService => 'Zasady użytkowania';

  @override
  String get settingsCommunityGuidelines => 'Wytyczne społeczności';

  @override
  String get settingsPrivacyPolicy => 'Polityka prywatności';

  @override
  String get searchInputHolder => 'Dokąd chcesz się udać?';
}
