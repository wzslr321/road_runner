import 'package:bloc/bloc.dart';
import 'package:freezed_annotation/freezed_annotation.dart';

part 'search_event.dart';

part 'search_state.dart';

part 'search_bloc.freezed.dart';

class SearchBloc extends Bloc<SearchEvent, SearchState> {
  SearchBloc() : super(const SearchState.initial()) {
    on<SearchStarted>((event, emit) {
      emit(const SearchState.started());
    });
    on<SearchSubmitted>((event, emit) {
      emit(const SearchState.submitted());
    });
  }
}
