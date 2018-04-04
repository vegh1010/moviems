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
    
Code Structure Design:
   1) For this project, 5 packages have been created.
      1) tools contains utility tools to support moviems.
      2) lib contains data structs and api libraries for moviems usage.
      3) Param contains a struct which is used to link common variables and initializing of values
      to different packages in Features and Handlers.
      4) Features contains object packages and functions to complete a process/task which is unique 
      to itself such as doing CRUD process for customer/contact. Both may have the same process but 
      they are handling different objects.
      5) Handlers contains handler function which is used complete a process/task and reference
      to a specific/multiple pattern routes defined. Each handler defined in Handlers could reflect
	  to Features packages.
   2) The handler packages are then defined in main where environment checking, initializing of 
   variables and routes are done and lastly start listen and serve to use the microservice.
   3) From my perspective, this structure design will allow for a more manageable and 
   readable of code.
   4) Further explanation has been included in the code. 