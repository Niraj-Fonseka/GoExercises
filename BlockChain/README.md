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


