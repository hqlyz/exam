public class FireAttack implements IAttackType {
    public Pokemon pokemon;

    public FireAttack(Pokemon p) {
        this.pokemon = p;
    }

    public void attack(Pokemon pokemonB) {
        if(pokemonB.type == "fire") {
            pokemonB.points -= this.pokemon.level + 1;
        } else if(pokemonB.type == "earth") {
            pokemonB.points -= this.pokemon.level * 3;
        } else if(pokemonB.type == "water") {
            pokemonB.points -= this.pokemon.level;
        }
    }
}