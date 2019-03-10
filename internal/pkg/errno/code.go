package errno

// common errors
var (
	// OK
	OK = &Errno{Code: 0, Message: "OK"}
	// InternalServerError
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	// ErrBind
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	// ErrValidation
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	// ErrDatabase
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}
	// ErrToken
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON web token."}
)

// user auth errors
var (
	// ErrEncrypt
	ErrEncrypt = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	// ErrUserNotFound
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
	// ErrTokenInvalid
	ErrTokenInvalid = &Errno{Code: 20103, Message: "The token was invalid."}
	// ErrPasswordIncorrect
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	// ErrSaveAvatar
	ErrSaveAvatar = &Errno{Code: 20105, Message: "Save file failed"}
)

// media errors
var (
	// ErrUploadFile
	ErrUploadFile = &Errno{Code: 20201, Message: "Upload file failed"}
	// ErrMediaNotFound
	ErrMediaNotFound = &Errno{Code: 20202, Message: "The media was not found."}
	// ErrTitleEmpty
	ErrTitleEmpty = &Errno{Code: 20203, Message: "Title can not be empty."}
)

// taxonomy errors
var (
	// ErrTypeEmpty
	ErrTypeEmpty = &Errno{Code: 20301, Message: "Taxonomy type can not be empty."}
	// ErrTermNotFount
	ErrTermNotFount = &Errno{Code: 20302, Message: "The term taxonomy was not found."}
	// ErrTaxonomyNameExist
	ErrTaxonomyNameExist = &Errno{Code: 20303, Message: "The term taxonomy name was already exist."}
	// ErrTaxonomyParentID
	ErrTaxonomyParentID = &Errno{Code: 20304, Message: "Error parent id."}
	// ErrTaxonomyParentCanNotSelf
	ErrTaxonomyParentCanNotSelf = &Errno{Code: 20305, Message: "Parent can not be itself."}
)

// Article errors
var (
	// ErrArticleNotFount
	ErrArticleNotFount = &Errno{Code: 20401, Message: "The article was not found."}
	// ErrArticleCreateFailed
	ErrArticleCreateFailed = &Errno{Code: 20402, Message: "Create article failed."}
)

// Page errors
var (
	// ErrPageNotFount
	ErrPageNotFount = &Errno{Code: 20501, Message: "The page was not found."}
	// ErrPageCreateFailed
	ErrPageCreateFailed = &Errno{Code: 20502, Message: "Create page failed."}
	// ErrSlugExist
	ErrSlugExist = &Errno{Code: 20503, Message: "The slug is already exist."}
)

// Option errors
var (
	// ErrSettingType
	ErrSettingType = &Errno{Code: 20601, Message: "Error option setting type."}
)

// Subject errors
var (
	// ErrSubjectNameExist
	ErrSubjectNameExist = &Errno{Code: 20701, Message: "The subject name was already exist."}
	// ErrSubjectNotFount
	ErrSubjectNotFount = &Errno{Code: 20702, Message: "The subject was not found."}
)
