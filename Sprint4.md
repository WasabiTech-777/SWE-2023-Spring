<!-- WORK TO COMPLETE:
Entire Team
    Make progress on issues uncompleted in Sprint 3, or new issues discovered during Sprint 3.
    Write test for new functionality implemented. 
     front-page readme that details requirements for running and using your application.
SUBMISSIONS:
    Submission Format: GitHub & Video Links (Use comments on submission page for multiple links)
    Narrated video presentation. Split the presentation such that each member of your team narrates a portion. Presentation should include:
    Demonstrate new functionality implemented.
    Show results of all unit tests (including those from Sprint 3).
    Finally, give an overview of your completed project as if you were pitching it to someone who has never seen it:
    Demonstrate all front-end functionality
    Detailed explanation of backend API
Sprint4.md
    Detail work you've completed in Sprint 4
    List frontend unit and Cypress tests
    List backend unit tests
    Show updated documentation for your backend API 

We will be checking individual commits. If you do not commit code, you will not receive any credit for this sprint. If you're having trouble contributing, speak with your TA sooner rather than later.
Rather than just checking whether or not you have contributed all, for this sprint we will also be comparing outputs with the rest of your team. A contribution ratio of 2:1 as compared with the top performing teammate is fine, but when we start seeing differences approaching 5:1, 10:1, etc, points will be docked. -->

## _Work Completed in Sprint 4_ ##

### [Video Demonstration for Sprint 4](https://youtu.be/nw9QVmq_TI8) ###
### [Alternate Link for Sprint 4 Video](https://clipchamp.com/watch/BSzmDGNW1P8)


### What issues the team planned to address + related user stories
* Account Security
    * As a typical user, I want my personal account information to be secure so that I am not vulnerable to a data breach affecting this service.

* Genre Choice for Typing Content
    * As a history enthusiast, I would like to type articles related to specific genres like moments in history so that I can be more engaged with the application.
### Which ones were successfully completed
* Account Security was completed for the time being; there was a small security bug that resulted in modified passwords not being hashed correctly, but that has been corrected. The tools used in our implementation also prevent XSS attacks. Additionally, tests have been added for user routes to ensure security of user data.

### Which ones didn't and why?
* Genre Choice for Typing Content required more time to implement. Working with the Wikipedia API requires a lot of time to review the documentation. 
## _Frontend Details_ ##

#### [articles-page](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/45eac36e5ac1ac549c8b33c09b5af34091ebdc37/src/app/articles-page/articles-page.component.html)
An articles page has been added that lists the articles that can by typed on the typing program.
#### [profile-page](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/45eac36e5ac1ac549c8b33c09b5af34091ebdc37/src/app/profile-page/profile-page.component.html)
The profile page has been updated so that the article count, WPM, and accuracy of the user is accurate.
#### [app-component.html](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/45eac36e5ac1ac549c8b33c09b5af34091ebdc37/src/app/app.component.html)
The top toolbar has been given functionality to show the user's username when logged in.

## _Backend Details_ ##

#### [user.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/routes/user.go) ####
A substantial security bug was fixed. Previously, users modifying their passwords would not have their new passwords properly hashed with the Go's BCrypt package. After much consideration for the best implementation to fix this issue, each password is rehashed when a user edits their details, regardless of the modificaiton of their password. There is a slightly ineffciency in hashing a password that has already been hashed for a user, but checking to see if the password was modified would result in a similar effect on performance, and could ultimately require hashing a new password anyway.  

#### [article.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/routes/article.go) [session.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/routes/session.go) ####
Routes were updated to include a full CRUD functionality for both the article, and session entity in the database. 

#### [Backend_Documentation.md](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/Backend_Documentation.md) ####
Updated changes to the back end and created an Entity Relationship (ER) diagram of the database organization.

#### [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/main_test.go) ####
The functionality of these tests in this file has been majorly improved by replacing hard-coded items with appropriate constants and variables.
11 new tests have been added for the new functionality in article.go and session.go. 

## _Cypress Tests and Frontend Unit Tests_
#### Cypress test in [register.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/register.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the register page and fills in the username and password fields, and finally clicks the sign up button to create a new account. The HTTP requests are checked for 200 OK responses.**
#### Cypress test in [demoPage.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d28ae8915ec5ed7c40697cf1399cac0636e6f18d/cypress/e2e/demoPage.spec.cy.ts)
* **Visits the home page (currently the login page), then navigates to the demo page and fills the text area and the resulting WPM, text length, # of correct keystrokes, and # of mistakes is correct. This is repeated 2 more times with different text inputs**
#### Cypress test in [login.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/456cdb1aa202ff37cd4707f5fbe604fa4c537fa0/cypress/e2e/login.spec.cy.ts)
* **Visits the home page (currently the login page), then inputs a valid username and password into the respective fields, and finally clicks the log in button to log into an account. It then navigates to the profile page via a successful login.**
#### Cypress test in [fulle2e.spec.cy.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/73cead4cbf45e9cad476f1fce75bd4a5e72d2ea7/cypress/e2e/fulle2e.spec.cy.ts) [NEW]
* **Performs a full E2E test of the application. It starts at the login page, logs in to redirect to the profile page, then navigates to the articles page, and then finally goes to the typing page and completes an article.**

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
* **Test for getUserInfo, which gets a user's data by username and returns their article count, WPM, and accuracy** [NEW]
## _Backend Unit Tests_ ##

#### Tests in [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/174bdd6c8f57efa12affba46334eadbfffbdb2a9/src/server/main_test.go)
* **TestLoadEnv**
* **TestConnect**
* **TestMigrate**
* **TestGET**
* **TestGetUsers**  
* **TestPOST**  [EDITED]
* **TestGetUser**   [EDITED]
* **TestGetUserFromName** 
* **TestPUT** [EDITED]
* **TestDELETE** [EDITED]
* **TestAuthenticateUser** 
* **TestValidateToken**
* **TestPostSession** [NEW]
* **TestGetSession**[NEW]
* **TestGetSessionFromUser** [NEW]
* **TestGetSessionFromArticle** [NEW]
* **TestPutSession** [NEW]
* **TestDeleteSession** [NEW]
* **TestPostArticle** [NEW]
* **TestGetArticle** [NEW]
* **TestGetBody** [NEW]
* **TestPutArticle** [NEW]
* **TestDeleteArticle** [NEW]

## _README with Requirements for Users Running Application_ ##
Please see [README.md](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d56f4462b4638e93bbfdde70142798d7d0f2109d/README.md)

## _Updated Backend Documentation_ ##
Please see [Backend_Documentation.md](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/d56f4462b4638e93bbfdde70142798d7d0f2109d/Backend_Documentation.md)
