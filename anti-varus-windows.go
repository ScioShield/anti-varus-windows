package main

import (
	"bufio"
	"bytes"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func vdbLoader() []string {
	// Loads the varus database and returns them as usable signatures
	appdataDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}
	vdbFile, err := os.Open(appdataDir + "\\anti-varus\\varus.vdb")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(vdbFile)
	scanner.Split(bufio.ScanLines)

	var sigs []string
	for scanner.Scan() {
		sigs = append(sigs, scanner.Text())
	}
	defer vdbFile.Close()
	return sigs
}

func searchFileLoader() (*os.File, []byte) {
	// Get the file to scan from the user
	fmt.Println("Enter what file you would like to scan (full path and filename):")
	var userSearchFile string
	fmt.Scanln(&userSearchFile)
	searchFile, err := os.Open(userSearchFile)
	if err != nil {
		log.Fatal(err)
	}

	searchData, err := ioutil.ReadAll(searchFile)
	if err != nil {
		log.Fatal(err)
	}
	defer searchFile.Close()
	return searchFile, searchData
}

func AVscan(sigs []string, searchData []byte, searchFile *os.File) {
	// main AV scan logic, compare the sigs to the content of the file to scan and return info to user
	for _, v := range sigs {
		decoded, err := hex.DecodeString(v)
		if err != nil {
			log.Fatal(err)
		}
		if bytes.Contains([]byte(searchData), []byte(decoded)) {
			fmt.Println("Match found!")
			fmt.Printf("Target hex: %s\n", v)
			fmt.Printf("Target string: %s\n", decoded)
			fmt.Printf("Target file: %s\n", searchFile.Name())
			// out, err := exec.Command("sha256sum", searchFile.Name()).Output()
			// if err != nil {
			//      log.Fatal(err)
			// }
			// fmt.Printf("File SHA256 Hash is: %s", out)
			out, err := exec.Command("cmd", "/C", "certutil -hashfile", searchFile.Name(), "SHA256").Output()
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("File SHA256 Hash is: %s", out)
		}
	}
}

func main() {
	sigs := vdbLoader()

	searchFile, searchData := searchFileLoader()

	AVscan(sigs, searchData, searchFile)
}
