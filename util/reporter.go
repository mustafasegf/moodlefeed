package util

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/google/go-cmp/cmp"
)

type DiffReporter struct {
	path  cmp.Path
	diffs []StructIndex
}

type StructIndex struct {
	Resource int
	Modules  int
}

func (r *DiffReporter) PushStep(ps cmp.PathStep) {
	r.path = append(r.path, ps)
}

func (r *DiffReporter) Report(rs cmp.Result) {
	if !rs.Equal() {
		re := regexp.MustCompile("\\{entity\\.Resource\\}\\.Resource\\[(.*)\\]\\.Modules\\[(.*)\\]\\.Description")
		match := re.FindStringSubmatch(fmt.Sprintf("%#v", r.path))
		if len(match) >= 3 {
			res, _ := strconv.Atoi(match[1])
			mod, _ := strconv.Atoi(match[2])
			data := StructIndex{
				Resource: res, 
				Modules: mod,
			}
			r.diffs = append(r.diffs, data)
		}
	}
}

func (r *DiffReporter) PopStep() {
	r.path = r.path[:len(r.path)-1]
}

func (r *DiffReporter) GetDiff() []StructIndex {
	return r.diffs
}
