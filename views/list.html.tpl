<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <meta http-equiv="X-UA-Compatible" content="ie=edge">
  <title>Document</title>
</head>
<body>
  <div style="width: 600px;">
    <img src="{{.User.Profile.ImageOriginal}}" alt="user image" style="display: block; width: 100%;">
  </div>
  <p>Name: {{.User.Profile.DisplayName}}</p>
</body>
</html>