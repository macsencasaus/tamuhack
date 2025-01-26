package messages

type ServerMessageType string

const (
	ServerMessageTypeHubGreeting  = "ServerMessageHubGreeting"
	ServerMessageTypeRoomGreeting = "ServerMessageRoomGreeting"
	ServerMessageCountdown        = "ServerMessageCountdown"
	ServerMessageClientJoined     = "ServerMessageClientJoined"
	ServerMessageClientLeft       = "ServerMessageClientLeft"
	ServerMessageRoundStart       = "ServerMessageRoundStart"
	ServerMessageRoundEnd         = "ServerMessageRoundEnd"
	ServerMessageTestPassed       = "ServerMessageTestPassed"
	ServerMessageTestFailed       = "ServerMessageTestFailed"
)

type ServerMessage interface {
	serverMessage()
}

type HubGreetingMessage struct {
	Type   ServerMessageType `json:"type"`
	Player PlayerInfo        `json:"player"`
}

func NewHubGreetingMessage(p PlayerInfo) HubGreetingMessage {
	return HubGreetingMessage{
		Type:   ServerMessageTypeHubGreeting,
		Player: p,
	}
}
func (m HubGreetingMessage) serverMessage() {}

type RoomGreetingMessage struct {
	Type         ServerMessageType `json:"type"`
	LobbyId      int               `json:"lobbyId"`
	OtherPlayers []PlayerInfo      `json:"otherPlayers"`
}

func NewRoomGreetingMessage(lobbyId int, otherPlayers []PlayerInfo) RoomGreetingMessage {
	return RoomGreetingMessage{
		Type:         ServerMessageTypeRoomGreeting,
		LobbyId:      lobbyId,
		OtherPlayers: otherPlayers,
	}
}
func (m RoomGreetingMessage) serverMessage() {}

type CountdownMessage struct {
	Type  ServerMessageType `json:"type"`
	Count int               `json:"count"`
}

func NewCountdownMessage(count int) CountdownMessage {
	return CountdownMessage{
		Type:  ServerMessageCountdown,
		Count: count,
	}
}
func (m CountdownMessage) serverMessage() {}

type ClientJoinedMessage struct {
	Type   ServerMessageType `json:"type"`
	Player PlayerInfo        `json:"player"`
}

func NewClientJoinedMessage(player PlayerInfo) ClientJoinedMessage {
	return ClientJoinedMessage{
		Type:   ServerMessageClientJoined,
		Player: player,
	}
}
func (m ClientJoinedMessage) serverMessage() {}

type ClientLeftMessage struct {
	Type   ServerMessageType `json:"type"`
	Player PlayerInfo        `json:"player"`
}

func NewClientLeftMessage(player PlayerInfo) ClientLeftMessage {
	return ClientLeftMessage{
		Type:   ServerMessageClientLeft,
		Player: player,
	}
}
func (m ClientLeftMessage) serverMessage() {}

type RoundStartMessage struct {
	Type             ServerMessageType        `json:"type"`
	Round            int                      `json:"round"`
	Time             int                      `json:"time"`
	Prompt           string                   `json:"prompt"`
	Templates        languageFunctionTemplate `json:"templates"`
	NumTestCases     int                      `json:"numTestCases"`
	VisibleTestCases []testCase               `json:"visibleTestCases"`
}

func NewRoundStartMessage(round int, sec int) RoundStartMessage {
	return RoundStartMessage{
		Type:  ServerMessageRoundStart,
		Round: round,
		Time:  sec,
        Prompt: "<p>Add two numbers.<p>",
        Templates: languageFunctionTemplate{
            Python: "def add(a, b):",
            Javascript: "function add(a, b){\n\n}",
            Cpp: "int add(int a, int b){\n\n}",
        },
        NumTestCases: 20,
        VisibleTestCases: []testCase{
            {
                Input: "1, 2",
                Output: "3",
            },
            {
                Input: "-1, 1",
                Output: "0",
            },
            {
                Input: "77, 33",
                Output: "110",
            },
        },
	}
}
func (m RoundStartMessage) serverMessage() {}

type RoundEndMessage struct {
	Type  ServerMessageType `json:"type"`
	Round int               `json:"round"`
}

func NewRoundEndMessage(round int) RoundEndMessage {
	return RoundEndMessage{
		Type:  ServerMessageRoundEnd,
		Round: round,
	}
}
func (m RoundEndMessage) serverMessage() {}

type TestPassedMessage struct {
	Type     ServerMessageType `json:"type"`
	Question string            `json:"question"`
}

func NewTestPassedMessage(question string) TestPassedMessage {
	return TestPassedMessage{
		Type:     ServerMessageTestPassed,
		Question: question,
	}
}
func (m TestPassedMessage) serverMessage() {}

type TestFailedMessage struct {
	Type     ServerMessageType `json:"type"`
	Question string            `json:"question"`
	// TODO: add failure reason
}

func NewTestFailedMessage(question string) TestFailedMessage {
	return TestFailedMessage{
		Type:     ServerMessageTestFailed,
		Question: question,
	}
}
func (m TestFailedMessage) servermessage() {}

type PlayerInfo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type languageFunctionTemplate struct {
	Python     string `json:"python"`
	Javascript string `json:"javascript"`
	Cpp        string `json:"cpp"`
}

type testCase struct {
	Input  string `json:"input"`
	Output string `json:"output"`
}
