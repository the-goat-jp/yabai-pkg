package main

import (
	"context"
	"log"

	jwthelper "github.com/the-goat-jp/yabai-pkg/jwt_helper"
)

func main() {
	jwtInput := `eyJhbGciOiJSUzI1NiIsImtpZCI6ImQ2YTZmMzc1LThlZmMtNDY3My1iNjliLWNjODIyOTEyOGM3MiIsInR5cCI6IkpXVCJ9.eyJleHAiOjM1MTQ1Nzk2MjksImlhdCI6MTcxNDU3OTYyOSwiaXNzIjoieWFiYWktdXNlci1zZXJ2aWNlIiwibmFtZSI6IlZpYW4gWiIsInJlZiI6IjZSOHVKaSIsInN1YiI6IjUxMTUwNzkxODkzNjM0NzE0NCIsInVybCI6Imh0dHA6Ly9sb2NhbGhvc3Q6OTAwMC95YWJhaS1wcm9maWxlLWRldi81MTE1MDc5MTg5MzYzNDcxNDRfcHJvZmlsZV9pbWFnZS5qcGcifQ.V6ntuUB2-rg37x8u5RAihDv2xvZUmfJv861mvZxF5ilyrrGixMe5v6Ha2fHBJR81rXBauNA5mSntu4rs2mvULC0AFnEdzbxTgh3Y3xpTTyK638RmjaZj6iErErBq5_QbCvcpjC2XRSZ_g-aq3BzV21W0_IWiMp4g6Jvqem2zLpdYH8gnCOfYXEz2xuMxhGC1JP2xqLl9DpFVv7o8qaRQsDeb3qImUD5v9ni0edV6Q_80zNQ3N4Sbo0zVfKAgeI7dpTyNynuozoDQPunFPcL1xzeiNa13lWfRR4izGKVYwbbxnFvZjs02pMM1AiKFjqjPX_E8g-X5FImNoGHVSoWpYw`

	parser, err := jwthelper.NewJWTParser("http://localhost:8080")
	if err != nil {
		log.Panicln("failed to create jwt parser:", err)
	}

	jwtData, err := parser.Parse(context.Background(), jwtInput)
	if err != nil {
		log.Panicln("failed to parse jwt:", err)
	}

	log.Printf("jwtData:%+v", jwtData)
}
