package logging

import (
	"log"
	"os"
)

var (
	Info  *log.Logger = log.New(os.Stdout, "INFO: ", log.LstdFlags|log.Lshortfile)
	Error *log.Logger = log.New(os.Stdout, "ERROR: ", log.LstdFlags|log.Lshortfile)
)
