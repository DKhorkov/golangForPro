# Phonebook

## Server 

Run:

```shell
go run phonebook/server/cmd/main.go
```

## Client

List:

```shell
go run phonebook/client/cmd/main.go list
```

List in reversed order:

```shell
go run phonebook/client/cmd/main.go list -r
```

Search:

```shell
go run phonebook/client/cmd/main.go search -k "+7 (021) 703-77-56"
```

Insert:

```shell
go run phonebook/client/cmd/main.go insert -n Иван -s Романов -p "+7 (021) 703-77-56"
```

Delete:

```shell
go run phonebook/client/cmd/main.go delete -k "+7 (021) 703-77-56"
```
