package test
//
import (
	"TinyTikTok/dao"
	"fmt"
	"log"
	"testing"
)

func TestGetUploadFileURL(t *testing.T)  {
	url, err := dao.GetUploadFileURL("cb53aaff-ed29-4e4e-8a61-8199201cee4c.mp4")

	if err != nil {
	t.Errorf("%v",err)
		//log.Println("链接获取失败")
	}else {
		log.Println("url=",url)
	}
}
func TestGetUploadURL(t *testing.T)  {
	url, err := dao.GetUploadURL("cb53aaff-ed29-4e4e-8a61-8199201cee4c.mp4")
	if err != nil {
		t.Errorf("%v",err)
		//log.Println("链接获取失败")
	}else {
		log.Println("url=",url)
	}
}
//func TestShowVideoList(t *testing.T)  {
//
//	more, err := shared.QueryMore()
//	if err!=nil{
//		t.Errorf("%v",err)
//	}else{
//		fmt.Println(more)
//	}
//
//}

func TestGetImage(t *testing.T){
	err ,s:= dao.GetImage("a6086b70-005e-41e3-bf37-57238b97bb6b")
	if err!=nil{
		t.Errorf("%v",err)
	}else {
		fmt.Println("成功")
	}
	fmt.Println(s)
}