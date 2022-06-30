package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"

	"jcellar/settings"

	"fyne.io/fyne/v2/layout"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Jdata struct {
	APIKey          string
	InitQueryString bool
	PathRoot        string
}

func loadJsonData() []settings.Config {
	fmt.Println("Loading data from JSON file")

	input, err := ioutil.ReadFile("./jcellar.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	var config []settings.Config

	json.Unmarshal(input, &config)

	return config
}

func saveJsonData(data []settings.Config) {
	fmt.Println("Saving data to JSON file")
	jsonData, _ := json.Marshal(data)
	ioutil.WriteFile("data.json", jsonData, 0644)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("jCellar")

	data := loadJsonData()

	defer saveJsonData(data)

	list := widget.NewList(
		func() int { return len(data) },
		func() fyne.CanvasObject {
			return widget.NewLabel("Config")
		},
		func(lid widget.ListItemID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[lid].APIKey)
			o.(*widget.Label).SetText(data[lid].APIKey)

			// Loop data and set list labels
			for i := 0; i < reflect.ValueOf(data[lid]); i++ {
				o.(*widget.Label).SetText(data[lid].APIKey)
			}
		})

	list.OnSelected = func(id widget.ListItemID) {
		list.Unselect(id)
		d, _ := data.GetValue(id)
		w := myApp.NewWindow("Edit Data")

		itemName := widget.NewEntry()
		itemName.Text = d

		updateData := widget.NewButton("Update", func() {
			data.SetValue(id, itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		deleteData := widget.NewButton("Delete", func() {
			var newData []string
			dt, _ := data.Get()

			for index, item := range dt {
				if index != id {
					newData = append(newData, item)
				}
			}

			data.Set(newData)

			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, updateData, deleteData, cancel))
		w.Resize(fyne.NewSize(800, 200))
		w.CenterOnScreen()
		w.Show()
	}

	add := widget.NewButton("Add", func() {
		w := myApp.NewWindow("Add Data")

		itemName := widget.NewEntry()

		addData := widget.NewButton("Add", func() {
			data.Append(itemName.Text)
			w.Close()
		})

		cancel := widget.NewButton("Cancel", func() {
			w.Close()
		})

		w.SetContent(container.New(layout.NewVBoxLayout(), itemName, addData, cancel))
		w.Resize(fyne.NewSize(400, 200))
		w.CenterOnScreen()
		w.Show()
	})

	exit := widget.NewButton("Quit", func() {
		myWindow.Close()
	})

	myWindow.SetContent(container.NewBorder(nil, container.New(layout.NewVBoxLayout(), add, exit), nil, nil, list))
	myWindow.Resize(fyne.NewSize(600, 600))
	myWindow.SetMaster()
	myWindow.CenterOnScreen()
	myWindow.ShowAndRun()
}
