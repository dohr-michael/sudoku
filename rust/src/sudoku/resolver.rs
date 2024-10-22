use crate::helpers::{ErrorMessage, N};
use crate::sudoku::validator::Validator;

pub struct SimpleResolver {
    validator: Box<dyn Validator>,
}

impl SimpleResolver {
    pub fn new(validator: Box<dyn Validator>) -> Self {
        Self { validator }
    }

    fn update_board(board: &[[i32; N]; N], row: usize, col: usize, new_value: i32) -> [[i32; N]; N] {
        let mut new_board = board.clone();
        new_board[row][col] = new_value;
        new_board
    }
}

impl Resolver for SimpleResolver {
    fn resolve(&self, board: &[[i32; N]; N]) -> Result<([[i32; N]; N], bool), ErrorMessage> {
        let case_vide = 0;

        for row in 0..N {
            for col in 0..N {
                if board[row][col] == case_vide {
                    // Essaye les nombres de 1 à 9 pour la case vide
                    for num in 1..=9 {
                        if self.validator.is_valid(&board, row, col, num)? {
                            let new_board = SimpleResolver::update_board(board, row, col, num);

                            // Appel récursif pour continuer la résolution
                            let result_solved = self.resolve(&new_board)?;
                            if result_solved.1 {
                                return Ok((result_solved.0, true));
                            }
                            // Si échec, on remet la case à vide et on continue
                        }
                    }
                    // Si aucun nombre ne convient, retournez false
                    return Ok((board.clone(), false));
                }
            }
        }

        // Si aucune case vide n'a été trouvée, le Sudoku est résolu
        Ok((board.clone(), true))
    }
}

pub trait Resolver {
    fn resolve(&self, board: &[[i32; N]; N]) -> Result<([[i32; N]; N], bool), ErrorMessage>;

}