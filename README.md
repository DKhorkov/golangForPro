## Phonebook

Build:

```shell
go build -o phonebook_cli phonebook/cmd/main.go
```

Provide path to phonebook:

```shell
export PHONEBOOK_CSV=<path_to_phonebook>
```

List:

```shell
./phonebook_cli list
```

List in reversed order:

```shell
./phonebook_cli list reverse
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
