package main

import (
	"time"

	"github.com/jessevdk/go-assets"
)

var _Assets3b4e08d4c718d774efe29935a4a6842efab4f09b = "<!doctype html>\n<html lang=\"ja\">\n\n<head>\n  <title>Yet Another Hatena Bookmark Viewer</title>\n  <!-- Required meta tags -->\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n  <!-- Bootstrap CSS -->\n  <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css\" integrity=\"sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb\"\n    crossorigin=\"anonymous\">\n</head>\n\n<body>\n  <div class=\"list-group\">\n    <a href=\"sample\" target=\"_blank\" class=\"list-group-item list-group-item-action flex-column align-items-start\">\n      <h5 class=\"mb-1\">SAMPLE</h5>\n    </a>\n  </div>\n\n  <!-- Optional JavaScript -->\n  <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n  <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\"\n    crossorigin=\"anonymous\"></script>\n  <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js\" integrity=\"sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh\"\n    crossorigin=\"anonymous\"></script>\n  <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js\" integrity=\"sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ\"\n    crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n"
var _Assets952699f28389272e948c53cd61e012cd1078ec98 = "<!doctype html>\n<html lang=\"ja\">\n\n<head>\n  <title>Yet Another Hatena Bookmark Viewer</title>\n  <!-- Required meta tags -->\n  <meta charset=\"utf-8\">\n  <meta name=\"viewport\" content=\"width=device-width, initial-scale=1, shrink-to-fit=no\">\n\n  <!-- Bootstrap CSS -->\n  <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/css/bootstrap.min.css\" integrity=\"sha384-PsH8R72JQ3SOdhVi3uxftmaW6Vc51MKb0q5P2rRUpPvrszuE4W1povHYgTpBfshb\"\n    crossorigin=\"anonymous\">\n</head>\n\n<body>\n  <div class=\"list-group\">\n    {{range .}}\n    <a href=\"{{.Link}}\" target=\"_blank\" class=\"list-group-item list-group-item-action flex-column align-items-start\">\n      <div class=\"media\">\n        <img src=\"http://n.hatena.com/{{.Creator}}/profile/image.gif?type=face&size=64\" class=\"mr-3\">\n        <div class=\"media-body\">\n          <div class=\"d-flex w-100 justify-content-between\">\n            <h5 class=\"mb-1\">{{.Creator}}</h5>\n            <small>{{.Date}}</small>\n          </div>\n          <p class=\"mb-1\">{{.Description}}</p>\n          <small>{{.Title}}</small>\n        </div>\n      </div>\n    </a>\n    {{end}}\n  </div>\n\n  <!-- Optional JavaScript -->\n  <!-- jQuery first, then Popper.js, then Bootstrap JS -->\n  <script src=\"https://code.jquery.com/jquery-3.2.1.slim.min.js\" integrity=\"sha384-KJ3o2DKtIkvYIK3UENzmM7KCkRr/rE9/Qpg6aAZGJwFDMVNA/GpGFF93hXpG5KkN\"\n    crossorigin=\"anonymous\"></script>\n  <script src=\"https://cdnjs.cloudflare.com/ajax/libs/popper.js/1.12.3/umd/popper.min.js\" integrity=\"sha384-vFJXuSJphROIrBnz7yo7oB41mKfc8JzQZiCq4NCceLEaO4IHwicKwpJf9c9IpFgh\"\n    crossorigin=\"anonymous\"></script>\n  <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/4.0.0-beta.2/js/bootstrap.min.js\" integrity=\"sha384-alpBpkh1PFOepccYVYDB4do5UnbKysX5WZXm3XxPqe5iKTfUKjNkCk9SaVuEZflJ\"\n    crossorigin=\"anonymous\"></script>\n</body>\n\n</html>\n"

// Assets returns go-assets FileSystem
var Assets = assets.NewFileSystem(map[string][]string{"/": []string{"index.html", "template.html"}}, map[string]*assets.File{
	"/template.html": &assets.File{
		Path:     "/template.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1514813459, 1514813459730134713),
		Data:     []byte(_Assets952699f28389272e948c53cd61e012cd1078ec98),
	}, "/": &assets.File{
		Path:     "/",
		FileMode: 0x800001ed,
		Mtime:    time.Unix(1514814998, 1514814998308951705),
		Data:     nil,
	}, "/index.html": &assets.File{
		Path:     "/index.html",
		FileMode: 0x1a4,
		Mtime:    time.Unix(1514814998, 1514814998236763747),
		Data:     []byte(_Assets3b4e08d4c718d774efe29935a4a6842efab4f09b),
	}}, "")
