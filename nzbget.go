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
	nzbgetURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, err
	}
	return &NZBGet{
		client:   &http.Client{},
		baseURL:  nzbgetURL,
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

// Config returns the server configuration
func (n NZBGet) Config() (map[string]string, error) {
	config := map[string]string{}
	var configEntries []struct {
		Name  string
		Value string
	}
	err := n.get("config", &configEntries)
	if err != nil {
		return nil, err
	}
	for _, entry := range configEntries {
		config[entry.Name] = entry.Value
	}
	return config, nil
}

// FileGroup is summary information for each group (nzb-file).
type FileGroup struct {
	// ActiveDownloads is the number of active downloads in the group. With this
	// filed can be determined what group(s) is (are) being currently
	// downloaded. In most cases only one group is downloaded at a time however
	// more that one group can be downloaded simultaneously when the first group
	// is almost completely downloaded.
	ActiveDownloads int `json:"ActiveDownloads"`

	// Category is the category for group or empty string if none category
	// is assigned.
	Category string `json:"Category"`

	// CriticalHealth is the calculated critical health of the group, in
	// permille. 1000 means 100.0%. The critical health is calculated based on
	// the number and size of par-files. Lower values are better.
	CriticalHealth int `json:"CriticalHealth"`

	// DeleteStatus indicates if the download was deleted:
	//
	//    NONE - not deleted;
	//    MANUAL - the download was manually deleted by user;
	//    HEALTH - the download was deleted by health check;
	//    DUPE - the download was deleted by duplicate check;
	//    BAD - v14.0 the download was marked as BAD by a queue-script during
	//   		download;
	//    SCAN - v16.0 the download was deleted because the nzb-file could not
	//   		be parsed (malformed nzb-file);
	//    COPY - v16.0 the download was deleted by duplicate check because an
	//   		nzb-file with exactly same content exists in download queue or
	//  		in history.
	DeleteStatus string `json:"DeleteStatus"`

	// Deleted is deprecated, use DeleteStatus instead
	Deleted bool `json:"Deleted"`

	// DestDir is the destination directory for output file.
	DestDir string `json:"DestDir"`

	// DownloadTimeSec is the download time in seconds.
	DownloadTimeSec int `json:"DownloadTimeSec"`

	// DownloadedSizeHi is the amount of downloaded data for group in bytes,
	// High 32-bits of 64-bit value.
	DownloadedSizeHi int `json:"DownloadedSizeHi"`

	// DownloadedSizeLo is the amount of downloaded data for group in bytes,
	// Low 32-bits of 64-bit value.
	DownloadedSizeLo int `json:"DownloadedSizeLo"`

	// DownloadedSizeMB is the amount of downloaded data for group in megabytes.
	DownloadedSizeMB int `json:"DownloadedSizeMB"`

	// DupeKey is the duplicate key. See RSS.
	DupeKey string `json:"DupeKey"`

	// DupeMode is the duplicate mode. One of SCORE, ALL, FORCE. See RSS.
	DupeMode string `json:"DupeMode"`

	// DupeScore is the duplicate score. See RSS.
	DupeScore int `json:"DupeScore"`

	// ExParStatus is the indicates if the download was repaired using duplicate
	// par-scan mode (option ParScan=dupe):
	//
	//    RECIPIENT - repaired using blocks from other duplicates;
	//    DONOR - has donated blocks to repair another duplicate;
	ExParStatus string `json:"ExParStatus"`

	// ExtraParBlocks is the amount of extra par-blocks received from other
	// duplicates or donated to other duplicates, when duplicate par-scan mode
	// was used (option ParScan=dupe):
	//
	//    > 0 - has received extra blocks;
	//    < 0 - has donated extra blocks;
	ExtraParBlocks int `json:"ExtraParBlocks"`

	// FailedArticles is the number of failed article downloads.
	FailedArticles int `json:"FailedArticles"`

	// FileCount is the initial number of files in group.
	FileCount int `json:"FileCount"`

	// FileSizeHi is the initial size of all files in group in bytes, High
	// 32-bits of 64-bit value.
	FileSizeHi int `json:"FileSizeHi"`

	// FileSizeLo is the initial size of all files in group in bytes, Low
	// 32-bits of 64-bit value.
	FileSizeLo int `json:"FileSizeLo"`

	// FileSizeMB is the initial size of all files in group in megabytes.
	FileSizeMB int `json:"FileSizeMB"`

	// FinalDir is the final destination if set by one of post-processing
	// scripts. Can be set only for items in post-processing state.
	FinalDir string `json:"FinalDir"`

	// FirstID is deprecated, use NZBID instead.
	FirstID int `json:"FirstID"`

	// Health is teh current health of the group, in permille. 1000 means 100.0%.
	// The health can go down below this valued during download if more article
	// fails. It can never increase (unless merging of groups). Higher values
	// are better. See forum topic Download health monitoring.
	Health int `json:"Health"`

	// Kind is the kind of queue entry: NZB or URL.
	Kind string `json:"Kind"`

	// LastID is deprecated, use NZBID instead.
	LastID int `json:"LastID"`

	// Log is an array of structs with log-messages. For description of struct
	// see method log. Only for a group which is being currently post-processed.
	// The number of returned entries is limited by parameter NumberOfLogEntries.
	// Deprecated, use method loadlog instead.
	Log []interface{} `json:"Log"`

	// MarkStatus indicates if the download was marked by user:
	//
	//    NONE - not marked;
	//    GOOD - the download was marked as good by user using command Mark as
	//   		 good in history dialog;
	//    BAD - the download was marked as bad by user using command Mark as bad
	//   		in history dialog;
	MarkStatus string `json:"MarkStatus"`

	// MaxPostTime is the date/time when the newest file in the group was
	// posted to newsgroup (Time is in C/Unix format).
	MaxPostTime int `json:"MaxPostTime"`

	// MaxPriority is the priority of the group. “Max” in the field name has
	// historical reasons.
	MaxPriority int `json:"MaxPriority"`

	// MessageCount is the number of messages stored in the item log. Messages
	// can be retrieved with method loadlog.
	MessageCount int `json:"MessageCount"`

	// MinPostTime is the date/time when the oldest file in the group was posted
	// to newsgroup (Time is in C/Unix format).
	MinPostTime int `json:"MinPostTime"`

	// MinPriority is deprecated, use MaxPriority instead.
	MinPriority int `json:"MinPriority"`

	// MoveStatus is the result of moving files from intermediate directory into
	// final directory:
	//
	//    NONE - the moving wasn’t made because either the option InterDir is
	//   		 not in use or the par-check or unpack have failed;
	//    SUCCESS - files were moved successfully;
	//    FAILURE - the moving has failed.
	MoveStatus string `json:"MoveStatus"`

	// NZBFilename is the name of nzb-file, this file was added to queue from.
	// The filename could include fullpath (if client sent it by adding the file
	// to queue).
	NZBFilename string `json:"NZBFilename"`

	// NZBID is the ID of NZB-file.
	NZBID int `json:"NZBID"`

	// NZBName is the name of nzb-file without path and extension. Ready for
	// user-friendly output.
	NZBName string `json:"NZBName"`

	// NZBNicename is deprecated, use NZBName instead.
	NZBNicename string `json:"NZBNicename"`

	// ParStatus is the result of par-check/repair:
	//
	//    NONE - par-check wasn’t performed;
	//    FAILURE - par-check has failed;
	//    REPAIR_POSSIBLE - download is damaged, additional par-files were
	//    					downloaded but the download was not repaired. Either
	//   					the option ParRepair is disabled or the par-repair
	//  					was cancelled by option ParTimeLimit;
	//    SUCCESS - par-check was successful;
	//    MANUAL - download is damaged but was not checked/repaired because
	//   		   option ParCheck is set to Manual.
	ParStatus string `json:"ParStatus"`

	// ParTimeSec is the par-check time in seconds (incl. verification and
	// repair).
	ParTimeSec int `json:"ParTimeSec"`

	// Parameters is the post-processing parameters for group. An array of
	// structures with following fields:
	Parameters []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Parameters"`

	// PausedSizeHi is the size of all paused files in group in bytes, High
	// 32-bits of 64-bit value.
	PausedSizeHi int `json:"PausedSizeHi"`

	// PausedSizeLo is the size of all paused files in group in bytes, Low
	// 32-bits of 64-bit value.
	PausedSizeLo int `json:"PausedSizeLo"`

	// PausedSizeMB is the size of all paused files in group in megabytes.
	PausedSizeMB int `json:"PausedSizeMB"`

	// PostInfoText is the text with short description of current action in
	// post processor. For example: “Verifying file myfile.rar”. Only for a
	// group which is being currently post-processed.
	PostInfoText string `json:"PostInfoText"`

	// PostStageProgress represents the completing of current stage, in
	// permille. 1000 means 100.0%. Only for a group which is being currently
	// post-processed.
	PostStageProgress int `json:"PostStageProgress"`

	// PostStageTimeSec is the number of seconds the current stage is being
	// processed. Only for a group which is being currently post-processed.
	PostStageTimeSec int `json:"PostStageTimeSec"`

	// PostTotalTimeSec is the number of seconds this post-job is being
	// processed (after it first changed the state from PP-QUEUED). Only for a
	// group which is being currently post-processed.
	PostTotalTimeSec int `json:"PostTotalTimeSec"`

	// RemainingFileCount is the remaining (current) number of files in group.
	RemainingFileCount int `json:"RemainingFileCount"`

	// RemainingParCount is the remaining (current) number of par-files in group.
	RemainingParCount int `json:"RemainingParCount"`

	// RemainingSizeHi is the remaining size of all (remaining) files in group
	// in bytes, High 32-bits of 64-bit value.
	RemainingSizeHi int `json:"RemainingSizeHi"`

	// RemainingSizeLo remaining size of all (remaining) files in group in bytes,
	// Low 32-bits of 64-bit value.
	RemainingSizeLo int `json:"RemainingSizeLo"`

	// RemainingSizeMB is the remaining size of all (remaining) files in group
	// in megabytes.
	RemainingSizeMB int `json:"RemainingSizeMB"`

	// RepairTimeSec is the par-repair time in seconds.
	RepairTimeSec int `json:"RepairTimeSec"`

	// ScriptStatus is the accumulated result of all post-processing scripts.
	// One of the predefined text constants: NONE, FAILURE, SUCCESS. Also see
	// field ScriptStatuses.
	ScriptStatus string `json:"ScriptStatus"`

	// ScriptStatuses is the status info of each post-processing script.
	ScriptStatuses []interface{} `json:"ScriptStatuses"`

	// ServerStats is the per news-server download statistics.
	ServerStats []struct {
		// FailedArticles (int) - Number of failed articles.
		FailedArticles int `json:"FailedArticles"`
		// ServerID is the server number as defined in section “news servers”
		// of the configuration file.
		ServerID int `json:"ServerID"`
		// SuccessArticles is the number of successfully downloaded articles.
		SuccessArticles int `json:"SuccessArticles"`
	} `json:"ServerStats"`

	// Status is the status of the group:
	//
	//    QUEUED - queued for download;
	//    PAUSED - paused;
	//    DOWNLOADING - item is being downloaded;
	//    FETCHING - nzb-file is being fetched from URL (Kind=URL);
	//    PP_QUEUED - queued for post-processing (completely downloaded);
	//    LOADING_PARS - stage of par-check;
	//    VERIFYING_SOURCES - stage of par-check;
	//    REPAIRING - stage of par-check;
	//    VERIFYING_REPAIRED - stage of par-check;
	//    RENAMING - processed by par-renamer;
	//    UNPACKING - being unpacked;
	//    MOVING - moving files from intermediate directory into destination
	//   		   directory;
	//    EXECUTING_SCRIPT - executing post-processing script;
	//    PP_FINISHED - post-processing is finished, the item is about to be
	//     				moved to history.
	Status string `json:"Status"`

	// SuccessArticles is the number of successfully downloaded articles.
	SuccessArticles int `json:"SuccessArticles"`

	// TotalArticles is the total number of articles in all files of the group.
	TotalArticles int `json:"TotalArticles"`

	// URL is the URL where the NZB-file was fetched (Kind=NZB) or should be
	// fetched (Kind=URL).
	URL string `json:"URL"`

	// UnpackStatus is the result of unpack:
	//
	//    NONE - unpack wasn’t performed, either no archive files were found or
	//    		 the unpack is disabled for that download or globally;
	//    FAILURE - unpack has failed;
	//    SPACE - unpack has failed due to not enough disk space;
	//    PASSWORD - unpack has failed because the password was not provided or
	//   			 was wrong. Only for rar5-archives;
	//    SUCCESS - unpack was successful.
	UnpackStatus string `json:"UnpackStatus"`

	// UnpackTimeSec (int) - v14.0 Unpack time in seconds.
	UnpackTimeSec int `json:"UnpackTimeSec"`

	// UrlStatus is the result of URL-download:
	//
	//    NONE - that nzb-file were not fetched from an URL;
	//    SUCCESS - that nzb-file was fetched from an URL;
	//    FAILURE - the fetching of the URL has failed.
	//    SCAN_SKIPPED - The URL was fetched successfully but downloaded file
	//   				 was not nzb-file and was skipped by the scanner;
	//    SCAN_FAILURE - The URL was fetched successfully but an error occurred
	//   				 during scanning of the downloaded file. The downloaded
	//  				 file isn’t a proper nzb-file. This status usually means
	// 					 the web-server has returned an error page (HTML page)
	//					 instead of the nzb-file.
	URLStatus string `json:"UrlStatus"`
}

// FileGroups returns the list of all file groups
func (n NZBGet) FileGroups() ([]FileGroup, error) {
	var fileGroups []FileGroup
	err := n.get("listgroups", &fileGroups)
	if err != nil {
		return nil, err
	}
	return fileGroups, nil
}

type Status struct {
	// ArticleCacheHi is the current usage of article cache, in bytes. This
	// field contains the high 32-bits of 64-bit value
	ArticleCacheHi int `json:"ArticleCacheHi"`

	// ArticleCacheLo is the current usage of article cache, in bytes. This field
	// contains the low 32-bits of 64-bit value
	ArticleCacheLo int `json:"ArticleCacheLo"`

	// ArticleCacheMB is the current usage of article cache, in megabytes.
	ArticleCacheMB int `json:"ArticleCacheMB"`

	// AverageDownloadRate is the average download speed since server start, in
	// Bytes per Second.
	AverageDownloadRate int `json:"AverageDownloadRate"`

	// DaySizeHi is the amount of data downloaded since the start of this
	// day. This field contains the high 32-bits of 64-bit value
	DaySizeHi int `json:"DaySizeHi"`

	// DaySizeLo is the amount of data downloaded since the start of this
	// day. This field contains the low 32-bits of 64-bit value
	DaySizeLo int `json:"DaySizeLo"`

	// DaySizeMB is the amount of data downloaded since server start, in
	// megabytes.
	DaySizeMB int `json:"DaySizeMB"`

	// DownloadPaused is “True” if download queue is paused via first pause
	// register (soft-pause).
	Download2Paused bool `json:"Download2Paused"`

	// DownloadLimit is the current download limit, in Bytes per Second. The
	// limit can be changed via method “rate”. Be aware of different scales used
	// by the method rate (Kilobytes) and this field (Bytes).
	DownloadLimit int `json:"DownloadLimit"`

	// DownloadPaused is “True” if download queue is paused via first pause
	// register (soft-pause).
	DownloadPaused bool `json:"DownloadPaused"`

	// DownloadRate is the current download speed, in Bytes per Second.
	DownloadRate int `json:"DownloadRate"`

	// DownloadTimeSec is the server download time in seconds.
	DownloadTimeSec int `json:"DownloadTimeSec"`

	// DownloadedSizeHi is the amount of data downloaded since server start, in
	// bytes. This field contains the high 32-bits of 64-bit value
	DownloadedSizeHi int `json:"DownloadedSizeHi"`

	// DownloadedSizeLo is the amount of data downloaded since server start, in
	// bytes. This field contains the low 32-bits of 64-bit value
	DownloadedSizeLo int `json:"DownloadedSizeLo"`

	// DownloadedSizeMB is the amount of data downloaded since server start, in
	// megabytes.
	DownloadedSizeMB int `json:"DownloadedSizeMB"`

	// FeedActive is “True” if any RSS feed is being fetched right now.
	FeedActive bool `json:"FeedActive"`

	// ForcedSizeHi is the remaining size of entries with FORCE priority, in
	// bytes. This field contains the high 32-bits of 64-bit value
	ForcedSizeHi int `json:"ForcedSizeHi"`

	// ForcedSizeLo is the remaining size of entries with FORCE priority, in
	//  bytes. This field contains the low 32-bits of 64-bit value
	ForcedSizeLo int `json:"ForcedSizeLo"`

	// ForcedSizeMB is the remaining size of entries with FORCE priority, in
	// megabytes.
	ForcedSizeMB int `json:"ForcedSizeMB"`

	// FreeDiskSpaceHi is the free disk space on ‘DestDir’, in bytes. This field
	// contains the high 32-bits of 64-bit value
	FreeDiskSpaceHi int `json:"FreeDiskSpaceHi"`

	// FreeDiskSpaceLo is the free disk space on ‘DestDir’, in bytes. This field
	// contains the low 32-bits of 64-bit value
	FreeDiskSpaceLo int `json:"FreeDiskSpaceLo"`

	// FreeDiskSpaceMB is the free disk space on ‘DestDir’, in megabytes.
	FreeDiskSpaceMB int `json:"FreeDiskSpaceMB"`

	// MonthSizeHi is the amount of data downloaded since the start of this
	// month. This field contains the high 32-bits of 64-bit value
	MonthSizeHi int `json:"MonthSizeHi"`

	// MonthSizeLo is the amount of data downloaded since the start of this
	// month. This field contains the low 32-bits of 64-bit value
	MonthSizeLo int `json:"MonthSizeLo"`

	// MonthSizeMB is the amount of data downloaded since the start of this
	// month in megabytes.
	MonthSizeMB int `json:"MonthSizeMB"`

	// NewsServers is the status of news-servers
	NewsServers []struct {
		// Active is true if server is in active state (enabled).
		Active bool `json:"Active"`

		// ID is the server number in the configuration file
		ID int `json:"ID"`
	} `json:"NewsServers"`

	// ParJobCount is deprecated, use PostJobCount instead.
	ParJobCount int `json:"ParJobCount"`

	// PostJobCount is the number of Par-Jobs or Post-processing script jobs in
	// the post-processing queue (including current file).
	PostJobCount int `json:"PostJobCount"`

	// PostPaused is “True” if post-processor queue is currently in paused-state.
	PostPaused bool `json:"PostPaused"`

	// QueueScriptCount is the number of Par-Jobs or Post-processing script jobs
	// in the script queue (including current file).
	QueueScriptCount int `json:"QueueScriptCount"`

	// QuotaReached is “True” if the quota has been reached
	QuotaReached bool `json:"QuotaReached"`

	// RemainingSizeHi is the remaining size of all entries in download queue,
	// in bytes. This field contains the high 32-bits of 64-bit value
	RemainingSizeHi int `json:"RemainingSizeHi"`

	// RemainingSizeLo is the remaining size of all entries in download queue,
	// in bytes. This field contains the low 32-bits of 64-bit value
	RemainingSizeLo int `json:"RemainingSizeLo"`

	// RemainingSizeMB is the remaining size of all entries in download queue,
	// in megabytes.
	RemainingSizeMB int `json:"RemainingSizeMB"`

	// ResumeTime is the time to resume if set with method “scheduleresume”.
	// Time is in C/Unix format.
	ResumeTime int `json:"ResumeTime"`

	// ScanPaused (bool) - “True” if the scanning of incoming nzb-directory is
	// currently in paused-state.
	ScanPaused   bool `json:"ScanPaused"`
	ServerPaused bool `json:"ServerPaused"`

	// ServerStandBy (bool) - “False” - there are currently downloads running,
	// “True” - no downloads in progress (server paused or all jobs completed).
	ServerStandBy bool `json:"ServerStandBy"`

	// ServerTime (int) - Current time on computer running NZBGet. Time is in
	// C/Unix format (number of seconds since 00:00:00 UTC, January 1, 1970).
	ServerTime int `json:"ServerTime"`

	// ThreadCount (int) - Number of threads running. It includes all threads,
	// created by the program, not only download-threads.
	ThreadCount int `json:"ThreadCount"`

	// UpTimeSec (int) - Server uptime in seconds.
	UpTimeSec int `json:"UpTimeSec"`

	// UrlCount (int) - Number of URLs in the URL-queue (including current file).
	URLCount int `json:"UrlCount"`
}

// Status returns the current status of nzbget
func (n NZBGet) Status() (*Status, error) {
	var status Status
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

// ServerVolume represents download volume statistics per news-server
type ServerVolume struct {
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
	BytesPerSeconds []ByteRate `json:"BytesPerSeconds"`

	// BytesPerMinutes is the - Per-minute amount of data downloaded in last 60 minutes. See below.
	BytesPerMinutes []ByteRate `json:"BytesPerMinutes"`

	// BytesPerHours is the - Per-hour amount of data downloaded in last 24 hours. See below.
	BytesPerHours []ByteRate `json:"BytesPerHours"`

	// BytesPerDays is the - Per-day amount of data downloaded since program installation. See below.
	BytesPerDays []ByteRate `json:"BytesPerDays"`

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
func (n NZBGet) ServerVolumes() ([]ServerVolume, error) {
	var volumes []ServerVolume
	err := n.get("servervolumes", &volumes)
	if err != nil {
		return nil, err
	}
	return volumes, nil
}

// HistoricalEntry is the list of items in history-list.
type HistoricalEntry struct {
	// ID is deprecated, use NZBID instead
	ID int `json:"ID"`

	// Name is the the name of nzb-file or info-name of URL, without path and
	// extension. Ready for user-friendly output.
	Name string `json:"Name"`

	// RemainingFileCount is the number of parked files in group. If this number
	// is greater than “0”, the history item can be returned to download queue
	// using command “HistoryReturn” of method
	RemainingFileCount int `json:"RemainingFileCount"`

	// RetryData indicates something  -- not sure
	RetryData bool `json:"RetryData"`

	// HistoryTime is the date/time when the file was added to history (Time is
	// in C/Unix format).
	HistoryTime int `json:"HistoryTime"`

	// Status Total status of the download. One of the predefined text constants
	// such as SUCCESS/ALL or FAILURE/UNPACK
	Status string `json:"Status"`

	// Log is deprecated, was never really used
	Log []interface{} `json:"Log"`

	// NZBID is ID of NZB-file
	NZBID int `json:"NZBID"`

	// NZBFilename is the name of nzb-file, this file was added to queue from.
	// The filename could include fullpath (if client sent it by adding the file
	// to queue).
	NZBFilename string `json:"NZBFilename"`

	// DestDir is the destination directory for output files
	DestDir string `json:"DestDir"`

	// FinalDir is the final destination if set by one of post-processing scripts
	FinalDir string `json:"FinalDir"`

	// Category is the category for group or empty string if none category is
	// assigned
	Category string `json:"Category"`

	// ParStatus is the ParStatus - Result of par-check/repair:
	//
	//    NONE - par-check wasn’t performed;
	//    FAILURE - par-check has failed;
	//    REPAIR_POSSIBLE - download is damaged, additional par-files were
	//   					downloaded but the download was not repaired. Either
	//  					the option ParRepair is disabled or the
	//  					par-repair was cancelled by option ParTimeLimit;
	//    SUCCESS - par-check was successful;
	//    MANUAL - download is damaged but was not checked/repaired because
	//   		   option ParCheck is set to Manual
	ParStatus string `json:"ParStatus"`

	// ExParStatus indicates if the download was repaired using duplicate
	// par-scan mode (option ParScan=dupe):
	//
	//    RECIPIENT - repaired using blocks from other duplicates;
	//    DONOR - has donated blocks to repair another duplicate;
	ExParStatus string `json:"ExParStatus"`

	// UnpackStatus is the result of unpack:
	//
	//    NONE - unpack wasn’t performed, either no archive files were found or
	//   		 the unpack is disabled for that download or globally;
	//    FAILURE - unpack has failed;
	//    SPACE - unpack has failed due to not enough disk space;
	//    PASSWORD - unpack has failed because the password was not provided or
	//               was wrong. Only for rar5-archives;
	//    SUCCESS - unpack was successful.
	UnpackStatus string `json:"UnpackStatus"`

	// MoveStatus is the result of moving files from intermediate directory into
	// final directory:
	//
	//    NONE - the moving wasn’t made because either the option InterDir is
	//   		 not in use or the par-check or unpack have failed;
	//    SUCCESS - files were moved successfully;
	//    FAILURE - the moving has failed.
	MoveStatus string `json:"MoveStatus"`

	// ScriptStatus is the
	ScriptStatus string `json:"ScriptStatus"`

	// DeleteStatus is the indicates if the download was deleted:
	//
	//    NONE - not deleted;
	//    MANUAL - the download was manually deleted by user;
	//    HEALTH - the download was deleted by health check;
	//    DUPE - the download was deleted by duplicate check;
	//    BAD - v14.0 the download was marked as BAD by a queue-script during
	//   		download;
	//    SCAN - v16.0 the download was deleted because the nzb-file could not
	//    		 be parsed (malformed nzb-file);
	//    COPY - v16.0 the download was deleted by duplicate check because an
	//   		 nzb-file with exactly same content exists in download queue or
	//  		 in history.
	DeleteStatus string `json:"DeleteStatus"`

	// MarkStatus indicates if the download was marked by user:
	//
	//    NONE - not marked;
	//    GOOD - the download was marked as good by user using command Mark as
	//           good in history dialog;
	//    BAD - the download was marked as bad by user using command Mark as bad
	//   		in history dialog;
	MarkStatus string `json:"MarkStatus"`

	// URLStatus is the result of URL-download:
	//
	//    NONE - that nzb-file were not fetched from an URL;
	//    SUCCESS - that nzb-file was fetched from an URL;
	//    FAILURE - the fetching of the URL has failed.
	//    SCAN_SKIPPED - The URL was fetched successfully but downloaded file
	//   				 was not nzb-file and was skipped by the scanner;
	//    SCAN_FAILURE - The URL was fetched successfully but an error occurred
	//   				 during scanning of the downloaded file. The downloaded
	//  				 file isn’t a proper nzb-file. This status usually means
	//	 				 the web-server has returned an error page (HTML page)
	//	 				 instead of the nzb-file.
	URLStatus string `json:"UrlStatus"`

	// FileSizeLo is the initial size of all files in group in bytes, Low
	// 32-bits of 64-bit value
	FileSizeLo int `json:"FileSizeLo"`

	// FileSizeHi is the initial size of all files in group in bytes, High
	// 32-bits of 64-bit value
	FileSizeHi int `json:"FileSizeHi"`

	// FileSizeMB is the Initial size of all files in group in megabytes
	FileSizeMB int `json:"FileSizeMB"`

	// FileCount is the initial number of files in group
	FileCount int `json:"FileCount"`

	// MinPostTime is the date/time when the oldest file in the item was posted
	// to newsgroup (Time is in C/Unix format).
	MinPostTime int `json:"MinPostTime"`

	// MaxPostTime is the date/time when the newest file in the item was posted
	// to newsgroup (Time is in C/Unix format).
	MaxPostTime int `json:"MaxPostTime"`

	// TotalArticles is the total number of articles in all files of the group
	TotalArticles int `json:"TotalArticles"`

	// SuccessArticles is the number of successfully downloaded articles
	SuccessArticles int `json:"SuccessArticles"`

	// FailedArticles is the number of failed article downloads
	FailedArticles int `json:"FailedArticles"`

	// Health is the final health of the group, in permille. 1000 means 100.0%.
	// Higher values are better
	Health int `json:"Health"`

	// CriticalHealth is the calculated critical health of the group, in
	// permille. 1000 means 100.0%. The critical health is calculated based on
	// the number and size of par-files. Lower values are better.
	CriticalHealth int `json:"CriticalHealth"`

	// DupeKey is the unique key (a string) for each title. In a case of newznab
	// feeds the duplicate key (short: dupekey) is built by NZBGet automatically
	DupeKey string `json:"DupeKey"`

	// DupeScore is the confidence that NZBGet has that the article is a
	// duplicate
	DupeScore int `json:"DupeScore"`

	// DupeMode is the duplicate mode. One of SCORE, ALL, FORCE
	DupeMode string `json:"DupeMode"`

	// Deleted indicates if the entry was deleted
	Deleted bool `json:"Deleted"`

	// DownloadedSizeLo is the amount of downloaded data for group in bytes, Low
	// 32-bits of 64-bit value
	DownloadedSizeLo int `json:"DownloadedSizeLo"`

	// DownloadedSizeHi is the amount of downloaded data for group in bytes,
	// High 32-bits of 64-bit value
	DownloadedSizeHi int `json:"DownloadedSizeHi"`

	// DownloadedSizeMB is the amount of downloaded data for group in megabytes
	DownloadedSizeMB int `json:"DownloadedSizeMB"`

	// DownloadTimeSec is the download time in seconds.
	DownloadTimeSec int `json:"DownloadTimeSec"`

	// PostTotalTimeSec is the total post-processing time in seconds
	PostTotalTimeSec int `json:"PostTotalTimeSec"`

	// ParTimeSec is the par-check time in seconds (incl. verification and
	// repair).
	ParTimeSec int `json:"ParTimeSec"`

	// RepairTimeSec is the par-repair time in seconds
	RepairTimeSec int `json:"RepairTimeSec"`

	// UnpackTimeSec is the unpack time in seconds
	UnpackTimeSec int `json:"UnpackTimeSec"`

	// MessageCount is the number of messages stored in the item log. Messages
	// can be retrieved with method
	MessageCount int `json:"MessageCount"`

	// ExtraParBlocks is the amount of extra par-blocks received from other
	// duplicates or donated to other duplicates, when duplicate par-scan mode
	// was used (option ParScan=dupe):
	//
	//    > 0 - has received extra blocks;
	//    < 0 - has donated extra blocks;
	ExtraParBlocks int `json:"ExtraParBlocks"`

	// Parameters are the post-processing parameters for group
	Parameters []struct {
		Name  string `json:"Name"`
		Value string `json:"Value"`
	} `json:"Parameters"`

	// ScriptStatuses are the status info of each post-processing script
	ScriptStatuses []struct {
		Name   string `json:"Name"`
		Status string `json:"Status"`
	} `json:"ScriptStatuses"`

	// ServerStats are the per-server article completion statistics
	ServerStats []struct {
		ServerID        int `json:"ServerID"`
		SuccessArticles int `json:"SuccessArticles"`
		FailedArticles  int `json:"FailedArticles"`
	} `json:"ServerStats"`
}

func (n *NZBGet) History() ([]HistoricalEntry, error) {
	var history []HistoricalEntry
	err := n.get("history", &history)
	if err != nil {
		return nil, err
	}
	return history, nil
}

func (n NZBGet) get(endpoint string, responseObject interface{}) error {
	n.baseURL.Path = path.Join("jsonrpc", endpoint)
	req, err := http.NewRequest("GET", n.baseURL.String(), nil)
	if err != nil {
		return err
	}
	req.SetBasicAuth(n.user, n.password)
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
