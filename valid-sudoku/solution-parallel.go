func isValidSudoku(board [][]byte) bool {
    messages := make(chan bool)
    for _, num := range []byte("123456789") {
        go checkNum(board, num, messages)
    }
    count := 0
    for {
        res := <-messages
        if res == false {
            return false
        } else {
            count++
            if count == 9 {
                return true
            }
        }
    }
    
}

func checkNum(board [][]byte, num byte, messages chan bool) {
    for i := 0; i < 9; i++ {
        foundRow := false
        foundColumn := false
        for j := 0; j < 9; j++ {
            // check row
            if board[i][j] == num {
                if foundRow {
                    messages <- false
                    return
                } else {
                    foundRow = true
                }
            }
            // check column 
            if board[j][i] == num {
                if foundColumn {
                    messages <- false
                    return
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
                        messages <- false
                        return
                    } else {
                        foundSquare = true
                    }
                }
            }
        }
    }
    messages <- true
    return
}