import 'dart:typed_data';
import 'package:flutter/material.dart';
import 'package:image_picker/image_picker.dart';
import 'dart:io';
import 'package:awesome_notifications/awesome_notifications.dart';
import 'package:app/app/controllers/local_notificaiton.dart';
import 'package:http/http.dart' as http;
import 'package:path/path.dart' as path;
import 'dart:convert';

class Camera extends StatefulWidget {
  @override
  _CameraState createState() => _CameraState();
}

class _CameraState extends State<Camera> {
  @override
  void initState() {
    AwesomeNotifications().setListeners(
        onActionReceivedMethod: NotificationController.onActionReceivedMethod,
        onNotificationCreatedMethod:
            NotificationController.onNotificationCreatedMethod,
        onNotificationDisplayedMethod:
            NotificationController.onNotificationDisplayedMethod,
        onDismissActionReceivedMethod:
            NotificationController.onDismissActionReceivedMethod);
    super.initState();
  }

  final ImagePicker _picker = ImagePicker();
  XFile? _image;
  Uint8List? _receivedImage;

  Future<void> _openCamera() async {
    final XFile? image = await _picker.pickImage(source: ImageSource.camera);
    if (image != null) {
      setState(() {
        _image = image;
        _receivedImage = null;  
      });
    }
  }

  Future<void> _uploadImage(File imageFile) async {
    var uri = Uri.parse("http://10.0.2.2/image/upload");
    var request = http.MultipartRequest("POST", uri);

    var multipartFile = await http.MultipartFile.fromPath(
      'file',
      imageFile.path,
      filename: path.basename(imageFile.path),
    );

    request.files.add(multipartFile);

    var response = await request.send();

    if (response.statusCode == 200) {
      var responseData = await http.Response.fromStream(response);
      setState(() {
        _receivedImage = base64Decode(responseData.body);
      });
      AwesomeNotifications().createNotification(
        content: NotificationContent(
            id: 1,
            channelKey: "basic_channel",
            title: "Imagem editada",
            body: "Sua imagem foi editada."),
      );
    } else {
      var responseData = await response.stream.bytesToString();
      print("Erro: $responseData");
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Abrir Câmera'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            _receivedImage != null
                ? Image.memory(_receivedImage!)
                : (_image != null
                    ? Image.file(File(_image!.path))
                    : Text('Nenhuma imagem selecionada.')),
            SizedBox(height: 20),
            Container(
              width: 200,
              height: 50,
              child: ElevatedButton(
                onPressed: _openCamera,
                style: ButtonStyle(
                  backgroundColor: MaterialStateProperty.all<Color>(
                      Colors.blue.shade600),
                ),
                child: const Text(
                  'Abrir Câmera',
                  style: TextStyle(
                    color: Colors.black87,
                    fontSize: 16,
                  ),
                ),
              ),
            ),
            SizedBox(height: 20),
            _image == null
                ? Container()
                : Container(
                    width: 200,
                    height: 50,
                    child: ElevatedButton(
                      onPressed: () => _uploadImage(File(_image!.path)),
                      style: ButtonStyle(
                        backgroundColor: MaterialStateProperty.all<Color>(
                            Colors.green.shade600),
                      ),
                      child: const Text(
                        'Editar foto',
                        style: TextStyle(
                          color: Colors.black87,
                          fontSize: 16,
                        ),
                      ),
                    ),
                  ),
          ],
        ),
      ),
    );
  }
}
