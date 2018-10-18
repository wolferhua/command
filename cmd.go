package command

import (
	"flag"
	"sync"
	"time"
)

//cmd 接口
type CmdInterface interface {
	GetName() string                 //命令行名称
	GetDesc() string                 //命令行描述
	SetFlags()                       //参数设置
	GetFlagUsage() map[string]string //参数设置
	Init()                           //初始化
	Run()                            //运行
	SetToHelpMod()                   //运行
	GetHelpMod() bool                //运行
}

//cmd
type Cmd struct {
	helpMod    bool
	flagUsages map[string]string
	mu         sync.Mutex
}

//获取名称
func (c *Cmd) GetName() string {
	return ""
}

//获取数据
func (c *Cmd) GetDesc() string {
	return ""
}

//设置参数
func (c *Cmd) SetFlags() {

}

//设置参数
func (c *Cmd) GetFlagUsage() map[string]string {
	return c.flagUsages
}

//初始化
func (c *Cmd) Init() {

}

//运行入口
func (c *Cmd) Run() {

}

func (c *Cmd) SetToHelpMod() {
	c.helpMod = true
}

func (c *Cmd) GetHelpMod() bool {
	return c.helpMod
}

func (c *Cmd) SetUsage(name, usage string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.flagUsages == nil {
		c.flagUsages = map[string]string{}
	}
	c.flagUsages[name] = usage
}

func (c *Cmd) Bool(name string, value bool, usage string) *bool {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Bool(name, value, usage)
}
func (c *Cmd) BoolVar(p *bool, name string, value bool, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.BoolVar(p, name, value, usage)
}
func (c *Cmd) Duration(name string, value time.Duration, usage string) *time.Duration {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Duration(name, value, usage)
}
func (c *Cmd) DurationVar(p *time.Duration, name string, value time.Duration, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.DurationVar(p, name, value, usage)
}
func (c *Cmd) Float64(name string, value float64, usage string) *float64 {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Float64(name, value, usage)
}
func (c *Cmd) Float64Var(p *float64, name string, value float64, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.Float64Var(p, name, value, usage)
}
func (c *Cmd) Int(name string, value int, usage string) *int {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Int(name, value, usage)
}
func (c *Cmd) Int64(name string, value int64, usage string) *int64 {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Int64(name, value, usage)
}
func (c *Cmd) Int64Var(p *int64, name string, value int64, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.Int64Var(p, name, value, usage)
}
func (c *Cmd) IntVar(p *int, name string, value int, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.IntVar(p, name, value, usage)
}
func (c *Cmd) String(name string, value string, usage string) *string {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.String(name, value, usage)
}
func (c *Cmd) StringVar(p *string, name string, value string, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.StringVar(p, name, value, usage)
}
func (c *Cmd) Uint(name string, value uint, usage string) *uint {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Uint(name, value, usage)
}
func (c *Cmd) Uint64(name string, value uint64, usage string) *uint64 {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return nil
	}
	return flag.Uint64(name, value, usage)
}
func (c *Cmd) Uint64Var(p *uint64, name string, value uint64, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.Uint64Var(p, name, value, usage)
}
func (c *Cmd) UintVar(p *uint, name string, value uint, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.UintVar(p, name, value, usage)
}
func (c *Cmd) Var(value flag.Value, name string, usage string) {
	c.SetUsage(name, usage)
	if c.GetHelpMod() {
		return
	}
	flag.Var(value, name, usage)
}
