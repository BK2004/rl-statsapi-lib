package listener

import (
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/bk2004/rl-statsapi-lib/parser/internal/events"
	_ "github.com/bk2004/rl-statsapi-lib/parser/internal/events"
	"github.com/bk2004/rl-statsapi-lib/parser/internal/publisher"
)

const (
	PORT        = 49123
	RETRY_DELAY = 5
)

type Response struct {
	Event string          `json:"Event"`
	Data  json.RawMessage `json:"Data"`
}

type Connected struct {
	Port int
}

type EventPublishers struct {
	// START EVENT PUBLISHER STRUCT
	UpdateState         publisher.Publisher[events.UpdateStateData]
	BallHit             publisher.Publisher[events.BallHitData]
	ClockUpdatedSeconds publisher.Publisher[events.ClockUpdatedSecondsData]
	CountdownBegin      publisher.Publisher[events.CountdownBeginData]
	CrossbarHit         publisher.Publisher[events.CrossbarHitData]
	GoalReplayEnd       publisher.Publisher[events.GoalReplayEndData]
	GoalReplayStart     publisher.Publisher[events.GoalReplayStartData]
	GoalReplayWillEnd   publisher.Publisher[events.GoalReplayWillEndData]
	GoalScored          publisher.Publisher[events.GoalScoredData]
	MatchCreated        publisher.Publisher[events.MatchCreatedData]
	MatchInitialized    publisher.Publisher[events.MatchInitializedData]
	MatchDestroyed      publisher.Publisher[events.MatchDestroyedData]
	MatchEnded          publisher.Publisher[events.MatchEndedData]
	MatchPaused         publisher.Publisher[events.MatchPausedData]
	MatchUnpaused       publisher.Publisher[events.MatchUnpausedData]
	PodiumStart         publisher.Publisher[events.PodiumStartData]
	ReplayCreated       publisher.Publisher[events.ReplayCreatedData]
	RoundStarted        publisher.Publisher[events.RoundStartedData]
	StatfeedEvent       publisher.Publisher[events.StatfeedEventData]
	// END EVENT PUBLISHER STRUCT
	Connected publisher.Publisher[Connected]
}

func publish(eventName string, data json.RawMessage, publishers *EventPublishers) {
	var stringData string
	if err := json.Unmarshal(data, &stringData); err == nil {
		// Data was wrapped in a string - unmarshal again
		data = json.RawMessage(stringData)
	}
	switch eventName {
	// START LISTEN EVENT SWITCH
	case "UpdateState":
		var parsed events.UpdateStateData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal UpdateState data: %v.", err)
			return
		}
		publishers.UpdateState.Publish(parsed)
	case "BallHit":
		var parsed events.BallHitData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal BallHit data: %v.", err)
			return
		}
		publishers.BallHit.Publish(parsed)
	case "ClockUpdatedSeconds":
		var parsed events.ClockUpdatedSecondsData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal ClockUpdatedSeconds data: %v.", err)
			return
		}
		publishers.ClockUpdatedSeconds.Publish(parsed)
	case "CountdownBegin":
		var parsed events.CountdownBeginData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal CountdownBegin data: %v.", err)
			return
		}
		publishers.CountdownBegin.Publish(parsed)
	case "CrossbarHit":
		var parsed events.CrossbarHitData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal CrossbarHit data: %v.", err)
			return
		}
		publishers.CrossbarHit.Publish(parsed)
	case "GoalReplayEnd":
		var parsed events.GoalReplayEndData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayEnd data: %v.", err)
			return
		}
		publishers.GoalReplayEnd.Publish(parsed)
	case "GoalReplayStart":
		var parsed events.GoalReplayStartData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayStart data: %v.", err)
			return
		}
		publishers.GoalReplayStart.Publish(parsed)
	case "GoalReplayWillEnd":
		var parsed events.GoalReplayWillEndData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayWillEnd data: %v.", err)
			return
		}
		publishers.GoalReplayWillEnd.Publish(parsed)
	case "GoalScored":
		var parsed events.GoalScoredData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalScored data: %v.", err)
			return
		}
		publishers.GoalScored.Publish(parsed)
	case "MatchCreated":
		var parsed events.MatchCreatedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchCreated data: %v.", err)
			return
		}
		publishers.MatchCreated.Publish(parsed)
	case "MatchInitialized":
		var parsed events.MatchInitializedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchInitialized data: %v.", err)
			return
		}
		publishers.MatchInitialized.Publish(parsed)
	case "MatchDestroyed":
		var parsed events.MatchDestroyedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchDestroyed data: %v.", err)
			return
		}
		publishers.MatchDestroyed.Publish(parsed)
	case "MatchEnded":
		var parsed events.MatchEndedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchEnded data: %v.", err)
			return
		}
		publishers.MatchEnded.Publish(parsed)
	case "MatchPaused":
		var parsed events.MatchPausedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchPaused data: %v.", err)
			return
		}
		publishers.MatchPaused.Publish(parsed)
	case "MatchUnpaused":
		var parsed events.MatchUnpausedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchUnpaused data: %v.", err)
			return
		}
		publishers.MatchUnpaused.Publish(parsed)
	case "PodiumStart":
		var parsed events.PodiumStartData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal PodiumStart data: %v.", err)
			return
		}
		publishers.PodiumStart.Publish(parsed)
	case "ReplayCreated":
		var parsed events.ReplayCreatedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal ReplayCreated data: %v.", err)
			return
		}
		publishers.ReplayCreated.Publish(parsed)
	case "RoundStarted":
		var parsed events.RoundStartedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal RoundStarted data: %v.", err)
			return
		}
		publishers.RoundStarted.Publish(parsed)
	case "StatfeedEvent":
		var parsed events.StatfeedEventData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal StatfeedEvent data: %v.", err)
			return
		}
		publishers.StatfeedEvent.Publish(parsed)
		// END LISTEN EVENT SWITCH for scripts/api-to-go.js
	default:
		fmt.Printf("Unknown event: %s\n", eventName)
	}
}

func decode(decoder *json.Decoder, publishers *EventPublishers) {
	var res Response
	err := decoder.Decode(&res)
	if err != nil {
		fmt.Printf("Decode error: %v.\n", err)
		return
	}
	publish(res.Event, res.Data, publishers)
}

func connect(port int, quit chan struct{}, publishers *EventPublishers) {
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", PORT))
	if err != nil {
		fmt.Printf("Failed to connect to Rocket League StatsAPI. Retrying connection in %d.\n", RETRY_DELAY)
		time.Sleep(RETRY_DELAY * time.Second)
		return
	}
	defer conn.Close()
	publishers.Connected.Publish(Connected{Port: port})
	decoder := json.NewDecoder(conn)
	for {
		select {
		case <-quit:
			return
		default:
			decode(decoder, publishers)
		}
	}
}

func Listen() (*EventPublishers, chan struct{}) {
	quit := make(chan struct{})
	fmt.Printf("Listening for Rocket League StatsAPI on port %d.\n", PORT)
	var publishers EventPublishers = EventPublishers{
		// START INIT PUBLISHER STRUCT
		UpdateState:         publisher.New[events.UpdateStateData]("UpdateState"),
		BallHit:             publisher.New[events.BallHitData]("BallHit"),
		ClockUpdatedSeconds: publisher.New[events.ClockUpdatedSecondsData]("ClockUpdatedSeconds"),
		CountdownBegin:      publisher.New[events.CountdownBeginData]("CountdownBegin"),
		CrossbarHit:         publisher.New[events.CrossbarHitData]("CrossbarHit"),
		GoalReplayEnd:       publisher.New[events.GoalReplayEndData]("GoalReplayEnd"),
		GoalReplayStart:     publisher.New[events.GoalReplayStartData]("GoalReplayStart"),
		GoalReplayWillEnd:   publisher.New[events.GoalReplayWillEndData]("GoalReplayWillEnd"),
		GoalScored:          publisher.New[events.GoalScoredData]("GoalScored"),
		MatchCreated:        publisher.New[events.MatchCreatedData]("MatchCreated"),
		MatchInitialized:    publisher.New[events.MatchInitializedData]("MatchInitialized"),
		MatchDestroyed:      publisher.New[events.MatchDestroyedData]("MatchDestroyed"),
		MatchEnded:          publisher.New[events.MatchEndedData]("MatchEnded"),
		MatchPaused:         publisher.New[events.MatchPausedData]("MatchPaused"),
		MatchUnpaused:       publisher.New[events.MatchUnpausedData]("MatchUnpaused"),
		PodiumStart:         publisher.New[events.PodiumStartData]("PodiumStart"),
		ReplayCreated:       publisher.New[events.ReplayCreatedData]("ReplayCreated"),
		RoundStarted:        publisher.New[events.RoundStartedData]("RoundStarted"),
		StatfeedEvent:       publisher.New[events.StatfeedEventData]("StatfeedEvent"),
		// END INIT PUBLISHER STRUCT
		Connected: publisher.New[Connected]("Connected"),
	}
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				connect(PORT, quit, &publishers)
			}
		}
	}()
	return &publishers, quit
}
