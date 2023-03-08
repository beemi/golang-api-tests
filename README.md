# golang-postcode-latlong

This repository contains API tests for postcode.io using [go](https://go.dev/) and [ginkgo](https://onsi.github.io/ginkgo/)

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

# :e-mail: Contact
Owner: [beemi.raja@gmail.com](beemi.raja@gmail.com)
