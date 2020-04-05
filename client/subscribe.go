package main

import (
	"fmt"
	stomp "github.com/drawdy/stomp-ws-go"
)

func DoSubscribe(sc stomp.STOMPConnector) (err error) {
	mdCh, err := sc.Subscribe(stomp.Headers{
		stomp.HK_DESTINATION, "/topic/greeting.back",
		stomp.HK_ID, stomp.Uuid(),
	})
	if err != nil {
		return fmt.Errorf("failed to suscribe greeting message: %v", err)
	}

	err = sc.Send(stomp.Headers{
		stomp.HK_DESTINATION, "/app/greeting",
		stomp.HK_ID, stomp.Uuid(),
	}, "hello STOMP!")
	if err != nil {
		return fmt.Errorf("failed to send greeting message: %v", err)
	}

	md := <-mdCh
	if md.Error != nil {
		return fmt.Errorf("receive greeting message caught error: %v", md.Error)
	}

	fmt.Printf("----> receive new message: %v\n", md.Message.BodyString())
	return
}
