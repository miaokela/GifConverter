// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

mod utils;

use log::info;
use tauri::Manager;
use utils::common::*;

fn main() {
    info!("Start up app!");

    // 触发保存文件
    tauri::Builder::default()
        .setup(|app|{
            let window = app.get_window("main").unwrap();
            window.set_always_on_top(true).expect("could not set window always on top");
            Ok(())
        })
        .invoke_handler(tauri::generate_handler![
            get_result,
        ])
        .run(tauri::generate_context!())
        .expect("error while running tauri application");
}
