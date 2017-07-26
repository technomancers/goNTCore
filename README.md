# ntcore

> Currently under development. Help is wanted. Read TODO list below for a better understanding what needs to get done.

GoLang implementation of WPILibs ntcore package. This repository should stay private until it has been fully tested and implemented properly.

## Network Table V3

To view the standards that this package follows read [NetworkTable3](https://github.com/technomancers/goNTCore/blob/master/networktables3.adoc)

## TODO

- [x] Create all the possible Entry types
  - Include a way to Marshal and Unmarshal them
- [x] Create all the possible Message types
  - Include a way to Marshal and Unmarshal them
- [x] Create a NetworkTable interface that should be used for anything to talk on the Network
- [x] Create a Data interface for NetworkTable interface to use for manipulating data
- [x] Create a NetworkTable struct that implements the NetworkTable interface but consumes the Data interface
- [x] Create a Server that can listen to multiple clients as well as broadcast to them
  - Cleanup old connections
- [x] Create a Client that can listen to a server for messages
  - Should be able to initiate handshake
- [ ] Create a data struct to implament Data that is connected to a Key Value store
  - Has to handle keys, IDs and SNs correctly
  - Should handle persistent flag correctly
- [ ] Server needs to have NetworkTabler methods so that it can send the appropriate messages
- [ ] Client needs to have NetworkTabler methods so that it can send the appropriate messages
- [ ] Fix the Server handshake to return EntryAssign for all entries it knows about
- [ ] Fix the Client handshake to return what is different between what the server has and what it has
- [ ] Modify Server handler method to manipulate data on the server side
- [ ] Modify Client handler method to manipulate data on the client side
