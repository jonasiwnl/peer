use display::Display;
use data::connect;

fn main() {
    let conn = connect().unwrap();

    let dis = Display::new();
    dis.run();
}
