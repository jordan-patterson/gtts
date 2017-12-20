package gtts

import (
	"fmt"
	"net/http"
	"os"
	"net/url"
	"io/ioutil"
)
type Gtts struct{
	text string
	lang string
}
func (tts *Gtts) Get() ([]byte,error){
	//prepare text so it can be placed in query
	tts.text=url.QueryEscape(tts.text)
	url:=fmt.Sprintf("http://translate.google.com/translate_tts?ie=UTF-8&q=%s&tl=%s&client=gtx",tts.text,tts.lang)
	
	//do the request
	resp,err:=http.Get(url)
	if err!=nil{
		return nil,err
	}
	defer resp.Body.Close()
	content,err:=ioutil.ReadAll(resp.Body)
	if err!=nil{
		return nil,err
	}
	return []byte(content),nil
	
}
func (tts *Gtts) Save(filename string) (error){
	//create new file
	file,err:=os.Create(filename)
	if err!=nil{
		return err
	}
	defer file.Close()
	bytes,err:=tts.Get()
	if err!=nil{
		return err
	}
	_,err=file.Write(bytes)
	if err!=nil{
		return err
	}
return nil
}
func main(){
	gtts:=Gtts{"This is a test","en"}
	gtts.Save("test.mp3")
}