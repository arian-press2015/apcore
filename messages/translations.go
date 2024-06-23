package messages

var EnglishMessages = map[string]string{
	// error
	"NotFound":               "resource not found",
	"Unauthorized":           "unauthorized access",
	"InternalServerError":    "internal server error",
	"InvalidEmailOrPassword": "invalid email or password",
	"NoAuthHeader":           "authorization token not provided",
	"Feature not found":      "feature not found",
	"Feature is disabled":    "requested feature is disabled, please contact support",
	"Method is not allowed":  "requested method is not allowed",
	// success
	"Successful": "operation finished successfully",
}

var FarsiMessages = map[string]string{
	// error
	"NotFound":               "اطلاعات مورد نظر یافت نشد",
	"Unauthorized":           "دسترسی غیر مجاز",
	"InternalServerError":    "خطای داخلی سرور",
	"InvalidEmailOrPassword": "ایمیل یا گذرواژه نامعتبر است",
	"NoAuthHeader":           "هدر امنیتی ارائه نشده است",
	"Feature not found":      "قابلیت مورد نظر یافت نشد",
	"Feature is disabled":    "قابلیت مورد نظر غیرفعال است،لطفا با پشتیبانی تماس بگیرید",
	// success
	"Successful": "عملیات با موفقیت انجام شد",
}
