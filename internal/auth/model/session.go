package model

type Session struct {
    Name string
    Email string

    IssuedAt int
    ExpiresAt int
}

func NewSession(name, email string, issuedAt, expiresAt int) (*Session, error) {
    session := Session{
        Name: name,
        Email: email,
        IssuedAt: issuedAt,
        ExpiresAt: expiresAt,
    }

    return &session, nil
}

