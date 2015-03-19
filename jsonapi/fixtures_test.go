package jsonapi

import (
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"gopkg.in/guregu/null.v2/zero"
)

type Magic struct {
	ID MagicID
}

func (m Magic) GetID() string {
	return m.ID.String()
}

type MagicID string

func (m MagicID) String() string {
	return "This should be visible"
}

type Comment struct {
	ID   int `json:"-"`
	Text string
}

func (c Comment) GetID() string {
	return fmt.Sprintf("%d", c.ID)
}

func (c *Comment) SetID(stringID string) error {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}

	c.ID = id

	return nil
}

type User struct {
	ID       int
	Name     string
	Password string `json:"-"`
}

func (u User) GetID() string {
	return fmt.Sprintf("%d", u.ID)
}

func (u *User) SetID(stringID string) error {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}

	u.ID = id

	return nil
}

type SimplePost struct {
	ID, Title, Text string
	Created         time.Time
}

func (s SimplePost) GetID() string {
	return s.ID
}

type Post struct {
	ID          int
	Title       string
	Comments    []Comment     `json:"-"`
	CommentsIDs []int         `json:"-"`
	Author      *User         `json:"-"`
	AuthorID    sql.NullInt64 `json:"-"`
}

func (c Post) GetID() string {
	return fmt.Sprintf("%d", c.ID)
}

func (c *Post) SetID(stringID string) error {
	id, err := strconv.Atoi(stringID)
	if err != nil {
		return err
	}

	c.ID = id

	return nil
}

func (c Post) GetReferences() []Reference {
	return []Reference{
		{
			Type: "comments",
			Name: "comments",
		},
		{
			Type: "users",
			Name: "author",
		},
	}
}

func (c *Post) SetReferencedIDs(ids []ReferenceID) error {
	return nil
}

func (c Post) GetReferencedIDs() []ReferenceID {
	result := []ReferenceID{}

	if c.Author != nil {
		authorID := ReferenceID{Type: "users", Name: "author", ID: c.Author.GetID()}
		result = append(result, authorID)
	} else if c.AuthorID.Valid {
		authorID := ReferenceID{Type: "users", Name: "author", ID: fmt.Sprintf("%d", c.AuthorID.Int64)}
		result = append(result, authorID)
	}

	if len(c.Comments) > 0 {
		for _, comment := range c.Comments {
			result = append(result, ReferenceID{Type: "comments", Name: "comments", ID: comment.GetID()})
		}
	} else if len(c.CommentsIDs) > 0 {
		for _, commentID := range c.CommentsIDs {
			result = append(result, ReferenceID{Type: "comments", Name: "comments", ID: fmt.Sprintf("%d", commentID)})
		}
	}

	return result
}

func (c Post) GetReferencedStructs() []MarshalIdentifier {
	result := []MarshalIdentifier{}
	if c.Author != nil {
		result = append(result, c.Author)
	}
	for key := range c.Comments {
		result = append(result, c.Comments[key])
	}
	return result
}

func (c *Post) SetReferencedStructs(references []UnmarshalIdentifier) error {
	return nil
}

type AnotherPost struct {
	ID       int
	AuthorID int   `json:"-"`
	Author   *User `json:"-"`
}

func (p AnotherPost) GetID() string {
	return fmt.Sprintf("%d", p.ID)
}

func (p AnotherPost) GetReferences() []Reference {
	return []Reference{
		{
			Type: "users",
			Name: "author",
		},
	}
}

func (p AnotherPost) GetReferencedIDs() []ReferenceID {
	result := []ReferenceID{}

	if p.AuthorID != 0 {
		result = append(result, ReferenceID{ID: fmt.Sprintf("%d", p.AuthorID), Name: "author", Type: "users"})
	}

	return result
}

type ZeroPost struct {
	ID    string
	Title string
	Value zero.Float
}

func (z ZeroPost) GetID() string {
	return z.ID
}

type ZeroPostPointer struct {
	ID    string
	Title string
	Value *zero.Float
}

func (z ZeroPostPointer) GetID() string {
	return z.ID
}

type Question struct {
	ID                  string
	Text                string
	InspiringQuestionID sql.NullString `json:"-"`
	InspiringQuestion   *Question      `json:"-"`
}

func (q Question) GetID() string {
	return q.ID
}

func (q Question) GetReferences() []Reference {
	return []Reference{
		{
			Type: "questions",
			Name: "inspiringQuestion",
		},
	}
}

func (q Question) GetReferencedIDs() []ReferenceID {
	result := []ReferenceID{}
	if q.InspiringQuestionID.Valid {
		result = append(result, ReferenceID{ID: q.InspiringQuestionID.String, Name: "inspiringQuestion", Type: "questions"})
	}

	return result
}

func (q Question) GetReferencedStructs() []MarshalIdentifier {
	result := []MarshalIdentifier{}
	if q.InspiringQuestion != nil {
		result = append(result, *q.InspiringQuestion)
	}

	return result
}

type Identity struct {
	ID     int64    `json:"user_id"`
	Scopes []string `json:"scopes"`
}

func (i Identity) GetID() string {
	return fmt.Sprintf("%d", i.ID)
}

type Unicorn struct {
	UnicornID int64    `json:"unicorn_id"` //Annotations are ignored
	Scopes    []string `json:"scopes"`
}

func (u Unicorn) GetID() string {
	return "magicalUnicorn"
}

type CompleteServerInformation struct{}

func (i CompleteServerInformation) GetBaseURL() string {
	return "http://my.domain"
}

func (i CompleteServerInformation) GetPrefix() string {
	return "v1"
}
