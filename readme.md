# CMS pas à pas

## Initialisation du module

`go mod init github.com/flibustenet/cms`

création du fichier `main.go`

## Lecture du fichier de configuration au format json

- Création d'une **struct** `Conf` dans le **package** `app`
- Lecture du fichier **json**
- Utilisation d'une **interface** `io.Reader`
- **Test**
