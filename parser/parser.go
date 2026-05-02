// Package parser provides a struct 'Parser' that provides subscribers to Rocket League StatsAPI events
// (see https://www.rocketleague.com/en/developer/stats-api)
//
//go:generate godocdown -output=README.md
package parser

import (
	"github.com/bk2004/rl-statsapi-lib/parser/internal/events"
	"github.com/bk2004/rl-statsapi-lib/parser/internal/listener"
	_ "github.com/bk2004/rl-statsapi-lib/parser/internal/listener"
	"github.com/bk2004/rl-statsapi-lib/parser/internal/publisher"
)

// A Subscriber allows subscribers to a topic
type Subscriber[T any] interface {
	Subscribe() chan T
}

type subscriber[T any] struct {
	publisher publisher.Publisher[T]
}

// A Parser provides many [Subscriber]s to each individual StatsAPI event.
// It also serves a subscriber to Parser.Connected to detect when a connection is opened
// with the StatsAPI socket.
type Parser struct {
	quit chan struct{}
	// START PARSER SUBSCRIBERS

	UpdateState         Subscriber[events.UpdateStateData]
	BallHit             Subscriber[events.BallHitData]
	ClockUpdatedSeconds Subscriber[events.ClockUpdatedSecondsData]
	CountdownBegin      Subscriber[events.CountdownBeginData]
	CrossbarHit         Subscriber[events.CrossbarHitData]
	GoalReplayEnd       Subscriber[events.GoalReplayEndData]
	GoalReplayStart     Subscriber[events.GoalReplayStartData]
	GoalReplayWillEnd   Subscriber[events.GoalReplayWillEndData]
	GoalScored          Subscriber[events.GoalScoredData]
	MatchCreated        Subscriber[events.MatchCreatedData]
	MatchInitialized    Subscriber[events.MatchInitializedData]
	MatchDestroyed      Subscriber[events.MatchDestroyedData]
	MatchEnded          Subscriber[events.MatchEndedData]
	MatchPaused         Subscriber[events.MatchPausedData]
	MatchUnpaused       Subscriber[events.MatchUnpausedData]
	PodiumStart         Subscriber[events.PodiumStartData]
	ReplayCreated       Subscriber[events.ReplayCreatedData]
	RoundStarted        Subscriber[events.RoundStartedData]
	StatfeedEvent       Subscriber[events.StatfeedEventData]
	// END PARSER SUBSCRIBERS

	Connected Subscriber[Connected]
}

type Connected = listener.Connected

func (p *Parser) Quit() {
	close(p.quit)
}

func (s *subscriber[T]) Subscribe() chan T {
	return s.publisher.Subscribe()
}

func newSubscriber[T any](publisher publisher.Publisher[T]) Subscriber[T] {
	return &subscriber[T]{publisher: publisher}
}

// A Config changes the behavior of a [Parser]
// Set Config.Port to the port of the StatsAPI socket
type Config struct {
	Port int
}

// Creates a new [Parser], configurable via [Config]
func New(cfg Config) Parser {
	if cfg.Port == 0 {
		cfg.Port = listener.PORT
	}

	publishers, quit := listener.Listen()
	return Parser{
		quit: quit,
		// START INIT PARSER SUBSCRIBERS
		UpdateState:         newSubscriber(publishers.UpdateState),
		BallHit:             newSubscriber(publishers.BallHit),
		ClockUpdatedSeconds: newSubscriber(publishers.ClockUpdatedSeconds),
		CountdownBegin:      newSubscriber(publishers.CountdownBegin),
		CrossbarHit:         newSubscriber(publishers.CrossbarHit),
		GoalReplayEnd:       newSubscriber(publishers.GoalReplayEnd),
		GoalReplayStart:     newSubscriber(publishers.GoalReplayStart),
		GoalReplayWillEnd:   newSubscriber(publishers.GoalReplayWillEnd),
		GoalScored:          newSubscriber(publishers.GoalScored),
		MatchCreated:        newSubscriber(publishers.MatchCreated),
		MatchInitialized:    newSubscriber(publishers.MatchInitialized),
		MatchDestroyed:      newSubscriber(publishers.MatchDestroyed),
		MatchEnded:          newSubscriber(publishers.MatchEnded),
		MatchPaused:         newSubscriber(publishers.MatchPaused),
		MatchUnpaused:       newSubscriber(publishers.MatchUnpaused),
		PodiumStart:         newSubscriber(publishers.PodiumStart),
		ReplayCreated:       newSubscriber(publishers.ReplayCreated),
		RoundStarted:        newSubscriber(publishers.RoundStarted),
		StatfeedEvent:       newSubscriber(publishers.StatfeedEvent),
		// END INIT PARSER SUBSCRIBERS
		Connected: newSubscriber(publishers.Connected),
	}
}
