## _Progress in Sprint 2_
### [Sprint 2 Demo Video](https://youtu.be/PX8B1Ce_P0U) | [Alternative Sprint 2 Demo Video Link](https://clipchamp.com/watch/mU8L6z12Q7Z)
### What issues the team planned to address + related user stories
* Add functionality in the demo page to track progress, accuracy, and speed of the user.
  * As a career typist, I want to log on to an engaging web app for typing so that I can maximize my typing speed and track my improvement over time.
* Integrate front end with the back end, adding account functionality to the web app.
  * As a typical user, I want my personal account information to be secure so that I am not vulnerable to a data breach affecting this service.
### Which ones were successfully completed
* We implemented color-changing font based on accuracy (on the "Demo" page, entering a correct character turns that corresponding character green, while entering an incorrect character turns that character red). Color feedback allows for a more engaging product.
* We have accounted for CORS permissions for an end-to-end connection that results in a successful GET request, and a POST request to register a new user from the "Register" page.
* We adjusted the database model so that IDs are a primary key, meaning no 2 users can have the same ID.
### Which ones didn't and why?
* There were plans to create a profile page but progress was halted due to some difficulties integrating the front end and back end. 
* The login functionality is half-complete. The front end and back end functions have been implemented, but there is no visible result yet as the HTTP response has not been handled yet.
## _Cypress Test and Front-End Unit Tests_
#### Cypress test in [register.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/register.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the register page and fills in the username and password fields, and finally clicks the sign up button to create a new account. The HTTP requests are checked for 200 OK responses.**
#### Cypress test in [demoPage.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/demoPage.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the demo page and fills the text area and the resulting WPM, text length, # of correct keystrokes, and # of mistakes is correct. This is repeated 2 more times with different text inputs**
#### Tests in [app.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/app.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [registration.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/registration/registration.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [demo-page.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/demo-page/demo-page.component.spec.ts)
* **Test for DemoPage's wiki.name variable**
* **Test for onTimeStop() **
* **Test for compare(char1, char2, strIndex) **

## _Back-End Unit Tests_
#### Tests in [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main_test.go)
* **TestLoadEnv**
* **TestConnect**
* **TestMigrate**
* **TestGET**
* **TestPOST**
* **TestPUT**
* **TestDELETE**
* **TestAuthenticateUser**

#### Test in [user_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/routes/user_test.go)
* **TestGenerateHashedPassword with a Helper Function**
## _Documentation_

### _Quick Start for the Database_
> The .env file contains constants for the localhost port for testing (type localhost:PORT# into a browser when program is running to see output). The other constant is the dns string for connecting a database to the server. 
**NOTE:** If you have trouble connecting to the database server, use a DSN string of this format: "postgres://username:password@hostname/databasename", where "username", "password", "hostname", and "databasename" are variables that should be replaced with our [ElephantSQL](https://www.elephantsql.com/docs/index.html) database credentials. You can create a constant for this string in your local .env file; remember to confirm that the DSN is set correctly in the [database] intialize package. Find the line "dsn := os.Getenv("YOUR_DSN_STRING")", where YOUR_DSN_STRING is the name of the string in the .env file (NOTE: the string must be in quotes, even though is is referenced as a variable).

> [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go) 1) intializes the router for the Rest API, 2) initializes the database, and 3) migrates the database (which creates rows and tables for adding data). 
> Handlers for the Rest API are located in [User.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/models/User.go). 

> JSON Formats for all of the objects associated with the database can be found in the [models](https://github.com/WasabiTech-777/SWE-2023-Spring/tree/main/src/server/models) directory. 

### _Event Handlers_

> A **new user registering** on the website requires a POST request (routes.PostUser); in the database, **usernames are labeled "uname"** and **passwords are labeled "pass"**.

> An **existing user logging in** should call the custom route (routes.AuthenticateUser). 

> Getting information about a user requires a GET request to "users/{uid}" where uid is a user's ID assigned at the time their database entry is created. Currently, you cannot get a user with their username, but this should be addressed soon.

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
> A proxy server configuration is used for front-end-to-back-end communication testing. Configuration options can be found in the proxy.conf.json file. To run the proxy server, use command "ng serve --proxy-config proxy.conf.json". If your machine is not permitted to run scripts, run Windows Powershell as Administrator, then run command "Get-ExecutionPolicy" to get the current policy level. To escalate the policy, run the command "Set-ExecutionPolicy RemoteSigned" then type "Y" to confirm. Once finished running the proxy server, it is **highly recommended** that you run the command "Set-ExecutionPolicy Restricted" to eliminate security vulnerabilities on your local machine. 

> Proxy log levels can be configured in the proxy.conf.json file; is no log level is specified, the default option is "info", but the log level in our proxy file is currently set to "debug". Other log level options are: warn, error, and silent. Please see this [guide](https://angular.io/guide/build) for more information on setting up and running a proxy server in Angular.
