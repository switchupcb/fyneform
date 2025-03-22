package fyneform

import (
	"github.com/fyneform/example/app/domain"
)

// Copygen defines the functions that are generated.
type Copygen interface {
	// AccountForm represents the generated functions name which returns a &widget.Form.
	//
	// You can include multiple Go types (e.g., `domain.Account`) in a single form.
	//
	// WARNING: Do not define Go types with a pointer.
	AccountForm(domain.Account)
}
