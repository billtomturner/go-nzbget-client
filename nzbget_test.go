package nzbget_test

import (
	"testing"

	"github.com/SemanticallyNull/golandreporter"
	"github.com/billtomturner/go-nzbget-client"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"gopkg.in/h2non/gock.v1"
)

func TestNZBGet(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithCustomReporters(t, "NZBGet Suite", []Reporter{
		golandreporter.NewGolandReporter(),
	})
}

const (
	nzbgetURL = "http://localhost:6789"
	config    = `
{
  "version": "1.1",
  "result": [
    {
      "Name": "ConfigFile",
      "Value": "\/config\/nzbget.conf"
    },
    {
      "Name": "AppBin",
      "Value": "\/app\/nzbget\/nzbget"
    },
    {
      "Name": "AppDir",
      "Value": "\/app\/nzbget"
    },
    {
      "Name": "Version",
      "Value": "21.0"
    },
    {
      "Name": "MainDir",
      "Value": "\/srv"
    },
    {
      "Name": "WebDir",
      "Value": "\/app\/nzbget\/webui"
    },
    {
      "Name": "ConfigTemplate",
      "Value": "\/app\/nzbget\/share\/nzbget\/nzbget.conf"
    },
    {
      "Name": "TempDir",
      "Value": "\/intermediate\/tmp"
    },
    {
      "Name": "DestDir",
      "Value": "\/downloads"
    },
    {
      "Name": "InterDir",
      "Value": "\/intermediate"
    },
    {
      "Name": "QueueDir",
      "Value": "\/intermediate\/queue"
    },
    {
      "Name": "NzbDir",
      "Value": "\/intermediate\/nzb"
    },
    {
      "Name": "LockFile",
      "Value": "\/srv\/nzbget.lock"
    },
    {
      "Name": "LogFile",
      "Value": "\/srv\/nzbget.log"
    },
    {
      "Name": "ScriptDir",
      "Value": "\/srv\/ppscripts"
    },
    {
      "Name": "RequiredDir",
      "Value": ""
    },
    {
      "Name": "WriteLog",
      "Value": "append"
    },
    {
      "Name": "RotateLog",
      "Value": "3"
    },
    {
      "Name": "AppendCategoryDir",
      "Value": "yes"
    },
    {
      "Name": "OutputMode",
      "Value": "loggable"
    },
    {
      "Name": "DupeCheck",
      "Value": "yes"
    },
    {
      "Name": "DownloadRate",
      "Value": "0"
    },
    {
      "Name": "ControlIp",
      "Value": "0.0.0.0"
    },
    {
      "Name": "ControlUsername",
      "Value": "nzbget"
    },
    {
      "Name": "ControlPassword",
      "Value": "SECRET"
    },
    {
      "Name": "RestrictedUsername",
      "Value": ""
    },
    {
      "Name": "RestrictedPassword",
      "Value": ""
    },
    {
      "Name": "AddUsername",
      "Value": ""
    },
    {
      "Name": "AddPassword",
      "Value": ""
    },
    {
      "Name": "ControlPort",
      "Value": "6789"
    },
    {
      "Name": "FormAuth",
      "Value": "no"
    },
    {
      "Name": "SecureControl",
      "Value": "no"
    },
    {
      "Name": "SecurePort",
      "Value": "6791"
    },
    {
      "Name": "SecureCert",
      "Value": ""
    },
    {
      "Name": "SecureKey",
      "Value": ""
    },
    {
      "Name": "CertStore",
      "Value": "\/app\/nzbget\/cacert.pem"
    },
    {
      "Name": "CertCheck",
      "Value": "no"
    },
    {
      "Name": "AuthorizedIP",
      "Value": "127.0.0.1"
    },
    {
      "Name": "ArticleTimeout",
      "Value": "60"
    },
    {
      "Name": "UrlTimeout",
      "Value": "60"
    },
    {
      "Name": "RemoteTimeout",
      "Value": "90"
    },
    {
      "Name": "FlushQueue",
      "Value": "yes"
    },
    {
      "Name": "NzbLog",
      "Value": "yes"
    },
    {
      "Name": "RawArticle",
      "Value": "no"
    },
    {
      "Name": "SkipWrite",
      "Value": "no"
    },
    {
      "Name": "ArticleRetries",
      "Value": "3"
    },
    {
      "Name": "ArticleInterval",
      "Value": "10"
    },
    {
      "Name": "UrlRetries",
      "Value": "3"
    },
    {
      "Name": "UrlInterval",
      "Value": "10"
    },
    {
      "Name": "ContinuePartial",
      "Value": "yes"
    },
    {
      "Name": "UrlConnections",
      "Value": "4"
    },
    {
      "Name": "LogBuffer",
      "Value": "1000"
    },
    {
      "Name": "InfoTarget",
      "Value": "both"
    },
    {
      "Name": "WarningTarget",
      "Value": "both"
    },
    {
      "Name": "ErrorTarget",
      "Value": "both"
    },
    {
      "Name": "DebugTarget",
      "Value": "both"
    },
    {
      "Name": "DetailTarget",
      "Value": "both"
    },
    {
      "Name": "ParCheck",
      "Value": "auto"
    },
    {
      "Name": "ParRepair",
      "Value": "yes"
    },
    {
      "Name": "ParScan",
      "Value": "extended"
    },
    {
      "Name": "ParQuick",
      "Value": "yes"
    },
    {
      "Name": "PostStrategy",
      "Value": "balanced"
    },
    {
      "Name": "FileNaming",
      "Value": "auto"
    },
    {
      "Name": "ParRename",
      "Value": "yes"
    },
    {
      "Name": "ParBuffer",
      "Value": "16"
    },
    {
      "Name": "ParThreads",
      "Value": "0"
    },
    {
      "Name": "RarRename",
      "Value": "yes"
    },
    {
      "Name": "HealthCheck",
      "Value": "park"
    },
    {
      "Name": "DirectRename",
      "Value": "no"
    },
    {
      "Name": "ScriptOrder",
      "Value": ""
    },
    {
      "Name": "Extensions",
      "Value": ""
    },
    {
      "Name": "DaemonUsername",
      "Value": "abc"
    },
    {
      "Name": "UMask",
      "Value": "1000"
    },
    {
      "Name": "UpdateInterval",
      "Value": "200"
    },
    {
      "Name": "CursesNzbName",
      "Value": "yes"
    },
    {
      "Name": "CursesTime",
      "Value": "no"
    },
    {
      "Name": "CursesGroup",
      "Value": "no"
    },
    {
      "Name": "CrcCheck",
      "Value": "yes"
    },
    {
      "Name": "DirectWrite",
      "Value": "yes"
    },
    {
      "Name": "WriteBuffer",
      "Value": "0"
    },
    {
      "Name": "NzbDirInterval",
      "Value": "5"
    },
    {
      "Name": "NzbDirFileAge",
      "Value": "60"
    },
    {
      "Name": "DiskSpace",
      "Value": "150"
    },
    {
      "Name": "CrashTrace",
      "Value": "yes"
    },
    {
      "Name": "CrashDump",
      "Value": "no"
    },
    {
      "Name": "ParPauseQueue",
      "Value": "no"
    },
    {
      "Name": "ScriptPauseQueue",
      "Value": "no"
    },
    {
      "Name": "NzbCleanupDisk",
      "Value": "no"
    },
    {
      "Name": "ParTimeLimit",
      "Value": "0"
    },
    {
      "Name": "KeepHistory",
      "Value": "7"
    },
    {
      "Name": "Unpack",
      "Value": "yes"
    },
    {
      "Name": "DirectUnpack",
      "Value": "no"
    },
    {
      "Name": "UnpackCleanupDisk",
      "Value": "yes"
    },
    {
      "Name": "UnrarCmd",
      "Value": "unrar"
    },
    {
      "Name": "SevenZipCmd",
      "Value": "7z"
    },
    {
      "Name": "UnpackPassFile",
      "Value": ""
    },
    {
      "Name": "UnpackPauseQueue",
      "Value": "no"
    },
    {
      "Name": "ExtCleanupDisk",
      "Value": ".par2, .sfv, _brokenlog.txt"
    },
    {
      "Name": "ParIgnoreExt",
      "Value": ".sfv, .nzb, .nfo"
    },
    {
      "Name": "UnpackIgnoreExt",
      "Value": ".cbr"
    },
    {
      "Name": "FeedHistory",
      "Value": "7"
    },
    {
      "Name": "UrlForce",
      "Value": "yes"
    },
    {
      "Name": "TimeCorrection",
      "Value": "0"
    },
    {
      "Name": "PropagationDelay",
      "Value": "0"
    },
    {
      "Name": "ArticleCache",
      "Value": "0"
    },
    {
      "Name": "EventInterval",
      "Value": "0"
    },
    {
      "Name": "ShellOverride",
      "Value": ""
    },
    {
      "Name": "MonthlyQuota",
      "Value": "0"
    },
    {
      "Name": "QuotaStartDay",
      "Value": "1"
    },
    {
      "Name": "DailyQuota",
      "Value": "0"
    },
    {
      "Name": "ReorderFiles",
      "Value": "yes"
    },
    {
      "Name": "UpdateCheck",
      "Value": "stable"
    },
    {
      "Name": "Server1.Name",
      "Value": "news.newsgroup.ninja"
    },
    {
      "Name": "Server1.Host",
      "Value": "my.newserver.com"
    },
    {
      "Name": "Server1.Port",
      "Value": "443"
    },
    {
      "Name": "Server1.Username",
      "Value": "username"
    },
    {
      "Name": "Server1.Password",
      "Value": "password"
    },
    {
      "Name": "Server1.Active",
      "Value": "yes"
    },
    {
      "Name": "Server1.Level",
      "Value": "0"
    },
    {
      "Name": "Server1.Encryption",
      "Value": "yes"
    },
    {
      "Name": "Server1.Group",
      "Value": "0"
    },
    {
      "Name": "Server1.JoinGroup",
      "Value": "no"
    },
    {
      "Name": "Server1.Cipher",
      "Value": ""
    },
    {
      "Name": "Server1.Connections",
      "Value": "50"
    },
    {
      "Name": "Server1.IpVersion",
      "Value": "ipv4"
    },
    {
      "Name": "Server1.Notes",
      "Value": ""
    },
    {
      "Name": "Category1.Name",
      "Value": "Movies"
    },
    {
      "Name": "Category1.Unpack",
      "Value": "yes"
    },
    {
      "Name": "Category2.Name",
      "Value": "TV"
    },
    {
      "Name": "Category2.Unpack",
      "Value": "yes"
    },
    {
      "Name": "Category3.Name",
      "Value": "Music"
    },
    {
      "Name": "Category3.Unpack",
      "Value": "yes"
    },
    {
      "Name": "Server1.Optional",
      "Value": "no"
    },
    {
      "Name": "Server1.Retention",
      "Value": "4049"
    },
    {
      "Name": "Category1.DestDir",
      "Value": ""
    },
    {
      "Name": "Category1.Extensions",
      "Value": ""
    },
    {
      "Name": "Category1.Aliases",
      "Value": ""
    },
    {
      "Name": "Category2.DestDir",
      "Value": ""
    },
    {
      "Name": "Category2.Extensions",
      "Value": ""
    },
    {
      "Name": "Category2.Aliases",
      "Value": ""
    },
    {
      "Name": "Category3.DestDir",
      "Value": ""
    },
    {
      "Name": "Category3.Extensions",
      "Value": ""
    },
    {
      "Name": "Category3.Aliases",
      "Value": ""
    }
  ]
}
`
	serverVolumes = `
{
  "version": "1.1",
  "result": [
    {
      "ServerID": 0,
      "DataTime": 1589514272,
      "FirstDay": 18206,
      "TotalSizeLo": 312677052,
      "TotalSizeHi": 2430,
      "TotalSizeMB": 9953578,
      "CustomSizeLo": 312677052,
      "CustomSizeHi": 2430,
      "CustomSizeMB": 9953578,
      "CustomTime": 1573013231,
      "SecSlot": 32,
      "MinSlot": 44,
      "HourSlot": 3,
      "DaySlot": 191,
      "BytesPerSeconds": [
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerMinutes": [
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerHours": [
        { "SizeLo": 1688637931, "SizeHi": 0, "SizeMB": 1610 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 3053505747, "SizeHi": 0, "SizeMB": 2912 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 3053505747, "SizeHi": 0, "SizeMB": 2912 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerDays": [
        { "SizeLo": 2220387312, "SizeHi": 213, "SizeMB": 874565 },
        { "SizeLo": 1128569454, "SizeHi": 44, "SizeMB": 181300 },
        { "SizeLo": 3075069771, "SizeHi": 41, "SizeMB": 170868 },
        { "SizeLo": 2247836278, "SizeHi": 139, "SizeMB": 571487 },
        { "SizeLo": 3171255530, "SizeHi": 72, "SizeMB": 297936 },
        { "SizeLo": 308158014, "SizeHi": 25, "SizeMB": 102693 },
        { "SizeLo": 200802673, "SizeHi": 31, "SizeMB": 127167 },
        { "SizeLo": 1487688225, "SizeHi": 0, "SizeMB": 1418 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 2971753499, "SizeHi": 5, "SizeMB": 23314 }
      ]
    },
    {
      "ServerID": 1,
      "DataTime": 1589514272,
      "FirstDay": 18206,
      "TotalSizeLo": 312677052,
      "TotalSizeHi": 2430,
      "TotalSizeMB": 9953578,
      "CustomSizeLo": 312677052,
      "CustomSizeHi": 2430,
      "CustomSizeMB": 9953578,
      "CustomTime": 1573013231,
      "SecSlot": 32,
      "MinSlot": 44,
      "HourSlot": 3,
      "DaySlot": 191,
      "BytesPerSeconds": [
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerMinutes": [
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerHours": [
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 1688637931, "SizeHi": 0, "SizeMB": 1610 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 },
        { "SizeLo": 0, "SizeHi": 0, "SizeMB": 0 }
      ],
      "BytesPerDays": [
        { "SizeLo": 2220387312, "SizeHi": 213, "SizeMB": 874565 },
        { "SizeLo": 1128569454, "SizeHi": 44, "SizeMB": 181300 },
        { "SizeLo": 3075069771, "SizeHi": 41, "SizeMB": 170868 },
        { "SizeLo": 2247836278, "SizeHi": 139, "SizeMB": 571487 },
        { "SizeLo": 3171255530, "SizeHi": 72, "SizeMB": 297936 },
        { "SizeLo": 308158014, "SizeHi": 25, "SizeMB": 102693 },
        { "SizeLo": 200802673, "SizeHi": 31, "SizeMB": 127167 },
        { "SizeLo": 151930761, "SizeHi": 5, "SizeMB": 20624 },
        { "SizeLo": 2079480133, "SizeHi": 21, "SizeMB": 87999 },
        { "SizeLo": 2755081465, "SizeHi": 8, "SizeMB": 35395 },
        { "SizeLo": 2851265019, "SizeHi": 20, "SizeMB": 84639 }
      ]
    }
  ]
}
`
	listGroups = `{
  "version": "1.1",
  "result": [
    {
      "FirstID": 15780,
      "LastID": 15780,
      "RemainingSizeLo": 1264637546,
      "RemainingSizeHi": 1,
      "RemainingSizeMB": 5302,
      "PausedSizeLo": 354370348,
      "PausedSizeHi": 0,
      "PausedSizeMB": 337,
      "RemainingFileCount": 60,
      "RemainingParCount": 8,
      "MinPriority": 0,
      "MaxPriority": 0,
      "ActiveDownloads": 50,
      "Status": "DOWNLOADING",
      "NZBID": 15780,
      "NZBName": "My_File_Name_1",
      "NZBNicename": "My_File_Name_1",
      "Kind": "NZB",
      "URL": "",
      "NZBFilename": "My_File_Name_1.nzb",
      "DestDir": "\/intermediate\/My_File_Name_1-SA89.#15780",
      "FinalDir": "",
      "Category": "Files",
      "ParStatus": "NONE",
      "ExParStatus": "NONE",
      "UnpackStatus": "NONE",
      "MoveStatus": "NONE",
      "ScriptStatus": "NONE",
      "DeleteStatus": "NONE",
      "MarkStatus": "NONE",
      "UrlStatus": "NONE",
      "FileSizeLo": 3100157546,
      "FileSizeHi": 1,
      "FileSizeMB": 7052,
      "FileCount": 76,
      "MinPostTime": 1543567681,
      "MaxPostTime": 1543568146,
      "TotalArticles": 1934,
      "SuccessArticles": 478,
      "FailedArticles": 0,
      "Health": 1000,
      "CriticalHealth": 949,
      "DupeKey": "",
      "DupeScore": 0,
      "DupeMode": "ALL",
      "Deleted": false,
      "DownloadedSizeLo": 1892954988,
      "DownloadedSizeHi": 0,
      "DownloadedSizeMB": 1805,
      "DownloadTimeSec": 20,
      "PostTotalTimeSec": 0,
      "ParTimeSec": 0,
      "RepairTimeSec": 0,
      "UnpackTimeSec": 0,
      "MessageCount": 18,
      "ExtraParBlocks": 0,
      "Parameters": [
        {
          "Name": "drone",
          "Value": "b1b83444492a4bc8846641107070a3bf"
        },
        {
          "Name": "*Unpack:",
          "Value": "yes"
        }
      ],
      "ScriptStatuses": [
      ],
      "ServerStats": [
        {
          "ServerID": 1,
          "SuccessArticles": 478,
          "FailedArticles": 0
        }
      ],
      "PostInfoText": "NONE",
      "PostStageProgress": -1639621749,
      "PostStageTimeSec": 0,
      "Log": [
      ]
    },
    {
      "FirstID": 15781,
      "LastID": 15781,
      "RemainingSizeLo": 4239181076,
      "RemainingSizeHi": 0,
      "RemainingSizeMB": 4042,
      "PausedSizeLo": 451017973,
      "PausedSizeHi": 0,
      "PausedSizeMB": 430,
      "RemainingFileCount": 99,
      "RemainingParCount": 21,
      "MinPriority": 0,
      "MaxPriority": 0,
      "ActiveDownloads": 0,
      "Status": "QUEUED",
      "NZBID": 15781,
      "NZBName": "My_File_Name_2",
      "NZBNicename": "My_File_Name_2",
      "Kind": "NZB",
      "URL": "",
      "NZBFilename": "My_File_Name_2.nzb",
      "DestDir": "\/intermediate\/My_File_Name_2.#15781",
      "FinalDir": "",
      "Category": "TV",
      "ParStatus": "NONE",
      "ExParStatus": "NONE",
      "UnpackStatus": "NONE",
      "MoveStatus": "NONE",
      "ScriptStatus": "NONE",
      "DeleteStatus": "NONE",
      "MarkStatus": "NONE",
      "UrlStatus": "NONE",
      "FileSizeLo": 4239181076,
      "FileSizeHi": 0,
      "FileSizeMB": 4042,
      "FileCount": 99,
      "MinPostTime": 1477894248,
      "MaxPostTime": 1477894248,
      "TotalArticles": 5429,
      "SuccessArticles": 0,
      "FailedArticles": 0,
      "Health": 1000,
      "CriticalHealth": 880,
      "DupeKey": "",
      "DupeScore": 0,
      "DupeMode": "ALL",
      "Deleted": false,
      "DownloadedSizeLo": 0,
      "DownloadedSizeHi": 0,
      "DownloadedSizeMB": 0,
      "DownloadTimeSec": 0,
      "PostTotalTimeSec": 0,
      "ParTimeSec": 0,
      "RepairTimeSec": 0,
      "UnpackTimeSec": 0,
      "MessageCount": 2,
      "ExtraParBlocks": 0,
      "Parameters": [
        {
          "Name": "drone",
          "Value": "a30d329784524b31b7fca21c6bccdb70"
        },
        {
          "Name": "*Unpack:",
          "Value": "yes"
        }
      ],
      "ScriptStatuses": [
      ],
      "ServerStats": [
      ],
      "PostInfoText": "NONE",
      "PostStageProgress": -1639621749,
      "PostStageTimeSec": 0,
      "Log": [
      ]
    },
    {
      "FirstID": 15782,
      "LastID": 15782,
      "RemainingSizeLo": 4161601052,
      "RemainingSizeHi": 0,
      "RemainingSizeMB": 3968,
      "PausedSizeLo": 424895108,
      "PausedSizeHi": 0,
      "PausedSizeMB": 405,
      "RemainingFileCount": 100,
      "RemainingParCount": 22,
      "MinPriority": 0,
      "MaxPriority": 0,
      "ActiveDownloads": 0,
      "Status": "QUEUED",
      "NZBID": 15782,
      "NZBName": "My_File_Name_3",
      "NZBNicename": "My_File_Name_3",
      "Kind": "NZB",
      "URL": "",
      "NZBFilename": "My_File_Name_3.nzb",
      "DestDir": "\/intermediate\/My_File_Name_3.#15782",
      "FinalDir": "",
      "Category": "TV",
      "ParStatus": "NONE",
      "ExParStatus": "NONE",
      "UnpackStatus": "NONE",
      "MoveStatus": "NONE",
      "ScriptStatus": "NONE",
      "DeleteStatus": "NONE",
      "MarkStatus": "NONE",
      "UrlStatus": "NONE",
      "FileSizeLo": 4161601052,
      "FileSizeHi": 0,
      "FileSizeMB": 3968,
      "FileCount": 100,
      "MinPostTime": 1477895316,
      "MaxPostTime": 1477895316,
      "TotalArticles": 5332,
      "SuccessArticles": 0,
      "FailedArticles": 0,
      "Health": 1000,
      "CriticalHealth": 886,
      "DupeKey": "",
      "DupeScore": 0,
      "DupeMode": "ALL",
      "Deleted": false,
      "DownloadedSizeLo": 0,
      "DownloadedSizeHi": 0,
      "DownloadedSizeMB": 0,
      "DownloadTimeSec": 0,
      "PostTotalTimeSec": 0,
      "ParTimeSec": 0,
      "RepairTimeSec": 0,
      "UnpackTimeSec": 0,
      "MessageCount": 2,
      "ExtraParBlocks": 0,
      "Parameters": [
        {
          "Name": "drone",
          "Value": "3c985d62c8a1459d81712f7477f25017"
        },
        {
          "Name": "*Unpack:",
          "Value": "yes"
        }
      ],
      "ScriptStatuses": [
      ],
      "ServerStats": [
      ],
      "PostInfoText": "NONE",
      "PostStageProgress": -1639621749,
      "PostStageTimeSec": 0,
      "Log": [
      ]
    }
  ]
}`
	status = `{
  "version": "1.1",
  "result": {
    "ArticleCacheHi": 0,
    "ArticleCacheLo": 0,
    "ArticleCacheMB": 0,
    "AverageDownloadRate": 46812257,
    "DaySizeHi": 25,
    "DaySizeLo": 2156123406,
    "DaySizeMB": 104456,
    "Download2Paused": false,
    "DownloadedSizeHi": 41,
    "DownloadedSizeLo": 1933356398,
    "DownloadedSizeMB": 169779,
    "DownloadLimit": 0,
    "DownloadPaused": false,
    "DownloadRate": 0,
    "DownloadTimeSec": 3803,
    "FeedActive": false,
    "ForcedSizeHi": 0,
    "ForcedSizeLo": 0,
    "ForcedSizeMB": 0,
    "FreeDiskSpaceHi": 1022,
    "FreeDiskSpaceLo": 1309999104,
    "FreeDiskSpaceMB": 4187361,
    "MonthSizeHi": 47,
    "MonthSizeLo": 2286421136,
    "MonthSizeMB": 194692,
    "NewsServers": [
      {
        "ID": 1,
        "Active": true
      },
      {
        "ID": 2,
        "Active": false
      },
      {
        "ID": 3,
        "Active": false
      }
    ],
    "ParJobCount": 3,
    "PostJobCount": 3,
    "PostPaused": false,
    "QueueScriptCount": 0,
    "QuotaReached": false,
    "RemainingSizeHi": 0,
    "RemainingSizeLo": 0,
    "RemainingSizeMB": 0,
    "ResumeTime": 0,
    "ScanPaused": false,
    "ServerPaused": false,
    "ServerStandBy": true,
    "ServerTime": 1589687531,
    "ThreadCount": 52,
    "UpTimeSec": 1036715,
    "UrlCount": 0
  }
}`
	history = `{
  "version": "1.1",
  "result": [
    {
      "ID": 15846,
      "Name": "My_File_1",
      "RemainingFileCount": 0,
      "RetryData": false,
      "HistoryTime": 1589707990,
      "Status": "SUCCESS/HEALTH",
      "Log": [],
      "NZBID": 15846,
      "NZBName": "My_File_1",
      "NZBNicename": "My_File_1",
      "Kind": "NZB",
      "URL": "",
      "NZBFilename": "/My_File_1.nzb",
      "DestDir": "/Dest/My_File_1.nzb",
      "FinalDir": "/MyDir/My_File_1.nzb",
      "Category": "Comics",
      "ParStatus": "NONE",
      "ExParStatus": "NONE",
      "UnpackStatus": "NONE",
      "MoveStatus": "SUCCESS",
      "ScriptStatus": "NONE",
      "DeleteStatus": "NONE",
      "MarkStatus": "NONE",
      "UrlStatus": "NONE",
      "FileSizeLo": 31871884,
      "FileSizeHi": 0,
      "FileSizeMB": 30,
      "FileCount": 1,
      "MinPostTime": 1469263988,
      "MaxPostTime": 1469263988,
      "TotalArticles": 49,
      "SuccessArticles": 49,
      "FailedArticles": 0,
      "Health": 1000,
      "CriticalHealth": 1000,
      "DupeKey": "",
      "DupeScore": 0,
      "DupeMode": "SCORE",
      "Deleted": false,
      "DownloadedSizeLo": 31837060,
      "DownloadedSizeHi": 0,
      "DownloadedSizeMB": 30,
      "DownloadTimeSec": 3,
      "PostTotalTimeSec": 1,
      "ParTimeSec": 0,
      "RepairTimeSec": 0,
      "UnpackTimeSec": 0,
      "MessageCount": 11,
      "ExtraParBlocks": 0,
      "Parameters": [
        {
          "Name": "*Unpack:",
          "Value": "no"
        }
      ],
      "ScriptStatuses": [
      ],
      "ServerStats": [
        {
          "ServerID": 1,
          "SuccessArticles": 49,
          "FailedArticles": 0
        }
      ]
    },
    {
      "ID": 15845,
      "Name": "My_File_2",
      "RemainingFileCount": 0,
      "RetryData": false,
      "HistoryTime": 1589707977,
      "Status": "SUCCESS/HEALTH",
      "Log": [],
      "NZBID": 15845,
      "NZBName": "My_File_2",
      "NZBNicename": "My_File_2",
      "Kind": "NZB",
      "URL": "",
      "NZBFilename": "\/config\/mylar\/cache\/My_File_2.nzb",
      "DestDir": "/Dest/My_File_2",
      "FinalDir": "",
      "Category": "Comics",
      "ParStatus": "NONE",
      "ExParStatus": "NONE",
      "UnpackStatus": "NONE",
      "MoveStatus": "SUCCESS",
      "ScriptStatus": "NONE",
      "DeleteStatus": "NONE",
      "MarkStatus": "NONE",
      "UrlStatus": "NONE",
      "FileSizeLo": 35237544,
      "FileSizeHi": 0,
      "FileSizeMB": 33,
      "FileCount": 1,
      "MinPostTime": 1469286370,
      "MaxPostTime": 1469286370,
      "TotalArticles": 54,
      "SuccessArticles": 54,
      "FailedArticles": 0,
      "Health": 1000,
      "CriticalHealth": 1000,
      "DupeKey": "",
      "DupeScore": 0,
      "DupeMode": "SCORE",
      "Deleted": false,
      "DownloadedSizeLo": 35193998,
      "DownloadedSizeHi": 0,
      "DownloadedSizeMB": 33,
      "DownloadTimeSec": 3,
      "PostTotalTimeSec": 1,
      "ParTimeSec": 0,
      "RepairTimeSec": 0,
      "UnpackTimeSec": 0,
      "MessageCount": 11,
      "ExtraParBlocks": 0,
      "Parameters": [
        {
          "Name": "*Unpack:",
          "Value": "no"
        }
      ],
      "ScriptStatuses": [
      ],
      "ServerStats": [
        {
          "ServerID": 1,
          "SuccessArticles": 54,
          "FailedArticles": 0
        }
      ]
    }
  ]
}`
)

var _ = Describe("NZBGet", func() {

	Context("#Config", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/jsonrpc/config").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(config)
			})

			It("should return a config map", func() {
				client, err := nzbget.New(nzbgetURL, "user", "password")
				Expect(err).ToNot(HaveOccurred())
				config, err := client.Config()
				Expect(err).ToNot(HaveOccurred())
				Expect(config["ConfigFile"]).To(Equal(`/config/nzbget.conf`))
			})

		})
	})

	Context("#ServerVolumes", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/jsonrpc/servervolumes").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(serverVolumes)
			})

			It("should return the server volumes", func() {
				client, err := nzbget.New(nzbgetURL, "user", "password")
				Expect(err).ToNot(HaveOccurred())
				volumes, err := client.ServerVolumes()
				Expect(err).ToNot(HaveOccurred())
				Expect(len(volumes)).To(Equal(2))
			})

		})
	})

	Context("#FileGroups", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/jsonrpc/listgroups").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(listGroups)
			})

			It("should return the file groups", func() {
				client, err := nzbget.New(nzbgetURL, "user", "password")
				Expect(err).ToNot(HaveOccurred())
				fileGroups, err := client.FileGroups()
				Expect(err).ToNot(HaveOccurred())
				Expect(len(fileGroups)).To(Equal(3))
			})
		})
	})

	Context("#Status", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/jsonrpc/status").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(status)
			})

			It("should return the status", func() {
				client, err := nzbget.New(nzbgetURL, "user", "password")
				Expect(err).ToNot(HaveOccurred())
				status, err := client.Status()
				Expect(err).ToNot(HaveOccurred())
				Expect(*status).To(MatchFields(IgnoreExtras, Fields{
					"DownloadedSizeLo":    Equal(1933356398),
					"DownloadedSizeHi":    Equal(41),
					"DownloadedSizeMB":    Equal(169779),
					"MonthSizeLo":         Equal(2286421136),
					"MonthSizeHi":         Equal(47),
					"MonthSizeMB":         Equal(194692),
					"DaySizeLo":           Equal(2156123406),
					"DaySizeHi":           Equal(25),
					"DaySizeMB":           Equal(104456),
					"AverageDownloadRate": Equal(46812257),
					"DownloadLimit":       Equal(0),
					"ThreadCount":         Equal(52),
					"ParJobCount":         Equal(3),
					"PostJobCount":        Equal(3),
					"URLCount":            Equal(0),
					"UpTimeSec":           Equal(1036715),
					"DownloadTimeSec":     Equal(3803),
					"ServerPaused":        Equal(false),
					"DownloadPaused":      Equal(false),
					"Download2Paused":     Equal(false),
					"ServerStandBy":       Equal(true),
					"PostPaused":          Equal(false),
					"ScanPaused":          Equal(false),
					"QuotaReached":        Equal(false),
					"FreeDiskSpaceLo":     Equal(1309999104),
					"FreeDiskSpaceHi":     Equal(1022),
					"FreeDiskSpaceMB":     Equal(4187361),
					"ServerTime":          Equal(1589687531),
					"ResumeTime":          Equal(0),
					"FeedActive":          Equal(false),
					"QueueScriptCount":    Equal(0),
				}))
			})
		})
	})

	Context("#History", func() {
		Context("successful", func() {
			AfterEach(func() {
				gock.Off()
			})

			BeforeEach(func() {
				gock.New(nzbgetURL).
					Get("/jsonrpc/history").
					MatchParams(map[string]string{}).
					Reply(200).
					JSON(history)
			})

			It("should return the history", func() {
				client, err := nzbget.New(nzbgetURL, "user", "password")
				Expect(err).ToNot(HaveOccurred())
				history, err := client.History()
				Expect(err).ToNot(HaveOccurred())
				Expect(len(history)).To(Equal(2))
			})
		})
	})
})
