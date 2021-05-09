public class EarthAttack implements IAttackType {
    public Pokemon pokemon;

    public EarthAttack(Pokemon p) {
        this.pokemon = p;
    }

    public void attack(Pokemon pokemonB) {
        if(pokemonB.type == "earth") {
            pokemonB.points -= this.pokemon.level;
        } else {
            pokemonB.points -= this.pokemon.level + 2;
        }
    }
}