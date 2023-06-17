package consts

const BasePathOwlGram = "/var/www/owlgram/"
const BasePathOwlGramDCStatus = "/var/www/dc_status/"
const NotFoundFile = BasePathOwlGram + "html/404.html"
const ForbiddenFile = BasePathOwlGram + "html/403.html"
const ApiFile = BasePathOwlGramDCStatus + "cache/apiInfo.json"
const TrustedASNFile = BasePathOwlGramDCStatus + "config/trustedASN.json"
const TrustedIPFile = BasePathOwlGramDCStatus + "config/trustedIP.json"
const CacheDCFile = BasePathOwlGramDCStatus + "cache/tmpDcStatus.json"
const EnvFilePath = BasePathOwlGram + "server_api/.env"

const MaxRequestsPerIP int64 = 10
const MaxTime int64 = 30
const BanTime int64 = 60 * 30
const OwlGramApiEndPrivate = "https://app.owlgram.org/saveDcData"
