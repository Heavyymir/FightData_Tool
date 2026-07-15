package main

type MoveData struct {
	//Move name i.e 5a
	Name				string
	//Move Type check (Normal, special etc)
	MoveType	 		MoveType
	//Move command i.e 236a
	Command				string
	//Number of frames before the move is active
	Startup				int	
	//Number of frames the move is active (Can hit the opponent)	
	Active				int		
	//Number of frames the move is in recovery after the active frames
	Recovery			int
	//The number of invincibility frames a move has.
	InvulnFrames		*int
	//Base damage of the move without scaling.
	Damage				int
	//Pointer to a knockdown advantage value. If the move does not knock down, value will be nil.
	KnockdownAdv		*int
	//Declares whether the move can fatal or punish counter, dependent on the game.
	FatalCounter		bool
	//A slice of moves the target move will naturally gatling into. Nil is a valid value.
	Gatlings			[]string
	//A slice of moves the target move will link into (Street Fighter). Nil is a valid value.
	Linksinto			[]string
	//Check to see if a move is special cancellable
	SpecialCancellable	bool
	//Check to see if a move is Super cancellable
	SuperCancellable	bool	
	//Check if a move is in counterhit recovery while active
	CounterHitRec		bool
}

type MoveType struct {
	//Stuct to assist in determining move type
	Normal		bool
	Special		bool
	Distortion	bool
	Overdrive	bool
	Astral		bool
}
