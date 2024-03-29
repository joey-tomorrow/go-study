package sixiang

import (
	"fmt"
	"testing"
)

// 【基类】
//定义一个最基础的struct类MsgModel，里面包含一个成员变量msgId
type MsgModel struct {
	msgId   int
	msgType int
}

// MsgModel的一个成员方法，用来设置msgId
func (msg *MsgModel) SetId(msgId int) {
	msg.msgId = msgId
}

func (msg *MsgModel) SetType(msgType int) {
	msg.msgType = msgType
}

//【子类】
// 再定义一个struct为GroupMsgModel，包含了MsgModel，即组合，但是并没有给定MsgModel任何名字，因此是匿名组合
type GroupMsgModel struct {
	MsgModel

	// 如果子类也包含一个基类的一样的成员变量，那么通过子类设置和获取得到的变量都是基类的
	msgId int
}

func (group *GroupMsgModel) GetId() int {
	return group.msgId
}


//func (group *GroupMsgModel) SetId(msgId int) {
//group.msgId = msgId
//}


func Test_zuhe(t *testing.T) {
	group := &GroupMsgModel{}

	group.SetId(123)
	group.SetType(1)

	fmt.Println("group.msgId =", group.msgId, "\tgroup.MsgModel.msgId =", group.MsgModel.msgId)
	fmt.Println("group.msgType =", group.msgType, "\tgroup.MsgModel.msgType =", group.MsgModel.msgType)
}
