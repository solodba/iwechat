package impl_test

import "testing"

func TestChatTextBot(t *testing.T) {
	err := svc.ChatBot(ctx)
	if err != nil {
		t.Fatal(err)
	}
}
