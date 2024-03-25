package netsuiteodbc

type Config struct {
	// ConnectionString is the base connection string to use to connect to ODBC driver - excluding any credential information
	ConnectionString string
	// ConsumerKey from the integration used to connect to NetSuite
	ConsumerKey string
	// ConsumerKey from the integration used to connect to NetSuite
	ConsumerSecret string
	// TokenId from the integration used to connect to NetSuite
	TokenId string
	// TokenSecret from the integration used to connect to NetSuite
	TokenSecret string
	// AccountId to use to connect with.  i.e. 123456_SB1
	AccountId string
}
