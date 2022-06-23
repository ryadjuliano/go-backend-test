# Overview

This is a simple app with CRUD operations to manage stocks in an inventory. This app is made as a preliminary for DANA Innovation/Ventures backend squad interview.

## Guidelines

Please implement the answers to the following problem, preferably in the specified order, 
but can be in parallel if the solution that you propose for one problem also tackles one or more of the other problems.

Please submit your implementation in a private GitHub repository. 
For each solution to a problem (or a couple of problems), please open a pull request to `master`, and then merge it, before moving on to the next solution.

As a note, if you use IDEs or other tools that adds artifacts that are not relevant to the source code in the project directory, 
please add a `.gitignore` to avoid committing to your submission repository.

It is also much preferred if you can add a Dockerfile to run this app in a Docker container, using Makefile also preferable for the workflow.

Finally, please submit your answers within the agreed timeline. 
Don't worry about the percentage of completion! 
Before you submit, please override this `README` file with instructions on how to run the app locally 
and relevant information on your technical decisions to solve these problems.

### Problem 1

In this source code we see all the functionalities lumped in the main function. 
Please refactor the code into different packages with different responsibilities, in a project structure that you see fit.

Note: Feel free to add/change libraries used if needed

### Problem 2

We can also see that the Stock entities are stored simply in an in-app key-value storage. 
Please propose the usage of a database of your choice and the database schema, and add the corresponding database client code in this source code

Note: The ID used in the current implementation is a UUID. Feel free to propose a different format

### Problem 3

Please take a look at the condition check in `GET /stock/:id` that would cause the endpoint to return a `not found` error. 
Then, try to make a request to `POST /stock` with the following request body:

```
{
    "name":"",
    "price": 0,
    "availability": 0,
    "is_active": false
}
```

Please identify the problem and add the necessary means to prevent this from happening

Note: `availability` may be 0 (but not negative), but `price` may not

### Problem 4

Pick at least one of the following improvements to implement:
- unit tests for the domain logic
- structured logging
- tracing mechanism

### Problem 5

Pick at least one of the following features to implement:

##### Option 1

Please add an endpoint that allows updating the Stock's price. And then, create a script to simulate price updates randomly within the range of 50000±5% every 1s

Notes:
- Please ensure that the price read via the `GET /stock/:id` is consistent through all price changes
- If you feel more comfortable changing the web framework to one that you are more comfortable with, feel free to do so

##### Option 2

Let's say we want to differentiate the permission granted to different actors (i.e. users of the client that consumes the API served in this app). Some actors will be allowed to do all of create, update price, and read operations, while some actors will only be allowed to update the price and do read operations, some actors will only have read access, and unauthorized actors will have no access at all. Please propose and implement a mechanism to allow us to enforce this rule

##### Option 3

Let's say the actors want to allow Stock insertion as a bulk, not just one-by-one. Please propose and implement a mechanism to enable this requirement# go-backend-test
