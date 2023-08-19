mod display;
mod data;

use display::Display;

fn main() {
    let _conn = data::connect();

    let mut dis = Display::new();
    dis.run();
}
