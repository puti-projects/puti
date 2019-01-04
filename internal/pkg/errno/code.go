package errno

var (
	// common errors
	OK                  = &Errno{Code: 0, Message: "OK"}
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	ErrBind             = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	ErrDatabase   = &Errno{Code: 20002, Message: "Database error."}
	ErrToken      = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}

	// user errors
	ErrEncrypt           = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	ErrUserNotFound      = &Errno{Code: 20102, Message: "The user was not found."}
	ErrTokenInvalid      = &Errno{Code: 20103, Message: "The token was invalid."}
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	ErrSaveAvatar        = &Errno{Code: 20105, Message: "Save file failed"}

	// media errors
	ErrUploadFile    = &Errno{Code: 20201, Message: "Upload file failed"}
	ErrMediaNotFound = &Errno{Code: 20202, Message: "The media was not found."}
	ErrTitleEmpty    = &Errno{Code: 20203, Message: "Title can not be empty."}

	// taxonomy errors
	ErrTypeEmpty                = &Errno{Code: 20301, Message: "Taxonomy type can not be empty."}
	ErrTermNotFount             = &Errno{Code: 20302, Message: "The term taxonomy was not found."}
	ErrTaxonomyNameExist        = &Errno{Code: 20303, Message: "The term taxonomy name was already exist."}
	ErrTaxonomyParentId         = &Errno{Code: 20304, Message: "Error parent id."}
	ErrTaxonomyParentCanNotSelf = &Errno{Code: 20305, Message: "Parent can not be itself."}

	// Article errors
	ErrArticleNotFount     = &Errno{Code: 20401, Message: "The article was not found."}
	ErrArticleCreateFailed = &Errno{Code: 20402, Message: "Create article failed."}

	// Page errors
	ErrPageNotFount     = &Errno{Code: 20501, Message: "The page was not found."}
	ErrPageCreateFailed = &Errno{Code: 20502, Message: "Create page failed."}
	ErrSlugExist        = &Errno{Code: 20503, Message: "The slug is already exist."}

	// Option errors
	ErrSettingType = &Errno{Code: 20601, Message: "Error option setting type."}
)
