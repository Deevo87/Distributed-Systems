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

<h2>Lab 2 using Golang</h2>
 
 The zadanie_domowe directory contains a REST API that transfers currency (using 2 diffrencts external APIs) and displays the best and worst prices with a simple client (only html) where:
   - REST API sends resposnes in JSON
   - each call from the client is asynchronous
   - request to external APIs are multi-threded, every call has it's own thread, for communication between these calls I use channels
   - there is simple security mechanism - before making a call to the API you need to obtain the API key (localhost)([localhost](http://localhost:4000/getApiKey))
   - common http errors are handled (client is receiving appropriate http status)
   - to create this application I used the gin framework

<h3>Running the program:</h3>

I used CompileDaemon to run my application, and if you want to run my app, I recommend you using it (https://github.com/githubnemo/CompileDaemon):
1. Open 2 terminals (one for api, one for client, you can run more clients if you want)
2. Build main.go files in each directory
3. Run these commands in terminal:
  - CompileDaemon -command="./api" - for the api terminal
  - CompileDaemon -command="./client" - for the client terminal
4. Client is running on port 4000 and server on port 3000
5. To get API key go to ([localhost](http://localhost:4000/getApiKey))
6. Then go to ([localhost](http://localhost:4000)) and fill out the form


Start api and client:

![image](https://github.com/Deevo87/Distributed-Systems/assets/85259534/5c460c85-4cb7-4682-96e4-d0bf486d5e64)

Client view:

![image](https://github.com/Deevo87/Distributed-Systems/assets/85259534/811fb3e9-87a1-43a2-b0ff-fc6be9deec19)


