package main

import (
	"log"
	"os"
	"text/template"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

type Part struct {
	Name  string
	Count int
}

func executeTemplate(text string, data interface{}) {

	tmpl, err := template.New("test").Parse(text)
	check(err)
	err = tmpl.Execute(os.Stdout, data)
	check(err)

}
func main() {
	executeTemplate("Dot is: {{.}}!\n", "ABC") // замість . підставляється значення!!!
	executeTemplate("Dot is: {{.}}!\n", 123.5)

	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", true)
	executeTemplate("start {{if .}}Dot is true!{{end}} finish\n", false)

	templateText := "Name: {{.Name}}\nCount: {{.Count}}\n"
	executeTemplate(templateText, Part{Name: "Fuses", Count: 5})
	executeTemplate(templateText, Part{Name: "Cables", Count: 2})

}
