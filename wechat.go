package main

import (
	"axiom"
	"time"
)

type WeChatListener struct{}

func (w *WeChatListener) Handle() []*axiom.Listener {

	return []*axiom.Listener{
		{
			Regex: "time|时间|几点",
			HandlerFunc: func(c *axiom.Context) {
				layout := "2006-01-02 15:04:05"
				t := time.Now()
				c.Reply(c.Message.User + " > 现在时间: " + t.Format(layout))
			},
		},
	}
}