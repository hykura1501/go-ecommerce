package pkg

const (
	// Common
	ErrorGetData               = "error get data"
	ErrorCreateData            = "error create data"
	ErrorBindingData           = "error binding data"
	ErrorRecordNotFound        = "record not found"
	ErrorUploadImage           = "error upload image"
	ErrorUpdateData            = "error update data"
	ErrorDeleteData            = "error delete data"
	ErrorLoginProvider         = "error login provider"
	ErrorResetPasswordNotLocal = "reset password not local"

	// Auth
	ErrorUserNotExisted    = "user is not existed"
	ErrorPasswordIncorrect = "password is incorrect"
	ErrorGenerateToken     = "error generate token"
	ErrorUserExisted       = "user is existed"
	ErrorForbidden         = "forbidden"
	ErrorInvalidIdToken    = "invalid id token"
	ErrorCreateUser        = "error create user"
	ErrorExchangeToken     = "error exchange token"
	ErrorGetUserInfo       = "error get user info"

	// User
	ErrorUserNotFound = "user not found"

	//Cart
	ErrorAddToCart      = "error add to cart"
	ErrorUpdateCart     = "error update cart"
	ErrorDeleteCartItem = "error delete cart item"

	//Order
	ErrorCreateOrder               = "error create order"
	ErrorInsufficientStock         = "insufficient stock"
	ErrorInsertOrderDetail         = "error insert order detail"
	ErrorUpdateProductStock        = "error update product stock"
	ErrorGetOrderHistory           = "error get order history"
	ErrorGetOrders                 = "error get orders"
	ErrorGetRevenueStatistics      = "error get revenue statistics"
	ErrorGetBestSellersStatistics  = "error get best sellers statistics"
	ErrorGetTopCustomersStatistics = "error get top customers statistics"
	ErrorGetOrderDetail            = "error get order detail"

	// Manufacturer
	ErrorGetManufacturers = "error get manufacturers"
)
