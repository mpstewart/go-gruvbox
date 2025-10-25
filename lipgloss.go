package gruvbox

// LipglossPalette exposes Gruvbox colors as hex strings suitable for lipgloss.Color.
type LipglossPalette = Palette[string]

// Lipgloss returns the Gruvbox palette specialized for Charm's Lip Gloss styles.
func Lipgloss() LipglossPalette {
	return lipglossPalette
}

var lipglossPalette = LipglossPalette{
	BG: "#282828",
}
