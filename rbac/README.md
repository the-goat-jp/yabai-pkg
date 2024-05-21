# RBAC Helper

Centralized Role ID checker

```go
	jwtInput := `eyJhbGciOiJSUzI1NiIsImtpZCI6IjJiOTI4M2VlLTlhOTctNDA0Yy05OTQzLTcxOTUyYTU1MGNmOCIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RhY2NvdW50MUBnbWFpbC5jb20iLCJleHAiOjE3MTQzMjYzNTksImlhdCI6MTcxNDMyNjA1OSwiaWQiOjEsImlzcyI6InlhYmFpLXVzZXItc2VydmljZSIsIm5hbWUiOiJ0ZXN0YWNjb3VudDFAZ21haWwuY29tIiwicmVmIjoicGhKVnhwIiwidXJsIjoiIn0.RxNpK8Gh86W3uV1VgIILMRRtwikGjTPsvG335CE2BIDlClvejM2PHNVOovu69XaIMFI_y79ttviv81gOoLMsZm7wR3HmHdgyTaXUmD8oSM0qMuSi8a7-teXaSSees_CX7ORMzPD1lgFJ8NLpSpdumg7_IXmXG4r3cbWEO92aaiV3j8nNojhro4jyd2n5a1HPAE-Lx5vk0pJY9sPebc0mVWXB9Ro5WZKVFwjVZ1iY9TaN6nLko2vLKbBw4PrpmVe5spOWtqu5UUA9A_sjrRjnnQ9Fe7ufBUpIbmoAvmGGkOR-AZAEuc6xJJLX9_qxVvCayZdiAv6iqx_bXNQDUMmKYg`

	parser, err := helper.NewJWTParser("http://localhost:8080")
	if err != nil {
		log.Panicln("failed to create jwt parser:", err)
	}

	jwtData, err := parser.Parse(context.Background(), jwtInput)
	if err != nil {
		log.Panicln("failed to parse jwt:", err)
	}

	if IsPaymentTesterByJWTData(jwtData) {
		// use test mode payment system
	}
```
