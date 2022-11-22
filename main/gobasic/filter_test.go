package gobasic

import (
	"context"
	"fmt"
	"testing"
)

type NotifyTest interface {
	Filter(ctx context.Context) NotifyTest
	Notify(ctx context.Context, webhookURL string) error
}

type NotifyImplemention struct {

}

func (n *NotifyImplemention) Filter(ctx context.Context) NotifyTest {
	fmt.Println("filter")
	return n
}

func (n *NotifyImplemention) Notify(ctx context.Context, webhookURL string) error {
	fmt.Println("notify")
	return nil
}


type NotifyImplemention1 struct {

}

func (n NotifyImplemention1) Filter(ctx context.Context) NotifyTest {
	fmt.Println("filter1")
	return n
}

func (n NotifyImplemention1) Notify(ctx context.Context, webhookURL string) error {
	fmt.Println("notify1")

	return nil
}

func Test_filter(t *testing.T) {
	ctx := context.Background()
	implemention := NotifyImplemention{}
	implemention.Filter(ctx).Notify(ctx, "123")

	implemention1 := NotifyImplemention1{}
	implemention1.Filter(ctx).Notify(ctx, "123")
}
