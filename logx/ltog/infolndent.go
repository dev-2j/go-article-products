package ltog

import "encoding/json"

func InfoIndent(v ...any) {

	vx, _ := json.MarshalIndent(v, "", "  ")
	Infoln(string(vx))

}
