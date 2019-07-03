# gorm-snippets
Code snippets generator for gorm

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

## Snippets
* FindByID
* FindByParam
