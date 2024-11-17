package filehandler

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleFile(src string, des string) {

	sourceDir := src
	destinationDir := des

	dir, err := os.ReadDir(sourceDir)
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
			fileSplit := strings.Split(fileName, ".")
			if len(fileSplit) == 1 {
				continue
			}
			folderName := destinationDir + "/" + fileSplit[len(fileSplit)-1]
			if _, err := os.Stat(folderName); os.IsNotExist(err) {
				err := os.MkdirAll(folderName, 0755)
				if err != nil {
					log.Fatalf("FAILDED TO CREATe DIR. %s", err)
				}
			}
			destination := fmt.Sprintf("%s/%s", folderName, fileName)
			err = os.Rename(sourceDir+"/"+fileName, destination)
			if err != nil {
				log.Fatalf("FAILDED TO MOVE. %s", err)
			}

		}
	}
}
