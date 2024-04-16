# Running instructions

Developed using go version 1.19.1.

## Using docker:

`docker build -t transaction-processor .`

`docker run --env-file docker.env transaction-processor <recipient_email@example.com>`

Once the email is sent, you might want to check the spam folder!