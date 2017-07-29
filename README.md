# stakmachine/extendplate
[![GoDoc](https://godoc.org/stackmachine.com/extendplate?status.svg)](https://godoc.org/stackmachine.com/extendplate) [![Build Status](https://travis-ci.org/stackmachine/extendplate.svg?branch=master)](https://travis-ci.org/stackmachine/extendplate)

## Install

```
dep ensure stackmachine.com/extendplate
```

## Usage

Out sample `templates` directory has three templates in a nested structure.

```
teamplates/
├── base
│   ├── billing.html
│   ├── dashboard.html
└── base.html
```

Let's take a look at the three templates.

```
{{/* base.html */}}
<html>
  <head>
    <title>Go Web Programming</title>
  </head>
  <body>
    {{ template "content" }}
  </body>
</html>
```

```
{{/* billing.html */}}
{{define "content"}}
  This is billing
{{end}}
```

```
{{/* dashboard.html */}}
{{define "content"}}
  This is the dashboard
{{end}}
```

Extendplate ensures the billing and dashboard template inherit from the base
template, inferring the hierarchy from the folder layout.

```go
package main

import (
    "os"

    "stackamachine.com/extendplate"
)

func main() {
    ts, _ := extendplate.ParseDir("templates", "*.html", nil)
    tmpl := ts.Lookup("base/dashboard.html")
    tmpl.Execute(os.Stdout, nil)
    // <html>
    //   <head>
    //     <title>Go Web Programming</title>
    //   </head>
    //   <body>
    //     This is the dashboard
    //   </body>
    // </html>
}
```
