package parser

import (
	"flag"
	"testing"
	"time"

	"github.com/bk2004/rl-statsapi-lib/parser/internal/listener"
)

var testTimeoutSeconds = flag.Int("timeout", 30, "timeout of an event test (in seconds)")
var testAll = flag.Bool("all", false, "test all events")
var parser Parser

// START TEST EVENTS
var testUpdateState = flag.Bool("updatestate", false, "test event UpdateState")
func TestUpdateState(t* testing.T) {
	if !*testUpdateState && !*testAll {
		t.Skip("skipping test of UpdateState")
	}
	ch := parser.UpdateState.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive UpdateState event")
	}
}

var testBallHit = flag.Bool("ballhit", false, "test event BallHit")
func TestBallHit(t* testing.T) {
	if !*testBallHit && !*testAll {
		t.Skip("skipping test of BallHit")
	}
	ch := parser.BallHit.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive BallHit event")
	}
}

var testClockUpdatedSeconds = flag.Bool("clockupdatedseconds", false, "test event ClockUpdatedSeconds")
func TestClockUpdatedSeconds(t* testing.T) {
	if !*testClockUpdatedSeconds && !*testAll {
		t.Skip("skipping test of ClockUpdatedSeconds")
	}
	ch := parser.ClockUpdatedSeconds.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive ClockUpdatedSeconds event")
	}
}

var testCountdownBegin = flag.Bool("countdownbegin", false, "test event CountdownBegin")
func TestCountdownBegin(t* testing.T) {
	if !*testCountdownBegin && !*testAll {
		t.Skip("skipping test of CountdownBegin")
	}
	ch := parser.CountdownBegin.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive CountdownBegin event")
	}
}

var testCrossbarHit = flag.Bool("crossbarhit", false, "test event CrossbarHit")
func TestCrossbarHit(t* testing.T) {
	if !*testCrossbarHit && !*testAll {
		t.Skip("skipping test of CrossbarHit")
	}
	ch := parser.CrossbarHit.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive CrossbarHit event")
	}
}

var testGoalReplayEnd = flag.Bool("goalreplayend", false, "test event GoalReplayEnd")
func TestGoalReplayEnd(t* testing.T) {
	if !*testGoalReplayEnd && !*testAll {
		t.Skip("skipping test of GoalReplayEnd")
	}
	ch := parser.GoalReplayEnd.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive GoalReplayEnd event")
	}
}

var testGoalReplayStart = flag.Bool("goalreplaystart", false, "test event GoalReplayStart")
func TestGoalReplayStart(t* testing.T) {
	if !*testGoalReplayStart && !*testAll {
		t.Skip("skipping test of GoalReplayStart")
	}
	ch := parser.GoalReplayStart.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive GoalReplayStart event")
	}
}

var testGoalReplayWillEnd = flag.Bool("goalreplaywillend", false, "test event GoalReplayWillEnd")
func TestGoalReplayWillEnd(t* testing.T) {
	if !*testGoalReplayWillEnd && !*testAll {
		t.Skip("skipping test of GoalReplayWillEnd")
	}
	ch := parser.GoalReplayWillEnd.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive GoalReplayWillEnd event")
	}
}

var testGoalScored = flag.Bool("goalscored", false, "test event GoalScored")
func TestGoalScored(t* testing.T) {
	if !*testGoalScored && !*testAll {
		t.Skip("skipping test of GoalScored")
	}
	ch := parser.GoalScored.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive GoalScored event")
	}
}

var testMatchCreated = flag.Bool("matchcreated", false, "test event MatchCreated")
func TestMatchCreated(t* testing.T) {
	if !*testMatchCreated && !*testAll {
		t.Skip("skipping test of MatchCreated")
	}
	ch := parser.MatchCreated.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchCreated event")
	}
}

var testMatchInitialized = flag.Bool("matchinitialized", false, "test event MatchInitialized")
func TestMatchInitialized(t* testing.T) {
	if !*testMatchInitialized && !*testAll {
		t.Skip("skipping test of MatchInitialized")
	}
	ch := parser.MatchInitialized.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchInitialized event")
	}
}

var testMatchDestroyed = flag.Bool("matchdestroyed", false, "test event MatchDestroyed")
func TestMatchDestroyed(t* testing.T) {
	if !*testMatchDestroyed && !*testAll {
		t.Skip("skipping test of MatchDestroyed")
	}
	ch := parser.MatchDestroyed.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchDestroyed event")
	}
}

var testMatchEnded = flag.Bool("matchended", false, "test event MatchEnded")
func TestMatchEnded(t* testing.T) {
	if !*testMatchEnded && !*testAll {
		t.Skip("skipping test of MatchEnded")
	}
	ch := parser.MatchEnded.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchEnded event")
	}
}

var testMatchPaused = flag.Bool("matchpaused", false, "test event MatchPaused")
func TestMatchPaused(t* testing.T) {
	if !*testMatchPaused && !*testAll {
		t.Skip("skipping test of MatchPaused")
	}
	ch := parser.MatchPaused.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchPaused event")
	}
}

var testMatchUnpaused = flag.Bool("matchunpaused", false, "test event MatchUnpaused")
func TestMatchUnpaused(t* testing.T) {
	if !*testMatchUnpaused && !*testAll {
		t.Skip("skipping test of MatchUnpaused")
	}
	ch := parser.MatchUnpaused.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive MatchUnpaused event")
	}
}

var testPodiumStart = flag.Bool("podiumstart", false, "test event PodiumStart")
func TestPodiumStart(t* testing.T) {
	if !*testPodiumStart && !*testAll {
		t.Skip("skipping test of PodiumStart")
	}
	ch := parser.PodiumStart.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive PodiumStart event")
	}
}

var testReplayCreated = flag.Bool("replaycreated", false, "test event ReplayCreated")
func TestReplayCreated(t* testing.T) {
	if !*testReplayCreated && !*testAll {
		t.Skip("skipping test of ReplayCreated")
	}
	ch := parser.ReplayCreated.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive ReplayCreated event")
	}
}

var testRoundStarted = flag.Bool("roundstarted", false, "test event RoundStarted")
func TestRoundStarted(t* testing.T) {
	if !*testRoundStarted && !*testAll {
		t.Skip("skipping test of RoundStarted")
	}
	ch := parser.RoundStarted.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive RoundStarted event")
	}
}

var testStatfeedEvent = flag.Bool("statfeedevent", false, "test event StatfeedEvent")
func TestStatfeedEvent(t* testing.T) {
	if !*testStatfeedEvent && !*testAll {
		t.Skip("skipping test of StatfeedEvent")
	}
	ch := parser.StatfeedEvent.Subscribe()
	select {
	case <-ch:
		break
	case <-time.After(time.Duration(*testTimeoutSeconds) * time.Second):
		t.Errorf("Failed to receive StatfeedEvent event")
	}
}

// END TEST EVENTS

func TestMain(m *testing.M) {
	parser = New(Config{
		Port: listener.PORT,
	})
	flag.Parse()
	m.Run()
	parser.Quit()
}
