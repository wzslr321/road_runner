import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../application/search/search_bloc.dart';

class SearchElement extends StatelessWidget {
  const SearchElement({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return InkWell(
      onTap: () => context.read<SearchBloc>().add(const SearchStarted()),
      child: const Icon(
        Icons.search,
        color: Colors.white,
        size: 35,
      ),
    );
  }
}
