# Gram
Go RESTful authentication micro-service

# Purpose
Gram is a standalone authentication server that allows for the creation and management of users across a multitude of services, and creates signed JSON web tokens

# Instructions
1. Add GRAM_TOKEN_SECRET="<ANY RANDOM STRING>" to your environment
2. Add GRAM_ENCRYPTION_KEY="<ANY 16-BYTE STRING>" to your environment
3. Create a postgresql database (keep track of your credentials)
4. Open the root directory of the project in a terminal window
5. Use the command `$ go build` to build the project
6. Use the command `$ ./gram` to start the program
7. After following any indicated instructions, the server will start, and API requests can be made

# Command Line Arguments
--port number: Use a specific port instead of the default
--service name: Specify a specific service name you would like to use instead of the default. This allows for the server to manage user data for multiple services simultaneously
