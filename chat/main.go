package chat

import "fmt"

type Message struct {
	Name     string
	Content  string
	IsServer bool
}

type Chat struct {
	Messages []Message
}

func New() *Chat {
	return &Chat{}
}

// AddMessage adds a new message to the Chat. If the name field is an empty
// string, the message will be displayed as a server message.
func (cl *Chat) AddMessage(name string, message string) {
	if name != "" {
		cl.Messages = append(cl.Messages, Message{name, message, false})
	} else {
		cl.Messages = append(cl.Messages, Message{"", message, true})
	}
}

// GetMessages returns the history of chat messages, the number of messages
// retrieved is determined by the passed quantity.
func (chatLog *Chat) GetMessages(quantity int) string {
	messages := "── Chat log ──\n"
	if len(chatLog.Messages) < quantity {
		for i := 0; i < quantity-len(chatLog.Messages); i++ {
			messages += "\n"
		}
		quantity = len(chatLog.Messages)
	}
	for _, message := range chatLog.Messages[len(chatLog.Messages)-quantity:] {
		if message.IsServer {
			messages += fmt.Sprintf("\033[2K\033[0;33m%s\033[0m\n", message.Content)
		} else {
			messages += fmt.Sprintf("\033[2K%s: %s\n", message.Name, message.Content)
		}
	}
	messages += "──────────────\n\n"
	return messages
}
