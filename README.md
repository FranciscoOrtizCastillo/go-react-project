# Go y React, usando Fiber, Vitejs y Mongodb

https://docs.gofiber.io
https://pocketbase.io
https://excalidraw.com

# Creación de una aplicación web con Go y React
```
go version  

god mod init github.com/FranciscoOrtizCastillo/go-react-project

go mod tidy 

# Crear proyecto de React usando Vite (Para crear proyectos), es un reemplazo de webpack, mas rapido y eficiente
npm create vite
cd client
npm run build

```

# Ejecutar la aplicación

```
# Mongodb desde Docker
docker run -d -p 27017:27017 --name gomongodb mongo
docker ps 

# Clonar aplicacion desde github
git clone https://github.com/FranciscoOrtizCastillo/go-react-project.git

cd client
npm i
npm run build

cd ..
go run .
```

