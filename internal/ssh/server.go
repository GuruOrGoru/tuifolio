package handlers

import (
	"fmt"
	"net"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/ssh"
	"github.com/charmbracelet/wish"
	"github.com/charmbracelet/wish/logging"

	wishtea "github.com/charmbracelet/wish/bubbletea"
	"github.com/guruorgoru/tuifolio/internal/tui"
)

func NewSSHServer(port string, host string, signer []byte) (*ssh.Server, error) {
	s, err := wish.NewServer(
		wish.WithAddress(net.JoinHostPort(host, port)),
		wish.WithHostKeyPEM(signer),
		wish.WithMiddleware(
			logging.Middleware(),
			wishtea.Middleware(func(sess ssh.Session) (tea.Model, []tea.ProgramOption) {
				return tui.NewModel(), []tea.ProgramOption{
					tea.WithAltScreen(),
					tea.WithMouseCellMotion(),
				}
			}),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create server: %w", err)
	}
	return s, nil
}
