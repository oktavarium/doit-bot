package doiterr

import "errors"

var ErrEmptyToken = errors.New("empty token in config")
var ErrEmptyWebAppURL = errors.New("empty web app url in config")
var ErrEmptyParameters = errors.New("empty parameters")
