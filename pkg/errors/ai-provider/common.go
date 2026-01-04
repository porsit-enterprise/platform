package aiprovider

import "errors"

//──────────────────────────────────────────────────────────────────────────────────────────────────

var (
	ErrConnection                      = errors.New("خطا در برقرار ارتباط با سرویس")
	ErrRateLimit                       = errors.New("تعداد درخواست\u200cهای دریافتی زیاد می\u200cباشد، لطفا چند ثانیه دیگر درخواست ارسال نمایید")
	ErrAuthentication                  = errors.New("خطا در احراز هویت")
	ErrTimeout                         = errors.New("خطا در دریافت اطلاعات")
	ErrLoading                         = errors.New("خطا در سرویس هوش مصنوعی")
	ErrSummarizeConversationAPI        = errors.New("خطا در دریافت پاسخ از سرویس هوش مصنوعی، لطفا چند لحظه دیگر دوباره تلاش نمایید")
	ErrSummarizeConversationUnexpected = errors.New("خطا در سرویس هوش مصنوعی")
)
