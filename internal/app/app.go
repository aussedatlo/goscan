package app

import (
	"github.com/dtylman/gowd"
	"github.com/dtylman/gowd/bootstrap"
)

//App application type
type App struct {
	root   *gowd.Element
	main   *gowd.Element
	table  *bootstrap.Table
	navBar *navBarApp
}

type navBarApp struct {
	root  *gowd.Element
	input *gowd.Element
}

//CreateApp create application
func CreateApp() *App {
	a := &App{}
	a.root = bootstrap.NewContainer(true)
	a.main = bootstrap.NewContainer(true)
	a.createNavBar()
	a.createTable()

	a.root.AddElement(a.navBar.root)
	a.root.AddElement(a.main)

	return a
}

//RunApp run application
func RunApp(a *App) {
	gowd.Run(a.root)
}

func (a *App) createNavBar() {
	navbar := &navBarApp{}
	title := bootstrap.NewElement("div", classNavBarDivTitle)
	title.AddHTML(`GoScan`, nil)

	buttonOption :=
		bootstrap.NewButton(classNavBarButtonOption, "Op")
	divButtonOption :=
		bootstrap.NewElement("div", classNavBarDivButtonOption)
	divButtonOption.AddElement(buttonOption)

	button := bootstrap.NewButton(classNavBarButtonScan, "Scan")
	button.OnEvent(gowd.OnClick, a.btnClicked)
	divButtonScan := bootstrap.NewElement("div", classNavBarDivButtonScan)
	divButtonScan.AddElement(button)

	navbar.input = bootstrap.NewInput("ip address")
	navbar.input.SetClass(classNavBarInputIP)
	navbar.input.SetValue("192.168.1.0/24")
	navbar.input.SetAttribute("type", "text")
	navbar.input.SetAttribute("placeholder", "target(s)")
	divInput := bootstrap.NewElement("div", classNavBarDivInput)
	divInput.AddElement(navbar.input)

	nb := bootstrap.NewElement("nav", classNavBar)
	container := bootstrap.NewContainer(true)

	nb.AddElement(container)
	container.AddElement(title)
	container.AddElement(divInput)
	container.AddElement(divButtonScan)
	container.AddElement(divButtonOption)

	navbar.root = nb
	a.navBar = navbar
}

func (a *App) printTargetsUp(l []string) {

	for _, host := range l {
		row := a.table.AddRow()
		p, v := getProductVersion(host)
		row.AddCells(host, p, v)
		a.root.Render()
	}

	a.root.Render()
}

func (a *App) createTable() {
	a.table = bootstrap.NewTable(classMainTable)
	e := a.table.AddHeader("Host")
	e.SetClass("head-1")
	e = a.table.AddHeader("Product")
	e.SetClass("head-2")
	e = a.table.AddHeader("Version")
	e.SetClass("head-3")
	a.main.AddElement(a.table.Element)
	a.root.Render()
}

func (a *App) btnClicked(sender *gowd.Element, event *gowd.EventElement) {
	// adds a text and progress bar to the body
	sender.SetText("Working...")
	sender.Disable()
	a.main.RemoveElements()
	a.createTable()
	a.root.Render()
	res := getTargetsUp(a.navBar.input.GetValue())
	a.printTargetsUp(res)
	sender.SetText("Scan")
	sender.Enable()
	a.root.Render()
}
