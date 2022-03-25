package main

type GetAssetsRsp struct {
	Assets []Assets `json:"assets"`
}

type AssetContract struct {
	Address                     string      `json:"address"`
	AssetContractType           string      `json:"asset_contract_type"`
	CreatedDate                 string      `json:"created_date"`
	Name                        string      `json:"name"`
	NftVersion                  interface{} `json:"nft_version"`
	OpenseaVersion              string      `json:"opensea_version"`
	Owner                       int         `json:"owner"`
	SchemaName                  string      `json:"schema_name"`
	Symbol                      string      `json:"symbol"`
	TotalSupply                 interface{} `json:"total_supply"`
	Description                 string      `json:"description"`
	ExternalLink                interface{} `json:"external_link"`
	ImageURL                    interface{} `json:"image_url"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	DevBuyerFeeBasisPoints      int         `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     int         `json:"dev_seller_fee_basis_points"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  int         `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints int         `json:"opensea_seller_fee_basis_points"`
	BuyerFeeBasisPoints         int         `json:"buyer_fee_basis_points"`
	SellerFeeBasisPoints        int         `json:"seller_fee_basis_points"`
	PayoutAddress               interface{} `json:"payout_address"`
}

type DisplayData struct {
	CardDisplayStyle string `json:"card_display_style"`
}

type Collection struct {
	BannerImageURL              string      `json:"banner_image_url"`
	ChatURL                     interface{} `json:"chat_url"`
	CreatedDate                 string      `json:"created_date"`
	DefaultToFiat               bool        `json:"default_to_fiat"`
	Description                 string      `json:"description"`
	DevBuyerFeeBasisPoints      string      `json:"dev_buyer_fee_basis_points"`
	DevSellerFeeBasisPoints     string      `json:"dev_seller_fee_basis_points"`
	DiscordURL                  interface{} `json:"discord_url"`
	DisplayData                 DisplayData `json:"display_data"`
	ExternalURL                 string      `json:"external_url"`
	Featured                    bool        `json:"featured"`
	FeaturedImageURL            string      `json:"featured_image_url"`
	Hidden                      bool        `json:"hidden"`
	SafelistRequestStatus       string      `json:"safelist_request_status"`
	ImageURL                    string      `json:"image_url"`
	IsSubjectToWhitelist        bool        `json:"is_subject_to_whitelist"`
	LargeImageURL               string      `json:"large_image_url"`
	MediumUsername              interface{} `json:"medium_username"`
	Name                        string      `json:"name"`
	OnlyProxiedTransfers        bool        `json:"only_proxied_transfers"`
	OpenseaBuyerFeeBasisPoints  string      `json:"opensea_buyer_fee_basis_points"`
	OpenseaSellerFeeBasisPoints string      `json:"opensea_seller_fee_basis_points"`
	PayoutAddress               string      `json:"payout_address"`
	RequireEmail                bool        `json:"require_email"`
	ShortDescription            interface{} `json:"short_description"`
	Slug                        string      `json:"slug"`
	TelegramURL                 interface{} `json:"telegram_url"`
	TwitterUsername             interface{} `json:"twitter_username"`
	InstagramUsername           interface{} `json:"instagram_username"`
	WikiURL                     interface{} `json:"wiki_url"`
	IsNsfw                      bool        `json:"is_nsfw"`
}
type User struct {
	Username string `json:"username"`
}
type Owner struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
type Creator struct {
	User          User   `json:"user"`
	ProfileImgURL string `json:"profile_img_url"`
	Address       string `json:"address"`
	Config        string `json:"config"`
}
type Assets struct {
	ID                      int           `json:"id"`
	NumSales                int           `json:"num_sales"`
	BackgroundColor         interface{}   `json:"background_color"`
	ImageURL                string        `json:"image_url"`
	ImagePreviewURL         string        `json:"image_preview_url"`
	ImageThumbnailURL       string        `json:"image_thumbnail_url"`
	ImageOriginalURL        interface{}   `json:"image_original_url"`
	AnimationURL            interface{}   `json:"animation_url"`
	AnimationOriginalURL    interface{}   `json:"animation_original_url"`
	Name                    string        `json:"name"`
	Description             string        `json:"description"`
	ExternalLink            interface{}   `json:"external_link"`
	AssetContract           AssetContract `json:"asset_contract"`
	Permalink               string        `json:"permalink"`
	Collection              Collection    `json:"collection"`
	Decimals                interface{}   `json:"decimals"`
	TokenMetadata           interface{}   `json:"token_metadata"`
	IsNsfw                  bool          `json:"is_nsfw"`
	Owner                   Owner         `json:"owner"`
	SellOrders              interface{}   `json:"sell_orders"`
	Creator                 Creator       `json:"creator"`
	Traits                  []interface{} `json:"traits"`
	LastSale                interface{}   `json:"last_sale"`
	TopBid                  interface{}   `json:"top_bid"`
	ListingDate             interface{}   `json:"listing_date"`
	IsPresale               bool          `json:"is_presale"`
	TransferFeePaymentToken interface{}   `json:"transfer_fee_payment_token"`
	TransferFee             interface{}   `json:"transfer_fee"`
	TokenID                 string        `json:"token_id"`
}
