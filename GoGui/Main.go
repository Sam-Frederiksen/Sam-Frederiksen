package main

// #include <windows.h>
//  int WINAPI WinMain(HINSTANCE hInst ,HINSTANCE hPrevInst, LPSTR args,int ncmdshow)
//  {
//  MessageBox(NULL,"Hello","My First GUI",MB_OK);
//  return 0;
//  }
//import "C"
import (
	"time"
	"fmt"
)

func main() {
	time.Sleep(10000000)
	fmt.Println("Hello")
}
