// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package ucmd

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/MDGSF/utils"
)

func TestExecCopy1(t *testing.T) {
	os.MkdirAll("/tmp/anno_test_copy1", 0777)
	CmdCopy("/tmp/anno_test_copy2", "/tmp/anno_test_copy1")

	ret := utils.IsDir("/tmp/anno_test_copy1")
	if !ret {
		t.Fatal(ret)
	}

	ret = utils.IsDir("/tmp/anno_test_copy2")
	if !ret {
		t.Fatal(ret)
	}

	os.RemoveAll("/tmp/anno_test_copy1")
	os.RemoveAll("/tmp/anno_test_copy2")
}

func TestExecCopy2(t *testing.T) {
	err := CmdCopy("", "/tmp/anno_test_copy1")
	if err == nil {
		t.Fatal(err)
	}
}

func TestZipFile1(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".tar.gz"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	err := PackTarGz(tmpPath, tmpPackageName)
	if err != nil {
		t.Fatal(err)
	}

	ret := utils.IsFile(tmpzipPackagePathName)
	if !ret {
		t.Fatalf("tmpzipPackagePathName = %v", tmpzipPackagePathName)
	}

	os.RemoveAll(tmpPackagePathName)
	os.RemoveAll(tmpzipPackagePathName)
}

func TestZipFile2(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".zip"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	err := PackZip(tmpPath, tmpPackageName)
	if err != nil {
		t.Fatal(err)
	}

	ret := utils.IsFile(tmpzipPackagePathName)
	if !ret {
		t.Fatalf("tmpzipPackagePathName = %v", tmpzipPackagePathName)
	}

	os.RemoveAll(tmpPackagePathName)
	os.RemoveAll(tmpzipPackagePathName)
}

func TestZipFile3(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package_noexist"
	err := PackTarGz(tmpPath, tmpPackageName)
	if err == nil {
		t.Fatal(err)
	}
}

func TestZipFile4(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package_noexist"
	err := PackZip(tmpPath, tmpPackageName)
	if err == nil {
		t.Fatal(err)
	}
}

func TestUnZipFile1(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".tar.gz"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	err := PackTarGz(tmpPath, tmpPackageName)
	if err != nil {
		t.Fatal(err)
	}

	ret := utils.IsFile(tmpzipPackagePathName)
	if !ret {
		t.Fatal(ret)
	}

	defer os.RemoveAll(tmpPackagePathName)
	defer os.RemoveAll(tmpzipPackagePathName)

	outputdir := "/tmp/anno_test_unzip"
	os.MkdirAll(outputdir, 0777)
	err = ExtractTarGz(tmpzipPackagePathName, outputdir)
	if err != nil {
		t.Fatal(err)
	}

	ret = utils.IsDir("/tmp/anno_test_unzip/" + tmpPackageName)
	if !ret {
		t.Fatal(ret)
	}

	defer os.RemoveAll(outputdir)
}

func TestUnZipFile2(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".zip"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	err := PackZip(tmpPath, tmpPackageName)
	if err != nil {
		t.Fatal(err)
	}

	ret := utils.IsFile(tmpzipPackagePathName)
	if !ret {
		t.Fatal(ret)
	}

	defer os.RemoveAll(tmpPackagePathName)
	defer os.RemoveAll(tmpzipPackagePathName)

	outputdir := "/tmp/anno_test_unzip"
	os.MkdirAll(outputdir, 0777)
	err = ExtractZip(tmpzipPackagePathName, outputdir)
	if err != nil {
		t.Fatal(err)
	}

	ret = utils.IsDir("/tmp/anno_test_unzip/" + tmpPackageName)
	if !ret {
		t.Fatal(ret)
	}

	defer os.RemoveAll(outputdir)
}

func TestUnZipFile3(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".rar"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	outputdir := "/tmp/anno_test_unzip"
	os.MkdirAll(outputdir, 0777)
	err := ExtractZip(tmpzipPackagePathName, outputdir)
	if err == nil {
		t.Fatal(err)
	}
}

func TestUnZipFile4(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".tar.gz"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	outputdir := "/tmp/anno_test_unzip"
	os.MkdirAll(outputdir, 0777)
	err := ExtractTarGz(tmpzipPackagePathName, outputdir)
	if err == nil {
		t.Fatal(err)
	}
}

func TestUnZipFile5(t *testing.T) {
	tmpPath := "/tmp"
	tmpPackageName := "anno_test_package"
	suffixName := ".zip"
	tmpPackagePathName := filepath.Join(tmpPath, tmpPackageName)
	tmpzipPackagePathName := filepath.Join(tmpPath, tmpPackageName+suffixName)
	os.MkdirAll(tmpPackagePathName, 0777)

	outputdir := "/tmp/anno_test_unzip"
	os.MkdirAll(outputdir, 0777)
	err := ExtractZip(tmpzipPackagePathName, outputdir)
	if err == nil {
		t.Fatal(err)
	}
}
