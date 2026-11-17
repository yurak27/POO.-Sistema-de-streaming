Prueba 1 con futuras mejoras del proyecto académico desarrollado en *Go (Golang)* que implementa un sistema modular de gestión de streaming aplicando principios de arquitectura limpia y programación funcional (funciones puras, inmutabilidad y funciones de orden superior).
Con la siguiente Estructura
El código fuente está organizado en paquetes independientes:
* catalogo: Define la estructura de películas/series y contiene funciones de orden superior (como FiltrarContenido) para consultas inmutables.
* usuarios: Administra las entidades de usuario y constructores puros (control de edad y estado VIP).
* reproductor: Evalúa de manera pura las reglas de negocio para permitir o bloquear la reproducción de contenido.
* main.go: Orquesta la carga inicial del catálogo, la creación de usuarios y las pruebas del sistema.
