# Collider

A websocket-based signaling server in Go.

## Information

This package is copy from [AppRTC](https://github.com/webrtc/apprtc/tree/master/src/collider) with some modification.

Most modification is for golang package struction.

## Building

1. Install the Go tools and workspaces as documented at http://golang.org/doc/install and http://golang.org/doc/code.html

2. Checkout the `collider` repository

        git clone https://github.com/tihtw/collider.git

3. Enter `collidermain` directory

        cd collidermain

5. Build

        go build

6. Install `collidermain` (Optional)

        go install collidermain

## Running

    ./collidermain  -port=8089 -tls=true

or

    $GOPATH/bin/collidermain -port=8089 -tls=true

## Testing

    go test

## Deployment
These instructions assume you are using Debian 7/8 and Go 1.14.

1. Change [roomSrv](https://github.com/webrtc/apprtc/blob/master/src/collider/collidermain/main.go#L16) to your AppRTC server instance e.g.

```go
var roomSrv = flag.String("room-server", "https://your.apprtc.server", "The origin of the room server")
```

2. Then repeat all steps in the Building section.

### Configutarion

TODO...

### Certificates
If you are deploying this in production, you should use certificates so that you can use secure websockets. Place the `cert.pem` and `key.pem` files in `/cert/`. E.g. `/cert/cert.pem` and `/cert/key.pem`

### Auto restart
1\. Add a `/collider/start.sh` file:

```bash
#!/bin/sh -
/collider/collidermain 2>> /collider/collider.log
```

2\. Make it executable by running `chmod 744 start.sh`.

#### If using inittab otherwise jump to step 5:

3\. Add the following line to `/etc/inittab` to allow automatic restart of the Collider process (make sure to either add `coll` as an user or replace it below with the user that should run collider):
```bash
coll:2:respawn:/collider/start.sh
```
4\. Run `init q` to apply the inittab change without rebooting.

#### If using systemd:

5\. Create a service by doing `sudo nano /lib/systemd/system/collider.service` and adding the following:

```
[Unit]
Description=AppRTC signalling server (Collider)
 
[Service]
ExecStart=/collider/start.sh
StandardOutput=null
 
[Install]
WantedBy=multi-user.target
Alias=collider.service
```
6\. Enable the service: `sudo systemctl enable collider.service`

7\. Verify it's up and running: `sudo systemctl status collider.service`


#### Rotating Logs
To enable rotation of the `/collider/collider.log` file add the following contents to the `/etc/logrotate.d/collider` file:

```
/collider/collider.log {
    daily
    compress
    copytruncate
    dateext
    missingok
    notifempty
    rotate 10
    sharedscripts
}
```

The log is rotated daily and removed after 10 days. Archived logs are in `/collider`.
