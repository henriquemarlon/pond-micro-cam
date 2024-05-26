import 'package:flutter/material.dart';
import 'package:app/app/widgets/input.dart';
import 'package:flutter/widgets.dart';
import 'package:app/app/integration/login.dart';

class LoginPage extends StatefulWidget {
  @override
  _LoginPageState createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  String _email = '';
  String _password = '';

  void _onChangedEmail(String value) {
    setState(() {
      _email = value;
    });
  }

  void _onChangedPassword(String value) {
    setState(() {
      _password = value;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            const Center(
              child: Text(
                'Login',
                style: TextStyle(
                  fontSize: 30,
                  color: Colors.blue,
                ),
              ),
            ),
            Padding(
              padding: const EdgeInsetsDirectional.all(20.0),
              child: MyInput(
                  onChanged: _onChangedEmail,
                  label: 'Email',
                  icon: Icons.email),
            ),
            Padding(
              padding: const EdgeInsetsDirectional.all(20.0),
              child: MyInput(
                  onChanged: _onChangedPassword,
                  label: 'Password',
                  icon: Icons.lock),
            ),
            Container(
              height: 40, 
              width: 400, 
              child: ElevatedButton(
                onPressed: () {
                  login(context, _email, _password);
                },
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all<Color>(
                      Colors.blue.shade600), 
                  shape: MaterialStateProperty.all<RoundedRectangleBorder>(
                    RoundedRectangleBorder(
                      borderRadius:
                          BorderRadius.circular(20), 
                    ),
                  ),
                ),
                child: const Text(
                  'Enter',
                  style: TextStyle(
                    fontSize: 15,
                    color: Colors.black87,
                  )
                )
              ),
            ),
            const SizedBox(height: 20), 
            GestureDetector(
              onTap: () {
                Navigator.pushNamed(context, "/create_user");
              },
              child: const Text(
                'Create an account',
                style: TextStyle(
                  decoration: TextDecoration.underline, 
                  color: Colors.blue, 
                  fontSize: 20,
                ),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
