## Phonebook

Build:

```shell
go build -o phonebook_cli phonebook/cmd/main.go
```

Provide phonebook source:

```shell
export PHONEBOOK_SOURCE=<value>
```

Available source values: <b>csv</b>, <b>json</b>.

Provide path to phonebook:

```shell
export PHONEBOOK_PATH=<path_to_phonebook>
```

List:

```shell
./phonebook_cli list
```

List in reversed order:

```shell
./phonebook_cli list -r
```

Search:

```shell
./phonebook_cli search -k "+7 (021) 703-77-56"
```

Insert:

```shell
./phonebook_cli insert -n Иван -s Романов -p "+7 (021) 703-77-56"
```

Delete:

```shell
./phonebook_cli delete -k "+7 (021) 703-77-56"
```
