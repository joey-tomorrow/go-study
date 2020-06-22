/*
Package aixpublic 存放了共享技术部所有golang公共代码,
请勿再建立其他repo放置共用代码.


MR方式维护aixpublic公共库


请以Merge Request方式维护公共代码:

0. 检出功能分支（`git checkout -b feature/xxx`）, 如果在非submodule模式下，sh sbin/gomod.sh创建go.mod;

1. 公开导出的类型和函数需要写文档注释，注意是"文档"注释，目录下需要维护doc.go文件, 便于直接用godoc生成文档;

2. 编写测试用例，用测试保障函数功能稳定性；

3. 运行 `make test` 查看测试、覆盖率、linter执行情况；

4. 提交分支， 创建MR ，等待code review.


Go doc语法说明


1. 如果注释后面紧跟着需要说明的对象，并且注释没有间断，则此注释被认为的对象的doc，否则为普通注释;

2. 建议包的说明放在doc.go中，请参考各个目录下的doc.go文件;
包的文档当中，从开始直到遇到英文句号(.)之前的语句，都被godoc认为是包的summary，summary将显示在树形文档索引界面;

3. 如果语句以大写字母开头并且结尾无标点符号，语句前后最少有一个空行，则godoc将其识别为标题;

4. godoc直到遇到空行才重新起一个段落, 否则将各条句子合并为一行;
这条语句既是一个示例，你可打开aixpublic根目录下的doc.go源码查看;

5. 缩进后的内容被godoc认为是代码块，比如:
	now := time.Now()
	return now

详见 https://blog.golang.org/godoc-documenting-go-code
*/
package cal
