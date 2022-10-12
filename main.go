// package main

// import (
// 	"github.com/dtylman/gowd"

// 	"fmt"
// 	"github.com/dtylman/gowd/bootstrap"
// 	"time"
// )

// var body *gowd.Element

// func main() {
// 	//creates a new bootstrap fluid container
// 	body = bootstrap.NewContainer(false)
// 	// add some elements using the object model
// 	div := bootstrap.NewElement("div", "well")
// 	row := bootstrap.NewRow(bootstrap.NewColumn(bootstrap.ColumnSmall, 3, div))
// 	body.AddElement(row)
// 	// add some other elements from HTML
// 	div.AddHTML(`<div class="dropdown">
// 	<button class="btn btn-primary dropdown-toggle" type="button" data-toggle="dropdown">Dropdown Example
// 	<span class="caret"></span></button>
// 	<ul class="dropdown-menu" id="dropdown-menu">
// 	<li><a href="#">HTML</a></li>
// 	<li><a href="#">CSS</a></li>
// 	<li><a href="#">JavaScript</a></li>
// 	</ul>
// 	</div>`, nil)
// 	// add a button to show a progress bar
// 	btn := bootstrap.NewButton(bootstrap.ButtonPrimary, "Start")
// 	btn.OnEvent(gowd.OnClick, btnClicked)
// 	row.AddElement(bootstrap.NewColumn(bootstrap.ColumnSmall, 3, bootstrap.NewElement("div", "well", btn)))
// 	//start the ui loop
// 	gowd.Run(body)
// }

// // happens when the 'start' button is clicked
// func btnClicked(sender *gowd.Element, event *gowd.EventElement) {
// 	// adds a text and progress bar to the body
// 	sender.SetText("Working...")
// 	text := body.AddElement(gowd.NewStyledText("Working...", gowd.BoldText))
// 	progressBar := bootstrap.NewProgressBar()
// 	body.AddElement(progressBar.Element)

// 	// makes the body stop responding to user events
// 	body.Disable()

// 	// clean up - remove the added elements
// 	defer func() {
// 		sender.SetText("Start")
// 		body.RemoveElement(text)
// 		body.RemoveElement(progressBar.Element)
// 		body.Enable()
// 	}()

// 	// render the progress bar
// 	for i := 0; i <= 123; i++ {
// 		progressBar.SetValue(i, 123)
// 		text.SetText(fmt.Sprintf("Working %v", i))
// 		time.Sleep(time.Millisecond * 20)
// 		// this will cause the body to be refreshed
// 		body.Render()
// 	}

// }
package main

import (
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v2"
)

func main() {

	// declare two map to hold the yaml content
	base := map[string]interface{}{}
	data1Map := map[string]interface{}{}
	data2Map := map[string]interface{}{}

	// read one yaml file
	data, _ := os.ReadFile("aide-rancher.yaml")
	if err := yaml.Unmarshal(data, &base); err != nil {
		panic(err)
	}

	// read another yaml file
	data1, _ := os.ReadFile("aide-oracle.yaml")
	if err := yaml.Unmarshal(data1, &data1Map); err != nil {
		panic(err)
	}

	data2, _ := os.ReadFile("aide-pi.yaml")
	if err := yaml.Unmarshal(data2, &data2Map); err != nil {
		panic(err)
	}

	base = mergeMaps(base, data1Map)
	base = mergeMaps(base, data2Map)

	output, _ := yaml.Marshal(base)
	if err := os.WriteFile("config", output, fs.FileMode(0777)); err != nil {
		panic(err)
	}
}

func mergeMaps(a, b map[string]interface{}) map[string]interface{} {
	for k, v := range b {
		if sv, ok := v.([]interface{}); ok {
			for _, v1 := range sv {
				if sa, ok := a[k].([]interface{}); ok {
					a[k] = append(sa, v1)
				}
			}
			fmt.Println("   append1()", k)
			continue
		}
		if _, ok := v.(map[interface{}]interface{}); ok {
			fmt.Println("   append2()", k)
			continue
		}
	}
	return a
}
