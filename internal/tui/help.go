package tui

import "github.com/charmbracelet/bubbles/help"

func RenderHelp(h help.Model, keys KeyMap) string {
	return h.View(keys)
}
