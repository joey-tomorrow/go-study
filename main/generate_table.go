package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var table = "ALTER TABLE `pipigaoxiao_sdk_user_mapping_%02d` DROP INDEX `uq_uid` ,ADD UNIQUE INDEX `uq_uid` (`uid` ASC, `third_game_id` ASC);"
var table1 = "ALTER TABLE `kuaishou_sdk_user_mapping_%02d` DROP INDEX `uq_uid` ,ADD UNIQUE INDEX `uq_uid` (`uid` ASC, `third_game_id` ASC);"

func main() {
	fmt.Println(GetAppPath())

}

func GetAppPath() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))

	return path[:index]
}
