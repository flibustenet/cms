# CMS pas à pas

## 1. Initialisation du module

`go mod init github.com/flibustenet/cms`

création du fichier `main.go`

## 2. Lecture du fichier de configuration au format json

- Création d'une **struct** `Conf` dans le **package** `app`
- Lecture du fichier **json**
- Utilisation d'une **interface** `io.Reader`
- **Test**


## 3. Templates et rendu d'une page

- **Template**
- interface **`io.Writer`**
- **`bytes.Buffer`** **`strings`**

## 4. Serveur HTTP et handler

- Rendu des pages sur un **serveur HTTP**
- **`DefaultServeMux http.ListenAndServe`**
- **`map`**

