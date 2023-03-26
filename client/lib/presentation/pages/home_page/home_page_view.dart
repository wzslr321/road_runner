import 'package:flutter/foundation.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:users_repository/users_repository.dart';

import '../../core/widgets/bottom_navbar/bottom_navbar.dart';
import 'widgets/ex_widgets.dart';

class HomePageView extends StatelessWidget {
  const HomePageView({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Stack(
        children: [
          const MapLocation(),
          const BottomNavBar(),
          Center(
            child: ElevatedButton(
                onPressed: () async {
                  final _result =
                      await context.read<UsersRepository>().getUser('antek');
                 final result = _result.getOrElse((l) => GetUserResponse());
                 if(kDebugMode){
                   debugPrint(result.toString());
                 }

                },
                child: const Text('click me')),
          )
        ],
      ),
    );
  }
}
