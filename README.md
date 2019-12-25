# SEC KINKAO🥩

sec kinkao is super application.

## Getting started

### Install Golang (go1.12+)

Please check the official golang installation guide before you start. [Official Documentation](https://golang.org/doc/install)
Also make sure you have installed a go1.12+ version.

### Environment Config

make sure your ~/.*shrc have those variable:

```bash
~  echo $GOPATH
/Users/wagyu/go
~  echo $GOROOT
/usr/local/go/
~  echo $PATH
...:/usr/local/go/bin:/Users/wagyu/test//bin:/usr/local/go/bin
```

For more info and detailed instructions please check this guide: [Setting GOPATH](https://github.com/golang/go/wiki/SettingGOPATH)

### Clone the repository

Clone this repository:

```bash
~ git clone https://gitlab.odds.team/internship/sec-kinkao/goback.git
```

Or simply use the following command which will handle cloning the repo:

```bash
~ go get -u -v gitlab.odds.team/internship/sec-kinkao/goback
```

Switch to the repository folder

```bash
~ cd $GOPATH/src/gitlab.odds.team/internship/sec-kinkao/goback
```

### Install dependencies

```bash
~ go mod download
```

### Run

```bash
~ go run server.go
```

### Build

```bash
~ go build
```

### Tests

```bash
~ go test ./...
```

