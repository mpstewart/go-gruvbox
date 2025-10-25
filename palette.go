package gruvbox

// Palette holds the Gruvbox color set for a given representation type.
type Palette[T any] struct {
	BG           T
	BG0H         T
	BG0          T
	BG0S         T
	BG1          T
	BG2          T
	BG3          T
	BG4          T
	FG0          T
	FG1          T
	FG2          T
	FG3          T
	FG4          T
	FG           T
	Gray         T
	GrayAlt      T
	Red          T
	RedBright    T
	Green        T
	GreenBright  T
	Yellow       T
	YellowBright T
	Blue         T
	BlueBright   T
	Purple       T
	PurpleBright T
	Aqua         T
	AquaBright   T
	Orange       T
	OrangeBright T
}
