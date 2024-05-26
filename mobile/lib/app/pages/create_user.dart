import 'package:flutter/material.dart';
import 'package:app/app/widgets/input.dart';
import 'package:flutter/widgets.dart';
import 'package:app/app/integration/login.dart';

class NewUserPage extends StatefulWidget {
  @override
  _NewUserPageState createState() => _NewUserPageState();
}

class _NewUserPageState extends State<NewUserPage> {
  String _name = '';
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

  void _onChangedName(String value) {
    setState(() {
      _name = value;
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: const Text('Criar nova conta'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            Padding(
              padding: const EdgeInsetsDirectional.all(20.0),
              child: MyInput(
                  onChanged: _onChangedName,
                  label: 'Name',
                  icon: Icons.person_2_outlined),
            ),
            Padding(
              padding: const EdgeInsetsDirectional.all(20.0),
              child: MyInput(
                  onChanged: _onChangedEmail,
                  label: 'E-mail',
                  icon: Icons.email_outlined),
            ),
            Padding(
              padding: const EdgeInsetsDirectional.all(20.0),
              child: MyInput(
                  onChanged: _onChangedPassword,
                  label: 'Senha',
                  icon: Icons.lock_outline_rounded,
                  obscureText: true),
            ),
            Container(
              width: 200, 
              height: 50, 
              child: ElevatedButton(
                onPressed: () {
                  create(_email, _name, _password);
                },
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all<Color>(
                      Colors.blue.shade600), 
                ),
                child: const Text(
                  'Create Account',
                  style: TextStyle(
                    color: Colors.black87,
                    fontSize: 16,
                  )
                )
              ),
            ),
          ],
        ),
      ),
    );
  }
}
