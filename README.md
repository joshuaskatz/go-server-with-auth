# Example Server (Golang)

This repo contains a very small and basic authentication module. The module contains the two router groups, both protected (`/api/admin`) and unprotected (`/api`).

For the sake of example, the `/albums` routes are included to demonstrate the authentication middleware using JWTs.

---

### The authentication routes include the following:

- `/register` - for registering a user
- `/verify/:id` - where `id` is the JWT sent to the user via the `email verification` middleware. This alters the `VERIFIED` column in the `users` table to `true`, allowing users to finally login.
- `/login` - logs in the user and returns an access token that the user can user to access protected routes.

_Note: There are also various authentication related middleware to check the validity of a users JWT, check the verification status of a user, and to send a verification email._

---

### Other features of the module include:

- `SQL` file parsing for purposes of querying and syncing/creating database tables.
- A function created to parse `.env` files for use throughout the module. Required environment variables include:

  ```
  JWT_SIGNING_KEY=

  DB_HOST=
  DB_PORT=
  DB_USER=
  DB_PASSWORD=
  DB_NAME=

  SERVER_URL=

  EMAIL_USER=
  EMAIL_PASSWORD=
  ```

* Use of the `gomail` module for `email verification` middleware.

---

## _NOTE: This is not meant for use in production, but is rather just an example module build with Golang._

### A few things I would add to make this more usable in a production context:

- As of now, logging in only returns a single short lived access token. I would either create a two step login process where the user exchanges a short lived token for a long lived token, or I would implement a refresh token to be used alongside the access token.
- Additionally, a function for properly handling database migrations would be very useful.
