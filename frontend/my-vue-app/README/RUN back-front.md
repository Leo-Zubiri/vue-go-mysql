# Iniciar el funcionamiento del BACKEND y FRONTEND

## Alternativa 1. Iniciar las construcciones finales de GO y VUE
Util cuando no te interesa estar trabajando con el frontend. Haces pruebas con el servidor, haces cambios y vuelve a ejecutarlo.

1. Asegurarse de construir la interfaz de vue a traves de su build. Para eso visita el readme en dado caso de no saber hacerlo. Si tenemos una construccion de la aplicacion deberemos tener una carpeta llamada ./dist dentro del proyecto.

2. Acceder a la ruta de la carpeta Go a traves de la terminal

3. Ejecutar:
```
go build
```
4. Si hay error acceder al archivo correspondiente en la carpeta README. Si no lo hay, desde la misma ruta de la terminal, escribir el nombre del .exe generado. Ejemplo:
```
mysqlgit.exe
```
5. Acceder en el navegador a la ruta del servidor que se acaba de iniciar
```
localhost:3000
```
6. Se debe cargar nuestra app en VUE al haber ejecutado el servidor en go.


## Alternativa 2. Iniciar edicion en tiempo real de VUE
Util para cuando se interesa estar modificando constantemente la interfaz. De igual manera se puede tener en segundo plano la ejecucion del backend en go, es decir, ejecutar el .exe del servidor go.

1. Acceder en la terminal: 
```
vue ui
```
2. Se abre el navegador con una interfaz para gestionar proyectos. Asegurarse de estar en el proyecto que se quiere editar. O bien importarlo.

Si no se abre automaticamente acceder desde el navegador a: 
```
localhost:8000
```
3. Entre las opciones acceder a Tareas/serve/    "Ejecutar Tarea"

4. Una vez terminado el proceso dar click en "Abrir aplicacion"

5. Ahora cada cambio del JS se reflejara en la edicion del frontend en Vue