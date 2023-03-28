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

## Work Completed in Sprint 3 ##

### [Video Demonstration for Sprint 3](tba.com) ###

For this sprint, we focused on enabling sessions/cookies for registered Wi-key users. We also started investigating ways of tracking some user statistics, including characters hit, characters missed, and the number of articles a user has completed. We also added more backend unit tests. 

### Frontend Details ###

### Backend Details ###

#### [user.go](https://github.com/WasabiTech-777/SWE-2023-Spring/src/server/routes/user.go) ####
Functionality was added to the "AuthenticateUser" function to generate a token after a user is successfully authenticated. This token is used to create a cookie. A new function "ValidateToken" was created to check the validity of a token; upon successful validation, a user is authorized.

#### [main.go](https://github.com/WasabiTech-777/SWE-2023-Spring/src/server/main.go) ####
Two new routes were added to access the functionality in user.go: "/token" for ValidateToken, and "/users/{uname}" for finding a user by their username (in progress). 

#### [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/src/server/main_test.go) ####
Multiple unit tests were added. "TestGetUsers" and "TestGetUser" were not completed in time for Sprint 2, but are now complete and integrated into an appropriate sequence with the other unit tests for user handlers. 
Unit tests for the new handlers were also added. 

### Frontend Unit Tests ###

#### Tests in [app.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/app.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [registration.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/registration/registration.component.spec.ts)
* **Test for helloWorld (getUsers)**
#### Tests in [demo-page.component.spec.ts](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/62a5d1f2e0e04d95f3aba275a97949f16396101f/src/app/demo-page/demo-page.component.spec.ts)
* **Test for DemoPage's wiki.name variable**
* **Test for onTimeStop()**
* **Test for compare(char1, char2, strIndex)**

### Backend Unit Tests ###

#### Tests in [main_test.go](https://github.com/WasabiTech-777/SWE-2023-Spring/blob/main/src/server/main_test.go)
* **TestLoadEnv**
* **TestConnect**
* **TestMigrate**
* **TestGET**
* **TestGetUsers**  [NEW]
* **TestPOST**
* **TestGetUser**   [NEW]
* **TestPUT**
* **TestDELETE**
* **TestAuthenticateUser**

### Updated Backend Documentation ###
Please see [Backend_Documentation.md](https://github.com/WasabiTech-777/SWE-2023-Spring/Backend_Documentation.md)