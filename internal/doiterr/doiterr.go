package doiterr

import "errors"

var ErrEmptyToken = errors.New("empty token in config")
var ErrEmptyDbURI = errors.New("empty db uri in config")
var ErrEmptyWebAppURL = errors.New("empty web app url in config")
var ErrEmptyParameters = errors.New("empty parameters")
var ErrEmptyEndpoint = errors.New("empty endpoint")
var ErrNotFound = errors.New("not found")
