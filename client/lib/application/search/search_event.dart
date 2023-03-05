part of 'search_bloc.dart';

@freezed
class SearchEvent with _$SearchEvent {
  const factory SearchEvent.searchStarted() = SearchStarted;
  const factory SearchEvent.searchSubmitted(String query) = SearchSubmitted;
}
