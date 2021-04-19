[![Go Reference](https://pkg.go.dev/badge/github.com/chaudharypraveen98/GoProfileAPI.svg)](https://pkg.go.dev/github.com/chaudharypraveen98/GoProfileAPI)

## Go Profile Rest API
This projects provide Rest API built on the top of the Gorm and Fiber using Golang.
It uses the json library for parsing the data from file and with the help of gorm it intialize the database.

### API Endpoints : -
1. Get `/api/v1/projects` --> Fetch all the projects
2. Get `/api/v1/projects/:id` --> Fetch a single project with id as a parameter
3. Post `/api/v1/projects/:id/update` --> Update the existing project with the help of id. Pass the body as a json.
        * Endpoint `http://127.0.0.1:3000/api/v1/projects/1/update`  
        * Json Body  
        
        ```
        {
            "title":"admin"
        }
        ```
4. Post `/api/v1/projects/create`
        * Endpoint `http://127.0.0.1:3000/api/v1/projects/create`
        * Example
        ```
        {
            "title": "praveen",
            "desc": "This project will shows the quotes. It has a collection of 60 quotes in initial release. It is bootstrapped with create-next-app.",
            "programming_language": "JavaScript",
            "stars": 0,
            "forks": 0,
            "last_updated": "21 hours ago",
            "link": "https://github.com/chaudharypraveen98/quotes-everywhere"
        }
        ```

5. Post `/api/v1/projects/:id/delete`
        * Endpoint `http://127.0.0.1:3000/api/v1/projects/1/delete`
        It doesn't require json body

#### Features
1. Json response
2. Separate urls for all projects and single projects
3. Clean Code
4. Implements the CRUD functionality.

#### Note
Feel free to contribute and suggest changes. PR's are most welcomed.

##### Last Update
19 April 2021
