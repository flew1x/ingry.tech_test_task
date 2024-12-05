# Testing

First of all you should relocate to the `task` folder with internal Makefile after you should input the command `make testing` and you can see results of all tests

# Launch

If you want to launch the project you should input `make run-docker-local` this command will start building the image with your app

# Migrate

The project ueses auto migrations with gorm but if you want to migrate data yourself you should input the command `make up`. This project uses the goose migrate tool.

# Swagger

You can find a public swagger endpoint: `http://localhost:8080/swagger/index.html`. Also if you want to update your old swagger file you can input `swag init --pd -g path-to-the-main-file` to create a new doc. 

# Configs

All configs placed in the `configs` folder and you can change it

The whole project can be improved, and for the test assignment I missed a lot of things to make the project more robust. 