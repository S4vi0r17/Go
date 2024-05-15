package main

import "time"

// file types
const (
	regularFile int = iota
	directoryFile
	executableFile
	compressedFile
	imageFile
	linkFile
)

// file extension
const (
	exe = ".exe"
	deb = ".deb"
	zip = ".zip"
	gz  = ".gz"
	tar = ".tar"
	rar = ".rar"
	png = ".png"
	jpg = ".jpg"
	gif = ".gif"
)

type file struct {
	name      string
	fileType  int
	isDir     bool
	isHidden  bool
	userName  string
	groupName string
	size      int64
	modTime   time.Time
	mode      string
}

type styleFileType struct {
	icon   string
	color  string
	symbol string
}

var fileTypes = map[int]styleFileType{
	regularFile:    {icon: "📄"},
	directoryFile:  {icon: "📁", color: "BLUE", symbol: "/"},
	executableFile: {icon: "🚀", color: "GREEN", symbol: "*"},
	compressedFile: {icon: "📦", color: "RED"},
	imageFile:      {icon: "🖼️", color: "MAGENTA"},
	linkFile:       {icon: "🔗", color: "CYAN"},
}
