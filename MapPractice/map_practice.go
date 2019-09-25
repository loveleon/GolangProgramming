package main

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq(fileName string){
	//open file
	f,_ := os.Open(fileName)

	//file info
	fI,_ := f.Stat()
	fmt.Println("file :",fI.Name())
	fmt.Println("size :",fI.Size())

	word := map[string]int{}
	scan := bufio.NewScanner(f)
	scan.Split(bufio.ScanWords)
	for scan.Scan(){
		word[string(scan.Bytes())]++
	}

	for str,count := range word{
		if count > 10 {
			fmt.Println(str,":",count)
		}

	}
}

func main(){
	str1 := "/root/install.log"
	wordfreq(str1)
}
