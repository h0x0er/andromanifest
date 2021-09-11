# Andromanifest
AndroidManifest.xml parser written in go

# Usage

```go

package main

import "io/ioutil"
import "fmt"
import "github.com/h0x0er/andromanifest"


func main(){

  mf, _ := androidmainfest.NewFromFile("/path/to/AndroidManifest.xml")
  fmt.Println(mf.Package) // prints the package name

}

```
