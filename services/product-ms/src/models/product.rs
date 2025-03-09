use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
pub struct Product {
    pub id: i32,
    pub name: String,
    pub price: f32,
}