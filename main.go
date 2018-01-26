package main

import (
	"fmt"
	"os"
	"os/exec"
	"log"
	"io/ioutil"
	"strconv"
	"strings"
	"path/filepath"
)

var isDebug bool = false;

var currentDir, inputDir, outputDir, ffExec string;

func main() {
	currentDir, _ = os.Getwd();
	inputDir = currentDir + "\\girdi\\";
	outputDir = currentDir + "\\cikti\\";
	ffExec = currentDir + "\\tool\\bin\\ffmpeg.exe";
	if isDebug {
		fmt.Println("Current : " + currentDir);
		fmt.Println("Input : " + inputDir);
		fmt.Println("Output : " + outputDir);
		fmt.Println("FFexec : " + ffExec)
		printFfmpegVersion();
	}

	files, err := ioutil.ReadDir(inputDir)
	if err != nil {
		log.Fatal(err)
	}


	for i, f := range files {
		if(isDebug){
			fmt.Println(strconv.Itoa(i)  + " " +  f.Name())
		}
		newName := strings.TrimSuffix(f.Name(), filepath.Ext(f.Name())) + ".mp3"
		fmt.Println(f.Name() + " -----> " + newName + " dosyasına çevriliyor. Çıktı klasörüne bak");
		convertMp3(inputDir + f.Name(),outputDir + newName)
	}

/*
c := exec.Command("cmd", "/C", "ls", "C://");
if cmdOut, err := c.Output(); err != nil{
	fmt.Println(os.Stderr,"Err : ",err)
}else{
	fmt.Println(string(cmdOut));
}



if err := c.Run(); err != nil {
	fmt.Println("Error: ", err)
}
*/
}

func printFfmpegVersion() {
	c := exec.Command("cmd", "/C", ffExec, "-version");
	if cmdOut, err := c.Output(); err != nil {
		fmt.Println(os.Stderr, "Err : ", err)
	} else {
		fmt.Println(string(cmdOut));
	}
}

func convertMp3(input string, output string) {
	c := exec.Command("cmd", "/C", ffExec, "-y", "-i", input, output)

	if cmdOut, err := c.CombinedOutput(); err != nil {
		if isDebug{
			log.Fatal(err.Error() + "Err : " + string(cmdOut));
		}
	} else {
		if isDebug {
			fmt.Println("Çıktı : " + string(cmdOut));
		}
	}
}
