// **************************************************************************
// AutoRouteGenerator
// **************************************************************************

// GENERATED CODE - DO NOT MODIFY BY HAND

// **************************************************************************
// AutoRouteGenerator
// **************************************************************************
//
// ignore_for_file: type=lint

part of 'router.dart';

class _$AppRouter extends RootStackRouter {
  _$AppRouter([GlobalKey<NavigatorState>? navigatorKey]) : super(navigatorKey);

  @override
  final Map<String, PageFactory> pagesMap = {
    SplashRoute.name: (routeData) {
      return MaterialPageX<SplashPage>(
        routeData: routeData,
        child: const SplashPage(),
      );
    },
    HomeRoute.name: (routeData) {
      return MaterialPageX<HomePage>(
        routeData: routeData,
        child: const HomePage(),
      );
    },
    SearchRoute.name: (routeData) {
      return MaterialPageX<SearchPage>(
        routeData: routeData,
        child: const SearchPage(),
      );
    },
    ProfileRoute.name: (routeData) {
      return MaterialPageX<ProfilePage>(
        routeData: routeData,
        child: const ProfilePage(),
      );
    },
    SearchResultsRoute.name: (routeData) {
      return MaterialPageX<SearchResultsPage>(
        routeData: routeData,
        child: const SearchResultsPage(),
      );
    },
  };

  @override
  List<RouteConfig> get routes => [
        RouteConfig(
          SplashRoute.name,
          path: '/',
        ),
        RouteConfig(
          HomeRoute.name,
          path: '/home-page',
        ),
        RouteConfig(
          SearchRoute.name,
          path: '/search-page',
        ),
        RouteConfig(
          ProfileRoute.name,
          path: '/profile-page',
        ),
        RouteConfig(
          SearchResultsRoute.name,
          path: '/search-results-page',
        ),
      ];
}

/// generated route for
/// [SplashPage]
class SplashRoute extends PageRouteInfo<void> {
  const SplashRoute()
      : super(
          SplashRoute.name,
          path: '/',
        );

  static const String name = 'SplashRoute';
}

/// generated route for
/// [HomePage]
class HomeRoute extends PageRouteInfo<void> {
  const HomeRoute()
      : super(
          HomeRoute.name,
          path: '/home-page',
        );

  static const String name = 'HomeRoute';
}

/// generated route for
/// [SearchPage]
class SearchRoute extends PageRouteInfo<void> {
  const SearchRoute()
      : super(
          SearchRoute.name,
          path: '/search-page',
        );

  static const String name = 'SearchRoute';
}

/// generated route for
/// [ProfilePage]
class ProfileRoute extends PageRouteInfo<void> {
  const ProfileRoute()
      : super(
          ProfileRoute.name,
          path: '/profile-page',
        );

  static const String name = 'ProfileRoute';
}

/// generated route for
/// [SearchResultsPage]
class SearchResultsRoute extends PageRouteInfo<void> {
  const SearchResultsRoute()
      : super(
          SearchResultsRoute.name,
          path: '/search-results-page',
        );

  static const String name = 'SearchResultsRoute';
}
