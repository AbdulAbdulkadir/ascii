# A.S.C.I.I.
Ascii Server Created In Instants
## Functionality

A.S.C.I.I. is an ASCII art generator. It is a small microservice that can return a random ASCII art when it is called by the client program. The ASCII art generator is a Go program that can store a database of random ASCII arts in Mongo. It uses a gRPC API that can be called by a client to randomly retrieve a random art piece. The Go Client is supplied so that you can compile it into a command-line program. The go client can then be hooked up to your own terminal and your server's terminals so that you can be presented with a fun ASCII art greeting when you log in.


### Commands


If you would to upload your own ascii art to the server then use the flag -upload= as shown below

```
ascii -upload=filePathToArt.txt
```
If you would like to change the connection to the server then use the flag -url= as shown below
```
ascii -url=https://address.com
```
