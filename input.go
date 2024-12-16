package input

import (
	"fmt"
	"log"

	"github.com/gospider007/blog"
)

func Input(tip ...any) (string, error) {
	fmt.Print(tip...)
	var val string
	_, err := fmt.Scanln(&val)
	return val, err
}
func Option(title, tip string, option map[string]string, defVals ...string) (string, error) {
	log.Print(blog.Color(1, title, "\r\n"))
	for k, v := range option {
		fmt.Print(blog.Color(3, k), " : ", v, "\r\n")
	}
	if len(defVals) > 0 {
		_, ok := option[defVals[0]]
		if ok {
			log.Print(blog.Color(2, "默认值为: ", defVals[0], "\r\n"))
		}
	}
	for {
		if val, err := Input(tip); err != nil {
			if err.Error() != "unexpected newline" {
				return val, err
			} else if len(defVals) > 0 {
				_, ok := option[defVals[0]]
				if ok {
					return defVals[0], nil
				}
			}
		} else if val == "" {
			if len(defVals) > 0 {
				_, ok := option[defVals[0]]
				if ok {
					return defVals[0], nil
				}
			}
		} else {
			_, ok := option[val]
			if ok {
				return val, nil
			}
		}
		log.Print(blog.Color(1, "请输入正确的选项！！！", "\r\n"))
	}
}
