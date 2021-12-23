# Disaster Communication System - Repeater

## Usage

`$ dis-communication-repeater [option] <tty>`

### Arguments

```
  tty
        TTY device path
        e.g.) /dev/ttyUSB0
```

### Option

```
  -i int
        wait specified amount of time before sending same packet (default 10000)
  -n int
        resend received packet specified number of times (default 3)
  -s int
        remember packet identifications specified number of times (default 256)
```
