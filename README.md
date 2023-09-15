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

## # Updating the Database with Prisma db push

This guide explains how to update your database schema using the Prisma `db push` command. Prisma is a powerful tool that allows you to define and manage your database schema using a Prisma schema file (`schema.prisma`), and `db push` is used to synchronize your database with the schema defined in this file.

1. **Prisma CLI**: Ensure that you have Prisma CLI installed globally. You can install it using npm or yarn:

   ```bash
   npm install -g prisma
   # or
   yarn global add prisma

2. **Run Prisma** 

   Open your terminal and run the following command to push your schema changes to the database:
  
   ```Run Prisma
   prisma db push

## # Clean Code

   Using 'go mod tidy' will automatically remove unused packages from the go.sum file
   
   ``` clean go sum and go mod
   go mod tidy