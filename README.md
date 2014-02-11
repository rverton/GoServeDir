GoServeDir
========

Serves the a static folder. Usage:
    
    GoServeDir  // Serves current folder at port 8080
    GoServeDir -port=8081 // Servers current folder at port 8081
    GoServeDir -port=8081 -path=/ //Servers / at 8081

Install:

    go get github.com/rverton/GoServeDir
    go install github.com/rverton/GoServeDir
  
Then run it via `GoServeDir`
