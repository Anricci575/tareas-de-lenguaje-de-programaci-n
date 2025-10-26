use clap::Parser;

mod models;
mod handler;

use models::{Cli, Comandos};
use handler::{agregar_tarea, listar_tareas, completar_tarea, mostrar_portada};


fn main() {
    let cli = Cli::parse();
    
    match cli.comando {
        Some(Comandos::Agregar { descripcion, fecha }) => agregar_tarea(descripcion, fecha), 
        Some(Comandos::Listar) => listar_tareas(),
        Some(Comandos::Hecho { id }) => completar_tarea(id),
        None => {
            mostrar_portada();
        }
    }
}