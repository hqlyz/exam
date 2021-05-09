public class WaterAttack implements IAttackType {
    public Pokemon pokemon;

    public WaterAttack(Pokemon p) {
        this.pokemon = p;
    }

    public void attack(Pokemon pokemonB) {
        if(pokemonB.type == "fire") {
            pokemonB.points -= this.pokemon.level * 5;
        } else if(pokemonB.type == "earth") {
            pokemonB.points -= this.pokemon.level + 2;
        } else if(pokemonB.type == "water") {
            pokemonB.points -= this.pokemon.level;
        }
    }
}