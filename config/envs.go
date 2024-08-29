package config

const (
	KEY_HMAC_SHA256       = `KEY_HMAC_SHA256`       // สำหรับเข้ารหัส HMAC_DOHOME_SECRET
	KEY_HMAC_QUICK_SHOP   = `KEY_HMAC_QUICK_SHOP`   // secret key for QuickShop , must be 256 bits, get from config, https://generate-secret.vercel.app/256
	KEY_HMAC_B2C          = `KEY_HMAC_B2C`          // secret key for B2C , must be 256 bits, get from config, https://generate-secret.vercel.app/256
	KEY_HMAC_B2B          = `KEY_HMAC_B2B`          // secret key for B2B , must be 256 bits, get from config, https://generate-secret.vercel.app/256
	KEY_HMAC_HOME_SERVICE = `KEY_HMAC_HOME_SERVICE` // secret key for B2B , must be 256 bits, get from config, https://generate-secret.vercel.app/256
)
