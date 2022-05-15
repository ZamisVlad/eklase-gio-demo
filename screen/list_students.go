package screen

import (
	"eklase/state"
	"log"
	"strings"

	"gioui.org/layout"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// AddStudent defines a screen layout for adding a new student.
func AddStudent(th *material.Theme, state *state.State) Screen {
	var (
		name    widget.Editor
		surname widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	enabledIfNameOK := func(w layout.Widget) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			name := strings.TrimSpace(name.Text())
			surname := strings.TrimSpace(surname.Text())
			if name == "" && surname == "" { // Either name or surname is OK.
				gtx = gtx.Disabled()
			}
			return w(gtx)
		}
	}
	editsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, material.Editor(th, &name, "First name").Layout),
			layout.Rigid(spacer.Layout),
			layout.Flexed(1, material.Editor(th, &surname, "Last name").Layout),
		)
	}
	buttonsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceStart}.Layout(gtx,
			layout.Rigid(material.Button(th, &close, "Close").Layout),
			layout.Rigid(spacer.Layout),
			layout.Rigid(enabledIfNameOK(material.Button(th, &save, "Save").Layout)),
		)
	}
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(editsRowLayout)),
			layout.Rigid(rowInset(buttonsRowLayout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		if save.Clicked() {
			err := state.AddStudent(
				strings.TrimSpace(name.Text()),
				strings.TrimSpace(surname.Text()),
			)
			if err != nil {
				// TODO: Show an error toast.
				log.Printf("unable to add student: %v", err)
			}
			return MainMenu(th, state), d
		}
		return nil, d
	}
}

//Alina + Valerija (team work)
func AddClass(th *material.Theme, state *state.State) Screen {
	var (
		year     widget.Editor
		modifier widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	enabledIfNameOK := func(w layout.Widget) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			year := strings.TrimSpace(year.Text())
			modifier := strings.TrimSpace(modifier.Text())
			if year == "" || modifier == "" {
				gtx = gtx.Disabled()
			} else if strings.ContainsAny(year, "qwertyuiopasdfghjklzxcvbnm") {
				gtx = gtx.Disabled()
			} else if strings.ContainsAny(modifier, "1234567890") {
				gtx = gtx.Disabled()
			}
			return w(gtx)
		}
	}
	editsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, material.Editor(th, &year, "Year").Layout),
			layout.Rigid(spacer.Layout),
			layout.Flexed(1, material.Editor(th, &modifier, "Modifier").Layout),
		)
	}
	buttonsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceStart}.Layout(gtx,
			layout.Rigid(material.Button(th, &close, "Close").Layout),
			layout.Rigid(spacer.Layout),
			layout.Rigid(enabledIfNameOK(material.Button(th, &save, "Save").Layout)),
		)
	}
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(editsRowLayout)),
			layout.Rigid(rowInset(buttonsRowLayout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		if save.Clicked() {
			err := state.AddClass(
				strings.TrimSpace(year.Text()),
				strings.TrimSpace(modifier.Text()),
			)
			if err != nil {
				log.Printf("unable to add class: %v", err)
			}
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
func AddGroup(th *material.Theme, state *state.State) Screen {
	var (
		student widget.Editor
		class   widget.Editor

		close widget.Clickable
		save  widget.Clickable
	)
	enabledIfNameOK := func(w layout.Widget) layout.Widget {
		return func(gtx layout.Context) layout.Dimensions {
			student := strings.TrimSpace(student.Text())
			class := strings.TrimSpace(class.Text())
			if student == "" || class == "" {
				gtx = gtx.Disabled()
			} else if strings.ContainsAny(student, "qwertyuiopasdfghjklzxcvbnm") {
				gtx = gtx.Disabled()
			} else if strings.ContainsAny(class, "1234567890") {
				gtx = gtx.Disabled()
			}
			return w(gtx)
		}
	}
	editsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal}.Layout(gtx,
			layout.Flexed(1, material.Editor(th, &student, "student").Layout),
			layout.Rigid(spacer.Layout),
			layout.Flexed(1, material.Editor(th, &class, "class").Layout),
		)
	}
	buttonsRowLayout := func(gtx layout.Context) layout.Dimensions {
		return layout.Flex{Axis: layout.Horizontal, Spacing: layout.SpaceStart}.Layout(gtx,
			layout.Rigid(material.Button(th, &close, "Close").Layout),
			layout.Rigid(spacer.Layout),
			layout.Rigid(enabledIfNameOK(material.Button(th, &save, "Save").Layout)),
		)
	}
	return func(gtx layout.Context) (Screen, layout.Dimensions) {
		d := layout.Flex{Axis: layout.Vertical}.Layout(gtx,
			layout.Rigid(rowInset(editsRowLayout)),
			layout.Rigid(rowInset(buttonsRowLayout)),
		)
		if close.Clicked() {
			return MainMenu(th, state), d
		}
		if save.Clicked() {
			err := state.AddClass(
				strings.TrimSpace(student.Text()),
				strings.TrimSpace(class.Text()),
			)
			if err != nil {
				log.Printf("unable to add class: %v", err)
			}
			return MainMenu(th, state), d
		}
		return nil, d
	}
}
