# Proyecto magneto-brain

Proyecto desarrollado en go usando serverless framework, que facilite el despliegue de aplicaciones en la nube de amazon web services.

## Configuracion

Información ampliada del proyecto en [este enlace](https://www.notion.so/jhonromerou/Proyecto-magneto-brain-ac63470d5e354837bb38f99c071078c1)

````git clone https://github.com/jhonromerou/magneto-brain.git````

## Despliegue del proyecto

Para realizar el despliegue se ejecuta el comando ````make deployer```` que utiliza docker para realizar el despliegue en aws. Si todo despliega correctamente, se muestran los 2 endpoints de creación y estadisticas.

## Coverage de testing

Para realizar el testing se ejecuta el comando ````make tester```` que utiliza docker para realizar el procedimiento de coverage.
