package uptrends

import (
	"strconv"
)

var uptrendsDashboards = []string{
	"UseAccountSpecifiedDashboard",
	"AccountOverview",
	"AlertLog",
	"AvailabilitySummary",
	"ErrorsOverview",
	"ProbeLog",
	"ProbeStatus",
	"Operations",
	"PerformanceDetails",
	"PerformanceSummary",
	"TotalTimeTable",
	"UptimePerProbe",
	"UptimeTable",
}

var browserToUserAgent = map[string]string{
	"chrome83":          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
	"chrome83_android":  "Mozilla/5.0 (Linux; Android 10) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.101 Mobile Safari/537.36",
	"chrome67":          "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/67.0.3396.79 Safari/537.36",
	"chrome49":          "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/49.0.2623.87 Safari/537.36",
	"firefox77":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:77.0) Gecko/20100101 Firefox/77.0",
	"firefox68_android": "Mozilla/5.0 (Android 10; Mobile; rv:68.0) Gecko/68.0 Firefox/68.0",
	"firefox60":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:60.0) Gecko/20100101 Firefox/60.0",
	"firefox55":         "Mozilla/5.0 (Windows NT 10.0; WOW64; rv:55.0) Gecko/20100101 Firefox/55.0",
	"ms_ie11":           "Mozilla/5.0 (Windows NT 10.0; Trident/7.0; rv:11.0) like Gecko",
	"ms_ie10":           "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.1; Trident/6.0)",
	"ms_edge83":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 Edg/83.0.478.45",
	"ms_edge18":         "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.102 Safari/537.36 Edge/18.18363",
	"opera68":           "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36 OPR/68.0.3618.125",
	"opera12":           "Opera/9.80 (Windows NT 6.1; U; en-US) Presto/2.9.181 Version/12.00",
	"safari13_macos":    "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.97 Safari/537.36",
	"safari13_iphone":   "Mozilla/5.0 (iPhone; CPU iPhone OS 13_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4103.88 Mobile/15E148 Safari/604.1",
	"safari13_ipad":     "Mozilla/5.0 (iPad; CPU OS 13_5 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/83.0.4103.88 Mobile/15E148 Safari/604.1",
}

// UserAgent accepts browser name or user agent and lookup user agent in map
func UserAgent(input string) string {
	if id, found := browserToUserAgent[input]; found {
		return id
	}
	return input
}

var validHttpCodes = []int{
	100, // Continue
	101, // Switching Protocol
	200, // OK
	201, // Created
	202, // Accepted
	203, // Non-Authoritative Information
	204, // No Content
	205, // Reset Content
	206, // Partial Content
	300, // Multiple Choice
	301, // Moved Permanently
	302, // Found
	303, // See Other
	304, // Not Modified
	305, // Use Proxy
	306, // unused
	307, // Temporary Redirect
	308, // Permanent Redirect
	400, // Bad Request
	401, // Unauthorized
	402, // Payment Required
	403, // Forbidden
	404, // Not Found
	405, // Method Not Allowed
	406, // Not Acceptable
	407, // Proxy Authentication Required
	408, // Request Timeout
	409, // Conflict
	410, // Gone
	411, // Length Required
	412, // Precondition Failed
	413, // Payload Too Large
	414, // URI Too Long
	415, // Unsupported Media Type
	416, // Range Not Satisfiable
	417, // Expectation Failed
	426, // Upgrade Required
	500, // Internal Server Error
	501, // Not Implemented
	502, // Bad Gateway
	503, // Service Unavailable
	504, // Gateway Timeout
	505, // HTTP Version Not Supported
}

var RegionIDToRegionName = map[string]string{
	"13":  "Austria",
	"20":  "Belgium",
	"53":  "Czech Republic",
	"56":  "Denmark",
	"256": "Bulgaria",
	"263": "Estonia",
	// Continent - Begin
	"1001": "Africa",
	"1002": "Asia",
	"1003": "Australia",
	"1004": "Europe",
	"1005": "North America",
	"1006": "South America",
	"1007": "Middle East",
	// Continent - End
	"1013": "Cyprus",
	"1014": "Bosnia and Herzegovina",
}

var RegionNameToRegionID = reverseMap(RegionIDToRegionName)

var checkpointIDToCheckpointName = map[string]string{
	"0":   "Amsterdam",
	"1":   "London",
	"2":   "San Antonio",
	"3":   "Sydney",
	"4":   "Tampa",
	"5":   "Nottingham",
	"6":   "Hannover",
	"7":   "Haarlem",
	"8":   "New York",
	"9":   "Las Vegas",
	"10":  "Los Angeles",
	"11":  "San Francisco",
	"12":  "Chicago",
	"13":  "Berlin",
	"14":  "Vancouver",
	"15":  "Toronto",
	"16":  "Bournemouth",
	"17":  "Gloucester",
	"18":  "Brussels",
	"19":  "Singapore",
	"20":  "Shanghai",
	"21":  "Stockholm",
	"22":  "Frankfurt",
	"23":  "Recife",
	"24":  "Rotterdam",
	"25":  "Madrid",
	"26":  "Denver",
	"27":  "Tel Aviv",
	"28":  "Portland",
	"29":  "Seattle",
	"30":  "Lausanne",
	"31":  "Dallas",
	"32":  "Lille",
	"33":  "Istanbul",
	"34":  "Rome",
	"35":  "Boston",
	"36":  "Mumbai",
	"37":  "Hong Kong",
	"38":  "Clifton",
	"39":  "Johannesburg",
	"40":  "Dublin",
	"41":  "Guadalajara",
	"42":  "Antwerp",
	"43":  "Düsseldorf",
	"44":  "Göteborg",
	"45":  "Orlando",
	"46":  "Melbourne",
	"47":  "Malmö",
	"48":  "Houston",
	"49":  "Tokyo",
	"50":  "Copenhagen",
	"51":  "Munich",
	"52":  "Hyderabad",
	"53":  "Nuremberg",
	"54":  "San Diego",
	"55":  "Atlanta",
	"56":  "Oslo",
	"57":  "Grand Rapids",
	"58":  "Manchester",
	"59":  "Budapest",
	"60":  "Kelowna",
	"61":  "São Paulo",
	"62":  "Washington, D.C.",
	"63":  "Warsaw",
	"64":  "Kuala Lumpur",
	"65":  "Strasbourg",
	"66":  "Lisbon",
	"67":  "Karlsruhe",
	"68":  "Kansas City",
	"69":  "Ghent",
	"70":  "Paris",
	"71":  "Leipzig",
	"72":  "Miami",
	"73":  "Philadelphia",
	"74":  "Karlskrona",
	"75":  "Helsinki",
	"76":  "Ashburn",
	"77":  "Beijing",
	"78":  "Manila",
	"79":  "Klagenfurt",
	"80":  "Panama City",
	"81":  "Phoenix",
	"82":  "Auckland",
	"83":  "Perth",
	"84":  "Salt Lake City",
	"85":  "Zurich",
	"86":  "Lelystad",
	"87":  "Prague",
	"88":  "Seoul",
	"89":  "New Delhi",
	"90":  "San Jose",
	"91":  "Buenos Aires",
	"92":  "Cologne",
	"93":  "Montreal",
	"94":  "Moscow",
	"95":  "Athens",
	"96":  "Bangkok",
	"97":  "Cairo",
	"98":  "Arnhem",
	"99":  "St. Louis",
	"100": "Vienna",
	"101": "Minneapolis",
	"102": "Groningen",
	"103": "Karlstad",
	"104": "Milan",
	"105": "Newark",
	"106": "Hamburg",
	"107": "Salzburg",
	"108": "Dubai",
	"109": "Osaka",
	"110": "Rio de Janeiro",
	"111": "Taipei",
	"112": "Bratislava",
	"113": "Marseille",
	"114": "Edinburgh",
	"115": "Palermo",
	"116": "Manassas",
	"117": "Bucharest",
	"118": "Eindhoven",
	"119": "Oulu",
	"120": "Columbus",
	"121": "Fayetteville",
	"122": "Brasília",
	"123": "Azure US West",
	"124": "Utrecht",
	"125": "Saint Petersburg",
	"126": "Luxembourg",
	"127": "Glasgow",
	"128": "Newcastle",
	"129": "Riyadh",
	"130": "Falkenstein",
	"131": "Brisbane",
	"132": "Bruges",
	"133": "Azure Central US",
	"134": "Valencia",
	"135": "Zhengzhou",
	"136": "Geneva",
	"137": "Amazon USA-West",
	"138": "Lyon",
	"139": "Leicester",
	"140": "Bergen",
	"141": "Adelaide",
	"142": "Vilnius",
	"143": "Canberra",
	"144": "Hangzhou",
	"145": "Jakarta",
	"146": "Cape Town",
	"147": "Sacramento",
	"148": "Santiago",
	"149": "Mexico City",
	"150": "St. Asaph",
	"151": "Medellín",
	"152": "Roubaix",
	"153": "Xaxim",
	"154": "Reims",
	"155": "Fez",
	"156": "Paramaribo",
	"157": "Shenzen",
	"158": "Bielefeld",
	"159": "Bremen",
	"160": "Sapporo",
	"161": "Tianjin",
	"162": "Bend",
	"163": "Redding",
	"164": "Honolulu",
	"165": "Barcelona",
	"166": "Nagano",
	"167": "Rennes",
	"168": "Hanoi",
	"169": "Sofia",
	"170": "Belgrade",
	"171": "Bordeaux",
	"172": "Charlotte",
	"173": "Lahore",
	"174": "Kiev",
	"175": "Lima",
	"176": "Nairobi",
	"177": "Riga",
	"178": "Tallinn",
	"179": "Kolkota",
	"180": "Indore",
	"181": "Bangalore",
	"182": "Buffalo",
	"183": "Arezzo",
	"184": "Albuquerque",
	"185": "Bern",
	"186": "Ottawa",
	"187": "Omaha",
	"188": "Paris Vitry-sur-Seine",
	"189": "Kristiansand",
	"190": "Hamina",
	"191": "Oklahoma City",
	"192": "Qingdao",
	"193": "Chengdu",
	"194": "Moncks Corner",
	"195": "Portsmouth",
	"196": "Graz",
	"197": "Hohhot",
	"198": "Chongqing",
	"199": "Detroit",
	"200": "Islamabad",
	"201": "Novosibirsk",
	"202": "Lupfig",
	"203": "Bursa",
	"204": "Sendai",
	"205": "Córdoba",
	"206": "New Orleans",
	"207": "Malta",
	"208": "Cyprus",
	"209": "Lenoir",
	"210": "Novi Travnik",
	"211": "Regensburg",
	"212": "Rende",
	"213": "Thessaloniki",
	"214": "Bahrain",
	"215": "Logroño",
	"216": "Baden-Baden",
	"217": "Gravelines",
	"218": "Campina Grande",
	"219": "Kassel",
	"220": "Santa Clara",
	"221": "Aalsmeer",
	"222": "Eppes",
	"223": "Turin",
	"224": "York",
	"225": "Malaga",
	"226": "The Hague",
}

var CheckpointNameToCheckpointID = reverseMap(checkpointIDToCheckpointName)

// CheckpointID accepts ID or checkpoint name and lookup ID
func CheckpointID(input string) string {
	if id, found := CheckpointNameToCheckpointID[input]; found {
		return id
	}
	return input
}

func DeduplicateCheckpointIDs(input []interface{}) *[]int32 {
	ids := make([]int32, 0)
	keys := make(map[int]bool)

	for _, v := range input {
		id, _ := strconv.Atoi(CheckpointID(v.(string)))
		if _, value := keys[id]; !value {
			keys[id] = true
			ids = append(ids, int32(id))
		}
	}

	return &ids
}

func RegionID(idOrName string) string {
	if id, found := RegionNameToRegionID[idOrName]; found {
		return id
	}
	return idOrName
}

func DeduplicateRegionIDs(input []interface{}) *[]int32 {
	ids := make([]int32, 0)
	keys := make(map[int]bool)

	for _, v := range input {
		id, _ := strconv.Atoi(RegionID(v.(string)))
		if _, value := keys[id]; !value {
			keys[id] = true
			ids = append(ids, int32(id))
		}
	}

	return &ids
}
