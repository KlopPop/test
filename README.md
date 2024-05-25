
before_install:
  - go get github.com/mattn/goveralls
  - go get github.com/kelseyhightower/envconfig
  - go get github.com/go-playground/validator/v10
  - go get golang.org/x/exp/slog
  - go get github.com/go-chi/render