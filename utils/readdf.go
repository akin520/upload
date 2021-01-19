package utils

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"time"
)

type Filetype struct {
	Name string
	//Time time.Time
	Time int64
	Type string
	Size int64
}

func (f Filetype) GetTime() string {
	return time.Unix(f.Time, 0).Format("2006-01-02 15:04")
}

func (f Filetype) GetSize() string {
	if f.Size < 1024 {
		return fmt.Sprintf("%d B", f.Size)
	} else if f.Size > 1024 && f.Size <= (1024*1024) {
		return fmt.Sprintf("%.2f KB", float64(f.Size)*1.0/1024)
	} else if f.Size > (1024*1024) && f.Size <= (1024*1024*1024) {
		return fmt.Sprintf("%.2f MB", float64(f.Size)*1.0/(1024*1024))
	} else {
		return fmt.Sprintf("%.2f GB", float64(f.Size)*1.0/(1024*1024*1024))
	}
}
func ReadDF(path string) ([]Filetype, []Filetype, error) {
	dirtype := make([]Filetype, 0, 100)
	filetype := make([]Filetype, 0, 100)
	fileInfoList, err := ioutil.ReadDir(path)
	// fmt.Println(path, fileInfoList, err)
	if err != nil {
		return dirtype, filetype, err
	}
	for _, i := range fileInfoList {
		if i.IsDir() {
			dir := Filetype{Name: i.Name(), Time: i.ModTime().Unix(), Type: "文件夹", Size: 0}
			dirtype = append(dirtype, dir)
		} else {
			suffix := strings.Split(i.Name(), ".")
			file := Filetype{Name: i.Name(), Time: i.ModTime().Unix(), Type: suffix[len(suffix)-1], Size: i.Size()}
			filetype = append(filetype, file)
		}
	}
	sort.Slice(dirtype, func(i, j int) bool {
		return dirtype[i].Time > dirtype[j].Time
	})
	sort.Slice(filetype, func(i, j int) bool {
		return filetype[i].Time > filetype[j].Time
	})
	return dirtype, filetype, nil
}

func ConvertToTime(int_time int64) string {
	return time.Unix(int_time, 0).Format("2006-01-02 15:04")
}
