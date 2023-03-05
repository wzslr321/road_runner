import 'package:auto_route/auto_route.dart';

import 'package:flutter/cupertino.dart';

// don't remove, the app won't compile without the above and below import, though IDE shows these are not used.
import 'package:flutter/material.dart';

import '../pages/ex_pages.dart';

part 'router.gr.dart';

@MaterialAutoRouter(
  replaceInRouteName: 'Page,Route',
  routes: <AutoRoute>[
    AutoRoute<SplashPage>(page: SplashPage, initial: true),
    AutoRoute<HomePage>(page: HomePage),
    AutoRoute<SearchPage>(page: SearchPage),
    AutoRoute<ProfilePage>(page: ProfilePage),
    AutoRoute<SearchResultsPage>(page: SearchResultsPage),
  ],
)
class AppRouter extends _$AppRouter {}
