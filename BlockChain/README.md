# Blockchain in Go 



## Proof of work 

Do hard work to gain something 
Adding blocks to the chain must me hard 
in blockchain this hard work is to find a hash for a block that meets certian requirements. This hash is the proof 
Proof of work algos must meet a requirement. Doing the work must be hard but veryving the proof should be easy
Checking for proof should be handled by someone else


#### Hashing 

Hasing is the process of obtaining hash for a specified data 
Hash function takes data of some size and produces a fixed hash

- original data cannot be restored from a hash. Hashing is not encrypting 
- certian data can have only one hash and the hash is unique 
- changing even a byte of data should change the hash


###### In Bitcoin 

HashCash algo steps

- Take some publicly known data (in case of email, it’s receiver’s email address; in case of Bitcoin, it’s block headers.
- Add a counter to it. The counter starts at 0.
- Get a hash of the data + counter combination.
- Check that the hash meets certain requirements.
- If it does, you’re done.
- If it doesn’t, increase the counter and repeat the steps 3 and 4.

Requirements 

- First 20 bits of hash must be 0s 
- In bitcoin the requirement is adjusted from time to time. A block must be generated every 10 minutes.

- The target is the upper boundry of a range. if the hash is lower than the target its a valid. 


#### Database 

Using BoltDB 

##### Database Structure 

- Blocks = Stores metadata describing all the blocks in the chain 
- chainstate = stores the state of the chain. Which is all currently unspent transaction outputs and some metadata

in bitcoin the blocks are stored as sperate files on the disk. ( for performence reasons )


In blocks the key --> value pairs 

```
'b' + 32-byte block hash -> block index record
'f' + 4-byte file number -> file information record
'l' -> 4-byte file number: the last block file number used
'R' -> 1-byte boolean: whether we're in the process of reindexing
'F' + 1-byte flag name length + flag name string -> 1 byte boolean: various flags that can be on or off
't' + 32-byte transaction hash -> transaction index record
```

In chainstate the key --> value pairs 

```
'c' + 32-byte transaction hash -> unspent transaction output record for that transaction
'B' -> 32-byte block hash: the block hash up to which the database represents the unspent transaction outputs
```