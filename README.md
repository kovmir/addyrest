# addyrest

[Addy RESTful API][3] client library.

[![builds.sr.ht status](https://builds.sr.ht/~kovmir/addyrest/commits/master/.build.yml.svg)](https://builds.sr.ht/~kovmir/addyrest/commits/master/.build.yml?)
[![Go Reference](https://pkg.go.dev/badge/git.sr.ht/~kovmir/addyrest.svg)](https://pkg.go.dev/git.sr.ht/~kovmir/addyrest)

# INSTALL

Enter your project directory with `go.mod` inside, and run `go get`:

```bash
go get git.sr.ht/~kovmir/addyrest
```

# USAGE

Go to your [account settings][2] to issue a new token, and then:

```go
package main

import (
	"fmt"

	"git.sr.ht/~kovmir/addyrest"
)

func main() {
	c := addyrest.NewClient("YOUR_TOKEN")

	// Get token name.
	details, err := c.TokenGetAPIDetails()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Token name: %s\n", details.Name)

	// Get first 5 active aliases.
	aliases, err := c.AliasesGet(&addyrest.AliasesGetArgs{
		Filter:   map[string]string{"active": "true"},
		PageSize: 5,
	})
	if err != nil {
		panic(err)
	}
	for i, v := range aliases.Data {
		fmt.Printf("%d. %s\n", i, v.Email)
	}

	// Create a new UUID alias.
	alias, err := c.AliasNew(&addyrest.AliasNewArgs{
		Desc:   "addy client test",
		Domain: "mailer.me",
		Format: addyrest.AliasFmtUUID,
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("alias %s created successfully\n", alias.Data.ID)
}
```

# DOCUMENTATION

The entire codebase resides in [`client.go`](client.go); the rest of the files
define methods and JSONs from the [Addy API reference][3].
[`types.go`](types.go) defines the return JSONs for all API methods. Each
method is more or less self-descriptive and has a URL pointing to the upstream
reference.

[2]: https://app.addy.io/settings/api
[3]: https://app.addy.io/docs/#
