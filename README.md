# Generate a Fyne GUI Form based on Go types

[![License](https://img.shields.io/github/license/switchupcb/fyneform.svg?style=for-the-badge)](https://github.com/switchupcb/fyneform/blob/main/LICENSE)

Use Fyneform to stop wasting time developing [Fyne GUI Forms](https://docs.fyne.io/widget/form) based on Go types.

## What is Fyneform?

Fyneform is a code generator which generates [Fyne GUI](https://github.com/fyne-io/fyne) form code based on Go types without adding a dependency to your project.

## Table of Contents

| Topic                                      | Category                                                                                                                                                        |
| :----------------------------------------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| [Using Fyneform](#how-do-you-use-fyneform) | [1. Define Go types](#step-1-define-go-types) <br> [2. Configure Copygen](#step-2-configure-copygen) <br> [3. Generate Fyne Forms](#step-3-generate-fyne-forms) |
| [View Output](#output)                     |

## How do you use Fyneform?

This demonstration (at [`example`](/example/)) generates a Fyneform for an `Account` Go type.

### Step 1. Define Go types

Go types are defined in a file.

`./domain/domain.go`

```go
// Account represents a user's account.
type Account struct {
    ID int
    Username string
    Password string
    Name string
}
```

### Step 2. Configure Copygen

[Copygen](https://github.com/switchupcb/copygen) is used to generate code based on Go types.

**1\)** [`example/setup/copygen/setup.yml`](/example/setup/copygen/setup.yml)

Configure the `setup.yml` file.

```yml
# Define where the code is generated.
generated:
  setup: ./setup.go
  output: ../../app/fyneform/fyneform.go
  template: ../fyneform/generate.go

matcher:
    skip: true
```

**2\)** [`example/setup/copygen/setup.go`](/example/setup/copygen/setup.go)

Configure the `setup.go` file.

```go
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
```

**3\)** [`example/setup/fyneform/generate.go`](/example/setup/fyneform/generate.go)

Configure the generator template _(e.g., copy the file)_.

_You can modify this file to customize how a form is generated._

### Step 3. Generate Fyne Forms

Install Copygen.

```
go install github.com/switchupcb/copygen@latest
```

Run the executable in a command line.

```
cd example/app
copygen -yml ../setup/copygen/setup.yml
```

_You must execute this demonstration from the `example/app` directory so the Go types defined in the example's Go module (i.e., `github.com/fyneform/example/app`) are resolved correctly by the Go module system._

### Output

[![Fyneform Example Application Output](/example/app/output.png)](https://github.com/switchupcb/fyneform/blob/main/LICENSE)

