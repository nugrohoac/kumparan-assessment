// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	kumparan_assessment "github.com/nugrohoac/kumparan-assessment"
	mock "github.com/stretchr/testify/mock"
)

// ArticleService is an autogenerated mock type for the ArticleService type
type ArticleService struct {
	mock.Mock
}

// Fetch provides a mock function with given fields: ctx, filter
func (_m *ArticleService) Fetch(ctx context.Context, filter kumparan_assessment.ArticleFilter) ([]kumparan_assessment.Article, error) {
	ret := _m.Called(ctx, filter)

	var r0 []kumparan_assessment.Article
	if rf, ok := ret.Get(0).(func(context.Context, kumparan_assessment.ArticleFilter) []kumparan_assessment.Article); ok {
		r0 = rf(ctx, filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]kumparan_assessment.Article)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kumparan_assessment.ArticleFilter) error); ok {
		r1 = rf(ctx, filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, article
func (_m *ArticleService) Store(ctx context.Context, article kumparan_assessment.Article) (kumparan_assessment.Article, error) {
	ret := _m.Called(ctx, article)

	var r0 kumparan_assessment.Article
	if rf, ok := ret.Get(0).(func(context.Context, kumparan_assessment.Article) kumparan_assessment.Article); ok {
		r0 = rf(ctx, article)
	} else {
		r0 = ret.Get(0).(kumparan_assessment.Article)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, kumparan_assessment.Article) error); ok {
		r1 = rf(ctx, article)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
