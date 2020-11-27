package main

import (
	"fmt"
	"strings"

	"github.com/jroimartin/gocui"
)

var (
	vm []string
)

func init() {
	for i := 0; i < 20; i++ {
		vm = append(vm, fmt.Sprintf("0100%03d", i+1))
	}
}

func main() {
	g, err := gocui.NewGui(gocui.OutputNormal)
	if err != nil {
		fmt.Println(err)
	}
	defer g.Close()

	g.SetManagerFunc(layout)

	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		fmt.Println(err)
	}

	if err := g.SetKeybinding("items", gocui.KeyTab, gocui.ModNone, sel); err != nil {
		fmt.Println(err)
	}

	if err := g.SetKeybinding("items", gocui.KeyCtrlS, gocui.ModNone, search); err != nil {
		fmt.Println(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		fmt.Println(err)
	}
}

func layout(g *gocui.Gui) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("items", 0, 0, 40, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Items"
		v.FgColor = gocui.ColorCyan
		v.Highlight = true
		v.SelBgColor = gocui.ColorBlue
		v.SelFgColor = gocui.ColorBlack
		v.SetCursor(0, 0)
		g.SetCurrentView("items")
		for _, vmCode := range vm {
			fmt.Fprintln(v, vmCode)
		}
	}

	if v, err := g.SetView("console", 41, 0, maxX-1, maxY-4); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Console"
		v.Highlight = false
		fmt.Fprintln(v, "Console......")
	}

	if v, err := g.SetView("input", 41, maxY-3, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Input"
		fmt.Fprintln(v, "$")
	}

	//if v, err := g.SetView("hello", maxX/2-7, maxY/2, maxX/2+7, maxY/2+2); err != nil {
	//if err != gocui.ErrUnknownView {
	//return err
	//}
	//fmt.Fprintln(v, "Hello world!")
	//}
	return nil
}

func quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

func sel(g *gocui.Gui, v *gocui.View) error {

	if _, err := g.SetCurrentView("items"); err != nil {
		return err
	}
	g.SetViewOnTop("items")
	v, err := g.View("items")
	if err != nil {
		return err
	}
	v.MoveCursor(0, 1, false)
	_, y := v.Cursor()
	l, err := v.Line(y)
	if l == "" {
		v.SetCursor(0, 0)
	}
	vv, _ := g.View("console")
	fmt.Fprintln(vv, l)

	return nil
}

func search(g *gocui.Gui, v *gocui.View) error {
	maxX, maxY := g.Size()
	if v, err := g.SetView("search", maxX/2-12, maxY/2, maxX/2+12, maxY/2+2); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Search"
		v.Editable = true
		g.Cursor = true
		if _, err := g.SetCurrentView("search"); err != nil {
			return err
		}
		if err := g.SetKeybinding("search", gocui.KeyEnter, gocui.ModNone, searchEnter); err != nil {
			return err
		}
	}
	return nil
}

func searchEnter(g *gocui.Gui, v *gocui.View) error {
	searVm := v.Buffer()
	searVm = strings.TrimSpace(searVm)
	index := findVm(searVm)
	if index >= 0 {
		vv, _ := g.View("items")
		vv.SetCursor(0, index)
	}
	g.Cursor = false
	g.DeleteKeybindings("search")
	if err := g.DeleteView(v.Name()); err != nil {
		return err
	}
	if _, err := g.SetCurrentView("items"); err != nil {
		return err
	}

	return nil
}

func findVm(vmCode string) int {
	for i := 0; i < len(vm); i++ {
		if vm[i] == vmCode {
			return i
		}
	}
	return -1
}
