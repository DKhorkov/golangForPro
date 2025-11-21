# Phonebook

## Server 

Run via IDE:

```shell
go run phonebook/server/cmd/main.go
```

Run via Docker:

```shell
task -d ./phonebook/server/scripts run
```

## Client

Build:

```shell
go build -o phonebook_client phonebook/client/cmd/main.go
```

List:

```shell
./phonebook_client -H <host> -P <port> list
```

List in reversed order:

```shell
./phonebook_client -H <host> -P <port> list -r
```

Search:

```shell
./phonebook_client -H <host> -P <port> search -k "+7 (021) 703-77-56"
```

Insert:

```shell
./phonebook_client -H <host> -P <port> insert -n Иван -s Романов -p "+7 (021) 703-77-56"
```

Delete:

```shell
./phonebook_client -H <host> -P <port> delete -k "+7 (021) 703-77-56"
```
