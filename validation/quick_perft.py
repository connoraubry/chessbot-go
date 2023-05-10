import chess
import subprocess
import sys
import csv


global_count = 0

results = []

def main():
    global results
    board = chess.Board(fen="r3k2r/p1ppqpb1/bn2pnp1/3PN3/1p2P3/2N2Q1p/PPPBBPPP/R3K2R w KQkq - 0 1")

    print(perft(board, int(sys.argv[1])))
    with open('res.csv','w') as out:
        csv_out=csv.writer(out)
        for row in results:
            csv_out.writerow(row)        

def perft(board, depth):
    global global_count
    global results 
    count = 0
    if depth == 0:
        # results.append((board.fen(), len(list(board.legal_moves))))
        global_count += 1
        if global_count < 2:
            sum = 0
            for legal_move in board.legal_moves:
                board.push(legal_move)
                inner_sum = 0
                for inner_legal_move in board.legal_moves:
                    board.push(inner_legal_move)
                    inner_sum += len(list(board.legal_moves))
                    board.pop()
                results.append((board.fen(), inner_sum))
                sum += inner_sum
                board.pop()
            return sum
        #     return 1

        # fen = board.fen()
        # ground_truth = len(list(board.legal_moves))
        # go_stdout = subprocess.run(["go", "run", "main.go", "-fen", fen], capture_output=True).stdout
        # go_result = int(str(go_stdout.decode('utf-8')).strip("\n"))
        # if ground_truth != go_result:
        #     print(fen, ground_truth, go_result)
        #     return -1 
        # if global_count % 500 == 0:
        #     print(global_count)
        return 0
    

    for move in board.legal_moves:
        board.push(move)

        rec = perft(board, depth-1)
        if rec == -1:
            return rec 
        count += rec 

        board.pop()

    return count 

# for move in board.legal_moves:
#     board.push(move)
#     ground_truth = len(list(board.legal_moves))
#     count += ground_truth
#     print(move, ground_truth)
#     fen = board.fen()
#     go_stdout = subprocess.run(["go", "run", "main.go", "-fen", fen], capture_output=True).stdout
#     go_result = int(str(go_stdout.decode('utf-8')).strip("\n"))

#     if ground_truth != go_result:
#         print(fen, ground_truth, go_result)


#     board.pop()

# print(count)
if __name__ == '__main__':
    main()