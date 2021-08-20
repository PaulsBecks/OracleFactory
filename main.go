package main

import (
	"fmt"
	"net/http"

	"github.com/PaulsBecks/OracleFactory/src/models"
	"github.com/PaulsBecks/OracleFactory/src/routes"
	"github.com/PaulsBecks/OracleFactory/src/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDB() {
	//db, err := sql.Open("sqlite", c.filePath)
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Check if table exists - if not create it
	db.AutoMigrate(&models.EventParameter{},
		&models.OutboundOracleTemplate{},
		&models.EventValue{},
		&models.OutboundOracle{},
		&models.EventParameter{},
		&models.Event{},
		&models.InboundOracleTemplate{},
		&models.InboundOracle{},
		&models.User{},
		&models.Filter{},
		&models.ParameterFilter{},
	)
	models.InitFilter(db)
	user := models.User{
		Email:                       "test@example.com",
		Password:                    utils.HashAndSalt([]byte("test")),
		EthereumPrivateKey:          "b28c350293dcf09cc5b5a9e5922e2f73e48983fe8d325855f04f749b1a82e0e6",
		EthereumAddress:             "ws://eth-test-net:8545/",
		HyperledgerOrganizationName: "Org1MSP",
		HyperledgerChannel:          "mychannel",
		HyperledgerConfig: `---
name: test-network-org1
version: 1.0.0
client:
  organization: Org1
  connection:
    timeout:
      peer:
        endorser: '300'
organizations:
  Org1:
    mspid: Org1MSP
    peers:
    - peer0.org1.example.com
    certificateAuthorities:
    - ca.org1.example.com
peers:
  peer0.org1.example.com:
    url: grpcs://localhost:7051
    tlsCACerts:
      pem: |
          -----BEGIN CERTIFICATE-----
          MIICVzCCAf6gAwIBAgIRAIM7bskNalteuIlNrv2fUSwwCgYIKoZIzj0EAwIwdjEL
          MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
          cmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHzAdBgNVBAMTFnRs
          c2NhLm9yZzEuZXhhbXBsZS5jb20wHhcNMjEwNjIwMTEzNzAwWhcNMzEwNjE4MTEz
          NzAwWjB2MQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UE
          BxMNU2FuIEZyYW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEfMB0G
          A1UEAxMWdGxzY2Eub3JnMS5leGFtcGxlLmNvbTBZMBMGByqGSM49AgEGCCqGSM49
          AwEHA0IABOzBX4OXFcbbntfKKy35rcfOzX8iGr9t/b7e3dx5hydP1iDZroKhUju6
          Ex0b8ItapD5/An4/yDC9irKNnTGo8FGjbTBrMA4GA1UdDwEB/wQEAwIBpjAdBgNV
          HSUEFjAUBggrBgEFBQcDAgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zApBgNV
          HQ4EIgQg2JgDo3wk6UtVNLItks9GYLYJd6pBnMYeEKpwHbRjBaAwCgYIKoZIzj0E
          AwIDRwAwRAIgdJQ6Z8XktWPnFN0BiQpBOzawJJku+2q+alc1hSDvm3ECIElVaGKp
          JYX9cTupGaWVDuAjBSMmZNzRzl6SlqCzKJaE
          -----END CERTIFICATE-----
          
    grpcOptions:
      ssl-target-name-override: peer0.org1.example.com
      hostnameOverride: peer0.org1.example.com
certificateAuthorities:
  ca.org1.example.com:
    url: https://localhost:7054
    caName: ca-org1
    tlsCACerts:
      pem: 
        - |
          -----BEGIN CERTIFICATE-----
          MIICUTCCAfegAwIBAgIQFaq5hC9hTI1KnZy3thUJ9zAKBggqhkjOPQQDAjBzMQsw
          CQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMNU2FuIEZy
          YW5jaXNjbzEZMBcGA1UEChMQb3JnMS5leGFtcGxlLmNvbTEcMBoGA1UEAxMTY2Eu
          b3JnMS5leGFtcGxlLmNvbTAeFw0yMTA2MjAxMTM3MDBaFw0zMTA2MTgxMTM3MDBa
          MHMxCzAJBgNVBAYTAlVTMRMwEQYDVQQIEwpDYWxpZm9ybmlhMRYwFAYDVQQHEw1T
          YW4gRnJhbmNpc2NvMRkwFwYDVQQKExBvcmcxLmV4YW1wbGUuY29tMRwwGgYDVQQD
          ExNjYS5vcmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE
          5CD1mcpY45dFV5hcnljG/dtviMDmVNaZHBDI6jdFTuT1yC2+C+twhZEq6MznehHg
          CoDfUafUSYMySJuPsyQs6aNtMGswDgYDVR0PAQH/BAQDAgGmMB0GA1UdJQQWMBQG
          CCsGAQUFBwMCBggrBgEFBQcDATAPBgNVHRMBAf8EBTADAQH/MCkGA1UdDgQiBCBC
          bSNDQTUmTPPo6LDFLM4w6LqDu1F69rThXnDfOKO/rTAKBggqhkjOPQQDAgNIADBF
          AiAcbbHuPmUyCChk3nPDTpvwTwyvWY4zvF5mX/u6esC3qQIhALW+0ri4JtD6V5aJ
          68WCYM2g2Sw67bxjt2g7E7S7nkJI
          -----END CERTIFICATE-----
          
    httpOptions:
      verify: false

`,
		HyperledgerCert: `-----BEGIN CERTIFICATE-----
MIICKzCCAdGgAwIBAgIRAKB3+P9722pCZP9nUeRrzO0wCgYIKoZIzj0EAwIwczEL
MAkGA1UEBhMCVVMxEzARBgNVBAgTCkNhbGlmb3JuaWExFjAUBgNVBAcTDVNhbiBG
cmFuY2lzY28xGTAXBgNVBAoTEG9yZzEuZXhhbXBsZS5jb20xHDAaBgNVBAMTE2Nh
Lm9yZzEuZXhhbXBsZS5jb20wHhcNMjEwNjIwMTEzNzAwWhcNMzEwNjE4MTEzNzAw
WjBsMQswCQYDVQQGEwJVUzETMBEGA1UECBMKQ2FsaWZvcm5pYTEWMBQGA1UEBxMN
U2FuIEZyYW5jaXNjbzEPMA0GA1UECxMGY2xpZW50MR8wHQYDVQQDDBZVc2VyMUBv
cmcxLmV4YW1wbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEYlmVKxL5
5TF6xmxzmk+ZmpF1/3y1BJJYZMR0s+BME4fl1cmKQqqbk7M22kzxCwQLZhT0rHWn
mDmvKbpsuxxCOKNNMEswDgYDVR0PAQH/BAQDAgeAMAwGA1UdEwEB/wQCMAAwKwYD
VR0jBCQwIoAgQm0jQ0E1Jkzz6OiwxSzOMOi6g7tReva04V5w3zijv60wCgYIKoZI
zj0EAwIDSAAwRQIhALu7KdaCRt8m1C5NdEXeEFL8HakaGbjLJY3vXGnbMBxuAiBw
q6zrYSyGMewuCV/xNad5oy1btMsMyYmqHjY6ngIDpw==
-----END CERTIFICATE-----
`,
		HyperledgerKey: `-----BEGIN PRIVATE KEY-----
MIGHAgEAMBMGByqGSM49AgEGCCqGSM49AwEHBG0wawIBAQQg5caOC5+nuPUGWby9
1D1hWj+cG91qHfuZaSuaAgTpFfmhRANCAARiWZUrEvnlMXrGbHOaT5makXX/fLUE
klhkxHSz4EwTh+XVyYpCqpuTszbaTPELBAtmFPSsdaeYOa8pumy7HEI4
-----END PRIVATE KEY-----
`,
	}

	db.Create(&user)
	oracleTemplate := models.OracleTemplate{
		BlockchainName:  "Hyperledger",
		EventName:       "TransferAsset",
		ContractAddress: "basic",
	}
	db.Create(&oracleTemplate)
	inboundOracleTemplate := models.InboundOracleTemplate{OracleTemplate: oracleTemplate}
	db.Create(&inboundOracleTemplate)
	oracle := models.Oracle{
		Name:   "Hyperledger Test",
		UserID: user.ID,
	}
	db.Create(&oracle)
	inboundOracle := models.InboundOracle{
		Oracle:                  oracle,
		InboundOracleTemplateID: inboundOracleTemplate.ID,
	}
	db.Create(&inboundOracle)
	eventParameter := models.EventParameter{
		Name:             "asset1",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameter)
	outboundOracleTemplate := models.OutboundOracleTemplate{
		OracleTemplate: oracleTemplate,
	}
	db.Create(&outboundOracleTemplate)
	eventParameterOut := models.EventParameter{
		Name:             "owner",
		Type:             "string",
		OracleTemplateID: oracle.ID,
	}
	db.Create(&eventParameterOut)
}

func auth(ctx *gin.Context) {
	db, err := gorm.Open(sqlite.Open("./OracleFactory.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	authHeader := ctx.GetHeader("Authorization")
	tokenString := authHeader[len("Bearer "):]
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if _, isvalid := t.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", t.Header["alg"])
		}
		return []byte(utils.JWT_SECRET), nil
	})
	if err == nil && token.Valid {
		claims := token.Claims.(jwt.MapClaims)
		userID := uint(claims["id"].(float64))
		var user models.User
		db.First(&user, userID)
		ctx.Set("user", user)
	} else {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}

func main() {
	app := gin.Default()
	middleware(app)
	InitDB()

	app.POST("/users/login", routes.Login)
	app.POST("/users/signup", routes.Register)

	authorized := app.Group("/", auth)
	{
		authorized.GET("/outboundOracles", routes.GetOutboundOracles)
		authorized.GET("/outboundOracles/:outboundOracleId", routes.GetOutboundOracle)
		authorized.PUT("/outboundOracles/:outboundOracleId", routes.UpdateOutboundOracle)
		authorized.DELETE("/outboundOracles/:outboundOracleId", routes.DeleteOutboundOracle)
		authorized.POST("/outboundOracles/:outboundOracleId/events", routes.PostOutboundOracleEvent)

		authorized.GET("/outboundOracleTemplates", routes.GetOutboundOracleTemplates)
		authorized.POST("/outboundOracleTemplates", routes.PostOutboundOracleTemplate)
		authorized.GET("/outboundOracleTemplates/:outboundOracleTemplateID", routes.GetOutboundOracleTemplate)
		authorized.POST("/outboundOracleTemplates/:outboundOracleTemplateID/outboundOracles", routes.PostOutboundOracle)
		authorized.POST("/outboundOracleTemplates/:outboundOracleTemplateID/eventParameters", routes.PostOutboundEventParameters)

		authorized.GET("/inboundOracles/:inboundOracleId", routes.GetInboundOracle)
		authorized.PUT("/inboundOracles/:inboundOracleId", routes.UpdateInboundOracle)
		authorized.POST("/inboundOracles/:inboundOracleID/events", routes.PostInboundOracleEvent)
		authorized.GET("/inboundOracleTemplates/:inboundOracleTemplateID", routes.GetInboundOracleTemplate)
		authorized.GET("/inboundOracles", routes.GetInboundOracles)

		authorized.POST("/inboundOracleTemplates/:inboundOracleTemplateID/inboundOracles", routes.PostInboundOracle)
		authorized.POST("/inboundOracleTemplates/:inboundOracleTemplateID/eventParameters", routes.PostInboundEventParameters)
		authorized.GET("/inboundOracleTemplates", routes.GetInboundOracleTemplates)
		authorized.POST("/inboundOracleTemplates", routes.PostInboundOracleTemplate)

		authorized.GET("/filters", routes.GetFilters)

		authorized.GET("/oracles/:oracleID/parameterFilters", routes.GetOracleParameterFilters)
		authorized.POST("/oracles/:oracleID/parameterFilters", routes.PostOracleParameterFilters)
		authorized.DELETE("/oracles/:oracleID/parameterFilters/:parameterFilterID", routes.DeleteOracleParameterFilter)

		authorized.GET("/oracleTemplates/:oracleTemplateID/eventParameters", routes.GetOracleTemplateEventParameters)

		authorized.GET("/user", routes.GetCurrentUserDetail)
		authorized.PUT("/user", routes.UpdateCurrentUser)
	}
	app.Run()
}
