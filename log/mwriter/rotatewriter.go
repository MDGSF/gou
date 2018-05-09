package mwriter

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type RotateWriter struct {
	lock     sync.Mutex
	filename string // should be set to the actual filename
	fp       *os.File

	count    int
	maxlsize int64
	maxrsize int64
}

// Make a new RotateWriter. Return nil if error occurs during setup.
func New(filename string, maxlsize int, maxrsize int) *RotateWriter {

	w := &RotateWriter{
		filename: filename,
		maxlsize: int64(maxlsize),
		maxrsize: int64(maxrsize),
	}

	var err error
	w.fp, err = os.Create(w.filename)
	if err != nil {
		return nil
	}

	return w
}

// Write satisfies the io.Writer interface.
func (w *RotateWriter) Write(output []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.write(output)
}

func (w *RotateWriter) write(output []byte) (int, error) {
	w.count++
	if w.count > 10000 {
		w.count = 0
		if w.fileIsTooBig() {
			w.reduceFileSize()
		}
	}

	return w.fp.Write(output)
}

func (w *RotateWriter) curFileSize() int64 {
	fileInfo, err := w.fp.Stat()
	if err != nil {
		fmt.Printf("err = %v", err)
		return 0
	}
	return fileInfo.Size()
}

func (w *RotateWriter) fileIsTooBig() bool {
	if w.curFileSize() > w.maxrsize {
		return true
	}
	return false
}

func (w *RotateWriter) reduceFileSize() {
	_, err := w.fp.Seek(w.curFileSize()-w.maxlsize, 0)
	if err != nil {
		fmt.Println("seek failed, err = ", err)
		return
	}

	b := make([]byte, w.maxlsize)
	_, err = w.fp.Read(b)
	if err != nil {
		fmt.Println("read failed, err = ", err)
		return
	}

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
	w.fp, err = os.Create(w.filename)
	if err != nil {
		fmt.Println("create file failed, err = ", err)
		return
	}

	_, err = w.fp.Write(b)
	if err != nil {
		fmt.Println("write to file failed, err = ", err)
		return
	}

	return

}

// Perform the actual act of rotating and reopening file.
func (w *RotateWriter) Rotate() (err error) {
	w.lock.Lock()
	defer w.lock.Unlock()
	return w.rotate()
}

func (w *RotateWriter) rotate() (err error) {
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
		if err != nil {
			return
		}
	}

	// Create a file.
	w.fp, err = os.Create(w.filename)
	return
}
