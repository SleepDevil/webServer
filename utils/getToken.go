package utils

import (
	"fmt"

	rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
)

func GenerateToken(channelName string, uidStr string) (toekn string) {

	appID := "a4e417b3d58e4642bac026c629a07e4d"
	appCertificate := "e38643dae4b44075be3376b528c57c80"

	result, err := rtctokenbuilder.BuildTokenWithUID(appID, appCertificate, channelName, 0, rtctokenbuilder.RoleAttendee, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Token with uid: %s\n", result)
	}
	return result
}
