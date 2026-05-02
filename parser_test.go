package parser

import (
	"flag"
	"fmt"
	"testing"
	"time"
)

const testTimeoutSeconds = 30

var testAll = flag.Bool("all", false, "Test all events")

// START TEST EVENTS

var testUpdateState = flag.Bool("updatestate", false, "test event UpdateState")
func TestUpdateState(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on UpdateState event")
	select {
	case <-UpdateState.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive UpdateState event")
	}
}

var testBallHit = flag.Bool("ballhit", false, "test event BallHit")
func TestBallHit(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on BallHit event")
	select {
	case <-BallHit.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive BallHit event")
	}
}

var testClockUpdatedSeconds = flag.Bool("clockupdatedseconds", false, "test event ClockUpdatedSeconds")
func TestClockUpdatedSeconds(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on ClockUpdatedSeconds event")
	select {
	case <-ClockUpdatedSeconds.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive ClockUpdatedSeconds event")
	}
}

var testCountdownBegin = flag.Bool("countdownbegin", false, "test event CountdownBegin")
func TestCountdownBegin(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on CountdownBegin event")
	select {
	case <-CountdownBegin.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive CountdownBegin event")
	}
}

var testCrossbarHit = flag.Bool("crossbarhit", false, "test event CrossbarHit")
func TestCrossbarHit(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on CrossbarHit event")
	select {
	case <-CrossbarHit.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive CrossbarHit event")
	}
}

var testGoalReplayEnd = flag.Bool("goalreplayend", false, "test event GoalReplayEnd")
func TestGoalReplayEnd(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on GoalReplayEnd event")
	select {
	case <-GoalReplayEnd.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive GoalReplayEnd event")
	}
}

var testGoalReplayStart = flag.Bool("goalreplaystart", false, "test event GoalReplayStart")
func TestGoalReplayStart(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on GoalReplayStart event")
	select {
	case <-GoalReplayStart.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive GoalReplayStart event")
	}
}

var testGoalReplayWillEnd = flag.Bool("goalreplaywillend", false, "test event GoalReplayWillEnd")
func TestGoalReplayWillEnd(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on GoalReplayWillEnd event")
	select {
	case <-GoalReplayWillEnd.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive GoalReplayWillEnd event")
	}
}

var testGoalScored = flag.Bool("goalscored", false, "test event GoalScored")
func TestGoalScored(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on GoalScored event")
	select {
	case <-GoalScored.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive GoalScored event")
	}
}

var testMatchCreated = flag.Bool("matchcreated", false, "test event MatchCreated")
func TestMatchCreated(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchCreated event")
	select {
	case <-MatchCreated.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchCreated event")
	}
}

var testMatchInitialized = flag.Bool("matchinitialized", false, "test event MatchInitialized")
func TestMatchInitialized(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchInitialized event")
	select {
	case <-MatchInitialized.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchInitialized event")
	}
}

var testMatchDestroyed = flag.Bool("matchdestroyed", false, "test event MatchDestroyed")
func TestMatchDestroyed(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchDestroyed event")
	select {
	case <-MatchDestroyed.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchDestroyed event")
	}
}

var testMatchEnded = flag.Bool("matchended", false, "test event MatchEnded")
func TestMatchEnded(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchEnded event")
	select {
	case <-MatchEnded.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchEnded event")
	}
}

var testMatchPaused = flag.Bool("matchpaused", false, "test event MatchPaused")
func TestMatchPaused(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchPaused event")
	select {
	case <-MatchPaused.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchPaused event")
	}
}

var testMatchUnpaused = flag.Bool("matchunpaused", false, "test event MatchUnpaused")
func TestMatchUnpaused(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on MatchUnpaused event")
	select {
	case <-MatchUnpaused.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive MatchUnpaused event")
	}
}

var testPodiumStart = flag.Bool("podiumstart", false, "test event PodiumStart")
func TestPodiumStart(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on PodiumStart event")
	select {
	case <-PodiumStart.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive PodiumStart event")
	}
}

var testReplayCreated = flag.Bool("replaycreated", false, "test event ReplayCreated")
func TestReplayCreated(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on ReplayCreated event")
	select {
	case <-ReplayCreated.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive ReplayCreated event")
	}
}

var testRoundStarted = flag.Bool("roundstarted", false, "test event RoundStarted")
func TestRoundStarted(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on RoundStarted event")
	select {
	case <-RoundStarted.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive RoundStarted event")
	}
}

var testStatfeedEvent = flag.Bool("statfeedevent", false, "test event StatfeedEvent")
func TestStatfeedEvent(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	fmt.Printf("Waiting on StatfeedEvent event")
	select {
	case <-StatfeedEvent.Subscribe():
		break
	case <-time.After(testTimeoutSeconds * time.Second):
		t.Errorf("Failed to receive StatfeedEvent event")
	}
}

// END TEST EVENTS

func TestMain(m *testing.M) {
	flag.Parse()
	m.Run()
}
