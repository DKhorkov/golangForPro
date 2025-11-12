## Phonebook

Build:

```shell
go build -o phonebook_cli phonebook/cmd/main.go
```

List:

```shell
./phonebook_cli list
```

Search:

```shell
./phonebook_cli search "+7 (021) 703-77-56"
```

Insert:

```shell
./phonebook_cli insert Иван Романов "+7 (021) 703-77-56"
```

Delete:

```shell
./phonebook_cli delete "+7 (021) 703-77-56"
```
