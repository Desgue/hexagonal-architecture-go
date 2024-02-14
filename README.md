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

## System Design Choices

### Hexagonal Architecture

I chose the Hexagonal Architecture, often known as Ports and Adapters, for several reasons:

1. **Separation of Concerns:** Separation between the core business logic and external concerns that Hexagonal Architecture provides, making the system more modular and maintainable.

2. **Testability:** This architecture facilitates unit testing of the core business logic without the need for external dependencies, allowing me to create more reliable and maintainable tests.

3. **Adaptability:** It makes easier for me to adapt to changes in external components (such as databases or APIs) without affecting the core application logic.

4. **Scalability:** Even in a small project, this architecture lays the groundwork for future scalability, ensuring a smoother transition as the project grows.

5. **Maintainability:** With a well-structured architecture, maintaining and evolving the system becomes more manageable, minimizing the risk of introducing bugs during updates.

### Test-Driven Development (TDD)

I follow the Test-Driven Development (TDD) methodology, which involves writing tests before the actual code. This approach offers several advantages:

1. **Early Detection of Issues:** Writing tests before implementation helps me catch potential issues early in the development process, ensuring a more robust codebase.

2. **Improved Code Quality:** TDD encourages me to write modular and testable code, leading to better code quality and maintainability.

3. **Refactoring Confidence:** With a comprehensive test suite, I can confidently refactor code knowing that existing functionality remains intact if the tests pass.

4. **Time Saving:** While it may seem counterintuitive, TDD often saves time in the long run. The upfront investment in writing tests pays off during refactoring, reducing the likelihood of introducing bugs and streamlining the debugging process.

5. **Documentation:** Test cases serve as living documentation, providing insights into the expected behavior of various components. This is especially valuable when working on a modular architecture like Hexagonal.

By combining Hexagonal Architecture with TDD, I aim to create a system that is flexible and maintainable and also thoroughly tested, ensuring a reliable and adaptable basis for future development.


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

The frontend is not yet in development but I am considering the following stack:

- **[React](https://react.dev/):** Main frontend framework
- **[Shadcn Components](https://ui.shadcn.com/):** Component library
- **[Tailwind CSS](https://tailwindcss.com/):** Utility-first CSS framework
- **[Next.js](https://nextjs.org/):** React framework for building web applications 

## Project Status

In the initial stages, development occurs during free time outside my current job commitments. The project aims to balance professionalism with personal exploration and learning.

Feel free to contribute or suggest improvements as the project evolves.
