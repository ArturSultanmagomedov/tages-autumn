package cache

//TODO: подключить настоящий кеш

var fileListCache []string

func GetFileListCache() []string {
	return fileListCache
}

func UpdateFileListCache(newValue []string) {
	fileListCache = newValue
}
