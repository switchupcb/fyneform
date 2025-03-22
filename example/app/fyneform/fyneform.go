// Code generated by github.com/switchupcb/copygen
// DO NOT EDIT.

package fyneform

import "fyne.io/fyne/v2/widget"

// AccountForm generates a UI form for the fields of a domain.Account.
func AccountForm() *widget.Form {
	formEntryAccountID := &widget.FormItem{
		Text:     "Account.ID",
		Widget:   widget.NewEntry(),
		HintText: "",
	}

	formEntryAccountUsername := &widget.FormItem{
		Text:     "Account.Username",
		Widget:   widget.NewEntry(),
		HintText: "",
	}

	formEntryAccountPassword := &widget.FormItem{
		Text:     "Account.Password",
		Widget:   widget.NewEntry(),
		HintText: "",
	}

	formEntryAccountName := &widget.FormItem{
		Text:     "Account.Name",
		Widget:   widget.NewEntry(),
		HintText: "",
	}

	formContainer := widget.NewForm(
		formEntryAccountID,
		formEntryAccountUsername,
		formEntryAccountPassword,
		formEntryAccountName,
	)

	return formContainer
}
