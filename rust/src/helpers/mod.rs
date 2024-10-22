pub type ErrorMessage = String;
pub const N: usize = 9;

pub fn print_board(board: &[[i32; N]; N]) {
    for row in 0..N {
        if row % 3 == 0 && row != 0 {
            println!("-----------");
        }
        for col in 0..N {
            if col % 3 == 0 && col != 0 {
                print!("|");
            }
            print!("{}", board[row][col]);
        }
        println!();
    }
}