package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"unicode/utf8"

	"github.com/optim-corp/cios-cli/utils/go_advance_type/convert"
	log "github.com/optim-corp/cios-cli/utils/loglog"

	"github.com/urfave/cli/v2"
	"gopkg.in/AlecAivazis/survey.v1"
)

func Question(question []*survey.Question, st interface{}) {
	err := survey.Ask(question, st)
	if err != nil {
		log.Emergency(err.Error())
		panic(err)
	}
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
	log.Debug(_dir)
	if err != nil {
		log.Error(err.Error())
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
		log.Error(err.Error())
	} else {
		result, err := convert.IndentJson(body)
		if err != nil {
			log.Error(err.Error())
		} else {
			Fprintln(result)
		}
	}
}
func OutStructJson(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		log.Error(err.Error())
	} else {
		result, err := convert.IndentJson(body)
		if err != nil {
			log.Error(err.Error())
		} else {
			println(result)
		}
	}
}
func FOutStructJsonSlim(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		log.Error(err.Error())
	} else {
		Fprintln(body)
	}
}
func OutStructJsonSlim(object interface{}) {
	body, err := StructToJsonStr(object)
	if err != nil {
		log.Error(err.Error())
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

func ListUtility(print func()) {
	fmt.Fprintln(Out, "\n********************************************************"+
		"********************************************************\n")
	print()
	fmt.Fprintln(Out, "\n********************************************************"+
		"********************************************************\n")
	Out.Flush()
}

func SpaceRight(val string, len int) string {
	valLen := utf8.RuneCountInString(val)
	for i := 1; 0 < (len - valLen); i++ {
		val += " "
		valLen = utf8.RuneCountInString(val)

	}
	return val
}

func StructToJsonStr(object interface{}) (string, error) {
	result, err := json.Marshal(object)
	return string(result), err
}
