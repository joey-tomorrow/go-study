package main_test

import "testing"

type Args struct {
	A, B int
}
type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func Test_RPC(t *testing.T) {
	args := &Args{7, 8}
	var reply int
	rpc.Client
	call := client.Go("Arith.Multiply", args, &reply, nil)
	<-call.Done
	if call.Error != nil {
		log.Fatal("arith error:", call.Error)
	}
	fmt.Printf("Arith: %d*%d=%d", args.A, args.B, reply)

}
