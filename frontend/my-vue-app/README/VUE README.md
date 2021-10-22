
## Vue JS
# Instalación
1. Tener node instalado
Abrir una terminal e introducir el siguiente comando:
```
npm install -g @vue/cli
```
Recomendado ejecutar ese comado dos veces para asegurarse que no falte ningún paquete.

# Generar nuevo proyecto
Ejecutar lo siguiente carga una interfaz en el navegador para empezar a crear proyectos:
```
vue ui
```
# Generar el proyecto para distribución
Se le indica a Go como backend que cargue el servidor, y junto a esto carga una interfaz ya terminada ubicada en ./dist
Al construir la versión del proyecto en vue, lo generado se guarda precisamente en esa carpeta ./dist y el backend apunta a esta.

Para empezar la construccion del proyecto:
```
npm run build
```