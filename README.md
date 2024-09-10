# Running instructions

Developed using go version 1.19.1.

## Using docker:

`docker build -t my-golang-app .`

`docker run --env-file docker.env my-golang-app <recipient_email@example.com>`

Once the email is sent, you might want to check the spam folder!