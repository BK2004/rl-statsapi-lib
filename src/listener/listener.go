package listener

import (
	"encoding/json"
	"fmt"
	"net"
	"rl-statsapi-parser/events"
	_ "rl-statsapi-parser/events"
	"rl-statsapi-parser/publisher"
	"time"
)

const (
	PORT        = 49123
	RETRY_DELAY = 5
)

var publishers map[string]publisher.Publisher[any]

type Response struct {
	Event string          `json:"Event"`
	Data  json.RawMessage `json:"Data"`
}

func publish(eventName string, data json.RawMessage) {
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
			fmt.Printf("Failed to unmarshal UpdateState data: %v\n.", err)
			return
		}
		events.UpdateState.Publish(parsed)
	case "BallHit":
		var parsed events.BallHitData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal BallHit data: %v.", err)
			return
		}
		events.BallHit.Publish(parsed)
	case "ClockUpdatedSeconds":
		var parsed events.ClockUpdatedSecondsData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal ClockUpdatedSeconds data: %v.", err)
			return
		}
		events.ClockUpdatedSeconds.Publish(parsed)
	case "CountdownBegin":
		var parsed events.CountdownBeginData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal CountdownBegin data: %v.", err)
			return
		}
		events.CountdownBegin.Publish(parsed)
	case "CrossbarHit":
		var parsed events.CrossbarHitData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal CrossbarHit data: %v.", err)
			return
		}
		events.CrossbarHit.Publish(parsed)
	case "GoalReplayEnd":
		var parsed events.GoalReplayEndData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayEnd data: %v.", err)
			return
		}
		events.GoalReplayEnd.Publish(parsed)
	case "GoalReplayStart":
		var parsed events.GoalReplayStartData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayStart data: %v.", err)
			return
		}
		events.GoalReplayStart.Publish(parsed)
	case "GoalReplayWillEnd":
		var parsed events.GoalReplayWillEndData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalReplayWillEnd data: %v.", err)
			return
		}
		events.GoalReplayWillEnd.Publish(parsed)
	case "GoalScored":
		var parsed events.GoalScoredData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal GoalScored data: %v.", err)
			return
		}
		events.GoalScored.Publish(parsed)
	case "MatchCreated":
		var parsed events.MatchCreatedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchCreated data: %v.", err)
			return
		}
		events.MatchCreated.Publish(parsed)
	case "MatchInitialized":
		var parsed events.MatchInitializedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchInitialized data: %v.", err)
			return
		}
		events.MatchInitialized.Publish(parsed)
	case "MatchDestroyed":
		var parsed events.MatchDestroyedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchDestroyed data: %v.", err)
			return
		}
		events.MatchDestroyed.Publish(parsed)
	case "MatchEnded":
		var parsed events.MatchEndedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchEnded data: %v.", err)
			return
		}
		events.MatchEnded.Publish(parsed)
	case "MatchPaused":
		var parsed events.MatchPausedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchPaused data: %v.", err)
			return
		}
		events.MatchPaused.Publish(parsed)
	case "MatchUnpaused":
		var parsed events.MatchUnpausedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal MatchUnpaused data: %v.", err)
			return
		}
		events.MatchUnpaused.Publish(parsed)
	case "PodiumStart":
		var parsed events.PodiumStartData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal PodiumStart data: %v.", err)
			return
		}
		events.PodiumStart.Publish(parsed)
	case "ReplayCreated":
		var parsed events.ReplayCreatedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal ReplayCreated data: %v.", err)
			return
		}
		events.ReplayCreated.Publish(parsed)
	case "RoundStarted":
		var parsed events.RoundStartedData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal RoundStarted data: %v.", err)
			return
		}
		events.RoundStarted.Publish(parsed)
	case "StatfeedEvent":
		var parsed events.StatfeedEventData
		if err := json.Unmarshal(data, &parsed); err != nil {
			fmt.Printf("Failed to unmarshal StatfeedEvent data: %v.", err)
			return
		}
		events.StatfeedEvent.Publish(parsed)
		// END LISTEN EVENT SWITCH for scripts/api-to-go.js
	default:
		fmt.Printf("Unknown event: %s\n", eventName)
	}
}

func listen() {
	fmt.Printf("Listening for Rocket League StatsAPI on port %d.\n", PORT)
	for {
		conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", PORT))
		if err != nil {
			fmt.Printf("Failed to connect to Rocket League StatsAPI. Retrying connection in %d.\n", RETRY_DELAY)
			time.Sleep(RETRY_DELAY * time.Second)
			continue
		}
		defer conn.Close()
		fmt.Printf("Connected to Rocket League StatsAPI.\n")

		var res Response
		decoder := json.NewDecoder(conn)
		for {
			err := decoder.Decode(&res)
			if err != nil {
				fmt.Printf("Decode error: %v.\n", err)
				break
			}
			publish(res.Event, res.Data)
		}
	}
}

func init() {
	go listen()
}
