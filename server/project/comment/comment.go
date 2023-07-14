package comment

import (
	"time"

	person "github.com/TLacault/TaskShield/server/project/person"
)

type Comment struct {
	content string
	date    string // dd/mm/yyyy-hh:mm:ss
	author  person.Person
}

/* ********************************************************************************************* */

func New(content string, author person.Person) Comment {
	comment := Comment{
		content: content,
		date:    time.Now().Format("02/01/2006-15:04:05"),
		author:  author,
	}
	return comment
}

/* ********************************************************************************************* */

func (c *Comment) GetContent() string       { return c.content }
func (c *Comment) GetDate() string          { return c.date }
func (c *Comment) GetAuthor() person.Person { return c.author }

/* ********************************************************************************************* */

func (c *Comment) ToString() string {
	return c.author.ToString() + ";" + c.date + ";" + c.content
}

func (c *Comment) ToStringFormat() string {
	return "[" + c.author.ToString() + "](" + c.date + "): " + c.content
}
