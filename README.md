# moviems
A microservice to get movie information.

Tools Used:

   1) go version 1.8.3
   2) bash
   2) PostMan
    
Assumptions:
    
   1) Microservice's api call requires header 'x-access-token'.
   2) The microservice will hold an access token for validation purpose.
   
Step to Step:
   1) 3 environment variables required which are URL, ACCESS_TOKEN and PORT.
      1) URL represents the external url for getting the movie's information.
      2) ACCESS_TOKEN represents access token needed to call the external url api.
      3) PORT represents the port used by moviems.
   2) Once run locally, you can access either through a web browser or PostMan by typing 
   127.0.0.1:{PORT}
   3) For more information, you call this api: 127.0.0.1:{PORT}/help
      1) This will return a list of api you can use and a short description on what 
      each of them does.
   4) Here's a bash script if you would like to run it using bash.
   
    #!/usr/bin/env bash
    echo "Launching MovieMS in Office Dev mode"
    
    export URL=<value>
    export ACCESS_TOKEN=<value>
    export PORT=<value>
    
    go run *.go