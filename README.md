RvPrxMx!
===========

This is `RvPrxMx`, a reverse socks5 proxy server as well as client.

It is by no means extensively tested or finished, **caveat emptor**.

Ideas, bug reports and patches are more than welcome.

Simple usage
------------

It's really simple:

1. Edit the `config.json` in `/srv`, build and run.

2. Edit the `cln.go` in `/cln` accordingly, build and run.

See how it makes a connection, open the specified HTTP service on your `/srv` machine
 to see a JSON representation of your client, then use a socks5 proxy software to connect
 and get your traffic tunneled. 


Todo & Ideas
---------

* TLS encryption for the tunnel connection
+ Command line options for the client
+ Stealth mode for the client


License
------------

Unlicense
