use std::{fs, io::{self, Write}};
use crate::models::{Tarea, Estado}; 
use chrono::{Utc, NaiveDate}; 

const RUTA_ARCHIVO: &str = "tareas.json"; 

fn actualizar_estado_vencido(tarea: &mut Tarea) {
    let hoy = Utc::now().date_naive();
    
    // Solo actualiza si el estado esta pendiente 
    if let Estado::Pendiente = tarea.estado {
        if tarea.fecha_limite < hoy {
            tarea.estado = Estado::Vencida;
        }
    }
}

// Carga las tareas desde el archivo JSON y actualiza estados de vencimiento
pub fn cargar_tareas() -> Vec<Tarea> { 
    let mut tareas = if let Ok(contenido) = fs::read_to_string(RUTA_ARCHIVO) {
        serde_json::from_str(&contenido).unwrap_or_else(|_| {
            eprintln!("Advertencia: El archivo de tareas parece estar corrupto. Se inicializará una nueva lista.");
            Vec::new()
        })
    } else {
        Vec::new()
    };

    // Actualiza el estado de las tareas cargadas
    for tarea in tareas.iter_mut() {
        actualizar_estado_vencido(tarea);
    }
    
    tareas
}

// Guarda las tareas en el archivo JSON
fn guardar_tareas(tareas: &Vec<Tarea>) -> io::Result<()> { 
    let json_string = serde_json::to_string_pretty(tareas)?;
    let mut archivo = fs::File::create(RUTA_ARCHIVO)?;
    archivo.write_all(json_string.as_bytes())
}

pub fn agregar_tarea(descripcion: String, fecha_limite: NaiveDate) { 
    let mut tareas = cargar_tareas();
    
    let siguiente_id = tareas.iter().map(|t| t.id).max().unwrap_or(0) + 1;

    let mut nueva_tarea = Tarea {
        id: siguiente_id,
        descripcion,
        estado: Estado::Pendiente, // Estado inicial siempre Pendiente
        fecha_limite, // Guardamos la fecha límite
    };

    actualizar_estado_vencido(&mut nueva_tarea); 

    tareas.push(nueva_tarea.clone());
    guardar_tareas(&tareas).expect("Error al guardar la tarea.");
    println!("Tarea '{}' agregada con ID: {}. Límite: {}.", 
        nueva_tarea.descripcion, siguiente_id, nueva_tarea.fecha_limite);
}


pub fn listar_tareas() { 
    let tareas = cargar_tareas(); 

    if tareas.is_empty() {
        println!("¡Tu lista de tareas está vacía! Usa 'cargo run -- agregar -d \"...\" -f YYYY-MM-DD' para añadir una.");
        return;
    }

    println!("\nLISTA DE TAREAS");
    println!("------------------------------------------------------------------");
    
    // Agrupa por estado
    let (pendientes, no_pendientes): (Vec<_>, Vec<_>) = tareas.into_iter().partition(|t| matches!(t.estado, Estado::Pendiente));
    let (vencidas, completadas): (Vec<_>, Vec<_>) = no_pendientes.into_iter().partition(|t| matches!(t.estado, Estado::Vencida));
    
    let mut grupos = vec![
        ("PENDIENTES", pendientes),
        ("VENCIDAS", vencidas),
        ("COMPLETADAS", completadas),
    ];

    for (titulo, lista) in grupos.iter_mut() {
        if !lista.is_empty() {
            println!("\n{}: ({})", titulo, lista.len());
            for tarea in lista.iter_mut() {
                println!("    [ID: {}] {} ({}) - Límite: {}", 
                         tarea.id, 
                         tarea.descripcion, 
                         tarea.estado, 
                         tarea.fecha_limite);
            }
        }
    }
    
    println!("\n------------------------------------------------------------------");
}


pub fn completar_tarea(id: usize) { 
    let mut tareas = cargar_tareas();
    
    let resultado_cambio: Option<String> = tareas.iter_mut()
        .find(|t| t.id == id)
        .and_then(|tarea| {
            if let Estado::Completada = tarea.estado {
                println!("La tarea con ID {} ya estaba completada.", id);
                None
            } else {
                let descripcion = tarea.descripcion.clone(); 
                tarea.estado = Estado::Completada; // Marcamos como completada
                Some(descripcion) 
            }
        });

    match resultado_cambio {
        Some(desc) => {
            guardar_tareas(&tareas).expect("Error al guardar la tarea completada.");
            println!("Tarea con ID {} ('{}') marcada como completada.", id, desc);
        }
        None => {
            if tareas.iter().all(|t| t.id != id) {
                 println!("Error: No se encontró ninguna tarea con el ID: {}.", id);
            }
        }
    }
}

pub fn mostrar_portada() { 
    println!("\n==================================================");
    println!("Gestor de tareas en rust");
    println!("==================================================");
    println!("Comandos disponibles (ejecuta 'cargo run -- help' para más detalles):");
    println!("  -> Añadir:    cargo run -- agregar -d \"Mi tarea\" -f YYYY-MM-DD");
    println!("  -> Listar:    cargo run -- listar");
    println!("  -> Completar: cargo run -- hecho <ID>");
    println!("==================================================\n");
}