package spider

import (
	"github.com/wolfogre/nest/internal/service/entity"
	"fmt"
	"io/ioutil"
	"time"
	"github.com/wolfogre/nest/internal/service/util/timeformat"
)

func saveJs(infos []*entity.Loupan, filepath string) error {
	temp := make([]*entity.Loupan, 0)
	for _, v := range infos {
		if v.Lat != 0 && v.Lng != 0 {
			temp = append(temp, v)
		}
	}
	infos = temp

	data := ""

	data += `
var map = new AMap.Map("container", {
    resizeEnable: true,
    mapStyle: "amap://styles/dark"
});

var heatmapData = [
`


	now := time.Now()
	for i, v := range infos {
		t := timeformat.ParseDate(v.StartDate)
		count := 60
		if t != nil {
			day := now.Sub(*t).Hours() / 24
			count -= int(day)
		}
		if count < 0 {
			count = 0
		}
		data += fmt.Sprintf(`{"lat":%v, "lng":%v, "count":%v}`, v.Lat, v.Lng, count)
		if i < len(infos) - 1 {
			data += ","
		}
		data += "\n"
	}

	data += `
];

markerList = [
`
	for i, v := range infos {
		data += fmt.Sprintf(`new AMap.Marker({map: map, visible: false, label:{content:'<a onclick="openWindow(\'%v\')">%v<\a>'}, position: [%v,%v]})`, v.Url, v.Name, v.Lng, v.Lat)
		if i < len(infos) - 1 {
			data += ","
		}
		data += "\n"
	}
	data += `
];
`
	return ioutil.WriteFile(filepath, []byte(data), 0666)
}
