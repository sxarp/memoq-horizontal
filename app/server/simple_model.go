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
}

func (s *Simple) SetKind(r Repo) {
	r.SetKind("Simple")

}

func (s *Simple) Save(r Repo) (*datastore.Key, error) {
	s.SetKind(r)

	key, err := r.Create(s)
	return key, err
}

func (s *Simple) Find(r Repo, k *datastore.Key) error {
	s.SetKind(r)

	return r.Get(k, s)
}

func (s *Simple) Destroy(r Repo, k *datastore.Key) error {
	s.SetKind(r)

	return r.Delete(k)
}

type Simples []Simple

func (ss *Simples) AllWithLimit(r Repo, limit int) error {
	(&Simple{}).SetKind(r)

	q := r.NewQuery().Limit(limit)

	_, err := r.GetAll(q, ss)

	return err
}
