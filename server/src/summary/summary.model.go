package summary

import (
	"github.com/ZombieHippie/test-gin/server/src/artifact"
	"github.com/ZombieHippie/test-gin/server/src/repo"
	"github.com/jinzhu/gorm"
	"time"
)

// Summary is a common object for reports to point at
type Summary struct {
	gorm.Model
	Repository repo.Repository
	BranchID   string
	BuildID    int
	Artifacts  []artifact.Artifact
	Commit     string

	Message string
	Author  string
	Success bool
	Created time.Time
}