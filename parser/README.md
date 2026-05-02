# parser
--
    import "."

Package parser provides a struct 'Parser' that provides subscribers to Rocket
League StatsAPI events (see https://www.rocketleague.com/en/developer/stats-api)

## Usage

#### type Config

```go
type Config struct {
	Port int
}
```

A Config changes the behavior of a [Parser] Set Config.Port to the port of the
StatsAPI socket

#### type Connected

```go
type Connected = listener.Connected
```


#### type Parser

```go
type Parser struct {
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

	Connected Subscriber[Connected]
}
```

A Parser provides many [Subscriber]s to each individual StatsAPI event. It also
serves a subscriber to Parser.Connected to detect when a connection is opened
with the StatsAPI socket.

#### func  New

```go
func New(cfg Config) Parser
```
Creates a new [Parser], configurable via [Config]

#### func (*Parser) Quit

```go
func (p *Parser) Quit()
```

#### type Subscriber

```go
type Subscriber[T any] interface {
	Subscribe() chan T
}
```

A Subscriber allows subscribers to a topic
