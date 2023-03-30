import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../application/search/search_bloc.dart';
import '../../../../utils/extensions.dart';

class ExtendedSearchElement extends StatelessWidget {
  const ExtendedSearchElement({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return TextFormField(
      decoration: InputDecoration(
        hintText: context.l10n.searchInputHolder,
        prefixIcon: const Icon(Icons.search),
      ),
      onFieldSubmitted: (value) => context.read<SearchBloc>().add(
            SearchSubmitted(value),
          ),
    );
  }
}
