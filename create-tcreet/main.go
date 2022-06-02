package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/hashgraph/hedera-sdk-go/v2"
	"github.com/joho/godotenv"
)

func CreateTestCreet(client *hedera.Client, acc hedera.AccountID, tknName, tknSym string, tknDec uint) {

	var tknInitTxn *hedera.TokenCreateTransaction = hedera.NewTokenCreateTransaction()

	// Admin Key
	admnPrivk, err := hedera.PrivateKeyGenerateEd25519()
	if err != nil {
		panic("Error generating new admin key.")
	}
	var admnPubk hedera.PublicKey = admnPrivk.PublicKey()

	// Treasury Key
	trsyPrivk, err := hedera.PrivateKeyGenerateEd25519()
	if err != nil {
		panic("Error generating treasury key.")
	}
	var trsyPubk hedera.PublicKey = trsyPrivk.PublicKey()

	// Set token transaction parameters.
	tknInitTxn.SetTokenName(tknName)
	tknInitTxn.SetTokenSymbol(tknSym)
	tknInitTxn.SetTreasuryAccountID(acc)
	tknInitTxn.SetDecimals(tknDec)
	tknInitTxn.SetInitialSupply(1000000000000000)
	tknInitTxn.SetTokenType(hedera.TokenTypeFungibleCommon)
	tknInitTxn.SetAdminKey(admnPubk)
	tknInitTxn.SetMaxTransactionFee(hedera.NewHbar(10))

	_, err = tknInitTxn.FreezeWith(client)
	if err != nil {
		panic("Error intiliasing transaction.")
	}

	txnResp, err := tknInitTxn.Sign(admnPrivk).Sign(trsyPrivk).Execute(client)
	if err != nil {
		panic(err)
	}

	rec, err := txnResp.GetReceipt(client)
	if err != nil {
		panic(err)
	}

	fmt.Println("Admin Private Key, ", admnPrivk)
	fmt.Println("Admin Public Key, ", admnPubk)
	fmt.Println("Treasury Private Key, ", trsyPrivk)
	fmt.Println("Treasury Public Key, ", trsyPubk)
	fmt.Printf("%s's (%s) TokenID is: %s\n", tknName, tknSym, rec.TokenID)
}

func main() {

	err := godotenv.Load("../.env")
	if err != nil {
		panic("Couldn't load .env file.")
	}

	acc, err := hedera.AccountIDFromString(os.Getenv("HEDERA_ACCOUNT_ID"))
	if err != nil {
		panic("Error parsing AccountID.")
	}

	privk, err := hedera.PrivateKeyFromString(os.Getenv("HEDERA_ACCOUNT_PRIVATE_KEY"))
	if err != nil {
		panic("Error parsing Account Private Key.")
	}

	/*pubk, err := hedera.PublicKeyFromString(os.Getenv("HEDERA_ACCOUNT_PUBLIC_KEY"))
	if err != nil {
		panic("Error parsing Account Public Key.")
	}
	tknIdStr := os.Getenv("HEDERA_CREET_TOKEN_ADDRESS")
	var tknId hedera.TokenID
	if len(tknIdStr) > 0 {
		tknId, err = hedera.TokenIDFromString(tknIdStr)
		if err != nil {
			panic("Error parsing CREET TokenID.")
		}
	}*/

	var tknName string = os.Getenv("HEDERA_CREET_TOKEN_NAME")
	var tknSym string = os.Getenv("HEDERA_CREET_TOKEN_SYMBOL")

	tknDec, err := strconv.ParseUint(os.Getenv("HEDERA_CREET_TOKEN_DECIMALS"), 10, 32)
	if err != nil {
		panic("Error parsing Token decimal precision.")
	}

	var client *hedera.Client = hedera.ClientForTestnet()
	client.SetOperator(acc, privk)

	CreateTestCreet(client, acc, tknName, tknSym, uint(tknDec))
}
