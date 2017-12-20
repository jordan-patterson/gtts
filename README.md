# gtts
Gtts client library for converting text to speech. 

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
    converter=gtts.Gtts{text:"To infinity and beyond",lang:"en"}
    //now save to mp3 file
    converter.Save("speech.mp3")
 }
```
