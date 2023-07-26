package input

import (
	"fmt"
	"log"

	"gitee.com/baixudong/blog"
)

func Print(val any, tip ...any) error {
	fmt.Print(tip...)
	_, err := fmt.Scanln(val)
	if err != nil {
		return err
	}
	return err
}
func Option(title, tip string, option map[string]string, defVals ...string) (string, error) {
	log.Print(blog.Color(1, 0, title, "\r\n"))
	for k, v := range option {
		fmt.Print(blog.Color(3, 0, k), " : ", v, "\r\n")
	}
	if len(defVals) > 0 {
		_, ok := option[defVals[0]]
		if ok {
			log.Print(blog.Color(2, 0, "默认值为: ", defVals[0], "\r\n"))
		}
	}
	var val string
	for {
		if err := Print(&val, tip); err != nil {
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
		log.Print(blog.Color(1, 0, "请输入正确的选项！！！", "\r\n"))
	}
}
