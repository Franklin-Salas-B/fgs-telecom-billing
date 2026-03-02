# FGS-Telecom-Billing

## Resumen

El presente repositorio contiene el desarrollo de la aplicación web **FGS-Telecom-Billing**, implementada en el lenguaje de programación **Go (Golang)**. El sistema fue diseñado como una solución académica orientada a la gestión de usuarios y como base para plataformas administrativas y de facturación en entornos empresariales, particularmente en el sector de telecomunicaciones.

---

## 1. Objetivo del Proyecto

El objetivo general del proyecto es diseñar y desarrollar una aplicación web concurrente que permita aplicar y demostrar los siguientes conceptos fundamentales de la ingeniería de software:

* Programación web backend utilizando Go
* Programación Orientada a Objetos (POO)
* Manejo de concurrencia mediante goroutines
* Arquitectura modular y mantenible
* Integración con bases de datos relacionales (MySQL)
* Aplicación de pruebas de software

---

## 2. Alcance del Sistema

La aplicación permite la gestión básica de usuarios a través de una interfaz web, sirviendo como base funcional para la incorporación futura de módulos adicionales como productos, facturación y control de roles.

---

## 3. Funcionalidades Principales

### 3.1 Gestión de Usuarios

* Registro de nuevos usuarios mediante formularios HTML
* Listado de usuarios almacenados en la base de datos
* Validación básica de información ingresada
* Persistencia de datos en MySQL

### 3.2 Servidor Web Concurrente

* Servidor HTTP desarrollado en Go
* Ejecución en el puerto **8081**
* Atención concurrente de múltiples solicitudes mediante goroutines

### 3.3 Rutas Implementadas

* `/` : Página de inicio
* `/usuarios` : Gestión de usuarios
* `/productos` : Módulo base de productos
* `/factura` : Módulo base de facturación

---

## 4. Arquitectura del Proyecto

El proyecto está estructurado de forma modular para facilitar su mantenimiento, escalabilidad y comprensión:

* **handlers/**: Controladores responsables de gestionar las solicitudes HTTP
* **models/**: Definición de estructuras de datos y acceso a la base de datos
* **templates/**: Vistas HTML de la aplicación
* **static/**: Archivos estáticos (CSS)
* **main.go**: Punto de entrada de la aplicación

Esta organización permite aplicar principios de encapsulación, separación de responsabilidades y mantenibilidad.

---

## 5. Concurrencia

La concurrencia es gestionada de forma nativa por el servidor HTTP de Go. Cada solicitud entrante es atendida por una goroutine independiente, lo que garantiza un manejo eficiente de múltiples peticiones simultáneas y un uso óptimo de los recursos del sistema.

---

## 6. Pruebas de Software

El proyecto contempla la aplicación de diferentes tipos de pruebas:

* Pruebas unitarias para validar funciones individuales
* Pruebas de integración para verificar la interacción entre backend y base de datos
* Pruebas de aceptación para confirmar el cumplimiento de los requisitos funcionales

Estas pruebas contribuyen a garantizar la calidad y estabilidad del sistema.

---

## 7. Estado Actual del Proyecto

* Servidor Go operativo
* Interfaz HTML/CSS funcional
* Conexión estable a base de datos MySQL
* Código versionado y documentado en GitHub

---

## 8. Información Académica

* **Autor:** Franklin Salas Bahamonde
* **Curso:** Programación Orientada a Objetos 2-CIB-3A
* **Institución:** Universidad Internacional del Ecuador UIDE
* **Fecha:** Marzo 2026

---

## 9. Trabajo Futuro

Como líneas de trabajo futuras se propone:

* Implementación de autenticación y control de roles
* Incremento de la cobertura de pruebas automatizadas
* Evaluación del sistema bajo escenarios de alta carga
* Despliegue en entornos de nube

---

## 10. Contenido del Repositorio

Este repositorio incluye:

* Código fuente completo de la aplicación
* Documentación técnica asociada al proyecto
* Archivos de configuración y estructura base del sistema

---

© 2026 – Franklin Salas
