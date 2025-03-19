package commonConfig

type Settings struct {
	Application struct {
		Name        string `toml:"name"`
		Prefix      string `toml:"prefix"`
		Home        string `toml:"home"`
		Description string `toml:"description"`
		Version     string `toml:"version"`
		Environment string `toml:"environment"`
		ReleaseDate string `toml:"releaseDate"`
		Copyright   string `toml:"copyright"`
		Author      string `toml:"author"`
		License     string `toml:"license"`
		Locale      string `toml:"locale"`
	} `toml:"Application"`

	Database struct {
		Version  int    `toml:"version"`
		Type     string `toml:"type"`
		Host     string `toml:"host"`
		Port     int    `toml:"port"`
		Name     string `toml:"name"`
		User     string `toml:"user"`
		Pass     string `toml:"pass"`
		Path     string `toml:"path"`
		PoolSize int    `toml:"poolSize"`
		Timeout  int    `toml:"timeout"`
	} `toml:"Database"`

	Server struct {
		Host        string `toml:"host"`
		Port        int    `toml:"port"`
		Protocol    string `toml:"protocol"`
		Environment string `toml:"environment"`
	} `toml:"Server"`

	Message struct {
		Keys struct {
			Type    string `toml:"type"`
			Title   string `toml:"title"`
			Content string `toml:"content"`
			Action  string `toml:"action"`
		} `toml:"Keys"`
	} `toml:"Message"`

	Translation struct {
		Host      string `toml:"host"`
		Port      int    `toml:"port"`
		Locale    string `toml:"locale"`
		Protocol  string `toml:"protocol"`
		Permitted struct {
			Origins []struct {
				Name string `toml:"name"`
			} `toml:"Origins"`
			Locales []struct {
				Key  string `toml:"key"`
				Name string `toml:"name"`
			} `toml:"Locales"`
		} `toml:"Permitted"`
	} `toml:"Translation"`

	Assets struct {
		Logo    string `toml:"logo"`
		Favicon string `toml:"favicon"`
	} `toml:"Assets"`

	Dates struct {
		Formats struct {
			DateTime     string `toml:"dateTime"`
			Date         string `toml:"date"`
			Time         string `toml:"time"`
			Backup       string `toml:"backup"`
			BackupFolder string `toml:"backupFolder"`
			Human        string `toml:"human"`
			DMY2         string `toml:"dmy2"`
			YMD          string `toml:"ymd"`
			Internal     string `toml:"internal"`
		} `toml:"Formats"`
	} `toml:"Dates"`

	History struct {
		MaxEntries int `toml:"maxEntries"`
	} `toml:"History"`
	Hosts []struct {
		Name string `toml:"name"`
		FQDN string `toml:"fqdn"`
		IP   string `toml:"ip"`
		Zone string `toml:"zone"`
	} `toml:"Hosts"`
	Security struct {
		Sessions struct {
			ExpiryPeriod int `toml:"expiryPeriod"`
			Keys         struct {
				Session      string `toml:"session"`
				UserKey      string `toml:"userKey"`
				UserCode     string `toml:"userCode"`
				Token        string `toml:"token"`
				ExpiryPeriod string `toml:"expiryPeriod"`
			} `toml:"Keys"`
		} `toml:"Sessions"`
		Service struct {
			UserUID  string `toml:"userUID"`
			UserName string `toml:"userName"`
		} `toml:"Service"`
	} `toml:"Security"`

	Display struct {
		Delimiter string `toml:"delim"`
	} `toml:"Display"`

	Status struct {
		UNKNOWN string `toml:"unknown"`
		ONLINE  string `toml:"online"`
		OFFLINE string `toml:"offline"`
		ERROR   string `toml:"error"`
		WARNING string `toml:"warning"`
	} `toml:"Status"`

	Communications struct {
		Pushover struct {
			UserKey  string `toml:"userKey"`
			APIToken string `toml:"apiToken"`
		} `toml:"Pushover"`
		Email struct {
			Host     string `toml:"host"`
			Port     int    `toml:"port"`
			User     string `toml:"user"`
			Password string `toml:"password"`
			From     string `toml:"from"`
			Footer   string `toml:"footer"`
			Admin    string `toml:"admin"`
		} `toml:"Email"`
	} `toml:"Communications"`
	Logging struct {
		Disable struct {
			General        string `toml:"general"`
			Timing         string `toml:"timing"`
			Service        string `toml:"service"`
			Audit          string `toml:"audit"`
			Translation    string `toml:"translation"`
			Trace          string `toml:"trace"`
			Warning        string `toml:"warning"`
			Event          string `toml:"event"`
			Security       string `toml:"security"`
			Database       string `toml:"database"`
			Api            string `toml:"api"`
			Import         string `toml:"import"`
			Export         string `toml:"export"`
			Communications string `toml:"comms"`
			All            string `toml:"all"`
		} `toml:"disable"`
		Defaults struct {
			MaxSize    string `toml:"maxSize"`
			MaxBackups string `toml:"maxBackups"`
			MaxAge     string `toml:"maxAge"`
			Compress   string `toml:"compress"`
		} `toml:"Defaults"`
	} `toml:"Logging"`

	Backups struct {
		RetainDays int `toml:"retainDays"`
	}
}
