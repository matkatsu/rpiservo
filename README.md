# rpiservo

## build
```
$ GOARM=6 GOARCH=arm GOOS=linux go build rpiservo.go
$ scp rpiservo pi:/home/pi
```

## run(on raspberry pi)
```
./rpiservo
```

## move(on raspberry pi)
```
$ curl -H 'Content-Type:application/json' 'http://localhost:3000/api/robots/servoBot/commands/move'
```
