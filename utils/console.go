package utils

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Question(question []*survey.Question, st interface{}) {
	err := survey.Ask(question, st)
	if err != nil {
		Log.Emergency(err.Error())
		panic(err)
	}
}
func FindFolderIncludeFile(path string, fileKey string) (*string, error) {
	path, err := filepath.Abs(path)
	if err != nil || Path(path).IsFile() {
		println(path)
		return nil, errors.New("No Dir String")
	}
	if dirs, err := ioutil.ReadDir(path); err != nil {
		return nil, err
	} else {
		item := []string{}
		for _, d := range dirs {
			if strings.Contains(d.Name(), fileKey) {
				return &path, nil
			}
			if d.IsDir() {
				item = append(item, d.Name())
			}
		}
		if len(item) == 0 {
			return nil, errors.New("No Store")
		}
		ans := struct{ Value string }{}
		Question([]*survey.Question{
			{
				Name: "value",
				Prompt: &survey.Select{
					Options: item,
					Message: "Pick a File or Dir",
				},
			},
		}, &ans)
		return FindFolderIncludeFile(filepath.Join(path, ans.Value), fileKey)
	}
}
func PickFile(path string, filter *string) ([]byte, error) {
	path, err := filepath.Abs(path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	if !fileInfo.IsDir() {
		file, err := ioutil.ReadFile(path)
		return file, err
	}
	dirs, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	item := []string{}
	for _, d := range dirs {
		if filter != nil {
			if !strings.Contains(d.Name(), *filter) && !d.IsDir() {
				break
			}
		}
		item = append(item, d.Name())
	}
	ans := struct{ Value string }{}
	Question([]*survey.Question{
		{
			Name: "value",
			Prompt: &survey.Select{
				Options: item,
				Message: "Pick a File or Dir",
			},
		},
	}, &ans)
	return PickFile(filepath.Join(path, ans.Value), filter)
}

func CliArgsForEach(c *cli.Context, fun func(val string)) {
	for i := 0; i < c.Args().Len(); i++ {
		fun(c.Args().Get(i))
	}
}
func CliArgs(c *cli.Context) []string {
	result := []string{}
	for i := 0; i < c.Args().Len(); i++ {
		result = append(result, c.Args().Get(i))
	}
	return result
}
func ListDirs(_dir string, indent string) {
	dirs, err := ioutil.ReadDir(_dir)
	Log.Debug(_dir)
	if err != nil {
		Log.Error(err.Error())
		return
	}
	for _, dir := range dirs {
		if dir.IsDir() {
			println(indent + dir.Name() + " <Directory>")
			ListDirs(_dir+"/"+dir.Name(), " "+indent+"-")
		} else {
			println(indent + dir.Name() + " <File>")
		}

	}
}

func FOutStructJson(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		Log.Error(err.Error())
	} else {
		result, err := IndentJson(body)
		if err != nil {
			Log.Error(err.Error())
		} else {
			Fprintln(result)
		}
	}
}
func OutStructJson(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		Log.Error(err.Error())
	} else {
		result, err := IndentJson(body)
		if err != nil {
			Log.Error(err.Error())
		} else {
			println(result)
		}
	}
}
func FOutStructJsonSlim(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		Log.Error(err.Error())
	} else {
		Fprintln(body)
	}
}
func OutStructJsonSlim(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		Log.Error(err.Error())
	} else {
		println(body)
	}
}

func GetConsoleMultipleLine(message string) string {
	ans := struct{ Body string }{}
	Question([]*survey.Question{
		{Name: "body", Prompt: &survey.Multiline{Message: message}},
	}, &ans)
	return ans.Body
}
