package main

import (
	"log"
	"fmt"
	"bytes"
	"os/exec"
)
 

// "github.com/json-iterator/go"
// var json = jsoniter.ConfigCompatibleWithStandardLibrary


const shellToUse = "bash"

func shellOut(command string)(string, string, error){
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	
	cmd := exec.Command(shellToUse, "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	
	return stdout.String(), stderr.String(), err
}

func lagoo () {
	out, errout, err := shellOut("ls -ltr")
	
	if err != nil {
		log.Printf("error: %v \n", err)
	}
	
	fmt.Println("---stdout---")
	fmt.Println(out)
	fmt.Println("---stderr---")
	fmt.Println(errout)
}

func customed() {
	type Tag string
	
	var tags []Tag
	
	tags = append(tags, "Jajo", "names")
	fmt.Println(tags)
}
