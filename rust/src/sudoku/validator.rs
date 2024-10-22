use crate::helpers::{ErrorMessage, N};

pub struct BasicValidator {}
impl Validator for BasicValidator {
    fn is_valid(&self, board: &[[i32; N]; N], row: usize, col: usize, new_value: i32) -> Result<bool, ErrorMessage> {
        let rs = self.is_valid_line(&board, row, new_value)? &&
            self.is_valid_column(&board, col, new_value)? &&
            self.is_valid_bloc(&board, row, col, new_value);
        Ok(rs)
    }
}

impl BasicValidator {

    fn is_valid_line(&self, board: &[[i32;N];N], row: usize, new_value: i32) -> Result<bool, ErrorMessage> {
        let lines = board.get(row).ok_or(ErrorMessage::from("row out of bounds"))?;
        Ok(!lines.contains(&new_value))
    }

    fn is_valid_column(&self, board: &[[i32;N];N], col: usize, new_value: i32) -> Result<bool, ErrorMessage> {
        let columns = board.iter().fold(Ok(vec![]), |acc: Result<Vec<i32>, ErrorMessage>, current| {
            let value_column = current.get(col).ok_or(ErrorMessage::from("column out of bounds"))?.clone();
            let t_acc = acc?;
            Ok([&t_acc[..], &[value_column]].concat())
        })?;
        Ok(!columns.contains(&new_value))
    }

    fn is_valid_bloc(&self, board: &[[i32;N];N], row: usize, col: usize, new_value: i32) -> bool {
        let start_row = row - row % 3;
        let start_col = col - col % 3;

        let result = !board[start_row..start_row + 3]
            .iter()
            .flat_map(|row| row[start_col..start_col + 3].iter().cloned().collect::<Vec<_>>())
            .collect::<Vec<_>>()
            .contains(&(new_value));
        result
    }
}


pub trait Validator {
    fn is_valid(&self, board: &[[i32; N]; N], row: usize, col: usize, new_value: i32) -> Result<bool, ErrorMessage>;
}