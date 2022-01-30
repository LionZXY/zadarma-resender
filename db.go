package main

import "context"

var ctx = context.TODO()

func wasAlreadyPosted(callId string) bool {
	return r.SAdd(ctx, "postedCalls", callId).Val() == 0
}
