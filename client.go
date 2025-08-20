package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Usage: client.exe <peerIP:port> <filePath>	")
		return
	}

	addr := os.Args[1]
	filePath := os.Args[2]

	sendFile(addr, filePath)
}

func sendFile(addr, filePath string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("Connection error:", err)
		return
	}
	defer conn.Close()

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("File open error:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(conn, "%s\n", getFileName(filePath))
	_, err = io.Copy(conn, file)
	if err != nil {
		fmt.Println("Error sending file:", err)
	} else {
		fmt.Println("File sent successfully!")
	}
}

func getFileName(path string) string {
	parts := strings.Split(path, string(os.PathSeparator))
	return parts[len(parts)-1]
}
