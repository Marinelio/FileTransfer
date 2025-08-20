package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

func main() {

	folder := "recieved"
	if _, err := os.Stat(folder); os.IsNotExist(err) {

		err := os.Mkdir(folder, 0755)
		if err != nil {
			fmt.Println("Error creating folder:", err)
			return
		}
		fmt.Println("Folder created:", folder)
	} else {
		fmt.Println("Folder already exists:", folder)
	}

	err := os.Chdir(folder)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}

	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current directory:", err)
		return
	}
	fmt.Println("Current directory is now:", cwd)

	listener, err := net.Listen("tcp", ":25557")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	fmt.Println("listening on port 25557...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Connection error:", err)
			continue
		}
		go handleIncoming(conn)
	}
}

func handleIncoming(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	file, err := os.Create("received_" + fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, reader)
	if err != nil {
		fmt.Println("Error receiving file:", err)
	} else {
		fmt.Println("Received file:", fileName)
	}
}
