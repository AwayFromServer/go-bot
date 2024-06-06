package bot

import (
	"testing"

	"github.com/bwmarrin/discordgo"
	"github.com/stretchr/testify/assert"
)

const TESTFILE = "../config_test.yaml"

func TestConfigs(t *testing.T) {
	testdata := []struct {
		name     string
		expected conf
		actual   conf
	}{
		{
			"no overrides",
			conf{"foo", "bar", "baz"},
			conf{},
		},
		{
			"overrides",
			conf{"qax", "qix", "qux"},
			conf{"qax", "qix", "qux"},
		}, //
		{
			"error thrown",
			conf{"pap", "pep", "pip"},
			conf{"pap", "pep", "pip"},
		},
		// {"", "", ""}, // new test case
	}

	for _, subtest := range testdata {
		t.Run(subtest.name, func(t *testing.T) {
			t.Setenv(BT, subtest.actual.BotToken)
			t.Setenv(TU, subtest.actual.BotTarget)
			t.Setenv(BP, subtest.actual.BotPrefix)

			var c conf
			c.getConf(TESTFILE)
			c.getOverrides()

			assert.Equal(t, subtest.expected.BotToken, c.BotToken)
			assert.Equal(t, subtest.expected.BotTarget, c.BotTarget)
			assert.Equal(t, subtest.expected.BotPrefix, c.BotPrefix)

		})
	}
}

func TestStartSession(t *testing.T) {
	var c conf
	c.getConf(CFGFILE)
	b := Bot{config: c}

	b.startSession()

	assert.NotEqual(t, nil, b.session)

	m := &discordgo.MessageCreate{}

	m.Content = ""
	m.Author.ID = b.session.State.SessionID
	err := b.newMessage(b.session, m)
	assert.Equal(t, nil, err)

}
