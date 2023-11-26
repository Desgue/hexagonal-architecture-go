# hexagonal-architecture-go-example

#Used Vehicle auto-stand management#
Making this project to pratice building a backend api in Go and further consolidate the concpets of the Hexagonal Architecture Pattern when developing software

Project is on the beggining, as i only work here when i am not busy in my current Job. 

## Objectives ##

My main goal with this project is to build a fully functional Car Management APP  production ready for a small Auto Stand business, to be used by a small (think 1-2 persons) group of people. 

### Features ###
1. Car Stock Management
2. Store Customer Information
3. Financial Dashboard for easily analyse car sales costs and profits
4. Admin log in only, no user registration.
5. One DB for all Admins as this is mainly for my personal use


## Technologies Used ## 

#### Backend ####
I chose Go because its the language I currently work with, but not only. Go is blazingly fast, easy to read and to write and has many features needed for building a backend within the standard library, making less dependent of outside librarys and frameworks. This makes developing in Go fast and straightfoward. Also Binaries! Go can compile to a single binary either for Mac, Linux or Windows. Making it easy to run and deploy no matter the host environment. 

#### Database ####
For images and static files, I am using AWS S3, its cheap, easy and safeway to store images (mostly my needs in this app), but also any static file. 

For the database I am using DynamoDB from Amazon. Its a serverless database, which means i dont need to host and config a db server by myself, i have an Amazon provided DynamoDB Docker image to develop my application locally with no need to worry about unwanted cloud costs and the Aws Sdk for go is easy to implement. If i decide to deploy this application with DynamoDB I can do it easily. And because of its serverless nature, i do not have to worry about server maintence costs, as Dynamo is a paid per use service (think 1 million read and writes in the free tire ;).

Also testing, testing with DynamoDB is a breeze.

#### Testing ####
Go has a built in testing library to make it easy to Unit test your applications, being one more reason to use go. So the only particular technology i use to test is the Testcontainer library and the LocalStack Cloud Emulator. 
The Testcontainer library makes it easy to work with docker from within your Go code. You can pull images and lift containers with a few lines of code, making integration testing super simple. LocalStack has an AWS container that emulates a bunch of the Amazon Services, I can use it to test my DynamoDB calls as much i can use it to test the S3 API or even some lambda functions. I super recomend to anyone interested in working in a cloud native envoirement. 

#### Frontend ####
Well mostly being a backend developer by choice (centering divs is painful), I still have to make an user interface or i wont have any users for my app. For that i chose the infamous React. Mainly because its a mainstream freamework(library? who knows). That means, plenty of resources and pre made components I can re use. I am thinking ChakraUI for the components as they come already estilized, but its a decision to be made when the project evolves and i will update this document accordingly. 




