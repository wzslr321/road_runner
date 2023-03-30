import 'package:auto_route/auto_route.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../utils/extensions.dart';
import '../../../../application/search/search_bloc.dart';
import '../../../router/router.dart';
import '../../constants/colors.dart';
import '../../constants/dimensions.dart';
import 'bottom_navbar_element.dart';
import 'extended_search_element.dart';
import 'search_element.dart';

class BottomNavBar extends StatelessWidget {
  const BottomNavBar({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    const radius = Radius.circular(30);

    return Positioned(
      bottom: 0,
      child: Container(
        width: context.screenSize.width * 1,
        height: context.screenSize.height * 0.125,
        decoration: const BoxDecoration(
          color: primary,
          borderRadius: BorderRadius.only(
            topLeft: radius,
            topRight: radius,
          ),
        ),
        child: Padding(
          padding: const EdgeInsets.symmetric(
            vertical: verticalPadding,
            horizontal: horizontalPadding,
          ),
          child: BlocConsumer<SearchBloc, SearchState>(
            listener: (context, state) {
              // TODO: implement listener
            },
            builder: (context, state) {
              if (state is Started) {
                return const ExtendedSearchElement();
              }
              if (state is Submitted) {
                context.router.push(const SearchResultsRoute());
              }
              return const Row(
                mainAxisAlignment: MainAxisAlignment.spaceAround,
                children: [
                  BottomNavbarElement(
                    icon: Icons.home,
                    redirectionRoute: HomeRoute(),
                  ),
                  SearchElement(),
                  BottomNavbarElement(
                    icon: Icons.person,
                    redirectionRoute: ProfileRoute(),
                  ),
                ],
              );
            },
          ),
        ),
      ),
    );
  }
}
