# Open Source Icons as Templ Components

-- still needs to be written .... At the moment only supports Heroicons (made by TailwindCss Creator )

# How to use ?

Import it as a module

```bash
go get "github.com/livghit/TempIcons"

```

Use the Icons

```go
package main
import "github.com/livghit/TempIcons/heroicons"

func main(){
  icon := heroicons.AcademicCap()
  icon.Render(context.Background(), os.Stdout)
}


// or just add it inside a templ file

package templates

import "github.com/livghit/TempIcons/heroicons"

templ Layout(title string){
  <html>
    <head>
      <title>{title}</title>
    </head>
    <body>
<button> {{ heroicons.AcademicCap() }} </button>
    </body>
  </html>
}

```
