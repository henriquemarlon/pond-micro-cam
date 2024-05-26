import 'package:http/http.dart' as http;
import 'package:flutter/material.dart';
import 'dart:convert';

Future<void> login(BuildContext context, String email, String password) async {
  const url = "http://10.0.2.2/api/v1/login";

  final response = await http.post(
    Uri.parse(url),
    headers: <String, String>{
      'Content-Type': 'application/json',
    },
    body: jsonEncode({
      'email': email,
      'password': password,
    }),
    );

  if (response.statusCode == 200) {
    Navigator.pushNamed(context, "/open_camera");
  } else {
    throw Exception("Erro ao logar.");
  }
}

Future<void> create(String email, String name, String password) async {
  const url = "http://10.0.2.2/api/v1/user";

  final response = await http.post(
    Uri.parse(url),
    headers: <String, String>{
      'Content-Type': 'application/json',
    },
    body: jsonEncode({
      'email': email,
      'name': name,
      'password': password,
    }),
    );

  if (response.statusCode == 200) {
    print("ok");
  } else {
    throw Exception("Erro ao criar usuário.");
  }
}




