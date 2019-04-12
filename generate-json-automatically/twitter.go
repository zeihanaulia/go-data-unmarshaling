package main

type StatusResponse struct {
	SearchMetadata StatusResponse_sub1    `json:"search_metadata"`
	Statuses       []StatusResponse_sub10 `json:"statuses"`
}

type StatusResponse_sub1 struct {
	CompletedIn float64 `json:"completed_in"`
	Count       int64   `json:"count"`
	MaxID       int64   `json:"max_id"`
	MaxIDStr    string  `json:"max_id_str"`
	NextResults string  `json:"next_results"`
	Query       string  `json:"query"`
	SinceID     int64   `json:"since_id"`
	SinceIDStr  string  `json:"since_id_str"`
}

type StatusResponse_sub10 struct {
	Contributors         interface{}         `json:"contributors"`
	Coordinates          interface{}         `json:"coordinates"`
	CreatedAt            string              `json:"created_at"`
	Entities             StatusResponse_sub4 `json:"entities"`
	FavoriteCount        int64               `json:"favorite_count"`
	Favorited            bool                `json:"favorited"`
	Geo                  interface{}         `json:"geo"`
	ID                   int64               `json:"id"`
	IDStr                string              `json:"id_str"`
	InReplyToScreenName  interface{}         `json:"in_reply_to_screen_name"`
	InReplyToStatusID    interface{}         `json:"in_reply_to_status_id"`
	InReplyToStatusIDStr interface{}         `json:"in_reply_to_status_id_str"`
	InReplyToUserID      interface{}         `json:"in_reply_to_user_id"`
	InReplyToUserIDStr   interface{}         `json:"in_reply_to_user_id_str"`
	IsQuoteStatus        bool                `json:"is_quote_status"`
	Lang                 string              `json:"lang"`
	Metadata             StatusResponse_sub5 `json:"metadata"`
	Place                interface{}         `json:"place"`
	PossiblySensitive    bool                `json:"possibly_sensitive"`
	RetweetCount         int64               `json:"retweet_count"`
	Retweeted            bool                `json:"retweeted"`
	Source               string              `json:"source"`
	Text                 string              `json:"text"`
	Truncated            bool                `json:"truncated"`
	User                 StatusResponse_sub9 `json:"user"`
}

type StatusResponse_sub9 struct {
	ContributorsEnabled            bool                `json:"contributors_enabled"`
	CreatedAt                      string              `json:"created_at"`
	DefaultProfile                 bool                `json:"default_profile"`
	DefaultProfileImage            bool                `json:"default_profile_image"`
	Description                    string              `json:"description"`
	Entities                       StatusResponse_sub8 `json:"entities"`
	FavouritesCount                int64               `json:"favourites_count"`
	FollowRequestSent              interface{}         `json:"follow_request_sent"`
	FollowersCount                 int64               `json:"followers_count"`
	Following                      interface{}         `json:"following"`
	FriendsCount                   int64               `json:"friends_count"`
	GeoEnabled                     bool                `json:"geo_enabled"`
	HasExtendedProfile             bool                `json:"has_extended_profile"`
	ID                             int64               `json:"id"`
	IDStr                          string              `json:"id_str"`
	IsTranslationEnabled           bool                `json:"is_translation_enabled"`
	IsTranslator                   bool                `json:"is_translator"`
	Lang                           string              `json:"lang"`
	ListedCount                    int64               `json:"listed_count"`
	Location                       string              `json:"location"`
	Name                           string              `json:"name"`
	Notifications                  interface{}         `json:"notifications"`
	ProfileBackgroundColor         string              `json:"profile_background_color"`
	ProfileBackgroundImageURL      string              `json:"profile_background_image_url"`
	ProfileBackgroundImageURLHTTPS string              `json:"profile_background_image_url_https"`
	ProfileBackgroundTile          bool                `json:"profile_background_tile"`
	ProfileBannerURL               string              `json:"profile_banner_url"`
	ProfileImageURL                string              `json:"profile_image_url"`
	ProfileImageURLHTTPS           string              `json:"profile_image_url_https"`
	ProfileLinkColor               string              `json:"profile_link_color"`
	ProfileSidebarBorderColor      string              `json:"profile_sidebar_border_color"`
	ProfileSidebarFillColor        string              `json:"profile_sidebar_fill_color"`
	ProfileTextColor               string              `json:"profile_text_color"`
	ProfileUseBackgroundImage      bool                `json:"profile_use_background_image"`
	Protected                      bool                `json:"protected"`
	ScreenName                     string              `json:"screen_name"`
	StatusesCount                  int64               `json:"statuses_count"`
	TimeZone                       string              `json:"time_zone"`
	TranslatorType                 string              `json:"translator_type"`
	URL                            string              `json:"url"`
	UtcOffset                      int64               `json:"utc_offset"`
	Verified                       bool                `json:"verified"`
}

type StatusResponse_sub8 struct {
	Description StatusResponse_sub6 `json:"description"`
	URL         StatusResponse_sub7 `json:"url"`
}

type StatusResponse_sub3 struct {
	DisplayURL  string  `json:"display_url"`
	ExpandedURL string  `json:"expanded_url"`
	Indices     []int64 `json:"indices"`
	URL         string  `json:"url"`
}

type StatusResponse_sub4 struct {
	Hashtags     []StatusResponse_sub2 `json:"hashtags"`
	Symbols      []interface{}         `json:"symbols"`
	Urls         []StatusResponse_sub3 `json:"urls"`
	UserMentions []interface{}         `json:"user_mentions"`
}

type StatusResponse_sub2 struct {
	Indices []int64 `json:"indices"`
	Text    string  `json:"text"`
}

type StatusResponse_sub5 struct {
	IsoLanguageCode string `json:"iso_language_code"`
	ResultType      string `json:"result_type"`
}

type StatusResponse_sub7 struct {
	Urls []StatusResponse_sub3 `json:"urls"`
}

type StatusResponse_sub6 struct {
	Urls []interface{} `json:"urls"`
}
