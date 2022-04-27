package discordac

import (
	"github.com/bwmarrin/discordgo"
	"github.com/sirupsen/logrus"
)

// CommandContext is a command execution context to execute it around
// It includes everything which is related and is useful to command execution
type CommandContext struct {
	// Session is a connection to Discord API
	Session *discordgo.Session
	// Interaction is what issued the command to execute
	Interaction *discordgo.Interaction
	// Options are provided command options to use during execution
	Options map[string]*discordgo.ApplicationCommandInteractionDataOption
}

// FollowupMessage TODO: move this inside CommandContext
type FollowupMessage struct {
	CommandContext *CommandContext
	Message        *discordgo.Message
}

// parseOptions adds all specified options to a list CommandContext.Options for them to be available to be easily used
// during command execution
func (cc *CommandContext) parseOptions() {
	options := cc.Interaction.ApplicationCommandData().Options
	cc.Options = make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		cc.Options[opt.Name] = opt
	}
}

func (cc *CommandContext) Respond(content string) {
	err := cc.Session.InteractionRespond(cc.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		logrus.WithError(err).Errorf("Couldn't respond inside context %v with a content: %v", cc, content)
	}
}

func (cc *CommandContext) EditResponse(content string) (err error) {
	_, err = cc.Session.InteractionResponseEdit(cc.Interaction, &discordgo.WebhookEdit{
		Content: content,
	})

	return
}

func (cc *CommandContext) DeleteResponse() (err error) {
	err = cc.Session.InteractionResponseDelete(cc.Interaction)

	return
}

func (cc *CommandContext) FollowupCreate(content string) (FollowupMessage, error) {
	msg, err := cc.Session.FollowupMessageCreate(cc.Interaction, true, &discordgo.WebhookParams{
		Content: content,
	})

	if err != nil {
		return FollowupMessage{Message: msg}, err
	}

	return FollowupMessage{CommandContext: cc, Message: msg}, nil
}

func (cc *CommandContext) GetOption(option string) (value *discordgo.ApplicationCommandInteractionDataOption, ok bool) {
	value, ok = cc.Options[option]

	return
}

func (fm FollowupMessage) Edit(content string) (err error) {
	fm.Message, err = fm.CommandContext.Session.FollowupMessageEdit(fm.CommandContext.Interaction, fm.Message.ID, &discordgo.WebhookEdit{
		Content: content,
	})

	return
}

func (fm FollowupMessage) Delete() (err error) {
	err = fm.CommandContext.Session.FollowupMessageDelete(fm.CommandContext.Interaction, fm.Message.ID)

	return
}
