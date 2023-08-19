use cursive::views::TextView;

struct Display {
    paused: bool,
}

impl Display {
    fn new() -> Self {
        Self { paused: false }
    }

    fn toggle_pause(&mut self) {
        self.paused = !self.paused;
    }

    fn run(&mut self) {
        let mut siv = cursive::default();

        siv.add_global_callback('q', |s| s.quit());
        siv.add_layer(TextView::new("Hello cursive! Press <q> to quit."));

        siv.run();
    }
}
