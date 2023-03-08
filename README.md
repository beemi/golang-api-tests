# golang-postcode-latlong

This repository contains API tests for postcode.io using [go](https://go.dev/) and [testing](https://pkg.go.dev/testing)

###### How to run tests

Set the below environment variables

```bash
export ENV=prod
```

Run tests from single go file

```bash
export ENV=prod
go test -run=TestTablePostCodeLatLong
```

Running specific tests or benchmarks

```bash
go test -run=TestTablePostCodeLatLong/"RM17 6EY"
```

## Other tools
https://transform.tools/json-to-go

# :e-mail: Contact
Owner: [beemi.raja@gmail.com](beemi.raja@gmail.com)
