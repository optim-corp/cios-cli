package utils

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

func Path(path string) FileService {
	return FileService{Path: path}
}
func (file FileService) ReadFile() ([]byte, error) {
	return ioutil.ReadFile(file.Path)
}
func (file FileService) ReadString() (string, error) {
	b, e := file.ReadFile()
	return string(b), e
}

func (file FileService) ReadDirInfo() ([]os.FileInfo, error) {
	return ioutil.ReadDir(file.Path)
}

func (file FileService) ReadDirRec() ([]DirByt, error) {
	var (
		result = []DirByt{}
		fn     func(path string, files []os.FileInfo)
		err    error
	)
	absPath, err := filepath.Abs(file.Path)
	files, err := file.ReadDirInfo()
	if err != nil {
		return nil, err
	}
	fn = func(path string, files []os.FileInfo) {
		for _, f := range files {
			_path := filepath.Join(path, f.Name())
			if f.IsDir() {
				fs, err := ioutil.ReadDir(_path)
				if err == nil {
					fn(_path, fs)
				}
			} else {
				Log.Debug(f.Name())
				byts, err := ioutil.ReadFile(_path)
				if err == nil {
					result = append(result, DirByt{
						Value:   byts,
						AbsPath: _path,
					})
				}
			}
		}
	}
	fn(absPath, files)
	return result, err
}
func (file FileService) ReadDir() ([]DirByt, error) {
	var (
		result = []DirByt{}
		err    error
	)
	absPath, err := filepath.Abs(file.Path)
	files, err := file.ReadDirInfo()
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		if !f.IsDir() {
			Log.Debug(f.Name())
			_path := filepath.Join(absPath, f.Name())
			byts, err := ioutil.ReadFile(_path)
			if err == nil {
				result = append(result, DirByt{
					Value:   byts,
					AbsPath: _path,
				})
			}
		}
	}
	return result, err
}

func (file FileService) WriteFileAsString(value string) error {
	if err := ioutil.WriteFile(file.Path, []byte(value), 0777); err != nil {
		return err
	}
	Log.Info("Saved "+file.Path)
	return nil
}

func (file FileService) CreateDir() {
	if _, err := os.Stat(file.Path); os.IsNotExist(err) {
		if err := os.Mkdir(file.Path, 0777); err != nil {
			panic(err)
		}
	}
}
func (file FileService) CreateFile() {
	f, err := os.OpenFile(file.Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func (file FileService) IsDir() bool {
	fileInfo, err := os.Stat(file.Path)
	if err != nil {
		panic(err)
	}
	return fileInfo.IsDir()
}
func (file FileService) IsFile() bool {
	return !file.IsDir()
}

func (file FileService) WriteJson(data interface{}) error {
	return WriteJson(file.Path, data)
}
func (file FileService) ReadJsonMap() (interface{}, bool) {
	return ReadJsonMap(file.Path)
}

func (file FileService) LoadJsonStruct(st interface{}) error {
	return LoadJsonStruct(file.Path, st)
}

func (file FileService) LoadYamlStruct(st interface{}) error {
	return LoadYamlStruct(file.Path, st)
}

func (file FileService) WriteYaml(st interface{}) error {
	return WriteYaml(file.Path, st)
}
