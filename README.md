# GoSimpleArtnet

GoSimpleArtnet is a lightweight and straightforward implementation of the ArtNet protocol written in Go. It is designed to help developers quickly integrate ArtNet functionalities with minimal setup.

ArtNet is a standard protocol for communicating DMX-512 data over IP networks, commonly used in lighting control systems for entertainment, architectural, and broadcast applications. This package simplifies working with ArtNet using the Go programming language.


## Example
The following example demonstrates how to create, configure, and use an ArtNet node to control a DMX universe:

```go
//Create a ArtNetNode that broadcast on the entire network
node, _ := node.NewArtNetNode("255.255.255.255:6454")
//Start the node worker
node.start()
//Get dmx universe 0
u := node.GetUniverse(0)
//Set channel 2 to 125
u.SetChannel(2, 125)
```

## License

GoSimpleArtnet is licensed under the _GNU Lesser General Public License v3.0_ License. 

---
