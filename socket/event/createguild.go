package event

import (
	"github.com/kataras/golog"
	"github.com/mitchellh/mapstructure"
	"github.com/thanhpk/randstr"
	"harmony-server/harmonydb"
	"harmony-server/socket"
)

type createGuildData struct {
	Token string `mapstructure:"token"`
	Guild string `mapstructure:"guild"`
}

func OnCreateGuild(ws *socket.Client, rawMap map[string]interface{}) {
	var data createGuildData
	if err := mapstructure.Decode(rawMap, &data); err != nil {
		return
	}
	userid := VerifyToken(data.Token)
	if userid == "" {
		deauth(ws)
		return
	}
	guildid := randstr.Hex(16)
	_, err := harmonydb.DBInst.Exec(`INSERT INTO guilds(guildid, guildname, picture, owner) VALUES($1, $2, $3, $4); 
										   INSERT INTO guildmembers(userid, guildid) VALUES($5, $6);
										   INSERT INTO channels(channelid, guildid, channelname) VALUES($7, $8, $9)`, guildid, data.Guild, "", userid, userid, guildid, randstr.Hex(16), guildid, "general")
	if err != nil {
		golog.Warnf("Error creating guild : %v", err)
		ws.Send(&socket.Packet{
			Type: "createguild",
			Data: map[string]interface{}{
				"message": "error creating guild",
			},
		})
		return
	}
	ws.Send(&socket.Packet{
		Type: "createguild",
		Data: map[string]interface{}{
			"guild": guildid,
		},
	})
}
