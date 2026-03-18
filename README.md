## Features du langage

L'objectif du langage est de répondre à deux problématiques :

- la validation des données entre les différentes fonctions
- la réduction de la quantité de dépendances nécessaires pour un projet

### Validation des données

J'aimerais que les fonctions pusisent déclarer des genres d'attributs indiquant quelle validation doit être effectuées pour un paramètre donné avant de le passer à la fonction.

Exemple de syntaxe :

```
valid LengthGreaterThan(val string|error, length int): string|error {
    return match val.(type) {
        error => val,
        string => String.length(val) > length ? val : Error.new('Example error messsage')
    }
}

@param email > LengthGreaterThan(5) > EmailFormat
fn getUserByEmail(email string): User {
    // récupère utilisateur
}
```

Si cette fonction est appelée en passant un "email" sous forme de chaine de caractère litérale (ex: getUserByEmail('mail@test.com')), alors la suite de fonctions de validation définie dans le "@param" est executée automatiquement avant le lancement de la fonction. 

En gros, ça donnerait ça une fois "compilé":
 ```
 // Code source
 getUserByEmail('mail@test.com');
 
 // Code une fois compilé
 val = EmailFormat(LengthGreaterThan('mail@test.com'))
 if val.(Error) {
    return val
 }
 getUserByEmail(val)
 ```
 
 ### Réduction de la taille des dépendances
 
 Le gros problème (à mon sens) avec les dépendances notamment quand il y en a beaucoup est que certaines fonctions définies dans une dépendances pourrraient être utilisées pour d'autres dépendances.
 
 Par exemple, si une librairie définit cette fonction en Go :
 
 ```
 package lib1
 
 func Add(a int, b int): int {
    retrun a + b
 }
 ```
 
 Et qu'une autre librairie définit cette fonction :
 
 ```
 package lib2
 
 func AddFloat(a float, b float): float {
    return a + b
 }
 ```
 
 Alors techniquement, la fonction "AddFloat" pourrait remplacer "Add" de "lib1" en convertissant uniquement les int en float lors de l'appel de la fonction et en convertissant le "float" en "int" pour la valeur de retour.
 
 De manière plus globale, l'idée est que si une fonction définit une liste de tests et qu'une autre fonction d'un autre package peut aussi valider cette suite de tests, alors autant utiliser cette seconde fonction et supprimer la première.
 
 Cela incite les librairies à créer des tests pour leur fonctions et cette fonctionnalités serait uniquement une optimisation optionnelle.
