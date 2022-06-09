# Lector de ADN para saber si eres mutante

## Funcionamiento
- Para saber si eres mutante se encuentrarán más de una secuencia de cuatro letras
iguales, de forma oblicua, horizontal o vertical.

- La aplicación funciona de la siguiente manera:

1. Deberás hacer un POST a ec2-18-118-132-37.us-east-2.compute.amazonaws.com:3000/mutant con un json que contenga la secuencia de ADN a revisar.

2. Puedes ejecutar el siguiente comando 
```
docker-compose -f docker-compose.yaml up
```
