# Sprint 1

## User stories

As an avid social media user, I want to message other users and compare my progress to theirs so that I can encourage myself.

As a career typist, I want to log on to an engaging web app for typing so that I can maximize my typing speed and track my improvement over time.

As a keyboarding enthusiast, I want to track my improvement over time on different keyboards, so that I can determine which keyboards maximize my typing speed and accuracy.

As a history enthusiast I would like to type articles related to specific genres like moments in history so that I can be more engaged with the application.

As a user new to computers, I would like to try out the web app before committing to creating an account so I do not have to waste time creating a login for a web app I may not use again.

As a visually disabled person, I would like to use a typing web app that works seamlessly with the text-to-speech apps I already use to make typing practice easier.

As a user with limited time on their hands, I would like to have control on how long typing sessions are on the app so that I can practice typing on my own schedule.

As a person with eclectic interests, I would like to only write the introduction section of the articles so that I can learn information from a broad range of topics while I practice typing. 

As a beginner typist, I would like to see a UI keyboard with live feedback for which buttons I pressed so that I do not have to look at the keyboard while I practice typing so that I can develop my typing muscle memory.

As a statistics enthusiast, I want to view the average typing speed of all users so that I can see the distribution of typing speeds.

As a statistics enthusiast, I want to be able to export my performance data locally so I can further analyze the data and make predictions about my future improvement.

As a typical user, I want my personal account information to be secure so that I am not vulnerable to a data breach affecting this service.

As a typical user I want to be able to change my account information such as my username so that I have control over my presence in the application.

## What issues your team planned to address

We planned to address the:
- Adding User Keyboard Input Feature
- Creating a Login Service
- Accessing the demo page as a guest
- Editing Account Information
- Account Security

## Which ones were successfully completed
We were successful, in some form, with:
### Adding User Keyboard Input Feature
- We implemented a simple-keyboard UI
- Users can now see what keys they are typing on the screen or type using the on-screen keyboard
### Creating a Login Service
- A new user's username and password can be saved for persistent future use
- UI for user login was created, functionality not yet implemented
### Accessing the demo page as a guest
- button created that leads you to the demo page
### Editing Account Information
- After creating an account, users can change details such as their username; these changes are saved to the database for persistent use.
### Account Security
- No plaintext passwords are saved in the database whatsoever. Passwords entered by users are immediately encrypted with the bcrypt hash, an industry-standard hash with a built-in library in Golang. 
- When signing in, users are first checked for existence in the database. If the user exists, their identity is authenticated by comparing the bcrypt hash of the password they provided with the stored bcrypt hash of the user in the database with the corresponding username. 
- BONUS: Using GORM for database design and correct inputs into GORM database queries also keeps our application's database safe from SQL injection attacks.

## Which ones didn't and why?
### Editing Account Information
- There is more information we would like to make available for editing, such as passwords. This was not completed because it was not concieved of until the very end of the Sprint. 
### General Observations 
- We accomplished precisely the issues we set out to accomplish, but working on issues sometimes reveals other standard features that should be implemented even if they do not fit into a user story. 
## Demo Video Links
- Front end: https://www.youtube.com/watch?v=MmqSiXthUYE
