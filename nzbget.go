package nzbget

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"path"
)

// New returns a new instance of an NZBGet client
func New(baseURL, user, password string) (*NZBGet, error) {
	url, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &NZBGet{
		client:   &http.Client{},
		baseURL:  url,
		user:     user,
		password: password,
	}, nil
}

// NZBGet is a client instance for NZBGet
type NZBGet struct {
	client   *http.Client
	baseURL  *url.URL
	user     string
	password string
}

type response struct {
	Result  json.RawMessage `json:"result"`
	Version string          `json:"version"`
}

type NZBFileGroup struct {
	ActiveDownloads  int           `json:"ActiveDownloads"`
	Category         string        `json:"Category"`
	CriticalHealth   int           `json:"CriticalHealth"`
	DeleteStatus     string        `json:"DeleteStatus"`
	Deleted          bool          `json:"Deleted"`
	DestDir          string        `json:"DestDir"`
	DownloadTimeSec  int           `json:"DownloadTimeSec"`
	DownloadedSizeHi int           `json:"DownloadedSizeHi"`
	DownloadedSizeLo int           `json:"DownloadedSizeLo"`
	DownloadedSizeMB int           `json:"DownloadedSizeMB"`
	DupeKey          string        `json:"DupeKey"`
	DupeMode         string        `json:"DupeMode"`
	DupeScore        int           `json:"DupeScore"`
	ExParStatus      string        `json:"ExParStatus"`
	ExtraParBlocks   int           `json:"ExtraParBlocks"`
	FailedArticles   int           `json:"FailedArticles"`
	FileCount        int           `json:"FileCount"`
	FileSizeHi       int           `json:"FileSizeHi"`
	FileSizeLo       int           `json:"FileSizeLo"`
	FileSizeMB       int           `json:"FileSizeMB"`
	FinalDir         string        `json:"FinalDir"`
	FirstID          int           `json:"FirstID"`
	Health           int           `json:"Health"`
	Kind             string        `json:"Kind"`
	LastID           int           `json:"LastID"`
	Log              []interface{} `json:"Log"`
	MarkStatus       string        `json:"MarkStatus"`
	MaxPostTime      int           `json:"MaxPostTime"`
	MaxPriority      int           `json:"MaxPriority"`
	MessageCount     int           `json:"MessageCount"`
	MinPostTime      int           `json:"MinPostTime"`
	MinPriority      int           `json:"MinPriority"`
	MoveStatus       string        `json:"MoveStatus"`
	NZBFilename      string        `json:"NZBFilename"`
	NZBID            int           `json:"NZBID"`
	NZBName          string        `json:"NZBName"`
	NZBNicename      string        `json:"NZBNicename"`
	ParStatus        string        `json:"ParStatus"`
	ParTimeSec       int           `json:"ParTimeSec"`
	Parameters       []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Parameters"`
	PausedSizeHi       int           `json:"PausedSizeHi"`
	PausedSizeLo       int           `json:"PausedSizeLo"`
	PausedSizeMB       int           `json:"PausedSizeMB"`
	PostInfoText       string        `json:"PostInfoText"`
	PostStageProgress  int           `json:"PostStageProgress"`
	PostStageTimeSec   int           `json:"PostStageTimeSec"`
	PostTotalTimeSec   int           `json:"PostTotalTimeSec"`
	RemainingFileCount int           `json:"RemainingFileCount"`
	RemainingParCount  int           `json:"RemainingParCount"`
	RemainingSizeHi    int           `json:"RemainingSizeHi"`
	RemainingSizeLo    int           `json:"RemainingSizeLo"`
	RemainingSizeMB    int           `json:"RemainingSizeMB"`
	RepairTimeSec      int           `json:"RepairTimeSec"`
	ScriptStatus       string        `json:"ScriptStatus"`
	ScriptStatuses     []interface{} `json:"ScriptStatuses"`
	ServerStats        []struct {
		FailedArticles  int `json:"FailedArticles"`
		ServerID        int `json:"ServerID"`
		SuccessArticles int `json:"SuccessArticles"`
	} `json:"ServerStats"`
	Status          string `json:"Status"`
	SuccessArticles int    `json:"SuccessArticles"`
	TotalArticles   int    `json:"TotalArticles"`
	URL             string `json:"URL"`
	UnpackStatus    string `json:"UnpackStatus"`
	UnpackTimeSec   int    `json:"UnpackTimeSec"`
	URLStatus       string `json:"UrlStatus"`
}

// FileGroups returns the list of all file groups
func (n NZBGet) FileGroups() (*NZBFileGroup, error) {
	var fileGroups NZBFileGroup
	err := n.get("listgroups", &fileGroups)
	if err != nil {
		return nil, err
	}
	return &fileGroups, nil
}

type NZBStatus struct {
	ArticleCacheHi      int  `json:"ArticleCacheHi"`
	ArticleCacheLo      int  `json:"ArticleCacheLo"`
	ArticleCacheMB      int  `json:"ArticleCacheMB"`
	AverageDownloadRate int  `json:"AverageDownloadRate"`
	DaySizeHi           int  `json:"DaySizeHi"`
	DaySizeLo           int  `json:"DaySizeLo"`
	DaySizeMB           int  `json:"DaySizeMB"`
	Download2Paused     bool `json:"Download2Paused"`
	DownloadLimit       int  `json:"DownloadLimit"`
	DownloadPaused      bool `json:"DownloadPaused"`
	DownloadRate        int  `json:"DownloadRate"`
	DownloadTimeSec     int  `json:"DownloadTimeSec"`
	DownloadedSizeHi    int  `json:"DownloadedSizeHi"`
	DownloadedSizeLo    int  `json:"DownloadedSizeLo"`
	DownloadedSizeMB    int  `json:"DownloadedSizeMB"`
	FeedActive          bool `json:"FeedActive"`
	ForcedSizeHi        int  `json:"ForcedSizeHi"`
	ForcedSizeLo        int  `json:"ForcedSizeLo"`
	ForcedSizeMB        int  `json:"ForcedSizeMB"`
	FreeDiskSpaceHi     int  `json:"FreeDiskSpaceHi"`
	FreeDiskSpaceLo     int  `json:"FreeDiskSpaceLo"`
	FreeDiskSpaceMB     int  `json:"FreeDiskSpaceMB"`
	MonthSizeHi         int  `json:"MonthSizeHi"`
	MonthSizeLo         int  `json:"MonthSizeLo"`
	MonthSizeMB         int  `json:"MonthSizeMB"`
	NewsServers         []struct {
		Active bool `json:"Active"`
		ID     int  `json:"ID"`
	} `json:"NewsServers"`
	ParJobCount      int  `json:"ParJobCount"`
	PostJobCount     int  `json:"PostJobCount"`
	PostPaused       bool `json:"PostPaused"`
	QueueScriptCount int  `json:"QueueScriptCount"`
	QuotaReached     bool `json:"QuotaReached"`
	RemainingSizeHi  int  `json:"RemainingSizeHi"`
	RemainingSizeLo  int  `json:"RemainingSizeLo"`
	RemainingSizeMB  int  `json:"RemainingSizeMB"`
	ResumeTime       int  `json:"ResumeTime"`
	ScanPaused       bool `json:"ScanPaused"`
	ServerPaused     bool `json:"ServerPaused"`
	ServerStandBy    bool `json:"ServerStandBy"`
	ServerTime       int  `json:"ServerTime"`
	ThreadCount      int  `json:"ThreadCount"`
	UpTimeSec        int  `json:"UpTimeSec"`
	URLCount         int  `json:"UrlCount"`
}

// Status returns the current status of nzbget
func (n NZBGet) Status() (*NZBStatus, error) {
	var status NZBStatus
	err := n.get("status", &status)
	if err != nil {
		return nil, err
	}
	return &status, nil
}

// ByteRate represents data transfer rates
type ByteRate struct {
	// SizeLo is the Amount of downloaded data, low 32-bits of 64-bit value.
	SizeLo int

	// SizeHi is the Amount of downloaded data, high 32-bits of 64-bit value.
	SizeHi int

	// SizeMB is the Amount of downloaded data, in megabytes.
	SizeMB int
}

// NZBServerVolumes represents download volume statistics per news-server
type NZBServerVolumes struct {
	// ServerID is the ID of news server
	ServerID int `json:"ServerID"`

	// DataTime is the Date/time when the data was last updated (time is in C/Unix format).
	DataTime int `json:"DataTime"`

	// TotalSizeLo is the Total amount of downloaded data since program installation, low 32-bits of 64-bit value.
	TotalSizeLo int `json:"TotalSizeLo"`

	// TotalSizeHi is the Total amount of downloaded data since program installation, high 32-bits of 64-bit value.
	TotalSizeHi int `json:"TotalSizeHi"`

	// TotalSizeMB is the Total amount of downloaded data since program installation, in megabytes.
	TotalSizeMB int `json:"TotalSizeMB"`

	// CustomSizeLo is the Amount of downloaded data since last reset of custom counter, low 32-bits of 64-bit value.
	CustomSizeLo int `json:"CustomSizeLo"`

	// CustomSizeHi is the Amount of downloaded data since last reset of custom counter, high 32-bits of 64-bit value.
	CustomSizeHi int `json:"CustomSizeHi"`

	// CustomSizeMB is the Amount of downloaded data since last reset of custom counter, in megabytes.
	CustomSizeMB int `json:"CustomSizeMB"`

	// CustomTime is the Date/time of the last reset of custom counter (time is in C/Unix format).
	CustomTime int `json:"CustomTime"`

	// BytesPerSeconds is the - Per-second amount of data downloaded in last 60 seconds. See below.
	BytesPerSeconds ByteRate `json:"BytesPerSeconds"`

	// BytesPerMinutes is the - Per-minute amount of data downloaded in last 60 minutes. See below.
	BytesPerMinutes ByteRate `json:"BytesPerMinutes"`

	// BytesPerHours is the - Per-hour amount of data downloaded in last 24 hours. See below.
	BytesPerHours ByteRate `json:"BytesPerHours"`

	// BytesPerDays is the - Per-day amount of data downloaded since program installation. See below.
	BytesPerDays ByteRate `json:"BytesPerDays"`

	// SecSlot is the The current second slot of field BytesPerSeconds the program writes into.
	SecSlot int `json:"SecSlot"`

	// MinSlot is the The current minute slot of field BytesPerMinutes the program writes into.
	MinSlot int `json:"MinSlot"`

	// HourSlot is the The current hour slot of field BytesPerHours the program writes into.
	HourSlot int `json:"HourSlot"`

	// DaySlot is the The current day slot of field BytesPerDays the program writes into.
	DaySlot int `json:"DaySlot"`

	// FirstDay is the Indicates which calendar day the very first slot of BytesPerDays corresponds to. Details see below.
	FirstDay int `json:"FirstDay"`
}

// ServerVolumes returns the current status of nzbget
func (n NZBGet) ServerVolumes() (*NZBServerVolumes, error) {
	var volumes NZBServerVolumes
	err := n.get("servervolumes", &volumes)
	if err != nil {
		return nil, err
	}
	return &volumes, nil
}

func (n NZBGet) get(endpoint string, responseObject interface{}) error {
	n.baseURL.Path = path.Join("jsonrpc", endpoint)
	req, err := http.NewRequest("GET", n.baseURL.String(), nil)
	req.SetBasicAuth(n.user, n.password)
	if err != nil {
		return err
	}
	result, err := n.client.Do(req)
	if err != nil {
		return err
	}
	defer result.Body.Close()
	var response response
	err = json.NewDecoder(result.Body).Decode(&response)
	if err != nil {
		log.Printf("error unmarshaling nzbget status: %v", err)
		return err
	}
	return json.Unmarshal(response.Result, &responseObject)
}
