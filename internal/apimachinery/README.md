# apimachinery
 Inspired by the kubernetes [internal repo](https://github.com/kubernetes/apimachinery), this package supports scheme, typing, encoding, decoding, and conversion packages for msgraph and msgraph-like API objects
 
 ## Purpose
 This library is a shared dependency for servers and clients to work with the msgraph API infrastructure without direct type dependencies. It's first consumer(s) are/is ```msgraph```.