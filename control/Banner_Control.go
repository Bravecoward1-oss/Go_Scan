package control

func bannerControl() {
	//Create icon
	banner = "▄████  ▒█████    ██████  ▄████▄   ▄▄▄       ███▄    █ \n " +
		"██▒ ▀█▒▒██▒  ██▒▒██    ▒ ▒██▀ ▀█  ▒████▄     ██ ▀█   █ \n" +
		"▒██░▄▄▄░▒██░  ██▒░ ▓██▄   ▒▓█    ▄ ▒██  ▀█▄  ▓██  ▀█ ██▒\n" +
		"░▓█  ██▓▒██   ██░  ▒   ██▒▒▓▓▄ ▄██▒░██▄▄▄▄██ ▓██▒  ▐▌██▒\n" +
		"░▒▓███▀▒░ ████▓▒░▒██████▒▒▒ ▓███▀ ░ ▓█   ▓██▒▒██░   ▓██░\n" +
		" ░▒   ▒ ░ ▒░▒░▒░ ▒ ▒▓▒ ▒ ░░ ░▒ ▒  ░ ▒▒   ▓▒█░░ ▒░   ▒ ▒ \n" +
		"  ░   ░   ░ ▒ ▒░ ░ ░▒  ░ ░  ░  ▒     ▒   ▒▒ ░░ ░░   ░ ▒░\n" +
		"░ ░   ░ ░ ░ ░ ▒  ░  ░  ░  ░          ░   ▒      ░   ░ ░ \n" +
		"      ░     ░ ░        ░  ░ ░            ░  ░         ░ \n" +
		"                          ░                             \n"
	simpleBanner = banner +
		"-h	help		Output to help\n" +
		"-v	version		Output version\n" +
		"-u	url		Url to scan\n" +
		"-d	dictionary	Choose the dictionary\n"

	helpBanner = banner +
		"-h	help		Output to help\n" +
		"-v	version		Output version\n" +
		"-u	url		Url to scan\n" +
		"-d	dictionary	-d {dictionary name}\n" +
		"			Choose the dictionary\n" +
		"			This program provides you with dictionaries\n" +
		"			{asp,aspx,backup,Comprehensive_risk,dir,fck,fingerprint,jsp,mdb,php,shell,top,type}\n" +
		"			Format such as -d *.txt" +
		"\n" +
		"	If you want to use your dictionary,\n" +
		"	put the dictionary in the dictionary folder\n" +
		"	Add the name and suffix of your dictionary after it"
}
