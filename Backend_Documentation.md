## Documentation

### Quick Start for the Database
> The .env file contains constants for the localhost port for testing (type localhost:PORT# into a browser when program is running to see output). The other constant is the dns string for connecting a database to the server. 
**NOTE:** If you have trouble connecting to the database server, use a DSN string of this format: "postgres://username:password@hostname/databasename", where "username", "password", "hostname", and "databasename" are variables that should be replaced with our [ElephantSQL](https://www.elephantsql.com/docs/index.html) database credentials. You can create a constant for this string in your local .env file; remember to confirm that the DSN is set correctly in the [database] intialize package. Find the line "dsn := os.Getenv("YOUR_DSN_STRING")", where YOUR_DSN_STRING is the name of the string in the .env file (NOTE: the string must be in quotes, even though is is referenced as a variable).
> [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go) 1) intializes the router for the Rest API, 2) initializes the database, and 3) migrates the database (which creates rows and tables for adding data). 
> Handlers for the Rest API are located in [User.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/models/User.go). 

> JSON Formats for all of the objects associated with the database can be found in the [models](https://github.com/WasabiTech-777/SWE-2023-Spring/tree/main/src/server/models) directory. 

### ER DIAGRAM ###
> Here we show the proposed back end design of the database with an entity relationship diagram. 
> 
![image(2)](https://user-images.githubusercontent.com/87680674/232930504-d2a6dbc1-922f-4810-acd3-0b0e132c6e9c.png)

### Event Handlers ###

#### USERS ####

* A **new user registering** on the website requires a POST request (routes.PostUser); in the database, **usernames are labeled "uname"** and **passwords are labeled "pass"**.

* An **existing user logging in** should call the custom route (routes.AuthenticateUser). Once authenticated, a user is authorized for a 24-hour period. The user's cookie for the new session is set on the response to the client. 

* **Getting information about a user from their ID** requires a GET request to "users/{uid}" where uid is a user's ID assigned at the time their database entry is created.

* **Getting information about a user from their username** is also possible; this requires a GET request to "uname/{uname}" where uname is the desired user's name. NOTE: this path starts with "uname" and not "users" to avoid undetermined behavior.

* **Editing a user** requires a PUT request to "/users/{uid}". 
    * To edit a user's username, the number of articles they have completed, the number of characters hit, or the number of characters missed, the json format of the request should be used:
    >{
        "uname": "name", //string /*NOTE: if querying the database for a particular name, use "name = ?"
        "pass": "password",  //string
        "articles": 1, //uint
        "charhit": 104, //uint
        "charmiss": 10 //uint
    }
 * To **edit the articles, charhit, or charmiss fields** correctly, the new data must be added to any previously existing data. The previously existing data could be accessed through the Get user request at "users/{uid}" above.

* **Deleting a user** requires a DELETE request to "/users/{uid}".

* **Getting all users** requires a GET request to "/users".

* **Checking if a user is authorized** requires a POST request to /token. 
   * The request body should be in this format:
     > {
      "token" : "long_cookie_string_in_header_of_/login_response"
       }

    * The response body returned will be in this format:
       > {
            "uname" : "user's_name",
            "exp: 1230983 //randomly generated uid
        }

#### SESSIONS ####
 * **Creating a new session:** POST request to "/session/"
 * **Deleting a user:** DELETE request to "/session/{sid}"
 * **Editing a session:** PUT request to "/session/{sid}"
 * **Getting a session:** GET request to "/session/{sid}"
 * **Getting a session from a user:** GET request to "/session/user/{uid}"
 * **Getting a session from an article:** GET request to "/session/article/{aid}"
 * **JSON FORMAT:**
 >	{
    "SessionID" : 1 //uint
	"ArticleID" : 1 //uint
	"UserID"   : 1 //uint
	"CharHit"   : 20 //uint
	"CharMiss"  : 2 //uint
	"Time"     : "2023-04-19 17:40:48.521" //string
 }
#### ARTICLES ####
 * **Creating an article:** POST request to "/article/"
 * **Deleting an article:** DELETE request to "/article/{aid}"
 * **Editing an article:** PUT request to "/article/{aid}"
 * **Getting an article:** GET request to "/article/{aid}"
 * **Getting an the body of an article:** "/article/body/{aid}"
 * **JSON FORMAT:**
 > 	{
    "ID"    : 1 //uint
	"Url"    : "https://wikipedia.com/cat" //string
	"Length" : 200 //uint
}

#### ACCESS CONTROL ####
> **Header Access Control** is located in [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go). 

### _Quick Start for Running the Project End-to-End_

> All of the necessary commands for running the project can be found in serve.sh. To run the file, use the command **"./serve.sh"** There is a know issue with running ./serve.sh while the server is already running. If you encounter the issue, the quickest fix is to type the commands found in ./serve.sh, with each command in a separate terminal window (or your desired command line interface).

> **Cypress Testing**: To run cypress tests, the server must already be running! The command for running cypress is **"npm run cypress:open"**.

### _Troubleshooting_
> If you are getting unexplained CORS errors, or any errors relating to ports already in use, use these commands to kill the process so you can restart the server:
> * On Windows: 
>   * **netstat -ano | findstr : port number**      //to get PID of the process running on the in-use-port
>   * **taskkill /PID typeyourPIDhere /F**
> 
> * On Mac:
>   * **sudo lsof -i :<port_number>**          //to get PID of the process running on the in-use-port
>   * **sudo kill <PID>**
>   * **sudo kill -9 <PID>**                   //if previous command did not successfully terminate the process, this will forcefully terminate the process


#### _Proxy Testing_
> NOTE: This has been deprecated, no proxy server is necessary for testing/demoing. The information about getting permission to run scripts may still be useful.
> A proxy server configuration is used for front-end-to-back-end communication testing. Configuration options can be found in the proxy.conf.json file. To run the proxy server, use command "ng serve --proxy-config proxy.conf.json". If your machine is not permitted to run scripts, run Windows Powershell as Administrator, then run command "Get-ExecutionPolicy" to get the current policy level. To escalate the policy, run the command "Set-ExecutionPolicy RemoteSigned" then type "Y" to confirm. Once finished running the proxy server, it is **highly recommended** that you run the command "Set-ExecutionPolicy Restricted" to eliminate security vulnerabilities on your local machine. 

> Proxy log levels can be configured in the proxy.conf.json file; is no log level is specified, the default option is "info", but the log level in our proxy file is currently set to "debug". Other log level options are: warn, error, and silent. Please see this [guide](https://angular.io/guide/build) for more information on setting up and running a proxy server in Angular.
