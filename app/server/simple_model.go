package main

import "cloud.google.com/go/datastore"

type Simple struct {
	Name string
	Age  int
}

type Repo interface {
	SetKind(string)
	Create(interface{}) (*datastore.Key, error)
	Get(*datastore.Key, interface{}) error
	Delete(*datastore.Key) error
	NewQuery() *datastore.Query
	GetAll(*datastore.Query, interface{}) ([]*datastore.Key, error)
	IDKey(int64) *datastore.Key
}

func (s *Simple) SetKind(r Repo) {
	r.SetKind("Simple")

}

func (s *Simple) IDKey(r Repo, id int) *datastore.Key {
	s.SetKind(r)

	return r.IDKey(int64(id))
}

func (s *Simple) Save(r Repo) (int, error) {
	s.SetKind(r)

	key, err := r.Create(s)

	if err != nil {
		return 0, err

	}

	return int(key.ID), nil
}

func (s *Simple) Find(r Repo, id int) error {
	s.SetKind(r)

	k := s.IDKey(r, id)

	return r.Get(k, s)
}

func (s *Simple) Destroy(r Repo, id int) error {
	s.SetKind(r)
	k := s.IDKey(r, id)

	return r.Delete(k)
}

type Simples []Simple

func (ss *Simples) AllWithLimit(r Repo, limit int) error {
	(&Simple{}).SetKind(r)

	q := r.NewQuery().Limit(limit)

	_, err := r.GetAll(q, ss)

	return err
}
