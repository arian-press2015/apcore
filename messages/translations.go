package messages

var EnglishMessages = map[string]string{
	// error
	"NotFound":               "resource not found",
	"Unauthorized":           "unauthorized access",
	"InternalServerError":    "internal server error",
	"InvalidEmailOrPassword": "invalid email or password",
	"NoAuthHeader":           "authorization token not provided",
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
	// success
	"Successful": "عملیات با موفقیت انجام شد",
}
