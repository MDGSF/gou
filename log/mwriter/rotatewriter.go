package mwriter

import (
	"fmt"
	"os"
	"sync"

	"github.com/MDGSF/utils/log"
)

type RotateWriter struct {
	lock     sync.Mutex
	filename string // should be set to the actual filename
	fp       *os.File
	maxsize  int64
}

// Make a new RotateWriter. Return nil if error occurs during setup.
func New(filename string, maxsize int) *RotateWriter {
	w := &RotateWriter{
		filename: filename,
		maxsize:  int64(maxsize),
	}
	w.createLogFile()
	return w
}

func (w *RotateWriter) createLogFile() {
	var err error
	w.fp, err = os.OpenFile(w.filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND|os.O_SYNC, 0666)
	if err != nil {
		log.Error("create log file [%v] failed, err = %v", w.filename, err)
		return
	}
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.write(output)
}

func (w *RotateWriter) write(output []byte) (int, error) {
	if w.needRotate() {
		w.reduceFileSize()
	}

	return w.fp.Write(output)
}

func (w *RotateWriter) reduceFileSize() {
	var err error

	// Close existing file if open
	if w.fp != nil {
		err = w.fp.Close()
		w.fp = nil
		if err != nil {
			return
		}
	}

	// Rename dest file if it already exists
	_, err = os.Stat(w.filename)
	if err == nil {
		//err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		err = os.Rename(w.filename, w.filename+".bak")
		if err != nil {
			return
		}
	}

	// Create a file.
	w.createLogFile()
	return
}

func (w *RotateWriter) curFileSize() int64 {
	fileInfo, err := w.fp.Stat()
	if err != nil {
		fmt.Printf("err = %v", err)
		return 0
	}
	return fileInfo.Size()
}

func (w *RotateWriter) needRotate() bool {
	if w.curFileSize() > w.maxsize {
		return true
	}
	return false
}
