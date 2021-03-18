![gosecure-logo](https://github.com/pmccau/gosecure-frontend/blob/master/public/gosecure_logo.png)

# gosecure

gosecure is a home security/automation hub. At the moment it features contact sensor notifications to sound a chime when doors or windows are opened. The intention is that it will subsequently support additional widgets, including an alarm system, notifications, and an exposed web interface. This runs on a Raspberry Pi.

## Getting Started

This is a work in progress.

### Prerequisites

Install local support for:

* [React](https://reactjs.org/docs/getting-started.html)
* [Golang](https://golang.org/doc/install)

### Installing

```bash
go get github.com/pmccau/gosecure
cd ~/go/src/github.com/pmccau/gosecure
git submodule update --init --recursive     # As needed
cd frontend
npm install
```

## Running with GPIO support

### Start the go backend
```bash
cd ~/go/src/github.com/pmccau/gosecure
go run main.go
```

### Start the React frontend
```bash
cd ~/go/src/github.com/pmccau/gosecure/frontend
npm start
```

## Running locally without GPIO support

Change line 14 of the routing in `router/router.go` from:
```go
router.HandleFunc("/api/pins", middleware.CheckPins)
```

to

```go
router.HandleFunc("/api/pins", middleware.TCheckPins)
```

Then, **comment out all of the code** with the exception of the package declaration in `middleware/contact_sensor.go
`. This is an inelegant solution, but will allow you to run the go server on a machine that does not have GPIO system
 calls.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details
