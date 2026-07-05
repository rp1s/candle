package core

/*
 * Represents style of error reporting
 */
type Style struct {
	ErrorPrefix        string
	ErrorPrefixColor   string
	TitleColor         string
	FileNameColor      string
	FileNameColon      string
	FileNameColonColor string
	LineNumberColor    string
	LineColor          string
	SpaceChar          string
	ArrowChar          string
	ArrowCharColor     string
	CaptionColor       string
}

// Default style
func DefaultStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   RedColor,
		TitleColor:         ResetColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: ResetColor,
		LineNumberColor:    GrayColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     RedColor,
		CaptionColor:       RedColor,
	}
}

// Minimal style with neutral colors.
func MinimalStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   ResetColor,
		TitleColor:         ResetColor,
		FileNameColor:      ResetColor,
		FileNameColon:      ":",
		FileNameColonColor: ResetColor,
		LineNumberColor:    ResetColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     ResetColor,
		CaptionColor:       ResetColor,
	}
}

// Colorful style with bright accents.
func ColorfulStyle() *Style {
	return &Style{
		ErrorPrefix:        "✖",
		ErrorPrefixColor:   RedColor,
		TitleColor:         YellowColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: YellowColor,
		LineNumberColor:    PurpleColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     LimeColor,
		CaptionColor:       GreenColor,
	}
}

// Dark style with muted colors.
func DarkStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   YellowColor,
		TitleColor:         GrayColor,
		FileNameColor:      BlueColor,
		FileNameColon:      ":",
		FileNameColonColor: GrayColor,
		LineNumberColor:    GrayColor,
		LineColor:          GrayColor,
		SpaceChar:          " ",
		ArrowChar:          "→",
		ArrowCharColor:     YellowColor,
		CaptionColor:       GrayColor,
	}
}

// Retro style with classic terminal colors.
func RetroStyle() *Style {
	return &Style{
		ErrorPrefix:        "!",
		ErrorPrefixColor:   YellowColor,
		TitleColor:         RedColor,
		FileNameColor:      GreenColor,
		FileNameColon:      ":",
		FileNameColonColor: RedColor,
		LineNumberColor:    BlueColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     YellowColor,
		CaptionColor:       RedColor,
	}
}

// Neon style with vivid highlights.
func NeonStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   PurpleColor,
		TitleColor:         LimeColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: LimeColor,
		LineNumberColor:    YellowColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     PurpleColor,
		CaptionColor:       LimeColor,
	}
}

// Soft style with pastel-like tones.
func SoftStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   BlueColor,
		TitleColor:         GrayColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: GrayColor,
		LineNumberColor:    GrayColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "-",
		ArrowCharColor:     BlueColor,
		CaptionColor:       BlueColor,
	}
}

// Bold style with stronger emphasis.
func BoldStyle() *Style {
	return &Style{
		ErrorPrefix:        "✖",
		ErrorPrefixColor:   RedColor,
		TitleColor:         RedColor,
		FileNameColor:      YellowColor,
		FileNameColon:      ":",
		FileNameColonColor: RedColor,
		LineNumberColor:    YellowColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     RedColor,
		CaptionColor:       RedColor,
	}
}

// Ocean style with blue-green tones.
func OceanStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   CyanColor,
		TitleColor:         BlueColor,
		FileNameColor:      GreenColor,
		FileNameColon:      ":",
		FileNameColonColor: BlueColor,
		LineNumberColor:    CyanColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "~",
		ArrowCharColor:     CyanColor,
		CaptionColor:       GreenColor,
	}
}

// Sunset style with warm colors.
func SunsetStyle() *Style {
	return &Style{
		ErrorPrefix:        "☼",
		ErrorPrefixColor:   YellowColor,
		TitleColor:         RedColor,
		FileNameColor:      PurpleColor,
		FileNameColon:      ":",
		FileNameColonColor: RedColor,
		LineNumberColor:    YellowColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     RedColor,
		CaptionColor:       YellowColor,
	}
}

// Forest style with green accents.
func ForestStyle() *Style {
	return &Style{
		ErrorPrefix:        "×",
		ErrorPrefixColor:   GreenColor,
		TitleColor:         LimeColor,
		FileNameColor:      CyanColor,
		FileNameColon:      ":",
		FileNameColonColor: LimeColor,
		LineNumberColor:    GreenColor,
		LineColor:          ResetColor,
		SpaceChar:          " ",
		ArrowChar:          "^",
		ArrowCharColor:     GreenColor,
		CaptionColor:       LimeColor,
	}
}
