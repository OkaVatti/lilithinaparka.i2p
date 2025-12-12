package games

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/okavatti/lilithinaparka.i2p/backend/m/v2/server/models"
	"gorm.io/gorm"
)

type GameEngine struct {
	DB *gorm.DB
}

func NewGameEngine(db *gorm.DB) *GameEngine {
	return &GameEngine{DB: db}
}

// Game types
type GameType string

const (
	GameTypeDice     GameType = "dice"
	GameTypePuzzle   GameType = "puzzle"
	GameTypeMatch3   GameType = "match3"
	GameTypeStrategy GameType = "strategy"
	GameTypeArcade   GameType = "arcade"
)

// Game state for multiplayer
type GameSession struct {
	ID        string                 `json:"id"`
	GameSlug  string                 `json:"game_slug"`
	HostID    string                 `json:"host_id"`
	Players   []Player               `json:"players"`
	State     map[string]interface{} `json:"state"`
	CreatedAt time.Time              `json:"created_at"`
	UpdatedAt time.Time              `json:"updated_at"`
	ExpiresAt time.Time              `json:"expires_at"`
	IsPrivate bool                   `json:"is_private"`
	Password  string                 `json:"-"`
}

type Player struct {
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Score    int       `json:"score"`
	Ready    bool      `json:"ready"`
	JoinedAt time.Time `json:"joined_at"`
}

type GameResult struct {
	Winner   string        `json:"winner"`
	Scores   []Player      `json:"scores"`
	Duration time.Duration `json:"duration"`
	Data     interface{}   `json:"data"`
}

// Dice Game Implementation
type DiceGame struct {
	Players       []string       `json:"players"`
	Scores        map[string]int `json:"scores"`
	CurrentPlayer int            `json:"current_player"`
	Round         int            `json:"round"`
	Dice          []int          `json:"dice"`
	RollsLeft     int            `json:"rolls_left"`
}

func (ge *GameEngine) CreateDiceGame(sessionID string, players []string) (*DiceGame, error) {
	game := &DiceGame{
		Players:       players,
		Scores:        make(map[string]int),
		CurrentPlayer: 0,
		Round:         1,
		Dice:          make([]int, 5),
		RollsLeft:     3,
	}

	for _, player := range players {
		game.Scores[player] = 0
	}

	return game, nil
}

func (dg *DiceGame) RollDice(keep []bool) ([]int, error) {
	if len(keep) != len(dg.Dice) {
		return nil, fmt.Errorf("invalid keep array length")
	}

	if dg.RollsLeft <= 0 {
		return nil, fmt.Errorf("no rolls left")
	}

	for i := range dg.Dice {
		if !keep[i] {
			dg.Dice[i] = rand.Intn(6) + 1
		}
	}

	dg.RollsLeft--
	return dg.Dice, nil
}

func (dg *DiceGame) ScoreCategory(category string, player string) (int, error) {
	score := dg.calculateScore(category)
	dg.Scores[player] += score

	dg.CurrentPlayer = (dg.CurrentPlayer + 1) % len(dg.Players)
	if dg.CurrentPlayer == 0 {
		dg.Round++
	}
	dg.RollsLeft = 3
	dg.Dice = make([]int, 5)

	return score, nil
}

func (dg *DiceGame) calculateScore(category string) int {
	switch category {
	case "ones":
		return dg.countDice(1) * 1
	case "twos":
		return dg.countDice(2) * 2
	case "threes":
		return dg.countDice(3) * 3
	case "fours":
		return dg.countDice(4) * 4
	case "fives":
		return dg.countDice(5) * 5
	case "sixes":
		return dg.countDice(6) * 6
	case "three_of_a_kind":
		if dg.hasNOfAKind(3) {
			return dg.sumDice()
		}
	case "four_of_a_kind":
		if dg.hasNOfAKind(4) {
			return dg.sumDice()
		}
	case "full_house":
		if dg.isFullHouse() {
			return 25
		}
	case "small_straight":
		if dg.isSmallStraight() {
			return 30
		}
	case "large_straight":
		if dg.isLargeStraight() {
			return 40
		}
	case "yahtzee":
		if dg.hasNOfAKind(5) {
			return 50
		}
	case "chance":
		return dg.sumDice()
	}
	return 0
}

func (dg *DiceGame) countDice(value int) int {
	count := 0
	for _, die := range dg.Dice {
		if die == value {
			count++
		}
	}
	return count
}

func (dg *DiceGame) sumDice() int {
	sum := 0
	for _, die := range dg.Dice {
		sum += die
	}
	return sum
}

func (dg *DiceGame) hasNOfAKind(n int) bool {
	counts := make(map[int]int)
	for _, die := range dg.Dice {
		counts[die]++
	}
	for _, count := range counts {
		if count >= n {
			return true
		}
	}
	return false
}

func (dg *DiceGame) isFullHouse() bool {
	counts := make(map[int]int)
	for _, die := range dg.Dice {
		counts[die]++
	}
	hasThree := false
	hasTwo := false
	for _, count := range counts {
		if count == 3 {
			hasThree = true
		} else if count == 2 {
			hasTwo = true
		}
	}
	return hasThree && hasTwo
}

func (dg *DiceGame) isSmallStraight() bool {
	unique := make(map[int]bool)
	for _, die := range dg.Dice {
		unique[die] = true
	}

	sequences := [][]int{
		{1, 2, 3, 4},
		{2, 3, 4, 5},
		{3, 4, 5, 6},
	}

	for _, seq := range sequences {
		found := true
		for _, num := range seq {
			if !unique[num] {
				found = false
				break
			}
		}
		if found {
			return true
		}
	}
	return false
}

func (dg *DiceGame) isLargeStraight() bool {
	sorted := make([]int, len(dg.Dice))
	copy(sorted, dg.Dice)

	for i := 0; i < len(sorted)-1; i++ {
		for j := 0; j < len(sorted)-i-1; j++ {
			if sorted[j] > sorted[j+1] {
				sorted[j], sorted[j+1] = sorted[j+1], sorted[j]
			}
		}
	}

	for i := 1; i < len(sorted); i++ {
		if sorted[i] != sorted[i-1]+1 {
			return false
		}
	}
	return true
}

// Match-3 Game Implementation
type Match3Game struct {
	Board       [][]int `json:"board"`
	Score       int     `json:"score"`
	MovesLeft   int     `json:"moves_left"`
	TimeLeft    int     `json:"time_left"`
	TargetScore int     `json:"target_score"`
}

func (ge *GameEngine) CreateMatch3Game(width, height int) (*Match3Game, error) {
	game := &Match3Game{
		Board:       make([][]int, height),
		Score:       0,
		MovesLeft:   30,
		TimeLeft:    180,
		TargetScore: 1000,
	}

	for i := range game.Board {
		game.Board[i] = make([]int, width)
		for j := range game.Board[i] {
			game.Board[i][j] = rand.Intn(7) + 1
		}
	}

	game.removeMatches()
	return game, nil
}

func (mg *Match3Game) Swap(x1, y1, x2, y2 int) (bool, [][][]int) {
	if !mg.isAdjacent(x1, y1, x2, y2) {
		return false, nil
	}

	mg.Board[y1][x1], mg.Board[y2][x2] = mg.Board[y2][x2], mg.Board[y1][x1]

	matches := mg.findMatches()
	if len(matches) == 0 {
		mg.Board[y1][x1], mg.Board[y2][x2] = mg.Board[y2][x2], mg.Board[y1][x1]
		return false, nil
	}

	mg.MovesLeft--
	return true, matches
}

func (mg *Match3Game) ProcessMatches(matches [][][]int) int {
	points := 0
	matchedCells := make(map[[2]int]bool)

	for _, match := range matches {
		for _, cell := range match {
			x, y := cell[0], cell[1]
			matchedCells[[2]int{x, y}] = true
		}

		baseScore := len(match) * 100
		if len(match) >= 4 {
			baseScore *= 2
		}
		if len(match) >= 5 {
			baseScore *= 3
		}

		points += baseScore
	}

	for cell := range matchedCells {
		x, y := cell[0], cell[1]
		mg.Board[y][x] = 0
	}

	mg.applyGravity()
	mg.fillEmptySpaces()
	mg.Score += points
	return points
}

func (mg *Match3Game) isAdjacent(x1, y1, x2, y2 int) bool {
	dx := abs(x1 - x2)
	dy := abs(y1 - y2)
	return (dx == 1 && dy == 0) || (dx == 0 && dy == 1)
}

func (mg *Match3Game) findMatches() [][][]int {
	height := len(mg.Board)
	width := len(mg.Board[0])
	var matches [][][]int

	// Check horizontal matches
	for y := 0; y < height; y++ {
		for x := 0; x < width; {
			value := mg.Board[y][x]
			if value == 0 {
				x++
				continue
			}

			length := 1
			for x+length < width && mg.Board[y][x+length] == value {
				length++
			}

			if length >= 3 {
				match := make([][]int, length)
				for i := 0; i < length; i++ {
					match[i] = []int{x + i, y}
				}
				matches = append(matches, match)
			}

			x += length
		}
	}

	// Check vertical matches
	for x := 0; x < width; x++ {
		for y := 0; y < height; {
			value := mg.Board[y][x]
			if value == 0 {
				y++
				continue
			}

			length := 1
			for y+length < height && mg.Board[y+length][x] == value {
				length++
			}

			if length >= 3 {
				match := make([][]int, length)
				for i := 0; i < length; i++ {
					match[i] = []int{x, y + i}
				}
				matches = append(matches, match)
			}

			y += length
		}
	}

	return matches
}

func (mg *Match3Game) removeMatches() {
	for {
		matches := mg.findMatches()
		if len(matches) == 0 {
			break
		}

		for i := range mg.Board {
			for j := range mg.Board[i] {
				mg.Board[i][j] = rand.Intn(7) + 1
			}
		}
	}
}

func (mg *Match3Game) applyGravity() {
	width := len(mg.Board[0])
	height := len(mg.Board)

	for x := 0; x < width; x++ {
		writeY := height - 1
		for readY := height - 1; readY >= 0; readY-- {
			if mg.Board[readY][x] != 0 {
				mg.Board[writeY][x] = mg.Board[readY][x]
				if writeY != readY {
					mg.Board[readY][x] = 0
				}
				writeY--
			}
		}
	}
}

func (mg *Match3Game) fillEmptySpaces() {
	width := len(mg.Board[0])
	height := len(mg.Board)

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			if mg.Board[y][x] == 0 {
				mg.Board[y][x] = rand.Intn(7) + 1
			}
		}
	}
}

// Helper functions
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Game session management
func (ge *GameEngine) CreateSession(gameSlug, hostID string, isPrivate bool, password string) (*GameSession, error) {
	sessionID := generateSessionID()

	session := &GameSession{
		ID:        sessionID,
		GameSlug:  gameSlug,
		HostID:    hostID,
		Players:   []Player{{ID: hostID, Name: "Host", JoinedAt: time.Now()}},
		State:     make(map[string]interface{}),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		ExpiresAt: time.Now().Add(2 * time.Hour),
		IsPrivate: isPrivate,
		Password:  password,
	}

	var game models.Game
	if err := ge.DB.Where("slug = ?", gameSlug).First(&game).Error; err != nil {
		return nil, err
	}

	switch game.Category {
	case "Dice":
		players := []string{hostID}
		diceGame, _ := ge.CreateDiceGame(sessionID, players)
		session.State["game"] = diceGame
	case "Match-3":
		match3Game, _ := ge.CreateMatch3Game(8, 8)
		session.State["game"] = match3Game
	}

	sessions[sessionID] = session
	return session, nil
}

func (ge *GameEngine) JoinSession(sessionID, playerID, playerName, password string) error {
	session, exists := sessions[sessionID]
	if !exists {
		return fmt.Errorf("session not found")
	}

	if session.IsPrivate && session.Password != password {
		return fmt.Errorf("invalid password")
	}

	for _, player := range session.Players {
		if player.ID == playerID {
			return fmt.Errorf("player already joined")
		}
	}

	session.Players = append(session.Players, Player{
		ID:       playerID,
		Name:     playerName,
		JoinedAt: time.Now(),
	})
	session.UpdatedAt = time.Now()

	return nil
}

func (ge *GameEngine) UpdateGameState(sessionID string, playerID string, action string, data interface{}) error {
	session, exists := sessions[sessionID]
	if !exists {
		return fmt.Errorf("session not found")
	}

	found := false
	for _, player := range session.Players {
		if player.ID == playerID {
			found = true
			break
		}
	}
	if !found {
		return fmt.Errorf("player not in session")
	}

	session.UpdatedAt = time.Now()
	return nil
}

func generateSessionID() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 12)
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}

var sessions = make(map[string]*GameSession)
