# Go Pi Sandbox

## Wiring

**NOTE!** When connecting and referencing pins, use the Arduino pin number! For example, use pin 2 a.k.a D2 in code
instead of physical pin 4.

## Cross Compiling

Change directories to where `main.go` resides and run `go build` with
environment variables for the target platform.

```shell
$ cd cmd/go-pi-sandbox
$ env GOOS=linux GOARCH=arm GOARM=5 go build
```

## Executing

On our computer:

1. Go into the Arduino project folder:
```
$ cd $~/dev/github.com/andycondon/pathfinder/src/main/arduino/Motion
```
2. Compile the Arduino code:
```
Arduino UI -> Sketch -> Export compiled binary
```
or
```
VS Code
```
3. Send the compiled `.hex` to the Raspberry pi for flashing:
```
$ scp Motion.ino.standard.hex pi@192.168.3.147:/home/p
```

On the Raspberry Pi:

1. Flash the `.hex` file:
```
sudo avrdude -P gpio -c gpio -p atmega328p -U flash:w:Motion.ino.standard.hex
```

2. Execute:
Make both motors go forward at 50%:
```
$ ./go-pi-sandbox 010180 // Motor 1
$ ./go-pi-sandbox 020180 // Motor 2
```

Set the cursor to row 0, column 0:
```
./go-pi-sandbox 00010000
```

Say Thorin's name:
t: 74 h: 68 o: 6f r: 72 i: 69 n: 6e
```
./go-pi-sandbox 000074686f72696e
```