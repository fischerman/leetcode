func isValidSudoku(board [][]byte) bool {
    for _, num := range []byte("123456789") {
        for i := 0; i < 9; i++ {
            foundRow := false
            foundColumn := false
            for j := 0; j < 9; j++ {
                // check row
                if board[i][j] == num {
                    if foundRow {
                        return false
                    } else {
                        foundRow = true
                    }
                }
                // check column 
                if board[j][i] == num {
                    if foundColumn {
                        return false
                    } else {
                        foundColumn = true
                    }
                }
            }
            foundSquare := false
            for j := (i / 3) * 3; j < (i / 3) * 3 + 3; j++ {
                for k := (i % 3) * 3; k < (i % 3) * 3 + 3; k++ {
                    // check square
                    if board[j][k] == num {
                        if foundSquare {
                            return false
                        } else {
                            foundSquare = true
                        }
                    }
                }
            }
        }
    }
    return true
}