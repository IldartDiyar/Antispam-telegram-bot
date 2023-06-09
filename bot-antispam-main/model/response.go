package model

type AutoGenerated struct {
	Data struct {
		Attributes struct {
			LastHTTPResponseContentSha256 string `json:"last_http_response_content_sha256"`
			HTMLMeta                      struct {
				Viewport []string `json:"viewport"`
			} `json:"html_meta"`
			LastHTTPResponseCode          int    `json:"last_http_response_code"`
			LastFinalURL                  string `json:"last_final_url"`
			LastHTTPResponseContentLength int    `json:"last_http_response_content_length"`
			URL                           string `json:"url"`
			LastAnalysisDate              int    `json:"last_analysis_date"`
			Tags                          []any  `json:"tags"`
			LastAnalysisResults           struct {
				Bkav struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Bkav"`
				CMCThreatIntelligence struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"CMC Threat Intelligence"`
				SnortIPSampleList struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Snort IP sample list"`
				VXVault struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"VX Vault"`
				ViriBack struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ViriBack"`
				PhishLabs struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"PhishLabs"`
				K7AntiVirus struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"K7AntiVirus"`
				CINSArmy struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"CINS Army"`
				Quttera struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Quttera"`
				BlockList struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"BlockList"`
				PrecisionSec struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"PrecisionSec"`
				OpenPhish struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"OpenPhish"`
				ZeroXSIF33D struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"0xSI_f33d"`
				FeodoTracker struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Feodo Tracker"`
				ArcSightThreatIntelligence struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ArcSight Threat Intelligence"`
				Scantitan struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Scantitan"`
				AlienVault struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"AlienVault"`
				Sophos struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Sophos"`
				Phishtank struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Phishtank"`
				Cyan struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Cyan"`
				Spam404 struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Spam404"`
				SecureBrain struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"SecureBrain"`
				Crdf struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"CRDF"`
				Rising struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Rising"`
				Fortinet struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Fortinet"`
				AlphaMountainAi struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"alphaMountain.ai"`
				Lionic struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Lionic"`
				Cyble struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Cyble"`
				Seclookup struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Seclookup"`
				XcitiumVerdictCloud struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Xcitium Verdict Cloud"`
				ArtistsAgainst419 struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Artists Against 419"`
				GoogleSafebrowsing struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Google Safebrowsing"`
				SafeToOpen struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"SafeToOpen"`
				ADMINUSLabs struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ADMINUSLabs"`
				ESTsecurity struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ESTsecurity"`
				JuniperNetworks struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Juniper Networks"`
				HeimdalSecurity struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Heimdal Security"`
				AutoShun struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"AutoShun"`
				Trustwave struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Trustwave"`
				AICCMONITORAPP struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"AICC (MONITORAPP)"`
				CyRadar struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"CyRadar"`
				DrWeb struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Dr.Web"`
				Emsisoft struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Emsisoft"`
				Abusix struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Abusix"`
				Webroot struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Webroot"`
				Avira struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Avira"`
				Securolytics struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"securolytics"`
				AntiyAVL struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Antiy-AVL"`
				AlphaSOC struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"AlphaSOC"`
				Acronis struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Acronis"`
				QuickHeal struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Quick Heal"`
				URLQuery struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"URLQuery"`
				ViettelThreatIntelligence struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Viettel Threat Intelligence"`
				DNS8 struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"DNS8"`
				BenkowCc struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"benkow.cc"`
				EmergingThreats struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"EmergingThreats"`
				ChongLuaDao struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Chong Lua Dao"`
				YandexSafebrowsing struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Yandex Safebrowsing"`
				Lumu struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Lumu"`
				Kaspersky struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Kaspersky"`
				SucuriSiteCheck struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Sucuri SiteCheck"`
				DesenmascaraMe struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"desenmascara.me"`
				CrowdSec struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"CrowdSec"`
				Cluster25 struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Cluster25"`
				URLhaus struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"URLhaus"`
				Prebytes struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"PREBYTES"`
				StopForumSpam struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"StopForumSpam"`
				Blueliv struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Blueliv"`
				Netcraft struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Netcraft"`
				ZeroCERT struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ZeroCERT"`
				PhishingDatabase struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Phishing Database"`
				MalwarePatrol struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"MalwarePatrol"`
				Sangfor struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Sangfor"`
				IPsum struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"IPsum"`
				Malwared struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Malwared"`
				BitDefender struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"BitDefender"`
				GreenSnow struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"GreenSnow"`
				GData struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"G-Data"`
				Vipre struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"VIPRE"`
				SCUMWAREOrg struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"SCUMWARE.org"`
				PhishFort struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"PhishFort"`
				MalwaresComURLChecker struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"malwares.com URL checker"`
				ForcepointThreatSeeker struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Forcepoint ThreatSeeker"`
				CriminalIP struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Criminal IP"`
				Certego struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Certego"`
				Eset struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ESET"`
				Threatsourcing struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Threatsourcing"`
				ThreatHive struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"ThreatHive"`
				BforeAiPreCrime struct {
					Category   string `json:"category"`
					Result     string `json:"result"`
					Method     string `json:"method"`
					EngineName string `json:"engine_name"`
				} `json:"Bfore.Ai PreCrime"`
			} `json:"last_analysis_results"`
			LastSubmissionDate      int      `json:"last_submission_date"`
			ThreatNames             []string `json:"threat_names"`
			LastHTTPResponseHeaders struct {
				CriticalCH       string `json:"Critical-CH"`
				ContentEncoding  string `json:"Content-Encoding"`
				TransferEncoding string `json:"Transfer-Encoding"`
				SetCookie        string `json:"Set-Cookie"`
				Expires          string `json:"Expires"`
				Vary             string `json:"Vary"`
				Connection       string `json:"Connection"`
				Server           string `json:"Server"`
				XAdblockKey      string `json:"X-Adblock-Key"`
				AcceptCH         string `json:"Accept-CH"`
				Pragma           string `json:"Pragma"`
				CacheControl     string `json:"Cache-Control"`
				Date             string `json:"Date"`
				ContentType      string `json:"Content-Type"`
			} `json:"last_http_response_headers"`
			Reputation int `json:"reputation"`
			Categories struct {
				DrWeb                  string `json:"Dr.Web"`
				Sophos                 string `json:"Sophos"`
				Webroot                string `json:"Webroot"`
				XcitiumVerdictCloud    string `json:"Xcitium Verdict Cloud"`
				ForcepointThreatSeeker string `json:"Forcepoint ThreatSeeker"`
				AlphaMountainAi        string `json:"alphaMountain.ai"`
			} `json:"categories"`
			Tld                  string `json:"tld"`
			LastModificationDate int    `json:"last_modification_date"`
			LastAnalysisStats    struct {
				Harmless   int `json:"harmless"`
				Malicious  int `json:"malicious"`
				Suspicious int `json:"suspicious"`
				Undetected int `json:"undetected"`
				Timeout    int `json:"timeout"`
			} `json:"last_analysis_stats"`
			TimesSubmitted      int `json:"times_submitted"`
			FirstSubmissionDate int `json:"first_submission_date"`
			TotalVotes          struct {
				Harmless  int `json:"harmless"`
				Malicious int `json:"malicious"`
			} `json:"total_votes"`
		} `json:"attributes"`
		Type  string `json:"type"`
		ID    string `json:"id"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
}
