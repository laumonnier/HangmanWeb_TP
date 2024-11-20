I. C'EST quoi Hangman-web ?  
Le Hangman en français est un "jeu du pendu" 
HANGMAN est un jeu qui combine réflexion, logique et suspense, offrant aux joueurs une expérience ludique et stimulante.  
Le PRINCIPE du jeu est très simple : les joueurs doivent deviner un mot en proposant les bonnes lettres avant d'épuiser toutes les tentatives.  
"Hangman Web" est une version en ligne du jeu classique du pendu. Il permet de jouer au pendu directement depuis un navigateur internet, souvent avec une interface plus agréable et interactive que l'interface console traditionnelle.
 

II. CE qu'il y a dans le projet :  
Le projet contient plusieurs modules qui permettent de gérer les différentes parties du jeu. 
Dans la structure du code : 

1. MAIN est la fonction principale (squelette) donc la fonction qui fait appel à tous les autres fonction pour faire fonctionner le jeu. 

2. GAME est la fonction chargée de la gestion du jeu au complet dont la gestion des tentatives, de la vérification de victoire ou de défaite, du choix des mots aléatoires et l'affichage de l'état actuel du jeu. 

3. INPUTVALIDATION est la fonction qui gère toutes les erreurs. 

4. WORD.TXT est le fichier contenant la liste des mots à deviner, un mot par ligne. 

5. HANGMAN.TXT contient les différentes étapes de dessin du pendu pour afficher au fur et à mesure les erreurs du joueur. 

6. PARTICULARITÉ DU HANGMAN-WEB / UTILISATION DU HTML - CSS : 
    L'HTML établit la structure et le contenu de la page, tandis que CSS s'assure que tout est beau et bien aligné. Ensemble, ils créent une interface utilisateur agréable et fonctionnelle pour le jeu Hangman Web.

        1. HTML (Hypertext Markup Language) :

Structure : HTML fournit la structure de base de la page web. Il définit les éléments tels que les titres, les paragraphes, les listes, les boutons, les champs de texte et les images.

Contenu : Il inclut le contenu statique de la page, comme le texte à afficher, les liens vers les ressources externes (comme les images et les styles CSS), et les sections du jeu (par exemple, la zone où les lettres sont affichées et l'espace pour le dessin du pendu).

        2. CSS :

Style et Apparence : CSS est utilisé pour appliquer des styles au HTML. Il contrôle la présentation visuelle du jeu, y compris les couleurs, les polices, les marges, les espacements et les alignements.

Disposition : Il aide à organiser la mise en page du jeu, en définissant comment les différents éléments sont positionnés et alignés les uns par rapport aux autres.

Réactivité : CSS permet aussi de rendre l'interface réactive, c'est-à-dire adaptée aux différentes tailles d'écran (ordinateur de bureau, tablette, smartphone), pour une expérience utilisateur optimale.


III. COMMENT fonctionne-t-il ?  
le jeu fonctionne en suivant une série de règles simples où le joueur doit deviner un mot en proposant des lettres, et ce avant de dépasser un nombre limité de tentatives incorrectes.  
Le jeu repose sur la capacité du joueur à deviner un mot en proposant des lettres une par une. 
Chaque mauvaise lettre entraîne une réduction de tentatives restantes et complète le dessin du pendu. 
L'objectif est de deviner toutes les lettres avant que le dessin du pendu soit terminé. 
Interface utilisateur : Vous avez une zone où les lettres du mot à deviner apparaissent sous forme de tirets. À côté, il y a l'image du pendu qui se dessine au fur et à mesure.

Voici quelques principes de fonctiionnement : 

Entrée utilisateur : Vous proposez des lettres soit en cliquant sur des boutons, soit en tapant sur votre clavier.
Vérification des lettres : Si la lettre est correcte, elle apparaît à sa place dans le mot. Si elle est incorrecte, une partie du dessin du pendu est ajoutée.
Progression du jeu : Le jeu continue jusqu'à ce que vous ayez deviné le mot complet ou que le dessin du pendu soit terminé.
Fin de la partie : Vous gagnez si vous devinez le mot avant que le dessin du pendu soit complet. Sinon, la partie se termine par une défaite.
