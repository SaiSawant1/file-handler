package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	dir, err := os.ReadDir("./source")
	if err != nil {
		log.Fatalf("FAILED TO READ DIR, %s", err)
	}
	if err != nil {
		log.Fatalf("FAILED TO READ DIR, %s", err)
	}

	for _, f := range dir {
		if !f.IsDir() {
			fileName := f.Name()
			fmt.Println(fileName)
			folderName := strings.Split(fileName, ".")
			if len(folderName) == 1 {
				continue
			}
			if _, err := os.Stat("./source/" + folderName[1]); os.IsNotExist(err) {

				err := os.Mkdir(folderName[1], 0755)
				if err != nil {
					log.Fatalf("FAILDED TO CREAT DIR. %s", err)
				}
			}
			destination := fmt.Sprintf("%s/%s", "./source/"+folderName[1], fileName)
			err = os.Rename(fileName, destination)
			if err != nil {
				log.Fatalf("FAILDED TO MOVE. %s", err)
			}

		}
	}

}
