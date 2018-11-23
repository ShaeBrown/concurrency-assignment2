# Byzantine Generals
In this implementation, traitors will always communicate the opposite command.  
The generals will RETREAT if there is a TIE vote.
The implementation and tests used where based on [this article](https://marknelson.us/posts/2007/07/23/byzantine.html)

## Recursive Function Call
`func om(m int, commander int, command Command, loyal []bool, orders map[string]Command, lieu mapset.Set, prefix string) Command`

**m** - level of recursion, or number of traitors  
**commander** - index of the commander  
**loyal** - array of whether each general is loyal or not  
**orders** - a map where messages passed from and to generals are stored  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;key:  the message id  
&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;value: ATTACK or RETREAT  
**lieu** - a set of index of the lietenants, in recursive calls the amount of lieutenants are decreased

**message id**- a message id, "023" would be a message sent from general 2 to general 3 where general 2 originally got the message from general 0. This message was sent in the second round. If m was originally 1 then m would be 0 when this message was sent.

At each recursive call, m is decreased, a new commander is chosen, and the new commander is removed from the list of lieutenants

## Run
Run the following to do the example in test `TestByzantine_m2_n7`
```
go build
./byzantine_main -m=2 -G=L,L,L,L,L,T,T -Oc=ATTACK
```
Run `./byzantine_main -h` to get the input arguments
```
Usage of ./byzantine_main:
  -G string
        The loyalty of the generals.Provide list in comma seperated values of "L" or "T" (default "L,L,L")
  -Oc string
        The order the commander gives: ATTACK or RETREAT (default "ATTACK")
  -m int
        The level of recursion
```
## Test
To run the tests
```
cd test
go test
```
There are examples with m=0,1,2