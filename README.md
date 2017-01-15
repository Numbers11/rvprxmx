RvPrxMx!
===========

This is `RvPrxMx`, a reverse socks5 proxy server as well as client.

It is by no means extensively tested or finished, **caveat emptor**.

Ideas, bug reports and patches are more than welcome.

Usage
------------

It's really simple:

1. Edit the `config.json` in `/srv`, execute `go build` in the directory and run.

2. Edit the `cln.go` in `/cln` accordingly, execute `go build` in the directory and run.

See how it makes a connection, open the specified HTTP service on your `/srv` machine
 to see a JSON representation of your client, then use a socks5 proxy software to connect
 and get your traffic tunneled. 


Todo & Ideas
---------

* Reconnect option for client
+ security for the web service via auth and/or whitelist
+ connection throttle / flood control/bandwith usage


License
------------

Unlicense
