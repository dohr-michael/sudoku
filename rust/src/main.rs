use sudoku_resolver_rust::helpers::{print_board, ErrorMessage, N};
use sudoku_resolver_rust::sudoku::resolver::{Resolver, SimpleResolver};
use sudoku_resolver_rust::sudoku::validator::BasicValidator;

fn main() -> Result<(), ErrorMessage> {
    let board: [[i32; 9]; 9] = [
        [3, 0, 6, 5, 0, 8, 4, 0, 0],
        [5, 2, 0, 0, 0, 0, 0, 0, 0],
        [0, 8, 7, 0, 0, 0, 0, 3, 1],
        [0, 0, 3, 0, 1, 0, 0, 8, 0],
        [9, 0, 0, 8, 6, 3, 0, 0, 5],
        [0, 5, 0, 0, 9, 0, 6, 0, 0],
        [1, 3, 0, 0, 0, 0, 2, 5, 0],
        [0, 0, 0, 0, 0, 0, 7, 4, 0],
        [0, 0, 5, 2, 0, 6, 3, 0, 0],
    ];

    let resolver = SimpleResolver::new(
        Box::new(BasicValidator {})
    );

    let start = std::time::Instant::now();
    let solution = resolver.resolve(&board)?;
    let elapsed = start.elapsed();

    if solution.1 {
        println!("Sudoku solved successfully in : {} ms", elapsed.as_millis());
        print_board(&solution.0);
    } else {
        println!("No solution exists for the given Sudoku.");
        print_board(&solution.0);
    }

    Ok(())
}
