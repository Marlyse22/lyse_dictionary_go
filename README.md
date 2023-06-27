# Dictionary Web API

This project implements a simple web API for a dictionary application. It allows users to add, retrieve, update, and delete words and their definitions.

## Features

- **List Words:** Retrieve a list of all words in the dictionary.
- **Get Definition:** Retrieve the definition of a specific word.
- **Add Word:** Add a new word and its definition to the dictionary.
- **Delete Word:** Delete a word and its definition from the dictionary.
- **Update Word:** Update the definition of an existing word in the dictionary.

## Technologies Used

- Go programming language
- Gin web framework
- FlashDB key-value store

## Setup Instructions

1. Clone the repository.
2. Install Go and set up your Go environment.
3. Install the required dependencies using `go get`.
4. Start the application by running `go run main.go`.
5. The API will be accessible at `http://localhost:8080`.

## API Endpoints

- **GET /list:** Retrieve a list of all words in the dictionary.
- **GET /word/{name}:** Retrieve the definition of a specific word.
- **POST /word:** Add a new word and its definition to the dictionary.
- **DELETE /delete/{name}:** Delete a word and its definition from the dictionary.
- **POST /update/{name}:** Update the definition of an existing word in the dictionary.

## Usage

- Use a REST client or cURL to send requests to the API endpoints.
- Refer to the API documentation or the code itself for request and response formats.

### Autor

Marlyse Saintich HANGAMALONGO M
