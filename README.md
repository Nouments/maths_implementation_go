
# Projet Maths & Algorithmes en Go

## 🚀 Présentation générale

Ce projet est une **bibliothèque modulaire** d'implémentations mathématiques et algorithmiques en Go, structurée par grands domaines :

- **Algèbre** : théorie des ensembles, espaces vectoriels, ...
- **Analyse** : équations différentielles, trigonométrie, ...
- **Arithmétique** : (à venir)
- **Autres modules** : extensions futures prévues

L’objectif est d’apprendre et d’expérimenter la modélisation mathématique en Go, en créant des modules indépendants, clairs, et réutilisables.

---

## 📂 Structure actuelle

.
├── algebre
│ ├── ensemble
│ │ └── ensemble.go # Opérations sur les ensembles (union, intersection, etc.)
│ └── espace_vectorielle
│ └── espace_vectorielle.go # Manipulation des espaces vectoriels
├── analyse
│ ├── equadiff
│ │ └── equadiff.go # Résolution équations différentielles
│ └── trigo
│ └── trigo.go # Fonctions trigonométriques
├── arithmetique # Module en préparation
├── go.mod # Gestionnaire de modules Go
└── main.go # Point d’entrée du projet (exemples/tests)

yaml
Copier
Modifier

---

## ✅ Modules fonctionnels actuels

- `algebre/ensemble` : manipulation d’ensembles mathématiques (union, intersection, complémentaire, différence)
- `algebre/espace_vectorielle` : opérations sur espaces vectoriels (à compléter)
- `analyse/equadiff` : résolution d’équations différentielles (prototype)
- `analyse/trigo` : fonctions trigonométriques de base

---

## 🎯 Objectifs & évolutions

- Finaliser et enrichir chaque module avec des fonctions mathématiques avancées
- Ajouter des modules d’arithmétique et autres branches des mathématiques
- Implémenter des tests unitaires automatisés pour garantir la fiabilité
- Proposer une interface commune pour faciliter l’utilisation des modules
- Eventuellement créer une interface CLI ou web pour interagir avec les fonctions

---

## 📚 Pourquoi ce projet ?

- Approfondir la compréhension des concepts mathématiques par la mise en œuvre algorithmique
- Maîtriser les bonnes pratiques de développement Go dans un contexte scientifique
- Bâtir une base modulaire évolutive utilisable dans divers projets

---

## 🛠️ Utilisation rapide

1. Cloner le projet
2. Compiler ou lancer `main.go` pour tester les fonctionnalités
3. Ajouter/modifier les modules selon les besoins

---

N’hésite pas à documenter chaque module au fur et à mesure, ainsi que les fonctions, pour garder une base claire et maintenable.

---
