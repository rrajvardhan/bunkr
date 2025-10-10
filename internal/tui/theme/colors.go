package theme

var Colors = struct {
	Primary, Secondary, Tertiary, Quaternary, Quinary, Senary,
	Selections, SelectionsNoAlpha, Foregrounds, Comments, Backgrounds string
}{
	Primary:           "#a277ff", // Main brand color, buttons, links, highlights
	Secondary:         "#61ffca", // Secondary actions, success states, accents
	Tertiary:          "#ffca85", // Warnings, notifications, secondary highlights
	Quaternary:        "#f694ff", // Special elements, tags, badges
	Quinary:           "#82e2ff", // Info states, links, interactive elements
	Senary:            "#ff6767", // Errors, danger states, alerts
	Selections:        "#3d375e", // Text selections, highlighted areas (with transparency)
	SelectionsNoAlpha: "#29263c", // Solid selection backgrounds, focus states
	Foregrounds:       "#edecee", // Main text, icons, content
	Comments:          "#6d6d6d", // Secondary text, captions, disabled states
	Backgrounds:       "#15141b", // Page backgrounds, containers, panels
}
