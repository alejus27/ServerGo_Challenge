# ServerGo_Challenge

Servidor TCP desarollado en Golang para gestionar archivos entre clientes.

Los clientes pueden suscribirse a canales para enviar y recibir archivos a través de linea de comandos.

(Canales predeterminados "ch_1", "ch_2", "ch_3").


# -- Ejecución --

# Server
Entrar en la carpeta del servidor y sigue los siguientes pasos:

  - Ejecutar: "go run ."
  
  - Iniciar: "server start"
  
  - Detener: "server stop"


# Client
Enter the client folder and follow the next steps.

  - Ejecutar: "go run ."
  
  - Suscribirse a canal: "subscribe channel:name"
  
  - Desuscribirse a canal: "unsubscribe channel:name"

  - Enviar archivo: "send channel:name file:path"
