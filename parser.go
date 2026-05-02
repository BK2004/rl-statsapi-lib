package parser

import (
	"github.com/bk2004/rl-statsapi-parser/internal/events"
	_ "github.com/bk2004/rl-statsapi-parser/internal/listener"
	"github.com/bk2004/rl-statsapi-parser/internal/publisher"
)

type Subscriber[T any] interface {
	Subscribe() chan T
}

type subscriber[T any] struct {
	publisher publisher.Publisher[T]
}

func (s *subscriber[T]) Subscribe() chan T {
	return s.publisher.Subscribe()
}

func newSubscriber[T any](publisher publisher.Publisher[T]) Subscriber[T] {
	return &subscriber[T]{publisher: publisher}
}

var (
	// START SUBSCRIBER EXPORT

	// Sent X amount of times per second based on the player's PacketSendRate preference.
	UpdateState = newSubscriber(events.UpdateState)
	// Sent one frame after the ball is hit.
	BallHit = newSubscriber(events.BallHit)
	// Sent when the in-game clock has changed.
	ClockUpdatedSeconds = newSubscriber(events.ClockUpdatedSeconds)
	// Sent at the start of each round when the countdown starts.
	CountdownBegin = newSubscriber(events.CountdownBegin)
	// Sent when the ball hits a crossbar.
	CrossbarHit = newSubscriber(events.CrossbarHit)
	// Sent when a goal replay ends.
	GoalReplayEnd = newSubscriber(events.GoalReplayEnd)
	// Sent when a goal replay starts.
	GoalReplayStart = newSubscriber(events.GoalReplayStart)
	// Sent when the ball explodes during a goal replay. If the replay is skipped this event will not fire.
	GoalReplayWillEnd = newSubscriber(events.GoalReplayWillEnd)
	// Sent when a goal is scored.
	GoalScored = newSubscriber(events.GoalScored)
	// Sent when all teams are created and replicated.
	MatchCreated = newSubscriber(events.MatchCreated)
	// Sent when the first countdown starts.
	MatchInitialized = newSubscriber(events.MatchInitialized)
	// Sent when leaving the game.
	MatchDestroyed = newSubscriber(events.MatchDestroyed)
	// Sent when the match ends and a winner is chosen.
	MatchEnded = newSubscriber(events.MatchEnded)
	// Sent when the game is paused by a match admin.
	MatchPaused = newSubscriber(events.MatchPaused)
	// Sent when the game is unpaused by a match admin.
	MatchUnpaused = newSubscriber(events.MatchUnpaused)
	// Sent when the game enters the podium state after the match ends.
	PodiumStart = newSubscriber(events.PodiumStart)
	// Sent when a replay is initialized. Does not pertain to goal replays, only replays you load via the Match History menu.
	ReplayCreated = newSubscriber(events.ReplayCreated)
	// Sent when the game enters the active state (after the countdown finishes).
	RoundStarted = newSubscriber(events.RoundStarted)
	// Sent when someone earns a stat.
	StatfeedEvent = newSubscriber(events.StatfeedEvent)

// END SUBSCRIBER EXPORT for scripts/api-to-go.js
)
