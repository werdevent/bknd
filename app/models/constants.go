package models

// DATABASE NAMES
const (
	// USERS is a constant that will define the name of the users database
	USERS_DB = "user"
	// SUPERADMINS is a constant that will define the name of the superadmins database
	SUPERADMINS_DB = "superAdmins"
)

// ROLES
const (
	// SUPER is a exclusive role that will allow the user to have full access to the app and code, it will have all the permission needed to performe any type of operation
	SUPER = 10
	// STAFF is a type of role that will allow user do certain decision on the code or to check some aspects of the code, but it mostly have access to not vulnerable parts of the app
	STAFF = 8
	// CREATOR is a type of role that will be use by the users of the app, this role will have permission to create events, issue tickets and accept payments. It also will have access to buy and browse other creators. It will have access to read and write certain parts if the database but the operations will be limited
	CREATOR = 5
	// SUB_ADMIN is a special type of role that will be able to only read and write certains parts of information in the database, this role is exclusive to just serve as a behalf of the main admin account
	SUB_ADMIN = 4
	// CONSUMER is a type of role that only have access to read and write certain and limited parts of the database, nothing outside the own account.
	CONSUMER = 1
)

//  MAIL CONSTANTS
const (
	// SENDER_MAILER is a constant for the mailer that will have the business email account
	SENDER_MAILER = "PROFESSIONAL"
	// RECEIVER_MAILER is a constant for the mailer that will no be the business email and its proupuse is to receive email
	RECEIVER_MAILER = "NON-PROFESSIONAL"
	// MAIL_SENDER_HOST is a constant that represent the host of the professional email
	MAIL_SENDER_HOST = "MAIL_SENDER_HOST"
	// MAIL_SENDER_PORT is a constant that represent the port of the professional email
	MAIL_SENDER_PORT = "MAIL_SENDER_PORT"
	// MAIL_SENDER_USER is a constant that represent the user of the professional email
	MAIL_SENDER_USER = "MAIL_SENDER_USER"
	// MAIL_SENDER_PASSWORD is a constant that represent the password of the professional email
	MAIL_SENDER_PASSWORD = "MAIL_SENDER_PASSWORD"
	// MAIL_RECEIVER_HOST is a constant that represent the host of the non-professional email
	MAIL_RECEIVER_HOST = "MAIL_RECEIVER_HOST"
	// MAIL_RECEIVER_PORT is a constant that represent the port of the non-professional email
	MAIL_RECEIVER_PORT = "MAIL_RECEIVER_PORT"
	// MAIL_RECEIVER_USER is a constant that represent the user of the non-professional email
	MAIL_RECEIVER_USER = "MAIL_RECEIVER_USER"
	// MAIL_RECEIVER_PASSWORD is a constant that represent the password of the non-professional email
	MAIL_RECEIVER_PASSWORD = "MAIL_RECEIVER_PASSWORD"
)
