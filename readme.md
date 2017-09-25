# HTTPNetworkTiming

> Timing Http Request 

## License 

MIT 

## Usage  

``` go run main.go https://example.com ```

## Example

```
PS C:\Users\tlv59\Code\go\src\github.com\sliverwing\HTTPNetworkTracing> go run main.go http://www.example.com
[DNS START] {Host:www.example.com}
[DNS INFO] {Addrs:[{IP:93.184.216.34 Zone:} {IP:2606:2800:220:1:248:1893:25c8:1946 Zone:}] Err:<nil> Coalesced:false}, COST: 15ms
[ConnectStart] Network: tcp, Addr: 93.184.216.34:80
[ConnectDone] Network: tcp, Addr: 93.184.216.34:80, COST: 238ms
[Got Conn] {Conn:0x124be010 Reused:false WasIdle:false IdleTime:0s}
[GotFirstResponseByte]: Total 494ms
```
