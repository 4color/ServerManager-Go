package globel

type TailSturct struct {
	Id   string
	Path string
	Cid  chan int
}

var TailFiles = []TailSturct{}

func TailFilesIsContain(ff string) bool {

	for _, eachItem := range TailFiles {
		if eachItem.Path == ff {
			return true
		}
	}
	return false
}

func TailFilesGet(id string) TailSturct {

	ta := TailSturct{}
	for _, eachItem := range TailFiles {
		if eachItem.Id == id {
			ta = eachItem
			return ta
		}
	}
	return ta
}

func TailFilesAdd(id string, ff string, cid chan int) {
	if !TailFilesIsContain(ff) {
		ta := TailSturct{}
		ta.Id = id
		ta.Path = ff
		ta.Cid = cid
		TailFiles = append(TailFiles, ta)
	}
}

func TailFilesDelet(id string) {
	for i := 0; i < len(TailFiles); i++ {
		if TailFiles[i].Id == id {
			TailFiles = append(TailFiles[:i], TailFiles[i+1:]...)
		}
	}
}
