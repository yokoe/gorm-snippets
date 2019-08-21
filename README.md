# gorm-snippets

[![Maintainability](https://api.codeclimate.com/v1/badges/0653c17e31664434a19a/maintainability)](https://codeclimate.com/github/yokoe/gorm-snippets/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/0653c17e31664434a19a/test_coverage)](https://codeclimate.com/github/yokoe/gorm-snippets/test_coverage)

Code snippet generator for GORM. Generates customizable DB helper functions work with GORM. Designed to be used by code generators that generates repository layer code.

## Usage

### Golang

```
s, _ := snippet.FindByParam("model.Book", "UUID", "string")
```

will generate

```
func findBookByUUID(db *gorm.DB, arg string) (*model.Book, error) {
        if len(arg) == 0 {
                return nil, xerrors.Errorf("UUID must be non-nil.")
        }
        var obj model.Book
        if err := db.Find(&obj, "`uuid` = ?", arg).Error; err != nil {
                return nil, xerrors.Errorf("failed to find model.Book by UUID: %w", err)
        }
        return &obj, nil
}
```

### Command line

TBD

## Snippets

- FindByID
- FindByParam
- BatchFindByID
