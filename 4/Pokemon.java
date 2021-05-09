public class Pokemon {

    public IAttackType attackType;

    public Pokemon(String type) {
        this.type = type;
        attackType = createAttackType(this.type);
    }

    public IAttackType createAttackType(String type) {
        if(type.Equals("fire")) {
            return new FireAttack(this);
        } else if(type.Equals("water")) {
            return new WaterAttack(this);
        } else if(type.Equals("earth")) {
            return new EarthAttack(this);
        }
    }

    public int level = 1;
    public int points = 20;
    public String type;

    public void fight(Pokemon anotherPokemon) {
        while (points > 0 && anotherPokemon.points > 0) {
            attackType.attack(anotherPokemon);
            if (anotherPokemon.points > 0) {
                IAttackType at2 = createAttackType(anotherPokemon.type); 
                at2.attack(this);
            }
        }
        if (anotherPokemon.points <= 0) {
            level += 1;
        } else {
            anotherPokemon.level += 1;
        }
    }
}