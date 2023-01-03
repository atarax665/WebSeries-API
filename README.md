<img src = "https://socialify.git.ci/atarax665/WebSeries-API/image?description=1&language=1&name=1&owner=1&stargazers=1&theme=Light" alt = "WebSeries-API" width = "640" height = "320" /> </a>


*An API to do CRUD operations on Web Series database.*

##  Available Methods
```
GetAllSeries - Get all series details in the database
GetSeries - Get a series based on the ID
AddSeries - Add a series to the database
UpdateSeries - Update a series entry in the database
DeleteSeries - Delete a series entry in the database
 ```

## Requirements (software):
* Golang
* Protobuf compiler

## Local Setup:
1. Drop a ⭐ on the Github Repository. 

2. Clone the Repo by going to your local Git Client and pushing in the command: 

```sh
git clone https://github.com/atarax665/WebSeries-API.git
```

3. Install the required packages: 
```sh
go mod tidy
```

4. In terminal do
```sh
go run server/main.go 
```

5. In another terminal do 
```sh
go run client/main.go
```

6. Uncomment functions in client/main.go to test various CRUD operations.

