package service

import "testing"

func TestValidateChannel(t *testing.T) {
	tests := map[string]Channel{
		"feishu":   ChannelFeishu,
		"Feishu":   ChannelFeishu,
		"FEISHU":   ChannelFeishu,
		"telegram": ChannelTelegram,
		"Telegram": ChannelTelegram,
		"TELEGRAM": ChannelTelegram,
	}
	for name, want := range tests {
		channel, err := ValidateChannel(name)
		if err != nil {
			t.Fatalf("ValidateChannel(%q) error = %v", name, err)
		}
		if channel != want {
			t.Fatalf("ValidateChannel(%q) = %q, want %q", name, channel, want)
		}
	}

	for _, name := range []string{"", "lark", "Lark", "wechat", "telegram-bot"} {
		if _, err := ValidateChannel(name); err == nil {
			t.Fatalf("ValidateChannel(%q) error = nil", name)
		}
	}
}
