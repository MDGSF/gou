package mwriter

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/MDGSF/utils"
	"github.com/MDGSF/utils/log"
)

type RotateWriter struct {
	lock            sync.Mutex
	filename        string // should be set to the actual filename
	fp              *os.File
	maxsize         int64
	maxFileDuration time.Duration
}

// Make a new RotateWriter. Return nil if error occurs during setup.
func New(filename string, maxsize int, maxFileDuration time.Duration) *RotateWriter {
	if len(filename) == 0 {
		panic("empty file name")
	}

	if maxsize <= 0 {
		panic(fmt.Sprintf("invalid maxsize (%v)", maxsize))
	}

	w := &RotateWriter{
		filename:        filename,
		maxsize:         int64(maxsize),
		maxFileDuration: maxFileDuration,
	}
	w.createLogFile()
	go w.autoClean()
	return w
}

func (w *RotateWriter) autoClean() {
	for {
		w.cleanExpiredFile()

		time.Sleep(time.Minute)
	}
}

func (w *RotateWriter) cleanExpiredFile() {

	logdir := filepath.Dir(w.filename)
	fileBaseName := filepath.Base(w.filename)

	files, err := ioutil.ReadDir(logdir)
	if err != nil {
		log.Error("read dir [%v] failed, err = %v", logdir, err)
		return
	}
	for _, file := range files {
		filename := file.Name()
		if strings.HasPrefix(filename, fileBaseName) {
			if filename == fileBaseName {
				continue
			}
			parts := strings.Split(filename, ".")
			if len(parts) == 0 {
				continue
			}
			strFileTimeStamp := parts[len(parts)-1]
			fileTimeStamp, err := time.Parse(time.RFC3339, strFileTimeStamp)
			if err != nil {
				log.Error("parse file [%v] timestamp failed, err = %v", filename, err)
				continue
			}

			dura := time.Now().Sub(fileTimeStamp)
			if dura <= 0 || dura > w.maxFileDuration {
				fileAbsolutePath := filepath.Join(logdir, filename)
				os.Remove(fileAbsolutePath)
				log.Info("remove expired log file: %v", fileAbsolutePath)
			}
		}
	}
}

func (w *RotateWriter) createLogFile() {
	var err error

	if utils.FileExists(w.filename) {
		err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		if err != nil {
			log.Error("rename file [%v] failed, err = %v", w.filename, err)
			return
		}
	}

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
		err = os.Rename(w.filename, w.filename+"."+time.Now().Format(time.RFC3339))
		//err = os.Rename(w.filename, w.filename+".bak")
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
