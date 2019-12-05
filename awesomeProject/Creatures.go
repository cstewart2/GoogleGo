/*
Program used to display how you can force Go to do object oriented programming
 */
package main

import (
	"fmt"
	"math/rand"
)

type Creature struct{//Creature Struct
	Name string
	class string
	health int
	attack string
	strength int

}
type ShadowMan struct{//ShadowMan Struct Extends Creature
	Creature
}
type Minotaur struct{//Minotaur Struct Extends Creature
	Creature
}
func main(){
	fmt.Println("Create Fighter 1. Select a type:")
	fmt.Println("Creature: standard creature.\nHealth: 30 Damage: 15\n Enter default")
	fmt.Println("Shadow Monster: a creature with the ability to turn transparent.\nHealth: 20 Damage: 10\nEnter shadow")
	fmt.Println("Minotaur: a hulking beast that is half man half bull.\nHealth: 60 Damage: 30\nEnter minotaur")
	var text string
	_, _ = fmt.Scanf("%s", &text)
	fmt.Println("Enter creature name:")
	var text2 string
	_, _ = fmt.Scanf("%s", &text2)

	var creature = Creature{}
	var creature2 = Creature{}
	if text == "shadow"{
		creature = NewShadowMan(text2)
	}else if text == "minotaur"{
		creature = NewMinotaur(text2)
	}else{
		creature = NewCreature(text2)
	}

	fmt.Println("Create Fighter 2. Select a type:")
	fmt.Println("Creature: standard creature.\nHealth: 30 Damage: 15\n Enter default")
	fmt.Println("Shadow Monster: a creature with the ability to turn transparent.\nHealth: 20 Damage: 10\nEnter shadow")
	fmt.Println("Minotaur: a hulking beast that is half man half bull.\nHealth: 60 Damage: 30\nEnter minotaur")
	var text3 string
	_, _ = fmt.Scanf("%s", &text3)
	fmt.Println("Enter creature name:")
	var text4 string
	_, _ = fmt.Scanf("%s", &text2)

	if text3 == "shadow"{
		creature2 = NewShadowMan(text4)
	}else if text3 == "minotaur"{
		creature2 = NewMinotaur(text4)
	}else{
		creature2 = NewCreature(text4)
	}
	Attack(Creature(creature), Creature(creature2))

}
func NewCreature(name string) Creature {//Creature Constructor
	something := Creature{}
	something = getName(something, name)
	something.class = "creature"
	something.health = 30
	something.strength = 15
	something.attack = something.Name +" strikes target."
	return something
}
func getName(creature Creature, name string) Creature {
	creature.Name = name
	return creature
}
func NewShadowMan(name string) Creature {//ShadowMan Constructor Extends Creature
	something := NewCreature(name)
	something.class = "shadowman"
	something.strength = 10
	something.health = 20
	something.attack = something.Name + " surrounds target in a veil of shadows. "
	return something
}
func NewMinotaur(name string) Creature{//Minotaur Constructor Extends Creature
	something := NewCreature(name)
	something.class = "minotaur"
	something.strength = 30
	something.health = 60
	something.attack = something.Name + " charges at target."
	return something
}
func Attack(attacker Creature, defender Creature){
	fmt.Println(attacker.Name + " attempts to attack " + defender.Name)
	TakeDamage(defender, attacker.strength, attacker)
}


func TakeDamage(defender Creature, damage int, attacker Creature) {
	var list[2] int
	list[0] = 1
	list[1] = 2
	rand.Shuffle(2,func(i int,j int){ list[i], list[j] = list[j], list[i] })
	var value = list[0]
	if defender.class == "shadowman"{
		if value == 1{
			defender.health = defender.health - damage
			fmt.Println(attacker.attack)
			if defender.health <= 0{
				fmt.Println(defender.Name + " is dead.")
			} else{
				fmt.Printf("%q 's Health: %b", defender.Name, defender.health )
			}
		}else{
			fmt.Println("The ShadowMan went transparent to avoid attack.")
		}
	}else{
		defender.health = defender.health -damage
		fmt.Println(attacker.attack)
		if defender.health <= 0{
			fmt.Println(defender.Name + " is dead.")
		} else{
			fmt.Printf("%q 's Health: %d", defender.Name, defender.health )
		}
	}

}

