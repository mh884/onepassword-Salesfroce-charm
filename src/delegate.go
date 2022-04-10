package main

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

func newItemDelegate(keys *delegateKeyMap) list.DefaultDelegate {
	d := list.NewDefaultDelegate()

	d.UpdateFunc = func(msg tea.Msg, m *list.Model) tea.Cmd {
		var title string
		var itemUUID string

		if i, ok := m.SelectedItem().(item); ok {
			title = i.Title()
			itemUUID = i.description
		} else {
			return nil
		}

		err := clipboard.Init()
		if err != nil {
			panic(err)
		}

		switch msg := msg.(type) {
		case tea.KeyMsg:
			switch {
			case key.Matches(msg, keys.login):
				// get the item from vault
				var vault = viperEnvVariable("VAULT")
				CredentialsItems := getCredentials(itemUUID, vault)
				// pass sleceted item to salesforcceLogin
				sessionId := GetSalesforceSession(CredentialsItems)
				openSalesforceBrowser(sessionId, CredentialsItems.InstnaceUrl)
				return m.NewStatusMessage(statusMessageStyle("You chose " + title))

			case key.Matches(msg, keys.getSession):
				// get the item from vault
				var vault = viperEnvVariable("VAULT")
				CredentialsItems := getCredentials(itemUUID, vault)
				// pass sleceted item to salesforcceLogin
				sessionId := GetSalesforceSession(CredentialsItems)
				clipboard.Write(clipboard.FmtText, []byte(sessionId))
				return m.NewStatusMessage(statusMessageStyle("session: " + sessionId))
			}
		}

		return nil
	}

	help := []key.Binding{keys.login, keys.getSession}

	d.ShortHelpFunc = func() []key.Binding {
		return help
	}

	d.FullHelpFunc = func() [][]key.Binding {
		return [][]key.Binding{help}
	}

	return d
}

type delegateKeyMap struct {
	login      key.Binding
	getSession key.Binding
}

// Additional short help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		d.login,
		d.getSession,
	}
}

// Additional full help entries. This satisfies the help.KeyMap interface and
// is entirely optional.
func (d delegateKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{
			d.login,
			d.getSession,
		},
	}
}

func newDelegateKeyMap() *delegateKeyMap {
	return &delegateKeyMap{
		login: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "choose"),
		),
		getSession: key.NewBinding(
			key.WithKeys("x", "backspace"),
			key.WithHelp("x", "get-session"),
		),
	}
}
