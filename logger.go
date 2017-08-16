package vidar

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

var (
	Debug   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {
	// TODO
	path := "service.log"

	file, err := os.OpenFile(Config.Log.Path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln("Failed to open %s: %s", path, err)
	}

	debugOutput := ioutil.Discard
	infoOutput := os.Stdout
	warningOutput := os.Stdout
	errorOutput := io.MultiWriter(file, os.Stderr)

	Debug = log.New(debugOutput, "[TRACE]: ", log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(infoOutput, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(warningOutput, "[INFO]: ", log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(errorOutput, "[ERROR]: ", log.Ldate|log.Ltime|log.Lshortfile)
}
