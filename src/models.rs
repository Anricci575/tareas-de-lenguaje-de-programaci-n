use serde::{Deserialize, Serialize};
use clap::{Parser, Subcommand};
use chrono::NaiveDate; 


#[derive(Debug, Serialize, Deserialize, Clone)]
pub enum Estado {
    Pendiente,
    Vencida,
    Completada,
}

impl std::fmt::Display for Estado {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        match self {
            Estado::Pendiente => write!(f, "Pendiente"),
            Estado::Vencida => write!(f, "Vencida"),
            Estado::Completada => write!(f, "Completada"),
        }
    }
}

#[derive(Debug, Serialize, Deserialize, Clone)]
pub struct Tarea {
    pub id: usize,
    pub descripcion: String, 
    pub estado: Estado, 
    pub fecha_limite: NaiveDate, 
}


#[derive(Parser)]
#[command(name = "gestor-tareas", author, version, about = "Administrador de Tareas por Línea de Comandos (CLI) en Rust.")]
pub struct Cli {
    #[command(subcommand)]
    pub comando: Option<Comandos>,
}

#[derive(Subcommand)]
pub enum Comandos { 

    Agregar { 
        ///la descripcion 
        #[arg(short, long)]
        descripcion: String,
        ///la fecha limite de la tarea (formato YYYY-MM-DD)
        #[arg(short, long)]
        fecha: NaiveDate, 
    },
    /// lista todas las tareas (pendientes, vencidas, completadas)
    Listar, 
    ///marca una tarea como completada usando su ID. Uso: hecho --id 1
    Hecho { 
        id: usize,
    },
}