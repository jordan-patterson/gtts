# gtts
Gtts client library for converting text to speech by using the Google Translate API.

### Installation
``` bash
$ go get github.com/jordan-patterson/gtts
```

### Usage
``` go
 package main
 
 import (
    "github.com/jordan-patterson/gtts"
 )
 func main() {
    converter := gtts.Gtts{Text:"To infinity and beyond",Lang:"en"}
    
    //now save to mp3 file
    converter.Save("speech.mp3")
 }
```
