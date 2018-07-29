package handler

import (
	"net/http"
)

type IndexHandler struct {
}

func (c *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(`<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>表情包制作神器</title>
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="renderer" content="webkit">
    <meta http-equiv="Cache-Control" content="no-siteapp"/>
    <link rel="icon" type="image/png" href="https://lattecake.com/static/favicon.ico">
    <meta name="mobile-web-app-capable" content="yes">
    <link rel="icon" sizes="192x192" href="https://lattecake.com/static/favicon.ico">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <meta name="apple-mobile-web-app-title" content="Amaze UI"/>
    <link rel="apple-touch-icon-precomposed" href="https://lattecake.com/static/favicon.ico">
    <meta name="msapplication-TileImage" content="https://lattecake.com/static/favicon.ico">
    <meta name="msapplication-TileColor" content="#0e90d2">
	<link href="https://cdn.bootcss.com/bootstrap/4.1.1/css/bootstrap.min.css" rel="stylesheet">
	<link href="https://cdn.bootcss.com/bootstrap-fileinput/4.4.8/css/fileinput.min.css" rel="stylesheet">
	<style>
		#preview {
      		width: 300px;
		  	height: 300px;
		  	overflow: hidden;
		}
		#preview img {
		  	width: 100%;
		  	height: 100%;
		}
	</style>
</head>
<body>
<form enctype="multipart/form-data" action="/upload/image" method="post">
  <input type="file" name="imageUpload"/>
  <div id="preview" style="width: 300px;height:300px;border:1px solid gray;"></div>
<button type="submit">Submit</button>
</form>

<script src="https://cdn.bootcss.com/jquery/3.3.1/jquery.min.js"></script>
<script src="https://cdn.bootcss.com/jquery.qrcode/1.0/jquery.qrcode.min.js"></script>
<script src="https://cdn.bootcss.com/bootstrap-fileinput/4.4.8/js/fileinput.min.js"></script>
<script src="/static/upload.js"></script>
<script src="/static/emoticon.js"></script>
<script type="text/javascript">
    function preview1(file) {
      var img = new Image(), url = img.src = URL.createObjectURL(file)
      var $img = $(img)
      img.onload = function() {
        URL.revokeObjectURL(url)
        $('#preview').empty().append($img)
      }
    }
    function preview2(file) {
      var reader = new FileReader()
      reader.onload = function(e) {
        var $img = $('<img>').attr("src", e.target.result)
        $('#preview').empty().append($img)
      }
      reader.readAsDataURL(file)
    }

    $(function() {
      $('[type=file]').change(function(e) {
        var file = e.target.files[0]
        preview1(file)
      })
    })
  </script>
</body>
</html>`))
}
