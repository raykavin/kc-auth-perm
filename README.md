# Keycloak Authentication and Permission Example

This project exemplifies the integration between Vue+Vite (Frontend) and Golang (Backend), using the OIDC library for authentication and authorization in SSO providers. The frontend is responsible for connecting to the SSO, collecting user information and obtaining the access token. When the frontend requests resources from one or more servers, the access token is sent to the backend, which verifies the validity of the token with the authorization provider (SSO). If the token signature is valid, the backend confirms that the user has the necessary permissions to access the API in the requested path and verb.

## Features

- Integration with Keycloak for user authentication.
- Management of permissions based on roles.
- User interface for login and viewing restricted data.

## Technologies Used

- **Golang**: The backend of the project.
- **Vue+Vite**: The frontend of the project.
