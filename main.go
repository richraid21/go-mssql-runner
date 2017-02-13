package main

import (
	"fmt"
	"strconv"
	"io/ioutil"
	"log"
	"strings"
	"os/exec"
	"os"
)

type ScriptExecution struct {
	instance string
	database string
	path string
	filename string
}

func (se *ScriptExecution) Cmd() string{
	return `sqlcmd -S `+se.instance+` -d `+se.database+` -i `+se.path
}

func main() {

	args := os.Args[1:]
	connectionString := os.Getenv("GOSQLCONNECTION")
	pair := strings.Split(connectionString,":")
	
	instanceString := pair[0]
	databaseName := pair[1]

	var folderPath string
	if len(args) == 0 {
		log.Fatal("Folder Path not provided")
	}
	folderPath = args[0]

	
	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		log.Fatal(err)
	}
	
	for _, file := range files {
	
		fileName := file.Name()
		
		if strings.HasSuffix(fileName,".sql") {
			se := ScriptExecution{instance: instanceString, database: databaseName, path: folderPath+`\`+fileName, filename: fileName}
			run(&se)
		}
		
	}

	numFiles := strconv.Itoa(len(files))
	fmt.Println(`Executed `+numFiles+` of `+numFiles+` files.`)
}

func run(eo *ScriptExecution){
	
	execCmd := eo.Cmd()	
	
	lsCmd := exec.Command("cmd", "/C", execCmd)
    lsOut, err := lsCmd.Output()
    if err != nil || strings.HasPrefix(string(lsOut),"Msg") {
        fmt.Println(`SQL Script: `+eo.filename+` encountered an error during execution.`)
		fmt.Println(string(lsOut))
		fmt.Println(err)
		log.Fatal(err)
    }
}

