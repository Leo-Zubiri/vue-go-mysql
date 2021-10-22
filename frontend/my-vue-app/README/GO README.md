#### SETUP 
### Golang
# Instalar golang de su página oficial
```
https://golang.org/
```
# Realizar la instalación, tener en cuenta ruta de instalación
Por defecto:
```
C:\Program Files\Go\
```
# Ejecutar desde la terminal
```
go version
```

# Verificar las variables de entorno y crear Nueva Variable
```
GOROOT
C:\Program Files\Go\
```

# Despues de tener nuestros archivos Go listos
1. Tener todo archivo dentro de una carpeta dedicada solo a GO
2. Asegurarse de tener un archivo main.go que contenga la función main y package main.
3. Acceder a esta ruta desde la terminal
4. Hacer esto solo una vez para indicar que es un módulo:
```
go mod init goserver
```
```
go mod tidy
```
## Arrancar nuestro servidor GO (Backend)
# Generar .exe de nuestro backend en golang
```
go build
```
# Para ejecutar nuestro backend en GO
```
goserver.exe
```

# Alternativa sin generar un ejecutable
Si son archivos separados, no tener el mismo package
# Para ejecutar nuestro backend en GO como prueba
```
go run main.go
```


