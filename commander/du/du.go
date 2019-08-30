package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress message")
var sema = make(chan struct{}, 20)
var done = make(chan struct{})

type TFileInfo struct {
	Dir      string
	FileInfo os.FileInfo
}

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	fileInfoChan := make(chan TFileInfo)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir(root, &n, fileInfoChan)
	}
	go func() {
		n.Wait()
		close(fileInfoChan)
	}()

	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileInfoChan {
			}
			return
		case fileInfo, ok := <-fileInfoChan:
			if !ok {
				break loop
			}

			size := fileInfo.FileInfo.Size()
			//if size > nbytes*10 && nbytes > 1000 {
			//	fmt.Printf("dir = %v, name = %v, size = %v", fileInfo.Dir, fileInfo.FileInfo.Name(), fileInfo.FileInfo.Size())
			//	time.Sleep(5 * time.Second)
			//}

			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files, %v bytes, %.1f GB\n", nfiles, nbytes, float64(nbytes)/1e9)
}

func walkDir(dir string, n *sync.WaitGroup, fileSizes chan<- TFileInfo) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, n, fileSizes)
		} else {
			fileSizes <- TFileInfo{Dir: dir, FileInfo: entry}
		}
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}
