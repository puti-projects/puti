package errno

// common errors
var (
	// OK no error
	OK = &Errno{Code: 0, Message: "OK"}
	// InternalServerError internal server error
	InternalServerError = &Errno{Code: 10001, Message: "Internal server error"}
	// ErrBind binding request struct error
	ErrBind = &Errno{Code: 10002, Message: "Error occurred while binding the request body to the struct"}

	// ErrValidation Validation failed error
	ErrValidation = &Errno{Code: 20001, Message: "Validation failed."}
	// ErrDatabase Database error
	ErrDatabase = &Errno{Code: 20002, Message: "Database error."}
	// ErrToken JSON view token error
	ErrToken = &Errno{Code: 20003, Message: "Error occurred while signing the JSON view token."}
)

// user auth errors
var (
	// ErrEncrypt encrypt error
	ErrEncrypt = &Errno{Code: 20101, Message: "Error occurred while encrypting the user password."}
	// ErrUserNotFound user not found error
	ErrUserNotFound = &Errno{Code: 20102, Message: "The user was not found."}
	// ErrTokenInvalid token invalid error
	ErrTokenInvalid = &Errno{Code: 20103, Message: "The token was invalid."}
	// ErrPasswordIncorrect password incorrect error
	ErrPasswordIncorrect = &Errno{Code: 20104, Message: "The password was incorrect."}
	// ErrSaveAvatar Save avatar failed error
	ErrSaveAvatar = &Errno{Code: 20105, Message: "Save file failed"}
)

// media errors
var (
	// ErrUploadFile upload media failed error
	ErrUploadFile = &Errno{Code: 20201, Message: "Upload file failed"}
	// ErrMediaNotFound media not found error
	ErrMediaNotFound = &Errno{Code: 20202, Message: "The media was not found."}
	// ErrTitleEmpty Title is empty error
	ErrTitleEmpty = &Errno{Code: 20203, Message: "Title can not be empty."}
)

// taxonomy errors
var (
	// ErrTypeEmpty taxonomy type empty error
	ErrTypeEmpty = &Errno{Code: 20301, Message: "Taxonomy type can not be empty."}
	// ErrTermNotFount term taxonomy not found error
	ErrTermNotFount = &Errno{Code: 20302, Message: "The term taxonomy was not found."}
	// ErrTaxonomyNameExist  term taxonomy name already exist
	ErrTaxonomyNameExist = &Errno{Code: 20303, Message: "The term taxonomy name was already exist."}
	// ErrTaxonomyParentID error parent id
	ErrTaxonomyParentID = &Errno{Code: 20304, Message: "Error parent id."}
	// ErrTaxonomyParentCanNotSelf parent can not be itself
	ErrTaxonomyParentCanNotSelf = &Errno{Code: 20305, Message: "Parent can not be itself."}
)

// Article errors
var (
	// ErrArticleNotFount article not found error
	ErrArticleNotFount = &Errno{Code: 20401, Message: "The article was not found."}
	// ErrArticleCreateFailed create article failed error
	ErrArticleCreateFailed = &Errno{Code: 20402, Message: "Create article failed."}
)

// Page errors
var (
	// ErrPageNotFount page was not found
	ErrPageNotFount = &Errno{Code: 20501, Message: "The page was not found."}
	// ErrPageCreateFailed create page failed
	ErrPageCreateFailed = &Errno{Code: 20502, Message: "Create page failed."}
	// ErrSlugExist slug is already exist
	ErrSlugExist = &Errno{Code: 20503, Message: "The slug is already exist."}
)

// Option errors
var (
	// ErrSettingType error option setting type
	ErrSettingType = &Errno{Code: 20601, Message: "Error option setting type."}
)

// Subject errors
var (
	// ErrSubjectNameExist subject name was already exist
	ErrSubjectNameExist = &Errno{Code: 20701, Message: "The subject name was already exist."}
	// ErrSubjectNotFount subject not found error
	ErrSubjectNotFount = &Errno{Code: 20702, Message: "The subject was not found."}
)
