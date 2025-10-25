package gruvbox

// LipglossPalette exposes Gruvbox colors as hex strings suitable for lipgloss.Color.
type LipglossPalette = Palette[string]

// Lipgloss returns the Gruvbox palette specialized for Charm's Lip Gloss styles.
func Lipgloss() LipglossPalette {
	return lipglossPalette
}

var lipglossPalette = LipglossPalette{
	BG:           "#282828",
	BG0H:         "#1d2021",
	BG0:          "#282828",
	BG0S:         "#32302f",
	BG1:          "#3c3836",
	BG2:          "#504945",
	BG3:          "#665c54",
	BG4:          "#7c6f64",
	FG0:          "#fbf1c7",
	FG1:          "#ebdbb2",
	FG2:          "#d5c4a1",
	FG3:          "#bdae93",
	FG4:          "#a89984",
	FG:           "#ebdbb2",
	Gray:         "#a89984",
	GrayAlt:      "#928374",
	Red:          "#cc241d",
	RedBright:    "#fb4934",
	Green:        "#98971a",
	GreenBright:  "#b8bb26",
	Yellow:       "#d79921",
	YellowBright: "#fabd2f",
	Blue:         "#458588",
	BlueBright:   "#83a598",
	Purple:       "#b16286",
	PurpleBright: "#d3869b",
	Aqua:         "#689d6a",
	AquaBright:   "#8ec07c",
	Orange:       "#d65d0e",
	OrangeBright: "#fe8019",
}
