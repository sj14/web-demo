package view

import (
	"html/template"
)

//Compile templates on start
var (
	Templates = template.Must(template.ParseFiles(
		"interfaces/web/view/html/header.html",
		"interfaces/web/view/html/navigation.html",
		"interfaces/web/view/html/footer.html",
		"interfaces/web/view/html/index.html",

		// User
		"interfaces/web/view/html/user/login.html",
		"interfaces/web/view/html/user/register.html",
		"interfaces/web/view/html/user/show.html",
		"interfaces/web/view/html/user/profile.html",
		"interfaces/web/view/html/user/edit.html",

		// HTTP Status
		"interfaces/web/view/html/status/400_BadRequest.html",
		"interfaces/web/view/html/status/401_Unauthorized.html",
		"interfaces/web/view/html/status/403_Forbidden.html",
		"interfaces/web/view/html/status/404_NotFound.html",
		"interfaces/web/view/html/status/500_InternalServerError.html", // Used on panics

		// Various
		"interfaces/web/view/html/various/panic.html",     // Used to test panics
		"interfaces/web/view/html/various/csrf_test.html", // Used to test CSRF
		"interfaces/web/view/html/various/privacy_statement.html",
		"interfaces/web/view/html/various/contact.html",
		"interfaces/web/view/html/various/imprint.html",
	))
)
