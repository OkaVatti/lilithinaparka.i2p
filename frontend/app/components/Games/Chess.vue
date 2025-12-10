<!-- frontend/app/components/Games/Chess.vue -->
<template>
    <div class="chess-game">
        <div class="game-header">
            <h2 class="text-2xl font-bold font-mono text-accent mb-2">Chess</h2>
            <p class="text-text-secondary text-sm mb-4">
                Classic chess with AI opponent
            </p>
        </div>

        <div class="game-info">
            <div class="player-info">
                <div class="player white">
                    <span class="player-icon">♔</span>
                    <span class="player-name">{{ whitePlayer }}</span>
                    <span class="captured-pieces">
                        <span v-for="piece in capturedBlack" :key="piece">{{ piece }}</span>
                    </span>
                </div>
                <div class="turn-indicator" v-if="currentTurn === 'white'">
                    Your Turn
                </div>
            </div>

            <div class="game-status">
                <div class="status-item">Move: {{ moveNumber }}</div>
                <div class="status-item" v-if="isCheck">Check!</div>
                <div class="status-item" v-if="isCheckmate">Checkmate!</div>
            </div>

            <div class="player-info">
                <div class="player black">
                    <span class="player-icon">♚</span>
                    <span class="player-name">{{ blackPlayer }}</span>
                    <span class="captured-pieces">
                        <span v-for="piece in capturedWhite" :key="piece">{{ piece }}</span>
                    </span>
                </div>
                <div class="turn-indicator" v-if="currentTurn === 'black'">
                    AI Thinking...
                </div>
            </div>
        </div>

        <div class="chess-board">
            <div class="board-coords left">
                <div v-for="n in 8" :key="n" class="coord">{{ 9 - n }}</div>
            </div>

            <div class="board-grid">
                <div v-for="(row, rowIndex) in board" :key="rowIndex" class="board-row">
                    <div v-for="(square, colIndex) in row" :key="colIndex" class="chess-square" :class="{
                        light: (rowIndex + colIndex) % 2 === 0,
                        dark: (rowIndex + colIndex) % 2 === 1,
                        selected:
                            selectedSquare &&
                            selectedSquare.row === rowIndex &&
                            selectedSquare.col === colIndex,
                        'valid-move': isValidMoveSquare(rowIndex, colIndex),
                        'last-move': isLastMoveSquare(rowIndex, colIndex),
                    }" @click="handleSquareClick(rowIndex, colIndex)">
                        <span v-if="square" class="chess-piece" :class="square.color">
                            {{ getPieceSymbol(square) }}
                        </span>
                        <div v-if="isValidMoveSquare(rowIndex, colIndex)" class="move-indicator"></div>
                    </div>
                </div>

                <div class="board-coords bottom">
                    <div v-for="letter in ['a', 'b', 'c', 'd', 'e', 'f', 'g', 'h']" :key="letter" class="coord">
                        {{ letter }}
                    </div>
                </div>
            </div>
        </div>

        <div class="game-controls">
            <button @click="newGame" class="control-button">New Game</button>
            <button @click="undoMove" class="control-button" :disabled="moveHistory.length === 0">
                Undo
            </button>
            <button @click="toggleHints" class="control-button">
                {{ showHints ? "Hide Hints" : "Show Hints" }}
            </button>
            <button @click="resign" class="control-button">Resign</button>
        </div>

        <div class="move-history">
            <h3 class="text-lg font-bold font-mono text-accent mb-2">Move History</h3>
            <div class="moves-list">
                <div v-for="(move, index) in moveHistory" :key="index" class="move-item">
                    <span class="move-number">{{ Math.floor(index / 2) + 1 }}.</span>
                    <span class="move-notation">{{ move.notation }}</span>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue'

interface ChessPiece {
    type: 'pawn' | 'knight' | 'bishop' | 'rook' | 'queen' | 'king'
    color: 'white' | 'black'
    hasMoved?: boolean
}

interface Square {
    row: number
    col: number
}

interface Move {
    from: Square
    to: Square
    piece: ChessPiece
    captured?: ChessPiece
    notation: string
}

const board = ref<(ChessPiece | null)[][]>([])
const selectedSquare = ref<Square | null>(null)
const validMoves = ref<Square[]>([])
const currentTurn = ref<'white' | 'black'>('white')
const moveNumber = ref(1)
const moveHistory = ref<Move[]>([])
const lastMove = ref<Move | null>(null)
const isCheck = ref(false)
const isCheckmate = ref(false)
const showHints = ref(false)

const whitePlayer = ref('You')
const blackPlayer = ref('AI (Easy)')

const capturedWhite = ref<string[]>([])
const capturedBlack = ref<string[]>([])

const pieceSymbols: Record<string, Record<string, string>> = {
    white: {
        pawn: '♙',
        knight: '♘',
        bishop: '♗',
        rook: '♖',
        queen: '♕',
        king: '♔'
    },
    black: {
        pawn: '♟',
        knight: '♞',
        bishop: '♝',
        rook: '♜',
        queen: '♛',
        king: '♚'
    }
}

const initBoard = () => {
    board.value = Array(8).fill(null).map(() => Array(8).fill(null))

    // Place pawns
    for (let i = 0; i < 8; i++) {
        board.value[1][i] = { type: 'pawn', color: 'black' }
        board.value[6][i] = { type: 'pawn', color: 'white' }
    }

    // Place other pieces
    const backRow: ('rook' | 'knight' | 'bishop' | 'queen' | 'king')[] =
        ['rook', 'knight', 'bishop', 'queen', 'king', 'bishop', 'knight', 'rook']

    backRow.forEach((piece, i) => {
        board.value[0][i] = { type: piece, color: 'black' }
        board.value[7][i] = { type: piece, color: 'white' }
    })

    currentTurn.value = 'white'
    moveNumber.value = 1
    moveHistory.value = []
    lastMove.value = null
    isCheck.value = false
    isCheckmate.value = false
    capturedWhite.value = []
    capturedBlack.value = []
}

const getPieceSymbol = (piece: ChessPiece): string => {
    return pieceSymbols[piece.color][piece.type]
}

const handleSquareClick = (row: number, col: number) => {
    if (isCheckmate.value) return
    if (currentTurn.value === 'black') return // AI turn

    const clickedPiece = board.value[row][col]

    // If no piece is selected
    if (!selectedSquare.value) {
        if (clickedPiece && clickedPiece.color === currentTurn.value) {
            selectedSquare.value = { row, col }
            validMoves.value = calculateValidMoves(row, col)
        }
        return
    }

    // If clicking on own piece
    if (clickedPiece && clickedPiece.color === currentTurn.value) {
        selectedSquare.value = { row, col }
        validMoves.value = calculateValidMoves(row, col)
        return
    }

    // Try to move
    const isValid = validMoves.value.some(move => move.row === row && move.col === col)
    if (isValid) {
        makeMove(selectedSquare.value, { row, col })
    }

    selectedSquare.value = null
    validMoves.value = []
}

const calculateValidMoves = (row: number, col: number): Square[] => {
    const piece = board.value[row][col]
    if (!piece) return []

    const moves: Square[] = []

    switch (piece.type) {
        case 'pawn':
            moves.push(...calculatePawnMoves(row, col, piece.color))
            break
        case 'knight':
            moves.push(...calculateKnightMoves(row, col, piece.color))
            break
        case 'bishop':
            moves.push(...calculateBishopMoves(row, col, piece.color))
            break
        case 'rook':
            moves.push(...calculateRookMoves(row, col, piece.color))
            break
        case 'queen':
            moves.push(...calculateQueenMoves(row, col, piece.color))
            break
        case 'king':
            moves.push(...calculateKingMoves(row, col, piece.color))
            break
    }

    return moves
}

const calculatePawnMoves = (row: number, col: number, color: string): Square[] => {
    const moves: Square[] = []
    const direction = color === 'white' ? -1 : 1
    const startRow = color === 'white' ? 6 : 1

    // Forward move
    if (!board.value[row + direction]?.[col]) {
        moves.push({ row: row + direction, col })

        // Double move from start
        if (row === startRow && !board.value[row + direction * 2]?.[col]) {
            moves.push({ row: row + direction * 2, col })
        }
    }

    // Captures
    for (const colOffset of [-1, 1]) {
        const newCol = col + colOffset
        const targetPiece = board.value[row + direction]?.[newCol]
        if (targetPiece && targetPiece.color !== color) {
            moves.push({ row: row + direction, col: newCol })
        }
    }

    return moves
}

const calculateKnightMoves = (row: number, col: number, color: string): Square[] => {
    const moves: Square[] = []
    const offsets = [
        [-2, -1], [-2, 1], [-1, -2], [-1, 2],
        [1, -2], [1, 2], [2, -1], [2, 1]
    ]

    for (const [rowOffset, colOffset] of offsets) {
        const newRow = row + rowOffset
        const newCol = col + colOffset

        if (newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8) {
            const targetPiece = board.value[newRow][newCol]
            if (!targetPiece || targetPiece.color !== color) {
                moves.push({ row: newRow, col: newCol })
            }
        }
    }

    return moves
}

const calculateBishopMoves = (row: number, col: number, color: string): Square[] => {
    const moves: Square[] = []
    const directions = [[-1, -1], [-1, 1], [1, -1], [1, 1]]

    for (const [rowDir, colDir] of directions) {
        let newRow = row + rowDir
        let newCol = col + colDir

        while (newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8) {
            const targetPiece = board.value[newRow][newCol]

            if (!targetPiece) {
                moves.push({ row: newRow, col: newCol })
            } else {
                if (targetPiece.color !== color) {
                    moves.push({ row: newRow, col: newCol })
                }
                break
            }

            newRow += rowDir
            newCol += colDir
        }
    }

    return moves
}

const calculateRookMoves = (row: number, col: number, color: string): Square[] => {
    const moves: Square[] = []
    const directions = [[-1, 0], [1, 0], [0, -1], [0, 1]]

    for (const [rowDir, colDir] of directions) {
        let newRow = row + rowDir
        let newCol = col + colDir

        while (newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8) {
            const targetPiece = board.value[newRow][newCol]

            if (!targetPiece) {
                moves.push({ row: newRow, col: newCol })
            } else {
                if (targetPiece.color !== color) {
                    moves.push({ row: newRow, col: newCol })
                }
                break
            }

            newRow += rowDir
            newCol += colDir
        }
    }

    return moves
}

const calculateQueenMoves = (row: number, col: number, color: string): Square[] => {
    return [
        ...calculateBishopMoves(row, col, color),
        ...calculateRookMoves(row, col, color)
    ]
}

const calculateKingMoves = (row: number, col: number, color: string): Square[] => {
    const moves: Square[] = []
    const offsets = [
        [-1, -1], [-1, 0], [-1, 1],
        [0, -1], [0, 1],
        [1, -1], [1, 0], [1, 1]
    ]

    for (const [rowOffset, colOffset] of offsets) {
        const newRow = row + rowOffset
        const newCol = col + colOffset

        if (newRow >= 0 && newRow < 8 && newCol >= 0 && newCol < 8) {
            const targetPiece = board.value[newRow][newCol]
            if (!targetPiece || targetPiece.color !== color) {
                moves.push({ row: newRow, col: newCol })
            }
        }
    }

    return moves
}

const makeMove = (from: Square, to: Square) => {
    const piece = board.value[from.row][from.col]
    const captured = board.value[to.row][to.col]

    if (!piece) return

    // Move piece
    board.value[to.row][to.col] = piece
    board.value[from.row][from.col] = null

    // Record captured piece
    if (captured) {
        const symbol = getPieceSymbol(captured)
        if (captured.color === 'white') {
            capturedWhite.value.push(symbol)
        } else {
            capturedBlack.value.push(symbol)
        }
    }

    // Create move notation
    const notation = createMoveNotation(piece, from, to, captured)

    // Record move
    const move: Move = { from, to, piece, captured, notation }
    moveHistory.value.push(move)
    lastMove.value = move

    // Switch turn
    currentTurn.value = currentTurn.value === 'white' ? 'black' : 'white'
    if (currentTurn.value === 'white') {
        moveNumber.value++
    }

    // AI move
    if (currentTurn.value === 'black') {
        setTimeout(makeAIMove, 500)
    }
}

const createMoveNotation = (
    piece: ChessPiece,
    from: Square,
    to: Square,
    captured?: ChessPiece | null
): string => {
    const pieceSymbol = piece.type === 'pawn' ? '' : piece.type[0].toUpperCase()
    const fromCol = String.fromCharCode(97 + from.col)
    const toCol = String.fromCharCode(97 + to.col)
    const toRow = 8 - to.row
    const captureSymbol = captured ? 'x' : ''
    return ${ pieceSymbol }${ fromCol }${ captureSymbol }${ toCol }${ toRow }
}
const makeAIMove = () => {
    // Simple AI: find all possible moves and pick random
    const allMoves: { from: Square, to: Square }[] = []
    for (let row = 0; row < 8; row++) {
        for (let col = 0; col < 8; col++) {
            const piece = board.value[row][col]
            if (piece && piece.color === 'black') {
                const moves = calculateValidMoves(row, col)
                moves.forEach(move => {
                    allMoves.push({ from: { row, col }, to: move })
                })
            }
        }
    }
    if (allMoves.length > 0) {
        const randomMove = allMoves[Math.floor(Math.random() * allMoves.length)]
        makeMove(randomMove.from, randomMove.to)
    }
}
const isValidMoveSquare = (row: number, col: number): boolean => {
    return validMoves.value.some(move => move.row === row && move.col === col)
}
const isLastMoveSquare = (row: number, col: number): boolean => {
    if (!lastMove.value) return false
    return (
        (lastMove.value.from.row === row && lastMove.value.from.col === col) ||
        (lastMove.value.to.row === row && lastMove.value.to.col === col)
    )
}
const newGame = () => {
    if (moveHistory.value.length > 0) {
        if (!confirm('Start a new game? Current game will be lost.')) return
    }
    initBoard()
}
const undoMove = () => {
    if (moveHistory.value.length === 0) return
    // Undo last two moves (player and AI)
    for (let i = 0; i < 2; i++) {
        const move = moveHistory.value.pop()
        if (!move) break

        board.value[move.from.row][move.from.col] = move.piece
        board.value[move.to.row][move.to.col] = move.captured || null

        if (move.captured) {
            const symbol = getPieceSymbol(move.captured)
            if (move.captured.color === 'white') {
                const index = capturedWhite.value.lastIndexOf(symbol)
                if (index > -1) capturedWhite.value.splice(index, 1)
            } else {
                const index = capturedBlack.value.lastIndexOf(symbol)
                if (index > -1) capturedBlack.value.splice(index, 1)
            }
        }
    }
    currentTurn.value = 'white'
    lastMove.value = moveHistory.value[moveHistory.value.length - 1] || null
}
const toggleHints = () => {
    showHints.value = !showHints.value
}
const resign = () => {
    if (confirm('Are you sure you want to resign?')) {
        alert('You resigned. AI wins!')
        newGame()
    }
}
// Initialize on mount
initBoard()
</script>

<style scoped>
.chess-game {
    @apply p-6 bg-bg-secondary rounded-xl border border-accent/20;
}

.game-header {
    @apply mb-4 text-center;
}

.game-info {
    @apply flex justify-between items-center mb-6;
}

.player-info {
    @apply flex flex-col gap-2;
}

.player {
    @apply flex items-center gap-2 font-mono;
}

.player.white {
    @apply text-text-primary;
}

.player.black {
    @apply text-text-secondary;
}

.player-icon {
    @apply text-2xl;
}

.player-name {
    @apply font-bold;
}

.captured-pieces {
    @apply flex gap-1 text-sm opacity-50;
}

.turn-indicator {
    @apply text-accent text-sm font-mono animate-pulse;
}

.game-status {
    @apply text-center font-mono;
}

.status-item {
    @apply text-text-secondary;
}

.chess-board {
    @apply flex gap-2 mb-6 justify-center;
}

.board-coords {
    @apply flex flex-col justify-around font-mono text-xs text-text-secondary;
}

.board-coords.left {
    @apply items-end pr-2;
}

.board-coords.bottom {
    @apply flex-row justify-around pl-6 pt-2;
}

.board-grid {
    @apply relative;
}

.board-row {
    @apply flex;
}

.chess-square {
    @apply w-12 h-12 flex items-center justify-center cursor-pointer relative transition-all duration-200;
}

.chess-square.light {
    @apply bg-gray-200;
}

.chess-square.dark {
    @apply bg-gray-400;
}

.chess-square.selected {
    @apply ring-2 ring-accent ring-inset;
}

.chess-square.valid-move {
    @apply ring-2 ring-green-500 ring-inset;
}

.chess-square.last-move {
    @apply bg-yellow-200;
}

.chess-square.last-move.dark {
    @apply bg-yellow-400;
}

.chess-piece {
    @apply text-3xl select-none;
}

.chess-piece.white {
    filter: drop-shadow(0 0 2px rgba(255, 255, 255, 0.5));
}

.chess-piece.black {
    filter: drop-shadow(0 0 2px rgba(0, 0, 0, 0.5));
}

.move-indicator {
    @apply absolute w-3 h-3 rounded-full bg-green-500 opacity-50;
}

.game-controls {
    @apply flex justify-center gap-4 mb-6;
}

.control-button {
    @apply px-4 py-2 bg-bg-primary text-text-secondary rounded-lg font-mono hover:text-accent hover:bg-accent/10 transition-all duration-200 disabled:opacity-50 disabled:cursor-not-allowed;
}

.move-history {
    @apply p-4 bg-bg-primary rounded-lg border border-accent/10 max-h-48 overflow-y-auto;
}

.moves-list {
    @apply grid grid-cols-2 gap-2 font-mono text-sm;
}

.move-item {
    @apply flex gap-2 text-text-secondary;
}

.move-number {
    @apply text-accent font-bold;
}

.move-notation {
    @apply text-text-primary;
}
</style>
