package filehandler

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func HandleFile(src string, des string, msg chan<- string) {

	msg <- "arranging files please wait..."

	sourceDir := src
	destinationDir := des

	dir, err := os.ReadDir(sourceDir)
	if err != nil {
		log.Fatalf("FAILED TO READ DIR, %s", err)
		msg <- err.Error()
	}
	if err != nil {
		log.Fatalf("FAILED TO READ DIR, %s", err)
		msg <- err.Error()
	}

	for _, f := range dir {
		if !f.IsDir() {
			fileName := f.Name()
			msg <- fileName
			fileSplit := strings.Split(fileName, ".")
			if len(fileSplit) == 1 {
				continue
			}
			folderName := destinationDir + "/" + fileSplit[len(fileSplit)-1]
			if _, err := os.Stat(folderName); os.IsNotExist(err) {
				err := os.MkdirAll(folderName, 0755)
				if err != nil {
					log.Fatalf("FAILDED TO CREATe DIR. %s", err)
					msg <- err.Error()
				}
			}
			destination := fmt.Sprintf("%s/%s", folderName, fileName)
			err = os.Rename(sourceDir+"/"+fileName, destination)
			if err != nil {
				log.Fatalf("FAILDED TO MOVE. %s", err)
				msg <- err.Error()
			}

		}
	}

	msg <- "Task completed!"
	close(msg)
}
