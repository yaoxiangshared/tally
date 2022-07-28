package config

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/urfave/cli/v2"
)

// Database drivers (sql dialects).
const (
	MySQL    = "mysql"
	MariaDB  = "mariadb"
	SQLite   = "sqlite3"
	Postgres = "postgres" // TODO: Requires GORM 2.0 for generic column data types
)

// Options provides a struct in which application configuration is stored.
// Application code must use functions to get config options, for two reasons:
//
// 1. Some options are computed and we don't want to leak implementation details (aims at reducing refactoring overhead).
//
// 2. Paths might actually be dynamic later (if we build a multi-user version).
//
// See https://github.com/photoprism/photoprism/issues/50#issuecomment-433856358
type Options struct {
	Name                  string  `json:"-"`
	Version               string  `json:"-"`
	Copyright             string  `json:"-"`
	Debug                 bool    `yaml:"Debug" json:"Debug" flag:"debug"`
	Test                  bool    `yaml:"-" json:"Test,omitempty" flag:"test"`
	Demo                  bool    `yaml:"Demo" json:"-" flag:"demo"`
	Sponsor               bool    `yaml:"-" json:"-" flag:"sponsor"`
	Public                bool    `yaml:"Public" json:"-" flag:"public"`
	ReadOnly              bool    `yaml:"ReadOnly" json:"ReadOnly" flag:"read-only"`
	Experimental          bool    `yaml:"Experimental" json:"Experimental" flag:"experimental"`
	ConfigPath            string  `yaml:"ConfigPath" json:"-" flag:"config-path"`
	ConfigFile            string  `json:"-"`
	AdminPassword         string  `yaml:"AdminPassword" json:"-" flag:"admin-password"`
	OriginalsPath         string  `yaml:"OriginalsPath" json:"-" flag:"originals-path"`
	OriginalsLimit        int64   `yaml:"OriginalsLimit" json:"OriginalsLimit" flag:"originals-limit"`
	ImportPath            string  `yaml:"ImportPath" json:"-" flag:"import-path"`
	StoragePath           string  `yaml:"StoragePath" json:"-" flag:"storage-path"`
	SidecarPath           string  `yaml:"SidecarPath" json:"-" flag:"sidecar-path"`
	TempPath              string  `yaml:"TempPath" json:"-" flag:"temp-path"`
	BackupPath            string  `yaml:"BackupPath" json:"-" flag:"backup-path"`
	AssetsPath            string  `yaml:"AssetsPath" json:"-" flag:"assets-path"`
	CachePath             string  `yaml:"CachePath" json:"-" flag:"cache-path"`
	Workers               int     `yaml:"Workers" json:"Workers" flag:"workers"`
	WakeupInterval        int     `yaml:"WakeupInterval" json:"WakeupInterval" flag:"wakeup-interval"`
	AutoIndex             int     `yaml:"AutoIndex" json:"AutoIndex" flag:"auto-index"`
	AutoImport            int     `yaml:"AutoImport" json:"AutoImport" flag:"auto-import"`
	DisableBackups        bool    `yaml:"DisableBackups" json:"DisableBackups" flag:"disable-backups"`
	DisableWebDAV         bool    `yaml:"DisableWebDAV" json:"DisableWebDAV" flag:"disable-webdav"`
	DisableSettings       bool    `yaml:"DisableSettings" json:"-" flag:"disable-settings"`
	DisablePlaces         bool    `yaml:"DisablePlaces" json:"DisablePlaces" flag:"disable-places"`
	DisableExifTool       bool    `yaml:"DisableExifTool" json:"DisableExifTool" flag:"disable-exiftool"`
	DisableFFmpeg         bool    `yaml:"DisableFFmpeg" json:"DisableFFmpeg" flag:"disable-ffmpeg"`
	DisableDarktable      bool    `yaml:"DisableDarktable" json:"DisableDarktable" flag:"disable-darktable"`
	DisableRawtherapee    bool    `yaml:"DisableRawtherapee" json:"DisableRawtherapee" flag:"disable-rawtherapee"`
	DisableSips           bool    `yaml:"DisableSips" json:"DisableSips" flag:"disable-sips"`
	DisableHeifConvert    bool    `yaml:"DisableHeifConvert" json:"DisableHeifConvert" flag:"disable-heifconvert"`
	DisableTensorFlow     bool    `yaml:"DisableTensorFlow" json:"DisableTensorFlow" flag:"disable-tensorflow"`
	DisableFaces          bool    `yaml:"DisableFaces" json:"DisableFaces" flag:"disable-faces"`
	DisableClassification bool    `yaml:"DisableClassification" json:"DisableClassification" flag:"disable-classification"`
	DetectNSFW            bool    `yaml:"DetectNSFW" json:"DetectNSFW" flag:"detect-nsfw"`
	UploadNSFW            bool    `yaml:"UploadNSFW" json:"-" flag:"upload-nsfw"`
	LogLevel              string  `yaml:"LogLevel" json:"-" flag:"log-level"`
	LogFilename           string  `yaml:"LogFilename" json:"-" flag:"log-filename"`
	PIDFilename           string  `yaml:"PIDFilename" json:"-" flag:"pid-filename"`
	CdnUrl                string  `yaml:"CdnUrl" json:"CdnUrl" flag:"cdn-url"`
	SiteUrl               string  `yaml:"SiteUrl" json:"SiteUrl" flag:"site-url"`
	SitePreview           string  `yaml:"SitePreview" json:"SitePreview" flag:"site-preview"`
	SiteTitle             string  `yaml:"SiteTitle" json:"SiteTitle" flag:"site-title"`
	SiteCaption           string  `yaml:"SiteCaption" json:"SiteCaption" flag:"site-caption"`
	SiteDescription       string  `yaml:"SiteDescription" json:"SiteDescription" flag:"site-description"`
	SiteAuthor            string  `yaml:"SiteAuthor" json:"SiteAuthor" flag:"site-author"`
	DatabaseDriver        string  `yaml:"DatabaseDriver" json:"-" flag:"database-driver"`
	DatabaseDsn           string  `yaml:"DatabaseDsn" json:"-" flag:"database-dsn"`
	DatabaseServer        string  `yaml:"DatabaseServer" json:"-" flag:"database-server"`
	DatabaseName          string  `yaml:"DatabaseName" json:"-" flag:"database-name"`
	DatabaseUser          string  `yaml:"DatabaseUser" json:"-" flag:"database-user"`
	DatabasePassword      string  `yaml:"DatabasePassword" json:"-" flag:"database-password"`
	DatabaseConns         int     `yaml:"DatabaseConns" json:"-" flag:"database-conns"`
	DatabaseConnsIdle     int     `yaml:"DatabaseConnsIdle" json:"-" flag:"database-conns-idle"`
	HttpHost              string  `yaml:"HttpHost" json:"-" flag:"http-host"`
	HttpPort              int     `yaml:"HttpPort" json:"-" flag:"http-port"`
	HttpMode              string  `yaml:"HttpMode" json:"-" flag:"http-mode"`
	HttpCompression       string  `yaml:"HttpCompression" json:"-" flag:"http-compression"`
	RawPresets            bool    `yaml:"RawPresets" json:"RawPresets" flag:"raw-presets"`
	DarktableBin          string  `yaml:"DarktableBin" json:"-" flag:"darktable-bin"`
	DarktableBlacklist    string  `yaml:"DarktableBlacklist" json:"-" flag:"darktable-blacklist"`
	RawtherapeeBin        string  `yaml:"RawtherapeeBin" json:"-" flag:"rawtherapee-bin"`
	RawtherapeeBlacklist  string  `yaml:"RawtherapeeBlacklist" json:"-" flag:"rawtherapee-blacklist"`
	SipsBin               string  `yaml:"SipsBin" json:"-" flag:"sips-bin"`
	HeifConvertBin        string  `yaml:"HeifConvertBin" json:"-" flag:"heifconvert-bin"`
	FFmpegBin             string  `yaml:"FFmpegBin" json:"-" flag:"ffmpeg-bin"`
	FFmpegEncoder         string  `yaml:"FFmpegEncoder" json:"FFmpegEncoder" flag:"ffmpeg-encoder"`
	FFmpegBitrate         int     `yaml:"FFmpegBitrate" json:"FFmpegBitrate" flag:"ffmpeg-bitrate"`
	FFmpegBuffers         int     `yaml:"FFmpegBuffers" json:"FFmpegBuffers" flag:"ffmpeg-buffers"`
	ExifToolBin           string  `yaml:"ExifToolBin" json:"-" flag:"exiftool-bin"`
	DetachServer          bool    `yaml:"DetachServer" json:"-" flag:"detach-server"`
	DownloadToken         string  `yaml:"DownloadToken" json:"-" flag:"download-token"`
	PreviewToken          string  `yaml:"PreviewToken" json:"-" flag:"preview-token"`
	ThumbFilter           string  `yaml:"ThumbFilter" json:"ThumbFilter" flag:"thumb-filter"`
	ThumbUncached         bool    `yaml:"ThumbUncached" json:"ThumbUncached" flag:"thumb-uncached"`
	ThumbSize             int     `yaml:"ThumbSize" json:"ThumbSize" flag:"thumb-size"`
	ThumbSizeUncached     int     `yaml:"ThumbSizeUncached" json:"ThumbSizeUncached" flag:"thumb-size-uncached"`
	JpegSize              int     `yaml:"JpegSize" json:"JpegSize" flag:"jpeg-size"`
	JpegQuality           int     `yaml:"JpegQuality" json:"JpegQuality" flag:"jpeg-quality"`
	FaceSize              int     `yaml:"-" json:"-" flag:"face-size"`
	FaceScore             float64 `yaml:"-" json:"-" flag:"face-score"`
	FaceOverlap           int     `yaml:"-" json:"-" flag:"face-overlap"`
	FaceClusterCore       int     `yaml:"-" json:"-" flag:"face-cluster-core"`
	FaceClusterDist       float64 `yaml:"-" json:"-" flag:"face-cluster-dist"`
	FaceMatchDist         float64 `yaml:"-" json:"-" flag:"face-match-dist"`
}

// NewOptions creates a new configuration entity by using two methods:
//
// 1. Load: This will initialize options from a yaml config file.
//
// 2. SetContext: Which comes after Load and overrides
//    any previous options giving an option two override file configs through the CLI.
func NewOptions(ctx *cli.Context) *Options {
	c := &Options{}
	//fmt.Printf("%#v\n", ctx)
	if ctx == nil {
		return c
	}

	fmt.Printf("%v\n", ctx.App.Name)
	c.Name = ctx.App.Name
	c.Copyright = ctx.App.Copyright
	c.Version = ctx.App.Version

	return c
}
