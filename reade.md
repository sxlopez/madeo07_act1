#

## Objetivo General.  Uso de docker en la creación de contenedores.

## Objetivos especificos. 

Construir y ejecutar tres contenedores en cuatros lenguajes: Node,js, Python, Java, C#.


## Desarrollo

### Parte 1
	- 3 Applicaciones
	- Cada uno en un puerto distinto
	- subir las imaganes a un repo de docker/aws
	
### Parte 2
	Ejecución de app basada en contendores que se ejecute con docker compose.
	

### Entrega
	Archivo comprimido zip con
	- Cuatro Docker construidos
	- Un informe PDF en
		- menos de 15 hojas
		- Calibri 11
		- interlineado 1,5


### Docker instalation

```sh
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
```

## Desarrollo

### Applicación con Node
1. Crear imagen con **docker build -t node-app:v0.0.2 .**
2. Crear contendor con **docker run --name my-node-app -d -p 3000:3000 node-app:v0.0.2**

### Applicación con Go
1. Crear imagen con **docker build -t my-go-app:v0.0.1 .**
2. Crear contendor con **docker run --name my-go-app -d -p 8080:8080 my-go-app:v0.0.1**

### Applicación con java y python
1. Crear imagen con **docker build -t my-java-python-app:v0.0.1 .**
2. Crear contendor con **docker run --name my-java-python-app -p 8080:8080 my-java-python-app:v0.0.1**

### DB docker compose
1. crear archivo de dependencias, el dockerfile y docker-compose
2. **sudo docker compose run web django-admin startproject djangopostgreapp .** 
![alt text](image.png)

3. Editar *./djangopostgreapp/settings.py* deteniendo el contenedor

```sh
DATABASES = {
	'default': {
		'ENGINE': 'django.db.backends.postgresql',
		'NAME': 'postgres',
		'USER': 'postgres',
		'PASSWORD': 'postgres',
		'HOST': 'db',
		'PORT': 5432
	}
}
```

además de agregar el '0.0.0.0' como host valido.

```sh
DATABASES = {
	'default': {
		'ENGINE': 'django.db.backends.postgresql',
		'NAME': 'postgres',
		'USER': 'postgres',
		'PASSWORD': 'postgres',
		'HOST': 'db',
		'PORT': 5432
	}
}
```

seguido de **sudo docker compose run web python manage.py migrate**, una forma de ver cuantos procesos en el cambio de base s de datos ha ocurrido, puede ejecutarse **sudo docker compose run web python manage.py makemigrations**.

y po ultimo **sudo docker compose up**

![alt text](image-1.png)


### Entorno de tratabajo

1. Agregar tu usuario al grupo docker

 

    Crea el grupo docker si no existe (esto puede que ya esté creado si Docker se instaló correctamente):

    sudo groupadd docker  

 
2. Agrega tu usuario al grupo docker (reemplaza tu_usuario con tu nombre de usuario):

sudo usermod -aG docker $USER  

O, si prefieres, puedes usar:

sudo usermod -aG docker tu_usuario  

 
3. Cierra la sesión y vuelve a iniciarla para que los cambios surtan efecto. Esto es necesario para que se apliquen los permisos del grupo. Alternativamente, puedes usar el siguiente comando para aplicar los cambios sin cerrar sesión:

newgrp docker  

 
2. Verificar la configuración

 
Para asegurarte de que todo esté funcionando correctamente, puedes ejecutar el siguiente comando para verificar que puedes usar Docker sin sudo:

docker ps  




## Referencias

### Oficiales
https://hub.docker.com/_node
https://docs.docker.com/engine/install/ubuntu/
https://go.dev/doc/install

### Complementarias
https://www.youtube.com/watch?v=cAFGcfIyKgI
https://www.youtube.com/watch?v=C5y-14YFs_8
