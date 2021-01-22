# Simple CRUD con GO
Este proyecto consiste en una API REST para la gestion de contactos

Preparar base de datos con el script en la carpeta docker run.sh

Instalación de GO, GORM y Gorilla Mux

```bash
export PATH=$PATH:/usr/local/go/bin
go get -u github.com/gorilla/mux
go get -u github.com/jinzhu/gorm
go mod init github.com/rbartolome/go_demo_fallabella
go run main.go
```
Estructura del proyecto
```
.
├── controllers/
  ├── ...
├── models/
  └── ...
├── utils/
  └── ...
└── main.go
```
http://localhost:4000/api/contacts/

ejemplo json

{
   "nombre":"Roberto",
   "edad":33,
   "telefono":"+584129634141",
   "direccion":"caracas",
   "email":"rbartolome@acl.com",
   "descripcion":"work"
}