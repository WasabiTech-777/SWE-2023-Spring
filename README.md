# Intro. to Software Eng. Team #35
Members: Saviely B.<sup>1</sup>, Kohki T.<sup>1</sup>, Marina T.<sup>2</sup>, Jonathan W.<sup>2</sup>
1. Front-end developer
2. Back-end developer

## ⌨️ Wi-Key ⌨️ An Engaging and Educational Typing Practice App

### _Problem_
>The need to make typing practice more engaging and accessible

### _Current services and limitations_
>There are many websites to help people improve their typing speed and accuracy. keybr.com is a very popular website that tracks user performance over time. However, it is not engaging as the example text has no context. typelit.io solves this issue by having users type from public-domain classic novels, but with these long-form texts, it can be difficult for users to form a habit of improving their typing.

### _Solution_
>A typing website/app that engages users by using articles from simple.wikipedia.com, a version of Wikipedia that uses simple English. This way, users can be engaged by learning about new concepts with short articles while also practicing typing. The application will store user metrics like pages read, friends, typing performance, and interests.

## Documentation

### Requirements for running this application
> Angular supports most recent browsers. This includes the following specific versions: Chrome, Firefox, Safari, iOS, Internet Explorer versions 9-11, and Edge

### Using this application

### _Quick Start for the Database_
> The .env file contains constants for the localhost port for testing (type localhost:PORT# into a browser when program is running to see output). The other constant is the dns string for connecting a database to the server. 
**NOTE:** If you have trouble connecting to the database server, use a DSN string of this format: "postgres://username:password@hostname/databasename", where "username", "password", "hostname", and "databasename" are variables that should be replaced with our [ElephantSQL](https://www.elephantsql.com/docs/index.html) database credentials. You can create a constant for this string in your local .env file; remember to confirm that the DSN is set correctly in the [database] intialize package. Find the line "dsn := os.Getenv("YOUR_DSN_STRING")", where YOUR_DSN_STRING is the name of the string in the .env file (NOTE: the string must be in quotes, even though is is referenced as a variable).

> [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go) 1) intializes the router for the Rest API, 2) initializes the database, and 3) migrates the database (which creates rows and tables for adding data). 
> Handlers for the Rest API are located in [User.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/models/User.go). 

> JSON Formats for all of the objects associated with the database can be found in the [models](https://github.com/WasabiTech-777/SWE-2023-Spring/tree/main/src/server/models) directory. 

### _Event Handlers_

> A **new user registering** on the website requires a POST request on the clientside on the route "/users", **usernames are labeled "uname"** and **passwords are labeled "pass"**.

> An **existing user logging in** should submit a POST request on the route "/login" with the username and password. If the user credentials are valid, the server will return a cookie to authenticate the user in the future.  

> Tokens can be authenticated submitting a POST request on the "/token" route with the JWT token. The writer returns the usernanme of the user if the request is the token is sucessfully validated. 

> Getting information about a user requires a GET request to "/users/{uid}" where uid is a user's ID assigned at the time their database entry is created. Currently, you cannot get a user with their username, but this should be addressed soon.

> Editing a user requires a PUT request to "/users/{uid}".

> Deleting a user requires a DELETE request to "/users/{uid}".

> Getting all users requires a GET request to "/users".

> Header Access Control is located in [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go). 

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
> A proxy server configuration is used for front-end-to-back-end communication testing. Configuration options can be found in the proxy.conf.json file. To run the proxy server, use command "ng serve --proxy-config proxy.conf.json". If your machine is not permitted to run scripts, run Windows Powershell as Administrator, run command "Get-ExecutionPolicy", then type "Y" to confirm. Once finished running the proxy server, it is **highly recommended** that you run the command "Set-ExecutionPolicy Restricted" for eliminate security vulnerabilities on your local machine. 

> Proxy log levels can be configured in the proxy.conf.json file; is no log level is specified, the default option is "info", but the log level in our proxy file is currently set to "debug". Other log level options are: warn, error, and silent. Please see this [guide](https://angular.io/guide/build) for more information on setting up and running a proxy server in Angular.
