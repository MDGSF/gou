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
	"fmt"
	"os/exec"
	"path/filepath"
)

// CmdCopy 用 cp 命令把 oldpath 拷贝到 newpath
func CmdCopy(newpath, oldpath string) error {
	cmd := exec.Command("bash", "-c", "cp -arf "+oldpath+" "+newpath)
	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	return nil
}

/*
PackZip 压缩文件为 zip
@param filePath: 图包所在目录的父目录
@param packagename: 图包名字，不带后缀
例子: cd /tmp;zip -r lane_20180609_11111.zip lane_20180609_11111;cd -

Zip command provides 10 levels of compression ( 0-9 ).
-6 is used as default compression level.
-0 is used for lowest level compression.
-9 is used for hightest level comression
因为图片压缩几乎没有什么效果，把 zip 的压缩 level 调整低一些会快很多。
*/
func PackZip(filePath, packagename string) error {
	zipPackageName := packagename + ".zip"
	cmd := exec.Command("bash", "-c", "cd \""+filePath+"\";"+"zip -0 -r \""+zipPackageName+"\" \""+packagename+"\";")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("bash zip file failed, filePath = %v, packagename = %v, err = %v", filePath, packagename, err)
	}
	return nil
}

/*
ExtractZip 解压 .zip 文件
@param absoluteZipFilePathName: 压缩包的绝对路径
@param outputdir: 解压到哪个目录下
需要处理 linux 下 shell 特殊字符，见 https://blog.csdn.net/u013063153/article/details/73838614
*/
func ExtractZip(absoluteZipFilePathName, outputdir string) error {
	cmd := exec.Command("bash", "-c", "unzip -d \""+outputdir+"\" \""+absoluteZipFilePathName+"\"")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unzip .zip file failed, absoluteZipFilePathName = %v, err = %v", absoluteZipFilePathName, err)
	}
	return nil
}

/*
PackTarGz 压缩文件为 .tar.gz
@param filePath: 图包所在目录的父目录
@param packagename: 图包名字，不带后缀
*/
func PackTarGz(filePath, packagename string) error {
	zipFilePathWithName := filepath.Join(filePath, packagename+".tar.gz")
	cmd := exec.Command("bash", "-c", "tar -zcvf \""+zipFilePathWithName+"\" -C \""+filePath+"\" \""+packagename+"\"")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("bash zip tar file failed, filePath = %v, packagename = %v, err = %v", filePath, packagename, err)
	}
	return nil
}

/*
ExtractTarGz 解压 .tar.gz 文件
@param absoluteZipFilePathName: 压缩包的绝对路径
@param outputdir: 解压到哪个目录下
*/
func ExtractTarGz(absoluteZipFilePathName, outputdir string) error {
	cmd := exec.Command("bash", "-c", "tar -zxvf \""+absoluteZipFilePathName+"\" -C \""+outputdir+"\"")
	_, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("unzip .tar.gz file failed, absoluteZipFilePathName = %v, err = %v", absoluteZipFilePathName, err)
	}
	return nil
}
