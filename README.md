# Firebase Token Generator

`Firebase Token Generator` is a simple application used to [create custom tokens](https://firebase.google.com/docs/auth/admin/create-custom-tokens) for existing users in your Firebase project

## Prerequisites

- Go
  - 1.22.5 was the version used to build this application; you can download it [here](https://go.dev/)
- Firebase Project
  - It's needed to have an existing [Firebase](https://firebase.google.com/) project with users registered

## Usage

- Create a new `.env` file by following the `.env.example`
- Replace the variable value with your Firebase credential
  - Note: usually a Firebase credential is a _JSON_ file; to make it work for this app, you need to convert it to a _string_ and encode it to _base64_

---

Inside the root folder, type the following command

```
make run
```

You can generate new tokens by passing an user `phone number` or `email`

## Example

```
-> make run

Type 1 for phone number, 2 for email: 2

Type the user email: user@email.com

Token: eyJhbaaG23ciOiJwSUzI1NiICJ9.eyJptZ2ktbI6Imh0MuY29tL2dvb2dsZS5pZGVv...
----------------------------------------------
Do you want to generate another token?
Type 1 for phone number, 2 for email:
```
