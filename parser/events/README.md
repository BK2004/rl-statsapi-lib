# Package [github.com/bk2004/rl-statsapi-lib/parser/events](https://pkg.go.dev/github.com/bk2004/rl-statsapi-lib/parser/events?tab=doc)

```go
import github.com/bk2004/rl-statsapi-lib/parser/events
```


## Types
### Type BallHitData
```go
type BallHitData struct {
	Players []struct {
		Name     string `json:"Name"`     // Display name.
		Shortcut int64  `json:"Shortcut"` // Spectator shortcut.
		TeamNum  int64  `json:"TeamNum"`  // Team index (0 = Blue, 1 = Orange).
	} `json:"Players"`
	Ball struct {
		PreHitSpeed  float64 `json:"PreHitSpeed"`  // Ball speed before the hit (Unreal Units/second).
		PostHitSpeed float64 `json:"PostHitSpeed"` // Ball speed after the hit (Unreal Units/second).
		Location     struct {
			X float64
			Y float64
			Z float64
		} `json:"Location"` // World position (X, Y, Z) of the ball at impact.
	} `json:"Ball"`
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent one frame after the ball is hit.


### Type ClockUpdatedSecondsData
```go
type ClockUpdatedSecondsData struct {
	TimeSeconds int64  `json:"TimeSeconds"` // Seconds remaining in the match.
	BOvertime   bool   `json:"bOvertime"`   // True if the game is in overtime.
	MatchGuid   string `json:"MatchGuid"`   // Only set for online or LAN matches.
}
```
Sent when the in-game clock has changed.


### Type CountdownBeginData
```go
type CountdownBeginData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent at the start of each round when the countdown starts.


### Type CrossbarHitData
```go
type CrossbarHitData struct {
	BallSpeed     float64 `json:"BallSpeed"`   // Ball speed on impact.
	ImpactForce   float64 `json:"ImpactForce"` // Impact force of the ball relative to the crossbar normal.
	BallLastTouch struct {
		Player struct {
			Name     string `json:"Name"`     // Display name.
			Shortcut int64  `json:"Shortcut"` // Spectator shortcut.
			TeamNum  int64  `json:"TeamNum"`  // Team index (0 = Blue, 1 = Orange).
		} `json:"Player"`
		Speed float64 `json:"Speed"` // Speed of the ball resulting from this hit.
	} `json:"BallLastTouch"`
	BallLocation struct {
		X float64
		Y float64
		Z float64
	} `json:"BallLocation"` // World position (X, Y, Z) of the ball when the impact occurred.
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the ball hits a crossbar.


### Type GoalReplayEndData
```go
type GoalReplayEndData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when a goal replay ends.


### Type GoalReplayStartData
```go
type GoalReplayStartData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when a goal replay starts.


### Type GoalReplayWillEndData
```go
type GoalReplayWillEndData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the ball explodes during a goal replay. If the replay is skipped
this event will not fire.


### Type GoalScoredData
```go
type GoalScoredData struct {
	GoalSpeed      float64 `json:"GoalSpeed"` // Speed of the ball (Unreal Units/second) when it crossed the goal line.
	GoalTime       float64 `json:"GoalTime"`  // Length of the previous round in seconds.
	ImpactLocation struct {
		X float64
		Y float64
		Z float64
	} `json:"ImpactLocation"` // World position (X, Y, Z) of the ball when the goal was scored.
	Scorer struct {
		Name     string `json:"Name"`     // Display name of the scorer.
		Shortcut int64  `json:"Shortcut"` // Spectator shortcut.
		TeamNum  int64  `json:"TeamNum"`  // Team index of the scorer.
	} `json:"Scorer"`
	BallLastTouch struct {
		Player struct {
			Name     string `json:"Name"`     // Name of the player who last touched the ball.
			Shortcut int64  `json:"Shortcut"` // Spectator shortcut.
			TeamNum  int64  `json:"TeamNum"`  // Team index.
		} `json:"Player"`
		Speed float64 `json:"Speed"` // Speed of the ball resulting from this touch.
	} `json:"BallLastTouch"`
	Assister struct {
	} `json:"Assister"`
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when a goal is scored.


### Type MatchCreatedData
```go
type MatchCreatedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when all teams are created and replicated.


### Type MatchDestroyedData
```go
type MatchDestroyedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when leaving the game.


### Type MatchEndedData
```go
type MatchEndedData struct {
	MatchGuid     string `json:"MatchGuid"`     // Only set for online or LAN matches.
	WinnerTeamNum int64  `json:"WinnerTeamNum"` // Team index of the winning team.
}
```
Sent when the match ends and a winner is chosen.


### Type MatchInitializedData
```go
type MatchInitializedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the first countdown starts.


### Type MatchPausedData
```go
type MatchPausedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the game is paused by a match admin.


### Type MatchUnpausedData
```go
type MatchUnpausedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the game is unpaused by a match admin.


### Type PodiumStartData
```go
type PodiumStartData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the game enters the podium state after the match ends.


### Type ReplayCreatedData
```go
type ReplayCreatedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when a replay is initialized. Does not pertain to goal replays,
only replays you load via the Match History menu.


### Type RoundStartedData
```go
type RoundStartedData struct {
	MatchGuid string `json:"MatchGuid"` // Only set for online or LAN matches.
}
```
Sent when the game enters the active state (after the countdown finishes).


### Type StatfeedEventData
```go
type StatfeedEventData struct {
	EventName  string `json:"EventName"` // Asset name of the StatEvent (e.g. "Demolish", "Save").
	Type       string `json:"Type"`      // Localized display label for the stat (e.g. "Demolition").
	MainTarget struct {
		Name     string `json:"Name"`     // Display name.
		Shortcut int64  `json:"Shortcut"` // Spectator shortcut.
		TeamNum  int64  `json:"TeamNum"`  // Team index (0 = Blue, 1 = Orange).
	} `json:"MainTarget"`
	MatchGuid       string `json:"MatchGuid"` // Only set for online or LAN matches.
	SecondaryTarget struct {
	} `json:"SecondaryTarget"`
}
```
Sent when someone earns a stat.


### Type UpdateStateData
```go
type UpdateStateData struct {
	Players []struct {
		Name          string  `json:"Name"`          // Display name.
		PrimaryId     string  `json:"PrimaryId"`     // Platform identifier in the format Platform|Uid|Splitscreen (e.g. "Steam|123|0", "Epic|456|0").
		Shortcut      int64   `json:"Shortcut"`      // Spectator shortcut number.
		TeamNum       int64   `json:"TeamNum"`       // Team index (0 = Blue, 1 = Orange).
		Score         int64   `json:"Score"`         // Total match score.
		Goals         int64   `json:"Goals"`         // Goals scored this match.
		Shots         int64   `json:"Shots"`         // Shot attempts this match.
		Assists       int64   `json:"Assists"`       // Assists earned this match.
		Saves         int64   `json:"Saves"`         // Saves made this match.
		Touches       int64   `json:"Touches"`       // Total ball touches.
		CarTouches    int64   `json:"CarTouches"`    // Touches by the car body (not ball).
		Demos         int64   `json:"Demos"`         // Demolitions inflicted.
		BHasCar       bool    `json:"bHasCar"`       // SPECTATORTrue if the player currently has a vehicle.
		Speed         float64 `json:"Speed"`         // SPECTATORVehicle speed in Unreal Units/second.
		Boost         int64   `json:"Boost"`         // SPECTATORBoost amount 0–100.
		BBoosting     bool    `json:"bBoosting"`     // SPECTATORTrue if the player is currently boosting.
		BOnGround     bool    `json:"bOnGround"`     // SPECTATORTrue if at least 3 wheels are touching the world.
		BOnWall       bool    `json:"bOnWall"`       // SPECTATORTrue if the vehicle is on a wall.
		BPowersliding bool    `json:"bPowersliding"` // SPECTATORTrue if the player is holding handbrake.
		BDemolished   bool    `json:"bDemolished"`   // SPECTATORTrue if the vehicle is currently destroyed.
		BSupersonic   bool    `json:"bSupersonic"`   // SPECTATORTrue if the vehicle is at supersonic speed.
		Attacker      struct {
			Name     string `json:"Name"`     // Name of the player who demolished this player.
			Shortcut int64  `json:"Shortcut"` // Spectator shortcut of the attacker.
			TeamNum  int64  `json:"TeamNum"`  // Team index of the attacker.
		} `json:"Attacker"`
	} `json:"Players"`
	Game struct {
		Teams []struct {
			Name           string `json:"Name"`           // Team name.
			TeamNum        int64  `json:"TeamNum"`        // Team index.
			Score          int64  `json:"Score"`          // Team goal count.
			ColorPrimary   string `json:"ColorPrimary"`   // Hex color code (no #) for the team’s primary color.
			ColorSecondary string `json:"ColorSecondary"` // Hex color code for the team’s secondary color.
		} `json:"Teams"`
		TimeSeconds int64 `json:"TimeSeconds"` // Seconds remaining in the match.
		BOvertime   bool  `json:"bOvertime"`   // True if the match is in overtime.
		Ball        struct {
			Speed   float64 `json:"Speed"`   // Current ball speed in Unreal Units/second.
			TeamNum int64   `json:"TeamNum"` // Index of the last team to touch the ball. 255 if the ball has not been touched.
		} `json:"Ball"`
		BReplay    bool   `json:"bReplay"`    // True if a goal replay or history replay is active.
		BHasWinner bool   `json:"bHasWinner"` // True if a team has won.
		Winner     string `json:"Winner"`     // Name of the winning team. Empty string if no winner yet.
		Arena      string `json:"Arena"`      // Asset name of the current map (e.g. "Stadium_P").
		BHasTarget bool   `json:"bHasTarget"` // True if the client is currently viewing a specific vehicle.
		Target     struct {
			Name     string `json:"Name"`     // Name of the player being viewed.
			Shortcut int64  `json:"Shortcut"` // Spectator shortcut of the viewed player.
			TeamNum  int64  `json:"TeamNum"`  // Team index of the viewed player.
		} `json:"Target"`
		Frame   int64   `json:"Frame"`   // CONDITIONALCurrent frame number if a replay is active.
		Elapsed float64 `json:"Elapsed"` // CONDITIONALSeconds elapsed since game start if a replay is active.
	} `json:"Game"`
}
```
Sent X amount of times per second based on the player's PacketSendRate
preference.





