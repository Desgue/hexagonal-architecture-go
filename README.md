# Used Vehicle Auto-Stand Management

## Overview

I am developing a backend API in Go to practice and consolidate the concepts of the Hexagonal Architecture Pattern. This project focuses on building a production-ready Car Management App suitable for a small auto stand business, intended for use by a small group (1-2 persons).

## Objectives

The main goal is to create a fully functional car management application with the following features:

- Car Stock Management
- Store Customer Information
- Financial Dashboard for analyzing car sales costs and profits
- Admin login only, no user registration
- Single DB for all admins (personal use)

## Technologies Used

### Backend

I chose Go for its speed, readability, and features within the standard library, reducing dependency on external libraries and frameworks. The ability to compile to a single binary for different platforms simplifies deployment.

- **Server Language: [Go](https://golang.org/)**
- **Authentication: [Amazon Cognito](https://aws.amazon.com/cognito/)**

### Database

- **AWS S3:** Used for storing images and static files.
- **DynamoDB (Amazon):** A serverless database, allowing easy development locally with a provided [Docker image](https://hub.docker.com/r/amazon/dynamodb-local/). It offers pay-per-use pricing and convenient testing

### Testing

- **Go Testing Library:** Built-in testing library in Go.
- **[Testcontainer](https://github.com/testcontainers/testcontainers-go):** Facilitates Docker integration testing. 
- **[LocalStack](https://github.com/localstack/localstack):** Emulates AWS services, making it easy to test DynamoDB calls and other AWS-related functionality. 

### Frontend

Considering my backend focus, the frontend is not yet in development. I am considering the following stack:

- **[React](https://react.dev/):** Main frontend framework
- **[Shadcn Components](https://ui.shadcn.com/):** Component library
- **[Tailwind CSS](https://tailwindcss.com/):** Utility-first CSS framework
- **[Next.js](https://nextjs.org/):** React framework for building web applications 

## Project Status

In the initial stages, development occurs during free time outside my current job commitments. The project aims to balance professionalism with personal exploration and learning.

Feel free to contribute or suggest improvements as the project evolves.
