package main

import (
	"html/template"
	"log"
	"path"
)

func loadTemplates(dir string) *template.Template {
	if dir == "" {
		return getTemplates()
	}
	log.Printf("using custom template directory %q", dir)
	t, err := template.New("").ParseFiles(path.Join(dir, "sign_in.html"), path.Join(dir, "error.html"))
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`{{define "sign_in.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
  <title>CoreOS Edge Security</title>
  <link href="//fonts.googleapis.com/css?family=Source+Sans+Pro:400,700" rel="stylesheet" type="text/css">
  <link rel="shortcut icon" href="//storage.googleapis.com/assets.coreos.systems/uberproxy/favicon.png">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <style>
  body {
    font-family: "Source Sans Pro", sans-serif;
    font-size: 18px;
    font-weight: normal;
    color: rgba(0, 0, 0, 0.75);
    background: url(//storage.googleapis.com/assets.coreos.systems/uberproxy/hexagons.png),linear-gradient(30deg,#2276ad,#6ec654);
  }
  .signin {
    background-image: url(//storage.googleapis.com/assets.coreos.systems/uberproxy/border.png);
    background-position: 0px 0px;
    background-repeat: repeat-x;
    background-color: #fff;
    position: absolute;
    left: 50%;
    top: 25%;
    transform: translate(-50%, -50%);
    box-shadow: 2px 2px 2px rgba(0, 0, 0, 0.75);
    width: 450px;
    border-radius: 8px;
    padding: 25px 25px 10px 25px;
    text-align: center;
  }
  .btn {
    width: 165px;
    height: 36px;
    border: 0px;
    cursor: pointer;
    outline: 0;
  }
  p { margin: 0px; }
  p.bold { font-weight: 700; }
  p.message { margin-bottom: 20px; }
  p.logo img { width: 40%; height: 40%; }
  p.footer { color: #888; font-size: 12px; margin-top: 10px; }
  p.footer a { color: inherit; }
  </style>
</head>
<body>
	<div class="signin center">
	<form method="GET" action="{{.ProxyPrefix}}/start">
	<input type="hidden" name="rd" value="{{.Redirect}}">
  <p class="logo"><img src="//storage.googleapis.com/assets.coreos.systems/uberproxy/coreos-wordmark-horiz-color.png"></p>
      <p class="bold">&#x261d; Ah ah ah! You didn't say the magic word!</p>
      <p class="message">Access to this system is restricted to CoreOS employees.</p>
      <input type="image" src="//storage.googleapis.com/assets.coreos.systems/uberproxy/Blue_signin_Long_normal_20dp_v3.png" class="btn">
      <p class="footer">Please contact infra-team with any issues or questions. <a href="https://coreos.com">Return to safety.</a></p>
    </form>
  </div>
</body>
</html>
{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}

	t, err = t.Parse(`{{define "error.html"}}
<!DOCTYPE html>
<html lang="en" charset="utf-8">
<head>
  <title>CoreOS Edge Security</title>
  <link href="//fonts.googleapis.com/css?family=Source+Sans+Pro:400,700" rel="stylesheet" type="text/css">
  <link rel="shortcut icon" href="//storage.googleapis.com/assets.coreos.systems/uberproxy/favicon.png">
  <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no">
  <style>
  body {
    font-family: "Source Sans Pro", sans-serif;
    font-size: 18px;
    font-weight: normal;
    color: rgba(0, 0, 0, 0.75);
    background: url(//storage.googleapis.com/assets.coreos.systems/uberproxy/hexagons.png),linear-gradient(30deg,#2276ad,#6ec654);
  }
  .signin {
    background-image: url(//storage.googleapis.com/assets.coreos.systems/uberproxy/border.png);
    background-position: 0px 0px;
    background-repeat: repeat-x;
    background-color: #fff;
    position: absolute;
    left: 50%;
    top: 25%;
    transform: translate(-50%, -50%);
    box-shadow: 2px 2px 2px rgba(0, 0, 0, 0.75);
    width: 450px;
    border-radius: 8px;
    padding: 25px 25px 10px 25px;
    text-align: center;
  }
  .btn {
    width: 165px;
    height: 36px;
    border: 0px;
    cursor: pointer;
    outline: 0;
  }
  p { margin: 0px; }
  p.bold { font-weight: 700; }
  p.message { margin-bottom: 20px; }
  p.logo img { width: 40%; height: 40%; }
  p.footer { color: #888; font-size: 12px; margin-top: 10px; }
  p.footer a { color: inherit; }
  </style>
</head>
<body>
  <div class="signin center">
  <p class="logo"><img src="//storage.googleapis.com/assets.coreos.systems/uberproxy/coreos-wordmark-horiz-color.png"></p>
  <p class="bold">{{.Title}}</p>
  <p class="message">{{.Message}}</p>
  <p class="footer">Please contact infra-team with any issues or questions. <a href="{{.ProxyPrefix}}/sign_in">Return to sign in.</a></p>
  </div>
</body>
</html>{{end}}`)
	if err != nil {
		log.Fatalf("failed parsing template %s", err)
	}
	return t
}
