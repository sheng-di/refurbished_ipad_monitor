package applebuyutil

type Info struct {
	Dictionaries        Dictionaries         `json:"dictionaries"`
	Dimensions          []Dimension          `json:"dimensions"`
	Products            []Product            `json:"products"`
	SelectedGridFilters []SelectedGridFilter `json:"selectedGridFilters"`
	SortData            SortData             `json:"sortData"`
	StaticAssets        StaticAssets         `json:"staticAssets"`
	Tiles               []Tile               `json:"tiles"`
}

type Image struct {
	Alt               string    `json:"alt"`
	Attrs             string    `json:"attrs"`
	DeferSrc          bool      `json:"deferSrc"`
	Height            string    `json:"height"`
	ImageName         string    `json:"imageName"`
	NoImage           bool      `json:"noImage"`
	OriginalImageName string    `json:"originalImageName"`
	ScaleFactor       string    `json:"scaleFactor"`
	SrcSet            Foo_sub16 `json:"srcSet"`
	Width             string    `json:"width"`
}

type Foo_sub19 struct {
	Amount    string `json:"amount"`
	RawAmount string `json:"raw_amount"`
}

type Price struct {
	BasePartNumber              string    `json:"basePartNumber"`
	ChooseDefaultPurchaseOption bool      `json:"chooseDefaultPurchaseOption"`
	CurrentPrice                Foo_sub19 `json:"currentPrice"`
	PartNumber                  string    `json:"partNumber"`
	PriceCurrency               string    `json:"priceCurrency"`
	PriceFeeDisclaimer          string    `json:"priceFeeDisclaimer"`
	RefurbProduct               bool      `json:"refurbProduct"`
	SeoPrice                    float64   `json:"seoPrice"`
	ShowDynamicFinancing        bool      `json:"showDynamicFinancing"`
	ShowItemPropAvailability    bool      `json:"showItemPropAvailability"`
	ShowItemPropPrice           bool      `json:"showItemPropPrice"`
	ShowPayPal                  bool      `json:"showPayPal"`
}

type OmnitureModel struct {
	BasePartNumber       string  `json:"basePartNumber"`
	CommitCodeID         float64 `json:"commitCodeId"`
	CustomerCommitString string  `json:"customerCommitString"`
	FeatureName          string  `json:"featureName"`
	LinkText             string  `json:"linkText"`
	PartNumber           string  `json:"partNumber"`
	SlotName             string  `json:"slotName"`
}

type SortData struct {
	Data Foo_sub13 `json:"data"`
}

type Dimensions struct {
	DimensionCapacity     DimensionCapacity     `json:"dimensionCapacity"`
	DimensionColor        DimensionColor        `json:"dimensionColor"`
	Dimensionconnectivity Dimensionconnectivity `json:"dimensionconnectivity"`
	RefurbClearModel      RefurbClearModel      `json:"refurbClearModel"`
}

type Foo_sub9 struct {
	DimensionCapacity     string `json:"dimensionCapacity"`
	DimensionColor        string `json:"dimensionColor"`
	Dimensionconnectivity string `json:"dimensionconnectivity"`
	RefurbClearModel      string `json:"refurbClearModel"`
}

type Dictionaries struct {
	Dimensions Dimensions `json:"dimensions"`
}

type Product struct {
	Dimensions Foo_sub9 `json:"dimensions"`
}

type Tile struct {
	Filters           Product       `json:"filters"`
	Image             Image         `json:"image"`
	Lob               string        `json:"lob"`
	OmnitureModel     OmnitureModel `json:"omnitureModel"`
	PartNumber        string        `json:"partNumber"`
	Price             Price         `json:"price"`
	ProductDetailsURL string        `json:"productDetailsUrl"`
	Sort              Sort          `json:"sort"`
	SortOrder         float64       `json:"sortOrder"`
	Title             string        `json:"title"`
}

type DimensionColor struct {
	Gold      Foo_sub1 `json:"gold"`
	Rosegold  Foo_sub1 `json:"rosegold"`
	Silver    Foo_sub1 `json:"silver"`
	SpaceGray Foo_sub1 `json:"space_gray"`
}

type RefurbClearModel struct {
	Ipad2017        Foo_sub1 `json:"ipad2017"`
	Ipad5gen        Foo_sub1 `json:"ipad5gen"`
	Ipadaccessories Foo_sub1 `json:"ipadaccessories"`
	Ipadmini4       Foo_sub1 `json:"ipadmini4"`
	Ipadmini5       Foo_sub1 `json:"ipadmini5"`
	Ipadpro10_5     Foo_sub1 `json:"ipadpro_10_5"`
	Ipadpro11       Foo_sub1 `json:"ipadpro_11"`
	Ipadpro12_9     Foo_sub1 `json:"ipadpro_12_9"`
	Ipadpro9_7      Foo_sub1 `json:"ipadpro_9_7"`
}

type Foo_sub13 struct {
	Key     string      `json:"key"`
	Legend  string      `json:"legend"`
	Options []Foo_sub12 `json:"options"`
}

type Dimension struct {
	Key       string  `json:"key"`
	Legend    string  `json:"legend"`
	SortOrder float64 `json:"sortOrder"`
}

type DimensionCapacity struct {
	One28gb  Foo_sub1 `json:"128gb"`
	OneTb    Foo_sub1 `json:"1tb"`
	Two56gb  Foo_sub1 `json:"256gb"`
	Three2gb Foo_sub1 `json:"32gb"`
	Five12gb Foo_sub1 `json:"512gb"`
	Six4gb   Foo_sub1 `json:"64gb"`
}

type Sort struct {
	PriceAsc  float64 `json:"priceAsc"`
	PriceDesc float64 `json:"priceDesc"`
}

type SelectedGridFilter struct {
	RefurbClearModel string `json:"refurbClearModel"`
}

type Foo_sub16 struct {
	ScaleParams1 string `json:"scaleParams1"`
	Src          string `json:"src"`
}

type Foo_sub12 struct {
	Selected bool   `json:"selected"`
	Text     string `json:"text"`
	Value    string `json:"value"`
}

type Foo_sub1 struct {
	SortOrder float64 `json:"sortOrder"`
	Text      string  `json:"text"`
}

type Dimensionconnectivity struct {
	Wifi     Foo_sub1 `json:"wifi"`
	Wificell Foo_sub1 `json:"wificell"`
}

type StaticAssets struct{}
