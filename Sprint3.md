<!-- WORK TO COMPLETE:
Entire Team

Make progress on issues uncompleted in Sprint 2, or new issues discovered during Sprint 2.
Write test for new functionality implemented. 
SUBMISSIONS:
Submission Format: GitHub & Video Links (Use comments on submission page for multiple links)
Narrated video presentation. Split the presentation such that each member of your team narrates a portion. Presentation should include:
Demonstrate new functionality implemented.
Show results of all unit tests (including those from Sprint 2).
Sprint3.md
Detail work you've completed in Sprint 3
List frontend unit tests
List backend unit tests
Show updated documentation for your backend API 
We will be checking individual commits. If you do not commit code, you will not receive any credit for this sprint. If you're having trouble contributing, speak with your TA sooner rather than later. -->

## _Work Completed in Sprint 3_ ##

### [Video Demonstration for Sprint 3](tba.com) ###

For this sprint, we focused on enabling sessions/cookies for registered Wi-key users. We also started investigating ways of tracking some user statistics, including characters hit, characters missed, and the number of articles a user has completed. We also added more backend unit tests. 
### What issues the team planned to address + related user stories
* Add functionality in the demo page to track progress, accuracy, and speed of the user.
  * As a career typist, I want to log on to an engaging web app for typing so that I can maximize my typing speed and track my improvement over time.
* Integrate front end with the back end, adding account functionality to the web app.
  * As a typical user, I want my personal account information to be secure so that I am not vulnerable to a data breach affecting this service.
### Which ones were successfully completed
* We have made great progress on the login system. Users are now able to login and a token is saved which allows the server and client to retrieve user information.
### Which ones didn't and why?
* 
## _Frontend Details_ ##

####  [log-in.component.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/log-in/log-in.component.ts)
The user can now log in by pressing the "Log In" button, which runs AuthenticateUser with the username and password fields as input.

####  [profile-page.component.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/profile-page/profile-page.component.ts)
A new page has been created which shows information about the user that is currently logged in. Currently it displays the username when logged in, and "Guest" when not.
## _Backend Details_ ##

#### [user.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/routes/user.go) ####
Functionality was added to the "AuthenticateUser" function to generate a token after a user is successfully authenticated. This token is used to create a cookie. A new function "ValidateToken" was created to check the validity of a token; upon successful validation, a user is authorized.

#### [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/main.go) ####
Two new routes were added to access the functionality in user.go: "/token" for ValidateToken, and "/uname/{uname}" for GetUserFromName. 

#### [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/main_test.go) ####
Multiple unit tests were added. "TestGetUsers" and "TestGetUser" were not completed in time for Sprint 2, but are now complete and integrated into an appropriate sequence with the other unit tests for user handlers. 
Unit tests for the new handlers were also added; these tests include "TestGetUserFromName" and "TestValidateToken"

## _Cypress Tests and Frontend Unit Tests_
#### Cypress test in [register.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/register.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the register page and fills in the username and password fields, and finally clicks the sign up button to create a new account. The HTTP requests are checked for 200 OK responses.**
#### Cypress test in [demoPage.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/demoPage.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the demo page and fills the text area and the resulting WPM, text length, # of correct keystrokes, and # of mistakes is correct. This is repeated 2 more times with different text inputs**
#### Cypress test in [login.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/login.spec.cy.ts)
* **Visits the home page (currently the login page), then inputs a valid username and password into the respective fields, and finally clicks the log in button to log into an account. It then navigates to the profile page via a successful login.**

#### Tests in [app.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/app.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [registration.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/registration/registration.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [demo-page.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/demo-page/demo-page.component.spec.ts)
* **Test for DemoPage's wiki.name variable**
* **Test for onTimeStop()**
* **Test for compare(char1, char2, strIndex)**
#### Tests in [log-in.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/log-in/log-in.component.spec.ts)
* **Test for login (AuthenticateUser)**
#### Tests in [account.service.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/5159d7ebe84904dd363ec3932b869f453e1244c5/src/app/account.service.spec.ts)
* **Test for validate (ValidateToken)**
* **Test for decodeToken (decodes the JWT token so user info can be retrieved)**
## _Backend Unit Tests_ ##

#### Tests in [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/main_test.go)
* **TestLoadEnv**
* **TestConnect**
* **TestMigrate**
* **TestGET**
* **TestGetUsers**  [NEW]
* **TestPOST**
* **TestGetUser**   [NEW]
* **TestGetUserFromName** [NEW]
* **TestPUT**
* **TestDELETE**
* **TestAuthenticateUser** [EXPANDED]
* **TestValidateToken** [NEW]

## _Updated Backend Documentation_ ##
Please see [Backend_Documentation.md](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/51e2cc30a810aa0da4dd435826402799daeee1ba/Backend_Documentation.md)
