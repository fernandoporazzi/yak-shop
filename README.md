# yak-shop

## Up and running (backend)

Install the required dependencies:

```sh
$ go get ./...
```

Build an executable:

```sh
$ go build
```

Run the binary:

```sh
# On Mac
$ ./yak-shop

# On Windows
$ ./yak-shop.exe
```

You should be prompted a few commands you can run. 

To start the HTTP server, execute the following:

```sh
$ ./yak-shop start-server
```

The above command will start a server on port `:8080`.

Possible endpoints:
- GET `/yak-shop/stock/:days`
- GET `/yak-shop/herd/:days`
- POST `/yak-shop/order/:days`

Also, it's possible to retrieve information about the stock using the command line:

```sh
$ ./yak-shop get-data -f herd.xml -d 14
```

Where `-f` is the name of the file and `-d` is the number of days

## Up and running (frontend)

`cd` into the `frontend` and Install the dependencies:

```sh
$ cd frontend
$ yarn install
```

Start the devlopment server:

```sh
$ yarn dev
```

The command above should start a server on port `:3000` and open a new tab on your browser.

## Todo

- Write unit tests (backend)
- Write unit tests (frontend)
- Create abstractions and remove duplicated code on services
- Create abstraction for file opener. Currently it has the logic in two places different places, `cmd/getData.go` and `app/app.go`
- Make typing more consistent
