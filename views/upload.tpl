<html>
<head>
	<meta charset="utf-8">
    <title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="https://127.0.0.1:8888/upload" method="post">
  <input type="file" name="uploadfile" />
  <input type="hidden" name="token" value="{{.}}"/>
  <input type="submit" value="upload" />
</form>
</body>
</html>