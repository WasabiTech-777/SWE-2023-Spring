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

### _Quick Start for the Server_
> The .env file contains constants for the localhost port for testing (type localhost:PORT# into a browser when program is running to see output). The other constant is the dns string for connecting a database to the server. 

> main.go intializes the router for the Rest API, initializes the database, and then migrates the database (which creates rows and tables for adding data). 
> Handlers for the Rest API are located in routes/user.go. 

> JSON Formats for all of the objects associated with the database can be found in the "models" directory. 

> **Registering** new users is a POST HTTP request (routes.PostUser); in the database, usernames are labeled "uname" and passwords are labaled "pass".

> Users **logging in** should call (routes.AuthenticateUser).

