package main

type USER_MAC struct{
	LINEID	string
	USER 	[]URL
	CRAWLER []KEYWORD
}

type URL struct{
	NAME 	string
	CONTENT string
	KEYWORD string
	TIME	int64
}

type KEYWORD struct{
	KEYWORD	string
	GOOGLE []CRAWLER_GOOGLE
}

type CRAWLER_GOOGLE struct{
	URL_NAME	string
	CONTENT string
	KEYWORD	string
	TIME	int64
}
