# Unbuffered channels
### In unbuffered channel there is no capacity to hold any value before it's received. In this type of channels both a sending and receiving goroutine to be ready at the same instant before any send or receive operation can complete. If the two goroutines aren't ready at the same instant, the channel makes the goroutine that performs its respective send or receive operation first wait. Synchronization is fundamental in the interaction between the send and receive on the channel. One can't happen without the other.

# Buffered channels
### In buffered channel there is a capacity to hold one or more values before they're received. In this types of channels don't force goroutines to be ready at the same instant to perform sends and receives. There are also different conditions for when a send or receive does block. A receive will block only if there's no value in the channel to receive. A send will block only if there's no available buffer to place the value being sent.

## Link: https://www.golangprograms.com/go-language/channels.html