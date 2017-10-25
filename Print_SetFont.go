package main

import(
	iconv "github.com/djimenez/iconv-go"
)

//set title font to utf-8
func Print_SetFont(str string)(str_cov string){
	str_cov,err := iconv.ConvertString(str,"BIG5","UTF-8")
	if err != nil{
		str_cov,err = iconv.ConvertString(str,"gb18030","UTF-8")
		if err != nil{
			str_cov,err = iconv.ConvertString(str,"GBK","UTF-8")
			// if err != nil{
			// 	fmt.Println("still have err")
			// }
		}
	}
	return
}
