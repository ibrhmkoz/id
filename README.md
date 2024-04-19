# ID Package

The ID package provides a simple and flexible way to generate and manage unique identifiers (IDs) in Go applications. It
leverages the power of ULIDs (Universally Unique Lexicographically Sortable Identifiers) and Go generics to create
strongly typed IDs for various entities in your application.

## Features

- Generate unique IDs using ULIDs
- Use Go generics as phantom types to prevent comparing IDs of different entities
- Restore IDs from their string representation
- Unwrap IDs to obtain their underlying string value
- Flexible and database-independent ID generation

## Rationale

The ID package is designed to address common challenges and needs in enterprise applications:

1. **Recurring Theme**: ID generation and management is a recurring theme in enterprise applications. Having a
   dedicated, reusable package for this purpose promotes code consistency and maintainability.

2. **Database Independence**: By generating IDs within the application rather than relying on the database, the ID
   package offers more flexibility. It allows you to generate IDs independently of the database, making it easier to
   work with different storage systems or migrate between them.

3. **K-Sortable IDs**: The package utilizes ULIDs, which are k-sortable. This means that the IDs are lexicographically
   sortable by time, providing a natural ordering for entities based on their creation timestamp. K-sortable IDs can be
   beneficial for various use cases, such as sorting, sharding, and indexing.

4. **Type Safety**: The use of Go generics as phantom types prevents accidentally comparing IDs of different entities.
   By associating each ID with a specific entity type, the package ensures type safety and helps catch potential bugs at
   compile time.

## Installation

To install the ID package, use the following command:

```
go get github.com/ibrhmkoz/id
```

## Usage

Here's a basic example of how to use the ID package:

```go
package main

import (
   "fmt"
   "github.com/ibrhmkoz/id"
)

type User struct{}

func main() {
   uid := id.NewID[User]()
   fmt.Println("Generated User ID:", uid)

   sid := uid.Unwrap()
   fmt.Println("String representation of User ID:", sid)

   rid, err := id.RestoreID[User](sid)
   if err != nil {
      fmt.Println("Error restoring User ID:", err)
   } else {
      fmt.Println("Restored User ID:", rid)
   }
}

```

