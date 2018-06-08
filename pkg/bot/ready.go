package bot

import (
	"github.com/bwmarrin/discordgo"
)

// HandleReadyEvent is called when the "ready" event is received from Discord.
func HandleReadyEvent(s *discordgo.Session, event *discordgo.Ready) {
	// Set playing status.
	s.UpdateStatus(0, "!aggronym")
}
