# Savannah-E-Shop Application
This is a Go web application for the Savannah e-Shop project, built using the standard Go and integrating with the Keycloak identity provider for authentication. The application also uses Africa's Talking for sending SMS notifications

## Features
- Authentication with Keycloak.
- Create and manage customer records.
- Create and manage order records.
- Demo callback endpoint for OAuth2 authentication.
- Send SMS notifications using Africa's Talking.

## Installation
To run this application locally, follow this step:
1. Clone the repository:
    - < git clone git@github.com:Tito-74/Savannah-e-shop.git >

2. Create a savannah.env file in the root directory and set the following environment variables:
    * KEYCLOAK_CONFIG_URL=<Keycloak Configuration URL>
    * KEYCLOAK_CLIENTID=<Your Keycloak Client ID>
    * KEYCLOAK_CLIENT_SECRET=<Your Keycloak Client Secret>
    * KEYCLOAK_REDIRECT_URL=<Your Redirect URL>
    * KEYCLOAK_STATE=<Your State>
    * AFRICASTALKING_API_KEY=<Your Africa's Talking API Key>
    * AFRICASTALKING_USERNAME=<Your Africa's Talking Username>
3. Install dependencies:
  - go mod tidy

4. Pull and run keycloak using:
  - docker run -e KEYCLOAK_USER=admin -e KEYCLOAK_PASSWORD=admin -p 8080:8080 jboss/keycloak

5. configure keycloak 
  - [https://www.keycloak.org/docs/latest/securing_apps/]

6. Run the application:
   - go run main.go



## Routes
- /token: Endpoint for handling access tokens and authentication with Keycloak.
- /customer: Create and manage customer records (HTTP POST).
- /order: Create and manage order records (HTTP POST).
- /demo/callback: Demo callback endpoint for OAuth2 authentication.

## Deployment to AWS
 This application has been deployed to AWS Elastic Beanstalk, a Platform as a Service (PaaS) offering that simplifies the deployment and management of web applications. The deployment includes the necessary configuration for running the application in the AWS cloud environment.


## Dependencies

  * OpenID Connect (OIDC): Library for OpenID Connect (OIDC) authentication.
  * Africa's Talking Go SDK: SDK for Africa's Talking SMS and other services.

## Authors
 [Langat Tito Kipkirui]




## License

[MIT](https://choosealicense.com/licenses/mit/)
