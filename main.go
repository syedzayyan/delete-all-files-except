package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/Bios-Marcel/wastebasket"
)

func deleteDirs(root string) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if !info.IsDir() && filepath.Ext(path) != ".tif" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fmt.Println(file + " Deleted")
		//os.Remove(file)
		fmt.Println(wastebasket.Trash(file))
	}
}

func chooseDirectory(w fyne.Window, h *widget.Label) {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		save_dir := "NoPathYet!"
		if err != nil {
			dialog.ShowError(err, w)
			return
		}
		if dir != nil {
			fmt.Println(dir.Path())
			save_dir = dir.Path() // here value of save_dir shall be updated!
		}
		fmt.Println(save_dir)
		h.SetText(save_dir)
		dialog.ShowConfirm("Hello", "Sure You Want To Delete? "+save_dir, func(b bool) {
			if b {
				deleteDirs(save_dir)
			}
		}, w)
	}, w)
}

func main() {
	a := app.New()
	w := a.NewWindow("FileDialogTest")

	hello := widget.NewLabel("Delete Everything But Tifs")
	w.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Go Get Directory!", func() {
			chooseDirectory(w, hello) // Text of hello updated by return value
		}),
	))
	w.Resize(fyne.NewSize(800, 800))
	w.ShowAndRun()
}
