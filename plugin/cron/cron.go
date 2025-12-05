package cron

import (
	"context"
	"errors"
	"sync"

	"github.com/robfig/cron/v3"
)

// CronManager 定时任务管理器
type CronManager struct {
	cron   *cron.Cron
	tasks  map[string]cron.EntryID
	mutex  sync.RWMutex
	ctx    context.Context
	cancel context.CancelFunc
}

// NewCronManager 创建定时任务管理器
func NewCronManager() *CronManager {
	ctx, cancel := context.WithCancel(context.Background())
	return &CronManager{
		cron:   cron.New(cron.WithSeconds()),
		tasks:  make(map[string]cron.EntryID),
		ctx:    ctx,
		cancel: cancel,
	}
}

// AddTask 添加定时任务
func (cm *CronManager) AddTask(name, spec string, task func()) error {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if _, exists := cm.tasks[name]; exists {
		return errors.New("任务已存在")
	}

	entryID, err := cm.cron.AddFunc(spec, task)
	if err != nil {
		return err
	}

	cm.tasks[name] = entryID
	return nil
}

// AddEvery10SecondsTask 添加每10秒执行的任务
func (cm *CronManager) AddEvery10SecondsTask(name string, task func()) error {
	return cm.AddTask(name, "*/50 * * * * *", task)
}

// Start 启动定时任务
func (cm *CronManager) Start() {
	cm.cron.Start()
}

// Stop 停止定时任务
func (cm *CronManager) Stop() {
	cm.cancel()
	cm.cron.Stop()
}

// RemoveTask 移除任务
func (cm *CronManager) RemoveTask(name string) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()

	if entryID, exists := cm.tasks[name]; exists {
		cm.cron.Remove(entryID)
		delete(cm.tasks, name)
	}
}
