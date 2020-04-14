package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("Title: \t\t%q\n" +
		"Author: \t\t%q\n" +
		"Published: \t\t%q\n", b.Title, b.Author, b.YearPublished )
}

var books = []Book {
	{
		ID:            1,
		Title:         "HitchHiker1",
		Author:        "Adam1",
		YearPublished: 1979,
	},
	{
		ID:            2,
		Title:         "HitchHiker2",
		Author:        "Adam2",
		YearPublished: 1979,
	},
	{
		ID:            3,
		Title:         "HitchHiker3",
		Author:        "Adam3",
		YearPublished: 1979,
	},
	{
		ID:            4,
		Title:         "HitchHiker4",
		Author:        "Adam4",
		YearPublished: 1979,
	},
	{
		ID:            5,
		Title:         "HitchHiker5",
		Author:        "Adam5",
		YearPublished: 1979,
	},
	{
		ID:            6,
		Title:         "HitchHiker6",
		Author:        "Adam6",
		YearPublished: 1979,
	},
	{
		ID:            7,
		Title:         "HitchHiker7",
		Author:        "Adam7",
		YearPublished: 1979,
	},
	{
		ID:            8,
		Title:         "HitchHiker8",
		Author:        "Adam8",
		YearPublished: 1979,
	},
	{
		ID:            9,
		Title:         "HitchHiker9",
		Author:        "Adam9",
		YearPublished: 1979,
	},
	{
		ID:            10,
		Title:         "HitchHiker10",
		Author:        "Adam10",
		YearPublished: 1979,
	},
}
