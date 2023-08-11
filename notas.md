# Notas

- Funcionar de manera modular (veamos el tema de crearla en un paquete / servicio aparte)
- Posibilidad de tomar los datos a actualizar desde cualquier fuente y no solamente de DB
    - Archivo, API, Storage, etc
- ¿Forzar refresco del cache?
- Pensar en el concepto de scheduler
    - Múltiples jobs
- Tener la capacidad de Start (re-start), Stop y Cancellation
- Requerimos que se ejecute en cada pod
- Apoyarnos de alguna librería externa vs desarrollo in-house
    - Qué tan complejo lo necesitamos?
    - Opciones externas
        - <https://github.com/go-co-op/gocron>
        - <https://awesome-go.com/job-scheduler/>
        - <https://www.reddit.com/r/golang/comments/t50280/distributed_job_scheduling_with_go/>
- Pensar en una interface

