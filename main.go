package main

import "fmt"

type Command func()

func main() {
	var commandMap map[string]Command
	commandMap = make(map[string]Command)

	commandMap["hello"] = hello
	commandMap["dev_info"] = dev_info

	for {
		var command string
		fmt.Print(">>>")
		fmt.Scan(&command)
		if command == "exit" {
			break
		}
		if command == "help" {
			var commandIndex int = 0
			for cmd := range commandMap {
				commandIndex++
				fmt.Println("CMD_", commandIndex, ":", cmd)
			}
			continue
		}
		commandFunc, ok := commandMap[command]
		if !ok {
			continue
		} else {
			commandFunc()
		}

	}
}

func hello() {
	fmt.Println("Hello world!\n")
}

func dev_info() {
	fmt.Println("Developed By Lighty The Light <3\n")
}
