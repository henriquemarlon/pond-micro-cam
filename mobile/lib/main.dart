import 'package:flutter/material.dart';
import 'package:app/app/pages/login_user.dart';
import 'package:app/app/pages/open_camera.dart';
import 'package:app/app/pages/create_user.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      initialRoute: "/",
      routes: {
        "/": (context) => LoginPage(),
        "/open_camera": (context) => Camera(),
        "/create_user": (context) => NewUserPage(),
      },
    );
  }
}

