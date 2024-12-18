package main //声明包名称
import (
	"flag"
	"fmt"
	"github.com/miekg/dns"
	"math/rand"
	"net"
	"sync"
	"time"
)

type DnsCheckDomainData struct {
	Domain    string `json:"domain"`
	Ip        string `json:"ip"`
	DnsServer string `json:"dnsServer"`
	Timer     string `json:"timer"`
}

func Timer() string { // 获取当前时间
	now := time.Now()
	formatted := now.Format("2006-01-02 15:04:05")
	return formatted
}

func RandomString() string {
	length := 32
	const charset = "abcdefghijklmnopqrstuvwxyz1234567890"
	var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func FanCheckWWW(list []DnsCheckDomainData, char string) *int {
	for num, c := range list {
		if c.Domain == char {
			return &num
		}
	}
	return nil
}

// CheckSubDomain 验证子域名
func CheckSubDomain(domain string) []DnsCheckDomainData {
	var result []DnsCheckDomainData
	var domains = []string{"www", "mail", "ftp", "smtp", "pop", "m", "webmail", "pop3", "imap", "localhost", "autodiscover", "admin", "bbs", "test", "mx", "en", "email", "wap", "blog", "oa", "ns1", "vpn", "ns2", "www2", "mysql", "webdisk", "dev", "old", "news", "calendar", "shop", "potala", "mobile", "web", "sip", "mobilemail", "ns", "cpanel", "www1", "whm", "new", "img", "search", "support", "mail2", "media", "files", "e", "video", "app", "secure", "my", "crm", "intranet", "portal", "demo", "api", "beta", "fax", "lyncdiscover", "dns", "images", "db", "staging", "info", "docs", "static", "ns3", "download", "forum", "cms", "cdn", "www3", "wiki", "pda", "dns1", "home", "mail1", "online", "sso", "lists", "webproxy", "office", "dns2", "get", "t", "gis", "proxy", "pic", "edu", "d", "ns4", "cs", "cn", "b2b", "store", "community", "start", "services", "wx", "service", "training", "remote", "health", "help", "vip", "soft", "finance", "photo", "apps", "owa", "login", "es", "s", "ads", "stats", "events", "forums", "sc", "tv", "data", "data.sql", "jobs", "survey", "it", "hr", "sms", "game", "stage", "i", "send", "member", "v", "a", "ww", "sz", "live", "im", "go", "chat", "3g", "gateway", "library", "ftp2", "dialin", "security", "meet", "upload", "w", "blogs", "de", "image", "msoid", "hk", "down", "gmail", "ssh", "fr", "english", "exchange", "so", "av", "cp", "erp", "cloud", "legacy", "ldap", "ad", "sites", "access", "archive", "job", "connect", "hq", "alumni", "downloads", "extranet", "lib", "tools", "careers", "wt", "ask", "student", "host", "ns5", "helpdesk", "u", "vc", "status", "direct", "rs", "hb", "sp", "mailhost", "uk", "netmang", "svn", "tz", "ms", "sx", "zb", "metrics", "mx2", "hd", "ss", "qa", "dj", "www4", "bm", "jp", "file", "project", "club", "dl", "feeds", "IN", "ent", "ws", "food", "update", "tj", "book", "as", "gb", "lab", "1", "dx", "js", "b", "sh", "cnc", "videos", "dns3", "outlook", "software", "auth", "tw", "preview", "hs", "git", "content", "press", "ir", "cq", "projects", "car", "monitor", "backup", "meeting", "c", "photos", "games", "radio", "gw", "public", "buy", "ssl", "mall", "research", "groups", "pt", "wwww", "forms", "music", "cx", "cj", "techmang", "bj", "math", "res", "mang", "open", "ntp", "w3", "biz", "ca", "map", "ru", "design", "share", "th", "relay", "house", "in", "vpn2", "x", "citrix", "labs", "pub", "education", "jw", "global", "tp", "card", "jk", "f", "gh", "sd", "www5", "rt", "reg", "us", "weather", "newsletter", "ticket", "server", "irc", "apple", "cache", "youth", "rms", "www0", "mx1", "feedback", "fz", "ams", "wh", "reports", "auto", "travel", "cm", "origin", "account", "site", "cc", "p", "vote", "bt", "ems", "manage", "pms", "dk", "sharepoint", "mssql", "partner", "spam", "lt", "link", "user", "tg", "sg", "business", "xx", "ly", "students", "ts", "fs", "vpn1", "dm", "uc", "digital", "cl", "pages", "abc", "brand", "event", "alpha", "sys", "assets", "0", "members", "money", "mdm", "sales", "stat", "local", "be", "tuan", "marketing", "bugs", "mail3", "dy", "time", "stream", "code", "partners", "view", "da", "g", "tr", "inside", "br", "phx", "st", "dms", "jj", "gallery", "wp", "shopping", "promo", "china", "social", "show", "union", "pm", "web1", "test2", "gc", "ja", "kb", "register", "sq", "pc", "maps", "android", "corp", "wireless", "pos", "ce", "rsc", "jira", "accounts", "customer", "jd", "list", "ec", "adm", "web2", "rtx", "corporate", "flash", "developer", "rss", "tour", "wl", "logger", "smtp2", "sns", "sf", "cas", "directory", "webdev", "temp", "ps", "q", "idc", "sm", "ks", "bc", "hy", "fx", "tech", "bb", "tao", "weixin", "journal", "updates", "agent", "art", "hg", "ic", "mobi", "cd", "usa", "cr", "hao", "desktop", "ex", "ra", "teacher", "r", "ag", "dc", "jwc", "mrtg", "tu", "sj", "ace", "space", "ft", "lp", "fw", "mailgw", "mis", "co", "investors", "bookstore", "sql", "hermes", "ip", "3", "eq", "sandbox", "client", "css", "catalog", "sports", "sentry", "kr", "internal", "ky", "sy", "investor", "mms", "exam", "transfer", "ga", "summer", "conference", "hotel", "dz", "pro", "doc", "eng", "cf", "registration", "whois", "life", "world", "ae", "sts", "mp", "staff", "id", "push", "enterprise", "bi", "webservices", "idp", "pe", "prod", "mailgate", "top", "yy", "zw", "lms", "edm", "linux", "vps", "nj", "mm", "net", "z", "ml", "webcast", "phone", "manager", "cz", "ids", "smtp1", "work", "mars", "listserv", "campus", "special", "ee", "n", "au", "cw", "pay", "dealer", "barracuda", "6", "pp", "acs", "chem", "nl", "w2", "audio", "family", "xa", "multimedia", "jabber", "click", "international", "ct", "resource", "cal", "mailman", "links", "pl", "2", "stock", "movie", "atlas", "fm", "shanghai", "keys", "mi", "love", "product", "ls", "sb", "ch", "wb", "fb", "qc", "intra", "em", "dw", "cert", "hn", "yp", "noc", "fc", "main", "md", "ww2", "streaming", "master", "ma", "daj", "ny", "print", "bg", "nas", "me", "kids", "archives", "wms", "webadmin", "stu", "xb", "h", "xml", "asp", "server1", "order", "sync", "ftp1", "zp", "asia", "hz", "pa", "power", "signup", "hh", "history", "learning", "password", "nc", "edge", "gg", "storage", "hub", "ess", "yun", "opac", "test1", "jupiter", "fms", "123", "xl", "cvs", "crl", "ocs", "bz", "lb", "newsroom", "pf", "webstats", "market", "radius", "cwc", "tk", "int", "dt", "acc", "rd", "jn", "post", "ys", "cis", "se", "ops", "one", "edit", "testing", "xt", "affiliate", "y", "train", "orion", "j", "mb", "build", "developers", "pr", "s1", "edi", "nb", "ds", "pj", "trac", "ac", "mkt", "ci", "gate", "csc", "mailer", "kf", "www7", "golf", "canada", "imp", "gm", "bugzilla", "development", "nanjing", "ar", "insight", "ns0", "zz", "profile", "ns01", "seo", "gs", "aaa", "law", "about", "act", "mercury", "star", "sam", "parking", "apollo", "dev1", "cgi", "police", "energy", "k", "websphere", "weblogic", "admissions", "gd", "sl", "sk", "mc", "www35", "xf", "free", "website", "faculty", "bio", "bh", "earth", "india", "xh", "ph", "wm", "study", "japan", "gov", "xs", "nagios", "conf", "zeus", "at", "smart", "bk", "win", "rc", "tl", "cps", "ice", "pan", "ah", "bak", "team", "tool", "gl", "ideas", "safety", "s2", "baby", "beijing", "db1", "nz", "ie", "dr", "UN", "athena", "bd", "ceshi", "no", "up", "is", "cg", "iphone", "welcome", "red", "domain", "affiliates", "control", "class", "voip", "gps", "cam", "network", "film", "by", "fi", "track", "nt", "arc", "pluto", "jf", "sdc", "www6", "traveler", "mt", "a1", "telecom", "voice", "arch", "redirect", "drive", "billing", "ccs", "dh", "blue", "arts", "its", "phoenix", "wwwdev", "l", "cy", "webtrends", "books", "redmine", "nic", "webcam", "play", "sport", "ims", "tc", "ns02", "engineering", "sun", "web3", "supplier", "syslog", "mg", "o", "www36", "mirror", "products", "bill", "sa", "tfs", "eas", "loopback", "talk", "mp3", "foundation", "tracker", "mail4", "pk", "college", "housing", "hp", "uploads", "www8", "router", "board", "aws", "gy", "oa1", "green", "fl", "mac", "jm", "chemistry", "mailserver", "plus", "group", "kh", "ricard", "beauty", "aa", "antivirus", "repo", "cat", "discover", "pg", "venus", "center", "dragon", "titan", "hi", "dg", "courses", "pki", "box", "portfolio", "vod", "boss", "jl", "astro", "tt", "solutions", "wwws", "solr", "un", "msg", "hybrid", "la", "mtest", "picard", "ios", "sw", "mas", "tel", "nms", "webapp", "km", "firewall", "bank", "recruit", "tms", "www9", "wsj", "source", "emergency", "read", "lm", "am", "gift", "www43", "nursing", "ht", "img2", "rp", "google", "water", "report", "dp", "iris", "analytics", "neptune", "lady", "guide", "img1", "eshop", "ww1", "db2", "csg", "man", "volunteer", "webserver", "pd", "phpmyadmin", "webmaster", "count", "jt", "eis", "mj", "marketplace", "delta", "computer", "java", "comm", "bs", "focus", "horizon", "jg", "pdf", "guest", "upgrade", "vision", "leo", "lg", "monitoring", "active", "psychology", "ariel", "hl", "nano", "cb", "ces", "haosf", "oss", "hc", "sky", "www37", "logs", "fj", "auction", "lw", "banner", "jr", "ep", "jenkins", "webservice", "happy", "mail01", "tx", "gold", "eps", "app1", "translate", "cdn2", "10", "sos", "saturn", "xg", "sis", "sr", "espanol", "form", "sn", "resources", "py", "gz", "ap", "mantis", "vcs", "wordpress", "solar", "chinese", "contact", "mobil", "www41", "planet", "5", "tm", "internet", "splunk", "bug", "physics", "gsa", "studio", "ty", "ccc", "mta", "4", "jc", "15", "cacti", "sale", "xian", "hpc", "ww3", "singapore", "wy", "com", "ren", "adv", "production", "tb", "mag", "audit", "01", "mo", "w1", "jpk", "loghost", "mse", "sus", "counter", "eagle", "geo", "sec", "mh", "acm", "check", "elasticsearch", "bo", "hf", "adserver", "hongkong", "hot", "vm", "struts", "php", "air", "kj", "feed", "02", "m2", "drm", "sbc", "disk", "trip", "mr", "pics", "ping", "accounting", "cars", "gp", "wow", "nn", "europe", "virtual", "ba", "spider", "core", "che", "dd", "gr", "users", "nib", "zend", "sps", "pin", "homer", "os", "registrar", "faq", "fy", "system", "passport", "matrix", "el", "credit", "ea", "ta", "kk", "safe", "join", "ge", "webaccess", "lc", "movies", "qq", "science", "analysis", "taobao", "mx3", "moa", "flow", "payment", "creative", "echo", "ua", "mini", "sas", "facilities", "cdn1", "8", "gk", "wifi", "bms", "dns4", "v2", "real", "news2", "orange", "jjc", "imail", "ras", "src", "publications", "communications", "ecommerce", "bgs", "ns6", "people", "ro", "nexus", "www01", "test3", "xmpp", "france", "ipv6", "oracle", "sem", "eu", "sou", "orders", "retail", "wc", "lx", "op", "xc", "ares", "find", "alert", "hm", "german", "case", "nginx", "wei", "iss", "engage", "hu", "mailbox", "www14", "tiger", "next", "pma", "gt", "nw", "commerce", "backend", "aj", "lh", "econ", "chicago", "fang", "museum", "ab", "si", "wf", "advertising", "nm", "paper", "elastic", "www34", "sv", "www10", "amp", "yt", "dag", "zs", "www12", "ismart", "moon", "wa", "technology", "school", "www11", "ka", "widget", "ck", "pharmacy", "csr", "taiwan", "president", "bus", "mw", "7", "fk", "jh", "lf", "french", "server5", "mob", "ok", "wd", "2013", "18", "biology", "ddh", "hydra", "reader", "city", "shop2", "kc", "om", "xxx", "galaxy", "gatekeeper", "expo", "wg", "rm", "spanish", "drupal", "ln", "database", "9", "voicemail", "statistics", "sanguo", "pv", "mcs", "af", "robot", "et", "ipad", "query", "fd", "traffic", "he", "mnews", "enroll", "s3", "classic", "m1", "www02", "http", "cmp", "cart", "max", "17", "bf", "action", "philosophy", "att", "smg", "psych", "mv", "ai", "channel", "16", "mailing", "li", "was", "publish", "fun", "20", "bl", "nk", "editor", "external", "smtp3", "rds", "eos", "df", "med", "apex", "grad", "sap", "ko", "vpn3", "jboss", "picwww", "emba", "ext", "manyi", "re", "pingan", "classifieds", "03", "korea", "expert", "knowledge", "po", "voyager", "legal", "devel", "newton", "speed", "lj", "pi", "rz", "touch", "fa", "ics", "quiz", "alerts", "srm", "www15", "style", "io", "soc", "index", "vn", "sol", "11", "ws1", "a2", "abs", "healthcare", "economics", "athletics", "applications", "idea", "td", "www31", "kiosk", "ha", "lian", "ssc", "dir", "junshi", "19", "vega", "sjc", "2012", "oc", "www13", "epaper", "kt", "bw", "zj", "pad", "pass", "academic", "14", "akamai", "cisco", "develop", "wk", "london", "static1", "mz", "eco", "ebook", "on", "sirius", "pegasus", "policy", "gemini", "bca", "artemis", "line3", "aries", "yule", "fp", "gaj", "dsp", "2014", "server2", "chang", "tmp", "demo2", "munin", "gjqx", "intl", "angel", "ecs", "mon", "hah", "helios", "config", "asset", "pulse", "rfb", "compass", "discovery", "dcs", "diamond", "charon", "eb", "procurement", "www30", "care", "msa", "bridge", "civil", "lv", "hw", "12", "fgw", "rsj", "aoe1", "sea", "admin2", "we", "wwwww", "13", "za", "gf", "ajax", "mf", "dns0", "casper", "gj", "tn", "teach", "wwwb", "pet", "gamma", "jump", "cds", "stores", "kjj", "er", "abacus", "isis", "young", "broadcast", "pool", "ec2", "www16", "esp", "al", "v2-ag", "visa", "advance", "wpad", "sss", "ali", "piwik", "africa", "callback", "bp", "www33", "img01", "thor", "icon", "allwww", "ao", "page", "zm", "graphics", "mail5", "spain", "success", "pre", "qqmail", "football", "express", "jcb", "noprefix", "christmas", "russian", "swt", "biotech", "verify", "vault", "printing", "rose", "myadmin", "dance", "wwyx", "demo1", "wwv", "today", "explore", "price", "www21", "fusion", "ld", "www17", "testapi", "trade", "wwwa", "wwwd", "wwxu", "caiwu", "aaron", "maths", "wwux", "apt", "aurora", "luna", "ww6", "www26", "error", "console", "static2", "best", "hj", "australia", "www22", "ocean", "paris", "xd", "mx4", "gw1", "silver", "cse", "moss", "fod", "wwvx", "soso", "iphone4s", "wwu", "ww5", "nh", "plan", "cdn1122", "wwvv", "193", "cbs", "chrome", "jy", "polaris", "magic", "www78", "point", "2011", "www32", "164", "cams", "t7", "zabbix", "koa", "dream", "zh", "helix", "adam", "sorry", "hero", "ifeng", "webgame", "bia", "www20", "www18", "darwin", "yahoo", "vs", "huan", "domino", "wwwh", "cpc", "base", "comic", "ww9", "neo", "jb", "long", "cap", "scs", "medicine", "wz", "4g", "cancer", "app2", "ips", "www-2", "match", "www27", "impact", "imc", "fin", "wwb", "germany", "jjw", "oas", "janus", "400", "www29", "adsl", "idm", "party", "ed", "cls", "hera", "led", "virus", "dss", "oma", "digi", "culture", "spark", "mx15", "amazon", "call", "ssp", "su", "isc", "talent", "storm", "csp", "brazil", "www24", "goto", "na", "wwxv", "yc", "ng", "cu", "cad", "fujian", "ws2", "zx", "zhidao", "webcdn", "omega", "apc", "az", "cme", "wwvy", "d9", "wwwf", "trace", "lion", "kz", "emp", "planning", "wwwc", "oem", "all", "wwyv", "wwxy", "export", "cec", "zone", "aps", "prism", "schedule", "fdc", "mexico", "alice", "qr", "wwvu", "il", "gopher", "swj", "quality", "wwuy", "lk", "dv", "odin", "tom", "qt", "chart", "russia", "cdn3", "vps3", "pilot", "wwwl", "b2c", "mmm", "merlin", "advantage", "triton", "aac", "lb1", "wwx", "signin", "static3", "ils", "www19", "notebook", "fire", "sub", "cea", "abracadabra", "np", "wlan", "line4", "bar", "smc", "tag", "wwa", "wwy", "cdc", "mcp", "panda", "wwwfilter", "sina", "sociology", "wwyu", "scm", "friends", "avatar", "vms", "aging", "wwuv", "extra", "shell", "adc", "poseidon", "fast", "wo", "ctc", "wac", "misc", "back", "spa", "v1", "te", "abel", "xz", "wwwv", "zt", "robin", "galileo", "iam", "ve", "wss", "wave", "2009", "wan", "author", "dhcp", "kw", "2010", "hercules", "img3", "test4", "agile", "www28", "messenger", "pb", "anywhere", "quote", "aas", "smtp4", "171", "ecom", "coop", "reseller", "honors", "www520", "guinv", "ethics", "kai", "ons", "image3", "king", "tibet", "ehs", "pandora", "www23", "politics", "lining", "fensike", "125125", "bme", "aba", "cerberus", "mzt", "company", "cts", "srv", "iso", "etc", "italy", "jx", "oasis", "tea", "images2", "bak78", "sds", "pds", "ece", "test5", "scan", "xn", "ais", "aim", "agora", "sme", "clubs", "www25", "das", "provost", "patch", "isp", "jason008", "135", "cognos", "wsc", "rainbow", "99comcn", "321", "rw", "360", "cma", "aq", "domains", "abraham", "baike", "wwwe", "888", "mn", "csm", "scholar", "grid", "psy", "wwwi", "outbound", "bak219", "common", "abragam", "vi", "homewo", "abra", "andromeda", "ui", "trans", "cluster", "tesla", "owl", "xy", "aero", "kangar", "trust", "jcc", "v3", "friend", "ada", "qd", "kp", "theatre", "oscar", "squid", "ota", "einstein", "zy", "mcu", "flight", "graphite", "wind", "187", "blackberry", "yx", "rec", "topaz", "m3", "bike", "challenge", "mail6", "csi", "inventory", "central", "abe", "jz", "review", "hospital", "lcs", "kronos", "coe", "alliance", "personal", "ti", "filter", "vista", "abrams", "bak204", "p2", "prime", "release", "now", "ez", "newyork", "spectrum", "download2", "quan", "pw", "yum", "carbon", "369", "mx11", "3d", "yanbak133", "mx5", "fish", "oms", "logo", "outreach", "adams", "byseg854", "gaia", "bbb", "yaho", "host2", "mx10", "sac", "api2", "sage", "template", "test6", "mu", "bamboo", "154", "wwwg", "summit", "mailhub", "nba", "eg", "management", "w4", "fly", "submit", "boston", "watch", "penguin", "application", "gitlab", "privacy", "premium", "phys", "lnmp", "wwvwv", "brain", "black", "npx", "ieee", "demo3", "coffee", "root", "you", "dns4512", "stmp", "adp", "houston", "weibo", "dali", "bak7", "thailand", "watson", "bn", "mx14", "future", "mwww", "xm", "aquarius", "hello", "advisor", "purchasing", "amber", "194", "xfz", "suppliers", "cntv", "stars", "try", "industry", "itc", "dns4527", "cs2", "ia", "adventure", "lu", "kd", "warehouse", "aag", "cee", "cdn4", "ebs", "budget", "cctv", "a01", "cai", "europa", "panel", "ias", "nova", "pay3", "agenda", "chi", "lotus", "falcon", "fashion", "v9", "url", "img0", "apis", "cafe", "text", "wwvvv", "ibm", "deploy", "bounce", "129", "platform", "snap", "im2", "uranus", "buzz", "apache", "fw1", "vendors", "aal", "holiday", "toolbox", "mmc", "austin", "tick", "fh", "s4", "bob", "berlin", "hyperion", "beacon", "bao", "cosmos", "tele", "do", "dds", "c1", "msn", "aw", "image1", "t2", "dns2138", "3w", "2008", "wedding", "cv", "dolphin", "168", "mai", "ews", "sugar", "merchant", "sqlserver", "232", "ucenter", "nat", "nav", "mails", "antares", "px", "intern", "mes", "f5", "dvd", "park", "html", "beian", "crs", "wall", "private", "ak", "tst", "platinum", "derek", "ux", "fg", "cmc", "dbs", "shop1", "cards", "d1", "wizard", "167", "parts", "vpn01", "server3", "lq", "asc", "4050", "mq", "netmon", "3com", "aqua", "t1", "dell", "zzz", "hades", "moe", "psp", "pss", "radar", "shs", "camera", "oa2", "ev", "isa", "yd", "ts1", "big", "daily", "animal", "admin1", "waimai", "snow", "smail", "tourism", "quantum", "demos", "cst", "hope", "ws3", "ccm", "medusa", "yl", "8u", "big5", "rx", "rr", "cns", "mss", "comet", "alt", "miami", "atlantis", "exmail", "mirrors", "dna", "key", "kg", "sta", "ars", "discuss", "gms", "none", "cs1", "img02", "score", "ii", "wsb", "172", "epc", "interactive", "san", "outlet", "coach", "keyserver", "vtc", "scp", "profiles", "ceres", "icp", "ppp", "crystal", "lynx", "nf", "rsa", "quest", "player", "notify", "rf", "160", "geography", "abraxas", "collection", "clock", "jia", "vp", "franklin", "wenwen", "bobo", "sakai", "women", "yuan", "batman", "nemo", "zjj", "ding", "ucc", "xcb", "pim", "mk", "geology", "c2", "4006", "oauth", "mango", "jazz", "metric", "alex", "turkey", "flv", "k2", "shadow", "news1", "s5", "ccb", "forest", "abbott", "x2", "i2", "sjb", "unix", "maple", "navigator", "xinjiang", "ke", "yj", "novel", "cce", "euler", "ups", "eval", "malaysia", "sci", "skb", "athens", "cygnus", "image2", "thunder", "techsupport", "qz", "moban", "medical", "asm", "ruby", "v7", "api1", "d2", "exp", "model", "gq", "ghost", "poc", "sigma", "wine", "resume", "sjz", "bear", "fox", "bell", "kiwi", "breeze", "metro", "dps", "dorm", "host1", "jerry", "b1", "apm", "cricket", "encore", "ttt", "esb", "jaguar", "cip", "aahl", "login2", "kerberos", "mx13", "dodo", "mds", "sxy", "sphinx", "diy", "run", "e11", "t3", "i1", "xj", "cube", "dn", "msp", "crc", "aardvark", "brown", "castor", "aardwolf", "hambur", "ylc", "agency", "baoming", "w7", "lucky", "cyber", "ben", "1111", "apitest", "wuyw", "spring", "xia", "climate", "dept", "testsite", "lz", "argentina", "asa", "jsw", "c3", "swf", "transport", "aachen", "mail7", "bx", "puma", "an", "mbs", "trs", "bing", "wcs", "extension", "aahz", "386", "jilin", "trial", "virgo", "raven", "s12", "nx", "drc", "arizona", "nd", "pdm", "162", "data1", "sitemap", "yw", "epic", "lin", "aallan", "muse", "som", "loki", "c4", "zk", "rank", "pac", "gx", "github", "dynamic", "lgb", "ll", "advice", "maint", "lead", "ses", "atm", "184", "if", "super", "pat", "pmp", "tlc", "host4", "365", "psa", "ad2", "atlanta", "esd", "ceo", "bsc", "du", "msc", "attach", "w5", "download1", "openapi", "vss", "bis", "oak", "cit", "aaguirre", "libra", "xyz", "rich", "stargate", "genome", "mail10", "light", "mod", "pod", "genesis", "seminar", "abcd", "qinghai", "texas", "mam", "iron", "aberdeen", "gauss", "indigo", "tts", "cmt", "emc", "eric", "activity", "mil", "wu", "lao", "peixun", "yunfu", "assist", "robotics", "sydney", "raptor", "p3", "zero", "rabbit", "eol", "hawk", "ax", "ccp", "netlab", "juno", "aapo", "abramson", "suzhou", "vvvww", "sep", "hudson", "shenzhen", "v5", "shanxi", "e10", "nv", "va", "washington", "sdk", "todo", "mark", "nike", "out", "imgs", "wvux", "jade", "qb", "r1", "bbc", "v4", "192", "mbox", "lan", "opinion", "hebei", "wap2", "wuxw", "cop", "attachments", "img4", "zr", "maxwell", "hhh", "aab", "lisa", "turing", "nanke", "seattle", "clothes", "wuvx", "or", "buffalo", "avalon", "subversion", "nntp", "banking", "zen", "urchin", "interface", "vps2", "esales", "ad1", "yn", "zc", "l2", "p2p", "guangdong", "plant", "maillist", "dls", "aage", "vps1", "tcs", "sell", "pz", "cca", "tianjin", "hx", "gaokao", "fantasy", "fred", "larry", "aadams", "h5", "dl1", "eye", "skynet", "mgt", "i4", "114", "taurus", "more", "token", "mps", "bulletin", "broker", "img7", "album", "svc", "ict", "pacific", "p4", "m4", "gems", "showcase", "f1", "cn2", "qhd", "xing", "windows", "opal", "111", "minerva", "aes", "177", "sample", "tim", "logos", "bigbrother", "zhao", "juniper", "9933", "godaddy", "dnstest", "farm", "jordan", "125", "166", "p1", "pack", "djh", "aragorn", "general", "qh", "img6", "imaging", "cobra", "smtp5", "mapi", "ningxia", "smtp10", "img5", "666", "award", "sga", "h3", "bdf", "chaos", "altair", "abo", "t5", "xk", "host3", "hal", "xq", "g7", "bbs2", "ring", "ur", "fishing", "qy", "humanities", "k3", "jss", "stc", "sunshine", "cie", "spock", "bbs1", "f2", "198", "court", "aims", "jcj", "lamp", "yz", "777", "bts", "ebh", "callisto", "t4", "elite", "ff", "r2", "offline", "vancouver", "setup", "mov", "zwj", "snoopy", "tokyo", "yh", "xiao", "to", "white", "wj", "148", "vl", "ldj", "perseus", "delphi", "1234", "e7", "bq", "chaxun", "fff", "mimi", "eva", "tss", "cn1", "mail9", "youjizz", "ibs", "pdc", "phobos", "tec", "bolg", "wvw", "xmail", "nemesis", "ipm", "yulanyou", "wechat", "jiangxi", "update2", "good", "ford", "land", "message", "seller", "ol", "micro", "xiu", "helper", "qm", "artist", "spf", "agri", "stem", "hc5", "147", "uni", "scc", "sex", "omni", "aasa", "aca", "logon", "a3", "trash", "fanli", "consulting", "lucy", "marvin", "yuyue", "181", "lis", "west", "ykt", "zixun", "um", "fortune", "m5", "img8", "qing", "story", "sim", "mft", "sat", "aap", "v6", "js1", "hainan", "tablet", "dmc", "aad", "xen", "trinity", "child", "install", "ib", "two", "login1", "s01", "purchase", "osiris", "acsvax", "800", "ditu", "smtp7", "chicken", "photon", "allen", "q1", "999", "ivr", "abby", "jie", "mail8", "picture", "david", "mgr", "casa", "socrates", "aacse", "mail12", "east", "jasper", "gi", "cname", "dba", "purple", "agnes", "topic", "baidu2014", "zt2", "124", "hubei", "tongji", "000", "transit", "person", "stone", "ip2", "n2", "rh", "zhaopin", "dcc", "smp", "port", "odyssey", "mail-4", "ipod", "ay", "anhui", "tap", "finaid", "rap", "drp", "e3", "at820", "bailefang", "vegas", "eclipse", "byby", "wuwv", "182", "mx6", "di", "ei", "smile", "camel", "gansu", "sign", "u1", "157", "nokia", "nfs", "soap", "b2", "wap1", "fuke", "test123", "010", "t6", "211", "e1", "cgs", "supply", "djj", "hlj", "a10", "custom", "argon", "titanium", "jingjia", "71", "bei", "197", "guangxi", "dot", "tf", "gandalf", "iq", "ets", "biochem", "lobby", "task", "wuxv", "dq", "tips", "sft", "psc", "think", "reach", "mx7", "dai", "tcm", "lyg", "daohang", "wuxy", "shu", "pptp", "neon", "file2", "1314", "uu", "charlotte", "ebay", "dean", "tuku", "columbus", "feng", "liu", "xin", "imgcdn", "euro", "tomcat", "turbine", "launch", "app01", "abalone", "yb", "jxc", "pinpai", "wvxy", "99", "ye", "samsung", "fu", "cpe", "a02", "g1", "w9", "c5", "wuwy", "149", "bee", "dtm", "ash", "guangzhou", "bach", "xray", "vb", "synergy", "s7", "metis", "chanel", "zeta", "zn", "eureka", "phantom", "sogou", "cha", "wuhan", "shouji", "depot", "ddd", "acp", "wuv", "mgame", "dao", "chef", "eee", "sbs", "fuck", "d3", "555", "988", "ha5", "hua", "proteus", "75", "tps", "zhu", "vis", "rl", "zf", "zg", "qs", "cmcc", "95", "a12", "wuwx", "hunan", "qj", "ya", "niu", "sfs", "honey", "joe", "viper", "cobalt", "cpa", "elvis", "202", "cba", "amd", "ait", "dove", "yang", "smtp6", "21", "100", "vortex", "vv", "laser", "hai", "abricot", "michaelyu", "404", "skin", "other", "tree", "chess", "dl2", "m6", "shoes", "rome", "qg", "gn", "ecc", "clarity", "emma", "monkey", "129979", "pix", "amsterdam", "mir", "acer", "rest", "systems", "esl", "smtp9", "dispatch", "lyris", "abell", "wuxi", "kehu", "piao", "gdsvr", "miao", "mail15", "gu", "srs", "olive", "director", "registry", "yoda", "sprout", "b5", "ut", "joke", "vnet", "broadband", "dam", "117", "guanli", "fanyi", "flower", "g1i8", "hg3", "wxtest", "john", "ku", "italian", "tmail", "biomed", "gem", "his", "alcor", "bd123002", "kaoshi"}
	var workWg sync.WaitGroup
	var semaphore = make(chan struct{}, 20)
	task := make(chan *DnsCheckDomainData, len(domains)+1)
	var ipFan = make(map[string]bool)
	fmt.Println("【SubDomain Burst】 泛解析测试")
	for i := 0; i <= 10; i++ {
		if r := DnsCheckDomain(fmt.Sprintf("%s.%s", RandomString(), domain)); r != nil {
			if ipFan[r.Ip] {
				continue
			}
			ipFan[r.Ip] = true
		} else {
			break
		}
	}
	for i := range ipFan {
		fmt.Println(fmt.Sprintf("【SubDomain Burst】 泛解析IP: %s", i))
	}
	fmt.Println()
	fmt.Println(fmt.Sprintf("【SubDomain Burst】 Count SubDomains Dict %d 个", len(domains)))
	workWg.Add(1)
	go threadSubDomain(domain, &workWg, semaphore, task)
	for _, sub := range domains {
		workWg.Add(1)
		go threadSubDomain(fmt.Sprintf("%s.%s", sub, domain), &workWg, semaphore, task)
	}
	workWg.Wait()
	close(task)
	fmt.Println(fmt.Sprintf("【SubDomain Burst】 %s 去除泛解析", domain))
	var tempRes = make(map[string][]DnsCheckDomainData)
	for r := range task {
		if r == nil || (r.Ip == "" || r.Domain == "") {
			continue
		}
		tempRes[r.Ip] = append(tempRes[r.Ip], *r)
	}
	fmt.Println(fmt.Sprintf("【SubDomain Burst】 %s 数据输出", domain))
	for vIp, v := range tempRes {
		if ipFan[vIp] && len(v) > 0 {
			var data DnsCheckDomainData
			if num := FanCheckWWW(v, fmt.Sprintf("www.%s", domain)); num != nil {
				data = *&v[*num]
			} else {
				data = *&v[0]
			}
			result = append(result, data)
			continue
		}
		for _, i := range v {
			result = append(result, i)
		}
	}
	return result
}

// threadSubDomain 多并发子域名检测
func threadSubDomain(domain string, wg *sync.WaitGroup, semaphore chan struct{}, result chan *DnsCheckDomainData) {
	semaphore <- struct{}{}
	defer func() { <-semaphore }()
	defer wg.Done()
	result <- DnsCheckDomain(domain)
}

// DnsCheckDomain DNS尝试域名解析
func DnsCheckDomain(domain string) *DnsCheckDomainData {
	fmt.Print(fmt.Sprintf("\x1b[K\r[#] "), domain)
	dnsServers := []string{"114.114.114.114", "8.8.8.8", "8.8.4.4", "114.114.115.115", "114.114.114.119", "114.114.115.119", "114.114.114.110", "114.114.115.110", "223.5.5.5", "223.6.6.6", "180.76.76.76", "119.29.29.29", "119.28.28.28", "182.254.116.116", "182.254.118.118", "1.2.4.8", "52.80.66.66", "117.50.10.10", "52.80.52.52", "117.50.60.30", "52.80.60.30", "210.2.4.8", "112.124.47.27", "114.215.126.16", "123.125.81.6", "117.50.11.11", "101.226.4.6", "218.30.118.6"}
	m := new(dns.Msg)
	m.SetQuestion(domain+".", dns.TypeA)
	c := new(dns.Client)
	server := dnsServers[rand.Intn(len(dnsServers))]
	if r, _, err := c.Exchange(m, net.JoinHostPort(server, "53")); err == nil {
		for _, ans := range r.Answer {
			if a, ok := ans.(*dns.A); ok {
				res := DnsCheckDomainData{Domain: domain, Ip: a.A.String(), DnsServer: server, Timer: Timer()}
				return &res
			}
		}
	}
	return nil
}

func main() {
	var d = flag.String("d", "", "输入单个子域名用于解析")
	var s = flag.String("s", "", "输入根域名，进行域名爆破")
	flag.Parse()
	if *d != "" {
		if res := DnsCheckDomain(*d); res != nil {
			fmt.Println("Domain:", res.Domain)
			fmt.Println("IPv4:", res.Ip)
			fmt.Println("Dns:", res.DnsServer)
			fmt.Println("Time:", res.Timer)
		}
	} else if *s != "" {
		res := CheckSubDomain(*s)
		for _, r := range res {
			fmt.Println("[+]", r.Timer, r.Domain, r.Ip, r.DnsServer)
		}
	}
}
