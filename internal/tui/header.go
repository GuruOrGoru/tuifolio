package tui

func RenderHeader(username string) string {
	if username != "" {
		return HeaderStyle.Render("Hi " + username + "! I am Siddhartha Dhakal(GuruOrGoru).")
	}
	return HeaderStyle.Render("GuruOrGoru TUI")
}
