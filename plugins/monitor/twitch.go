package monitor

import "github.com/fzxiao233/Vtb_Record/utils"

type Twitch struct {
	APIUrl string
}

func (t Twitch) getLiveStatus() error {
	_, err := utils.HttpGet(t.APIUrl, map[string]string{})
	if err != nil {
		return err
	}
	return nil
}

func (t Twitch) CheckLive(userConfig utils.UsersConfig) {
	t.APIUrl = "https://api.twitch.tv/helix/streams?user_login=" + userConfig.TargetId
}
