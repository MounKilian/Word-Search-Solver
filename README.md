## Word Scrabble Solver 🔎

Bienvenue dans Word Scrabble Solver, ce programme va vous permettre de trouver des mots contenus dans un tableau de lettre ! 🔤

# Prérequis :rewind:

Des connaissances dans les domaines ci dessous 🧠 : 

- Utilisation de Golang :page_facing_up:
- Utilisation de Git et GitHub pour la gestion de code :memo:

# Installation :wrench:

1. Commencez par cloner ce repository.
```bash
  git clone https://ytrack.learn.ynov.com/git/mkilian/word-scrabble-solver.git
```
2.  Accédez au répertoire word-scrabble-solver.
```bash
  cd word-scrabble-solver
```

3. Importez la bibliothèque github.
```bash
  go get github.com/01-edu/z01
```

# Démarrage :technologist:

1. Il suffit simplement de lancer la commande suivante : 
```bash
  go run main/main.go
```
2. Vous pouvez changer le tableau autant que vous voulez, les mots qui peuvent êtres reconnus sont reférencés dans le fichier words.txt

# Fabrication :hammer:

1. solver.go permet de faire le lien entre toutes les fonctions de recherches et les prints

2. horizontal.go permet de chercher tout les mots horizontalement dans le tableau a deux dimensions

3. vertical.go permet de chercher tout les mots verticalement dans le tableau a deux dimensions

4. diagonalDown.go et diagonalUp.gp permettent de chercher tout les mots en diagonale dans le tableau a deux dimensions

# Version :card_file_box:

golang 1.21.0

# Auteurs :money_with_wings:

1. Moun Kilian :beers:

2. Marciniak Lucas :see_no_evil: