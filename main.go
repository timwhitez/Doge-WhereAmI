package main

import (
	"fmt"
	"golang.org/x/sys/windows"
	"unicode/utf16"
	"unsafe"
)

func main() {
	peb := (*windows.PEB)(unsafe.Pointer(getPEB()))

	envSize := peb.ProcessParameters.EnvironmentSize
	fmt.Printf("Environment Size: %d\n", envSize)
	fmt.Println("Environment: ")
	printStringAtAddress(uintptr(peb.ProcessParameters.Environment), int(envSize))

}

func printStringAtAddress(nextEnvStringAddr uintptr, length int) {
	environmentEndAddr := nextEnvStringAddr + uintptr(length)
	for nextEnvStringAddr < environmentEndAddr {
		// 将字节数组转换为UTF16编码的字符串
		str := (*[1 << 30]uint16)(unsafe.Pointer(nextEnvStringAddr))[:]
		length := len(str)
		for i := 0; i < length; i++ {
			if str[i] == 0 {
				length = i
				break
			}
		}
		utf16Str := utf16.Decode(str[:length])

		// 将UTF16字符串转换为UTF8字符串并打印
		fmt.Printf("%s\n", string(utf16Str))

		// 计算下一个环境变量字符串的地址
		nextEnvStringAddr += uintptr(length+1) * 2
	}
}

func getPEB() uintptr
