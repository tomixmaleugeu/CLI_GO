# F1 CLI — Rechercher un pilote par son numéro

Programme en ligne de commande en Go : on saisit un numéro de pilote, le
programme interroge la [F1 Companion API](https://f1-companion-api-ba5k.onrender.com/drivers)
et affiche ses informations.

## Lancement

```bash
go run .
```

> Important : lance `go run .` (avec le point), pas `go run main.go`. Le projet
> est découpé en plusieurs fichiers : `main.go` tout seul ne contient pas les
> fonctions des autres fichiers, d'où l'erreur `undefined: ...`.

Pour créer un exécutable :

```bash
go build -o f1cli.exe .
./f1cli.exe
```

Puis on saisit un numéro (`44`), ou `all` pour voir toute la grille.

## Exemples

```
$ go run .
Entrez un numéro de pilote : 44

--- Pilote trouvé ---
Numéro      : 44
Nom         : Lewis Hamilton
Nationalité : British
Code        : HAM
```

```
$ go run .
Entrez un numéro de pilote : 99

Aucun pilote ne porte le numéro 99.
```

```
$ go run .
Entrez un numéro de pilote : abc
Erreur : "abc" n'est pas un numéro valide.
```

## Organisation des fichiers

| Fichier | Rôle |
| --- | --- |
| `main.go` | orchestration : saisie → validation → API → recherche → affichage |
| `saisie.go` | lecture de la saisie (`bufio.Scanner`) et validation (`all` ou entier > 0) |
| `api.go` | requête `GET /drivers` (timeout 30 s, contrôle du statut 200, décodage JSON) |
| `pilote.go` | structures `Pilote` / `Reponse` et recherche du pilote par numéro |
| `affichage.go` | affichage d'un pilote ou de la liste complète |

Les erreurs (réseau, timeout, statut ≠ 200, JSON invalide, saisie invalide) sont
retournées par chaque fonction puis affichées dans `main` : le programme se
termine proprement.

## Difficultés rencontrées et solutions

- Le champ `number` est une chaîne (`"44"`) et non un nombre : je convertis la
  saisie avec `strconv.Itoa` puis je compare les deux chaînes.
- L'API ne sait pas chercher par numéro : je récupère toute la liste et je la
  filtre moi-même avec une boucle.
- Le premier appel est lent (hébergement gratuit) : j'ai mis un `http.Client`
  avec un timeout de 30 s.
- Une saisie comme `abc` ou `0` doit être refusée : `strconv.Atoi` renvoie une
  erreur et je vérifie aussi que le numéro est positif.
- `fmt.Scanln` posait problème avec les espaces : j'utilise `bufio.Scanner`, plus
  fiable pour lire une ligne entière.
