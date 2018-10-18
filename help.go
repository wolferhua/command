package command

import (
	"flag"
	"fmt"
	"html/template"
	"strings"
)

type Help struct {
	Cmd
	Name string
}

func (c *Help) GetName() string {
	return "help"
}

func (c *Help) GetDesc() string {
	return "show commands usage, case:`help cmd_name`"
}

func (c *Help) SetFlags() {

}

func (c *Help) Init() {
	name := strings.TrimSpace(flag.Arg(0))
	if len(name) != 0 {
		c.Name = name
	} else {
		c.Name = ""
	}

}

func (c *Help) Run() {
	if c.Name == "" {
		c.ShowAll()
	} else {
		cmd, ok := Get(c.Name)
		if ok {
			c.Show(cmd)
		}
	}
}

func (c *Help) ShowAll() {
	cmds := GetAll()
	maxLen := 0
	f := ""
	a := []interface{}{}

	for name, cmd := range cmds {
		l := len(name)
		if l > maxLen {
			maxLen = l
		}
		f += "% {{.Len}}s\t%s\n"
		a = append(a, YellowBold(cmd.GetName()), Gray(cmd.GetDesc()))
	}

	//输出模板
	tpl := template.New("show-all")
	tpl.Parse(f)
	//绑定数据
	b := &strings.Builder{}
	//执行模板
	tpl.Execute(b, struct {
		Len int
	}{Len: maxLen + 4})
	//输出信息
	Info(fmt.Sprintf(b.String(), a...))
}

func (c *Help) Show(cmd CmdInterface) {
	cmd.SetToHelpMod()
	cmd.SetFlags()
	usages := cmd.GetFlagUsage()
	maxLen := 0
	f := Yellow(cmd.GetName()) + "\n"
	a := []interface{}{}

	if cmd.GetDesc() != "" {
		f += "    " + Gray(cmd.GetDesc()) + "\n\n"
	}
	for name, usage := range usages {
		l := len(name)
		if l > maxLen {
			maxLen = l
		}
		f += "\t% {{.Len}}s\t%s\n"
		a = append(a, YellowBold(name), usage)
	}

	//输出模板
	tpl := template.New("show")
	tpl.Parse(f)
	//绑定数据
	b := &strings.Builder{}
	//执行模板
	tpl.Execute(b, struct {
		Len int
	}{Len: maxLen + 4})
	//输出信息
	Info(fmt.Sprintf(b.String(), a...))
}
