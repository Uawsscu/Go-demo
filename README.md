# Go-demo

## # Create Local Database with Docker Compose

To create a local database using Docker Compose, follow these steps: 
1. **Install Docker Compose:** If you haven't already, make sure you have Docker Compose installed on your system. 


2. **Create a Docker Compose YAML File:** Create a docker-compose.yml file in your project directory
      ```
      version: "3.3"
      services:
      db:
         image: "postgres:14.2"
         restart:
            unless-stopped
         ports:
            - "5430:5432"
         environment:
            POSTGRES_USER: postgres
            POSTGRES_PASSWORD: postgres
            POSTGRES_DB: "go_demo"
      ```

3. **Start the Database Container:** In the same directory where your docker-compose.yml file is located,

   ```
   docker-compose up -d
   ```
   or using a Makefile.
   ```
   make up
   ```

## # Updating the Database with Gorm And run

   ```
   go run main.go
   ```
 

## # Clean Code

   Using 'go mod tidy' will automatically remove unused packages from the go.sum file
   
   ``` clean go sum and go mod
   go mod tidy