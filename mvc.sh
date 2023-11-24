#!/bin/sh


echo "Create MVC Folder ..." 
mkdir config views controller routes models test 

echo "Creating File Main..."
touch .env main.go

echo "Creating File Config"
cd config/ 
touch config.go
cd ..

echo "Creating File Views"
cd views/
mkdir layouts
touch index.gohtml
cd ..

echo "Creating File Controller"
cd controller/
touch controller.go
cd ..

echo "Creating Models"
cd models/
touch models.go
cd ..

echo "Creating Routes"
cd routes/
touch routes.go
cd ..

echo "Creating Testing"
cd test/
touch test.go
cd ..

# Lanjutkan menulis ke file .env
echo "Writing env"
echo "port=8080\nhost=localhost\n" >> .env

# Lanjutkan menulis ke file main.go
echo "Writing main"
echo "package main\n\nimport (\n\t\"fmt\"\n\t\"net/http\"\n)\n\nfunc main() {\n\t// Implement your main function here\n\tfmt.Println(\"Hello, MVC!\")\n}" >> main.go

# Lanjutkan menulis ke file config.go
echo "Writing Config"
echo   "package config\n\n// Implement your configuration here" >> config/config.go

# Lanjutkan menulis ke file controller.go
echo "Writing Controller"
echo   "package controller\n\n// Implement your controllers here" >> controller/controller.go

# Lanjutkan menulis ke file routes.go
echo "Writing Routes"
echo   "package routes\n\nimport (\n\t\"net/http\"\n\t\"your_project/controller\"\n)\n\nfunc SetupRoutes() {\n\t// Implement your routes here\n\thttp.HandleFunc(\"/\", controller.HomeHandler)\n}" >> routes/routes.go

# Lanjutkan menulis ke file index.gohtml
echo "Writing Index View"
echo   "<html>\n<head>\n\t<title>Your MVC App</title>\n</head>\n<body>\n\t<h1>Welcome to Your MVC App</h1>\n\t<!-- Implement your HTML content here -->\n</body>\n</html>" >> views/index.gohtml

# Lanjutkan menulis ke file test.go
echo "Writing Test"
echo   "package test\n\nimport (\n\t\"testing\"\n\t// Import your project packages here\n)\n\nfunc TestYourFunction(t *testing.T) {\n\t// Implement your tests here\n}" >> test/test.go


# echo "dowloading package godotenv & mysql"
# echo "Do you continue dowloading [y/n]"
# read input
# if [[ $input == "y"]]; then
#    echo "y"
# fi
# Tambahkan pesan selesai
echo "Finish !!"
