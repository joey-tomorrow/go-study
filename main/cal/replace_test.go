package cal

import (
	"fmt"
	"github.com/uniplaces/carbon"
	"regexp"
	"testing"
	"time"
)

func Test_Replace(t *testing.T) {
	s := "xyxvip:xyxvipp0ss@tcp(10.12.32.159:3307)/xyxvipdb?parseTime=true&charset=utf8mb4,utf8&loc=Asia%2FShanghai"
	compile, _ := regexp.Compile("autocommit=(false|true)")
	fmt.Println(compile.FindString(s) == "")

	now := time.Now()
	expireTime := time.Date(now.Year(), now.Month()-1, now.Day(), 0, 0, 0, 0, time.Local)
	fmt.Println(expireTime)

	month := carbon.Now().SubMonths(1)
	fmt.Println(month.String())
	fmt.Println(carbon.Now().IsLeapYear())

}
