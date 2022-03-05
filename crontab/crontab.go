// Package crontab 处理定时任务
package crontab

import "time"

type CronTab func() time.Duration

// StartCronTab 启动定时任务
func StartCronTab(crontab CronTab) {
	go func() {
		for {
			d := crontab()
			time.Sleep(d)
		}
	}()
}
