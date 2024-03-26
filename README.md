<h1>Distibuted Systems</h1>

This repository contains code from distributed systems labs, a class I have at AGH.

<h2>Lab 1 using Golang</h2>
 
 The zadanie_domowe directory contains the simple chat application where:
   - clients connect to the server via TCP protocol
   - the server accepts messages from each client and distributes them to the others along with the id of the sending client
   - the server is multi-threaded, each connection from a client have it's own thread
   - the server and each client open an additional UDP chanell (on the same port as TCP)
   - after typing the "U" command client sends a message (ASCII art) to the server via UDP, then the server distributes it to the oters (also using UDP)
     
 ![image](https://github.com/Deevo87/Distributed-Systems/assets/85259534/4f08251f-04f4-4e37-8846-5a11c8757857)
 ![image](https://github.com/Deevo87/Distributed-Systems/assets/85259534/763a5d96-8fb2-4ec7-9165-9a76294a396a)



