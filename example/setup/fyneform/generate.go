// DO NOT CHANGE PACKAGE

// Package template provides a template used by copygen to generate custom code.
package template

import (
	"strings"

	"github.com/switchupcb/copygen/cli/models"
)

// Generate generates code.
// GENERATOR FUNCTION.
// EDITABLE.
// DO NOT REMOVE.
func Generate(gen *models.Generator) (string, error) {
	var content strings.Builder

	content.WriteString(string(gen.Keep) + "\n")

	content.WriteString("import \"fyne.io/fyne/v2/widget\"\n")

	for i := range gen.Functions {
		function := gen.Functions[i]

		// Write the function signature to return the form.
		contentComment := "// " + function.Name + " generates a UI form for the fields of a "
		contentFuncSignature := "func " + function.Name + "() *widget.Form {\n"

		// Generate a form entry for each Go type field.
		var contentForm strings.Builder

		for j := range function.From {
			goType := function.From[j]

			contentComment += goType.Field.FullDefinition() + " "

			// Generate a form entry for each field.
			formEntryVariables := make([]string, 0, len(goType.Field.Fields))

			for k := range goType.Field.Fields {
				subfield := goType.Field.Fields[k]

				// formEntryVariable represents a form entry variable (e.g., `formEntryAccountID`).
				formEntryVariable := "formEntry" + goType.Field.Definition + subfield.Name

				formEntryVariables = append(formEntryVariables, formEntryVariable)

				contentForm.WriteString(formEntryVariable + " := &widget.FormItem{\n")
				contentForm.WriteString("Text: \"" + goType.Field.Definition + "." + subfield.Name + "\",\n")
				contentForm.WriteString("Widget:   widget.NewEntry(),\n")
				contentForm.WriteString("HintText: \"\",\n")
				contentForm.WriteString("}\n")
				contentForm.WriteString("\n")
			}

			// Generate a form container.
			contentForm.WriteString("formContainer := widget.NewForm(\n")

			for n := range formEntryVariables {
				contentForm.WriteString(formEntryVariables[n] + ",\n")
			}

			contentForm.WriteString(")\n")
		} // end for function.From

		// Write the function.
		content.WriteString(contentComment[:len(contentComment)-1] + ".\n")
		content.WriteString(contentFuncSignature)
		content.WriteString(contentForm.String())
		content.WriteString("\nreturn formContainer\n")
		content.WriteString("}") // contentFuncSignature end bracket
		content.WriteString("\n")

	} // end for function

	return content.String(), nil
}
