package console

import (
	"io/ioutil"
	"unicode/utf8"

	"github.com/AlecAivazis/survey/v2"
	cnv "github.com/fcfcqloow/go-advance/convert"
	"github.com/fcfcqloow/go-advance/log"
	"github.com/urfave/cli/v2"
)

func Q(question []*survey.Question, st interface{}) {
	Question(question, st)
}
func Question(question []*survey.Question, st interface{}) {
	if err := survey.Ask(question, st); err != nil {
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
		result, err := cnv.IndentJson(body)
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
		result, err := cnv.IndentJson(body)
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
	Fprintln("\n********************************************************" +
		"********************************************************\n")
	print()
	Fprintln("\n********************************************************" +
		"********************************************************\n")
	Flush()
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
	return cnv.Json(object)
}
