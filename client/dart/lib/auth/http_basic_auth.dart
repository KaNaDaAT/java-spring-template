//
// AUTO-GENERATED FILE, DO NOT MODIFY!
//
// @dart=2.0

// ignore_for_file: unused_element, unused_import
// ignore_for_file: always_put_required_named_parameters_first
// ignore_for_file: constant_identifier_names
// ignore_for_file: lines_longer_than_80_chars

part of projectname-client;

class HttpBasicAuth implements Authentication {
  String username;
  String password;

  @override
  void applyToParams(List<QueryParam> queryParams, Map<String, String> headerParams) {
    final credentials = '${username ?? ''}:${password ?? ''}';
    headerParams['Authorization'] = 'Basic ${base64.encode(utf8.encode(credentials))}';
  }
}
