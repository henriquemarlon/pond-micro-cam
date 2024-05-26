import 'package:flutter/material.dart';

class MyInput extends StatefulWidget {
  final String? label;
  final IconData? icon;
  final Function(String)? onChanged;
  final bool obscureText;

  const MyInput({
    Key? key,
    this.label,
    this.icon,
    this.onChanged,
    this.obscureText = false,
  }): super(key: key);


  @override
  _MyInputState createState() => _MyInputState();
}

class _MyInputState extends State<MyInput> {
  late TextEditingController _controller;

  @override
  void initState() {
    super.initState();
    _controller = TextEditingController(text: '');
  }

  @override
  void dispose() {
    _controller.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return TextField(
      obscureText: widget.obscureText, 
      controller: _controller,
      onChanged: widget.onChanged,
      decoration: InputDecoration(
        labelText: widget.label,
        prefixIcon: Icon(widget.icon),
      ),
    );
  }

}