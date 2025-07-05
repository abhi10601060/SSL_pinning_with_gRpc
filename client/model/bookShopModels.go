package model

type Book struct{
	Id int  
	Name string
}

type ListAllBookResponse struct{
	Books []Book
	Count int
}