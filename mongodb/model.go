package mongodb

type Stock struct {
	SimpleName               string
	Code                     string
	FullName                 string
	EnglishName              string
	OldName                  string
	PublishDate              int
	ForIndustry              string
	ForStockConcept          []string
	Area                     string
	LegalOwner               string
	IndependentDirector      []string
	AdvisoryOrga             string
	AccountingOrga           string
	SecuritiesRepresentative string
	Capital                  float32
	RegisterAddr             string
	RateOfTax                float32
	BusinessAddr             string
	MainProduct              string
	IssueDate                int
	OpenSaleDate             int
	WhichMarket              string
	SecurityType             string
	OutstandingCapitalStock  float32
	TotalCapitalStock        float32
	SallerAgent              string
	IssuePrice               float32
	OpenSaleOpenPrice        float32
	OpenSalePriceRate        float32
	OpenSaleExchangeRage     float32
	SpecialIssue             string
	IssuePE                  float32
	CurrentPE                float32
	Tel                      string
	Email                    string
	WebSite                  string
	ContactPerson            string
}

// ------------------------------

type HolderStruct struct {
	Name       string
	Vol        int64
	Percentage float32
}

type MainStruct struct {
	StopDate   int
	NumHolder  int
	EvenVol    int
	MainHolder []HolderStruct
}

type StockMainHolder struct {
	Code      string
	FetchDate int
	Main      []MainStruct
}

// ------------------------------

type PublicStruct struct {
	StopDate     int
	PublicHolder []HolderStruct
}

type StockPublicHolder struct {
	Code      string
	FetchDate int
	Public    []PublicStruct
}

// ------------------------------

type StockBigDeal struct {
	Code      string
	FetchDate int
	BigDeal   []BigDealStruct
}

type BigDealStruct struct {
	Date   int
	Deal   []float32 // price, vol(万), money(万)
	Buyer  string
	Saller string
}

// ------------------------------

type StockMarginTrading struct {
	Code          string
	FetchDate     int
	MarginTrading [][7]int64
}

// ------------------------------

type FundingStruct struct {
	Name string
	Type string
	Vol  int
}

type FundingDetailStruct struct {
	Name      string
	Code      string
	VolDetail []int
}

type StockFunding struct {
	Code             string
	Date             int
	FetchDate        int
	FundingCounting  []int
	VolCounting      []int
	TenFundingChange []FundingStruct
	FundingDetail    []FundingDetailStruct
}

// ------------------------------

type StockAccounting struct {
	Code            string
	Date            int
	IndustryClass   []string
	Ranking         []int
	StopDate        int
	EarningPerShare []float32 // 每股收益
	Naps            []float32 // 每股净资产
	SalesRevenue    []float32 // 主营收入
	Profit          []float32 // 净利润
	GrossProfitSale []float32 // 销售毛利率
	Others          []float32 // 每股资本公积金、每股未分配利润、净资产收益率
}

// ------------------------------

type FifteenMinuteTran struct {
	Time     string
	Count    []int // tranCount, priceChangeCount
	VolMoney []int64
}

type EachPriceTran struct {
	Time     string
	Price    float32
	VolMoney []int64
}

type StockTran struct {
	Date       int
	KeyPrices  []float32 // last, open, close, highest, lowest
	Count      []int     // priceCount, tranCount
	VolMoney   []int64
	Fifteen    []FifteenMinuteTran
	EachPrices []EachPriceTran
}
