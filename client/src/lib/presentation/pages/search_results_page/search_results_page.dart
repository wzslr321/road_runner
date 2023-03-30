import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../application/search/search_bloc.dart';
import 'search_results_page_view.dart';

class SearchResultsPage extends StatelessWidget {
  const SearchResultsPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return BlocProvider(
      create: (context) => SearchBloc(),
      child: const SearchResultsPageView(),
    );
  }
}
