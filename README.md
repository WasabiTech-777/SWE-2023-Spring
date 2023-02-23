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

### _Quick Start for the Database_
> The .env file contains constants for the localhost port for testing (type localhost:PORT# into a browser when program is running to see output). The other constant is the dns string for connecting a database to the server. 
**NOTE:** If you have trouble connecting to the database server, use a DSN string of this format: "postgres://username:password@hostname/databasename", where "username", "password", "hostname", and "databasename" are variables that should be replaced with our [ElephantSQL](https://www.elephantsql.com/docs/index.html) database credentials. You can create a constant for this string in your local .env file; remember to confirm that the DNS is set correctly in the [database] intialize package. Find the line "dsn := os.Getenv("YOUR_DSN_STRING")", where YOUR_DNS_STRING is the name of the string in the .env file (NOTE: the string must be in quotes, even though is is referenced as a variable).

> [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main.go) 1) intializes the router for the Rest API, 2) initializes the database, and 3) migrates the database (which creates rows and tables for adding data). 
> Handlers for the Rest API are located in [User.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/models/User.go). 

> JSON Formats for all of the objects associated with the database can be found in the [models](https://github.com/WasabiTech-777/SWE-2023-Spring/tree/main/src/server/models) directory. 

#### Event Handlers

> A **new user registering** on the website requires a POST HTTP request (routes.PostUser); in the database, **usernames are labeled "uname"** and **passwords are labeled "pass"**.

> An **existing user logging in** should call the custom route (routes.AuthenticateUser).

