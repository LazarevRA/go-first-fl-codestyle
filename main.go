package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func rnd(min, max int) int {

	return rand.Intn(max-min) + min
}

func attack(charName, charClass string) string {

	var damage int

	switch charClass {
	case "warrior":
		damage = rnd(8, 10)
	case "mage":
		damage = rnd(10, 15)
	case "healer":
		damage = rnd(2, 4)
	default:
		return "неизвестный класс персонажа"
	}
	return fmt.Sprintf("%s нанес урон противнику равный %d.", charName, damage)
}

func defence(charName, charClass string) string {

	var dmgBlocked int

	switch charClass {
	case "warrior":
		dmgBlocked = rnd(15, 20)
	case "mage":
		dmgBlocked = rnd(8, 12)
	case "healer":
		dmgBlocked = rnd(12, 15)
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s блокировал %d урона.", charName, dmgBlocked)
}

func special(charName, charClass string) string {

	var spcAbility string
	var spcAbilityPower int

	switch charClass {
	case "warrior":
		spcAbility = "Выносливость"
		spcAbilityPower = 80 + 25
	case "mage":
		spcAbility = "Атака"
		spcAbilityPower = 5 + 40
	case "healer":
		spcAbility = "Защита"
		spcAbilityPower = 10 + 30
	default:
		return "неизвестный класс персонажа"
	}

	return fmt.Sprintf("%s применил специальное умение `%s %d`", charName, spcAbility, spcAbilityPower)
}

func startTraining(charName, charClass string) string {

	switch charClass {
	case "warrior":
		fmt.Printf("%s, ты Воитель - отличный боец ближнего боя.\n", charName)
	case "mage":
		fmt.Printf("%s, ты Маг - превосходный укротитель стихий.\n", charName)
	case "healer":
		fmt.Printf("%s, ты Лекарь - чародей, способный исцелять раны.\n", charName)
	default:
		return "неизвестный класс персонажа"
	}

	fmt.Println("Потренируйся управлять своими навыками.")
	fmt.Println("Введи одну из команд:")
	fmt.Println("attack — чтобы атаковать противника,")
	fmt.Println("defence — чтобы блокировать атаку противника,")
	fmt.Println("special — чтобы использовать свою суперсилу.")
	fmt.Println("Если не хочешь тренироваться, введи команду skip.")

	var cmd string
	for cmd != "skip" {
		fmt.Print("Введи команду: ")
		fmt.Scanf("%s\n", &cmd)
		switch {
		case cmd == "attack":
			fmt.Println(attack(charName, charClass))
		case cmd == "defence":
			fmt.Println(defence(charName, charClass))
		case cmd == "special":
			fmt.Println(special(charName, charClass))
		default:
			fmt.Println("неизвестная команда")
		}
	}
	return "тренировка окончена"
}

func choiseCharClass() string {

	var approveChoise string
	var charClass string

	for approveChoise != "y" {
		fmt.Print("Введи название персонажа, за которого хочешь играть: Воитель — warrior, Маг — mage, Лекарь — healer: ")
		fmt.Scanf("%s\n", &charClass)

		switch charClass {
		case "warrior":
			fmt.Println("Воитель — дерзкий воин ближнего боя. Сильный, выносливый и отважный.")
		case "mage":
			fmt.Println("Маг — находчивый воин дальнего боя. Обладает высоким интеллектом.")
		case "healer":
			fmt.Println("Лекарь — могущественный заклинатель. Черпает силы из природы, веры и духов.")
		default:
			fmt.Println("неизвестный класс персонажа")
		}

		fmt.Print("Нажми (Y), чтобы подтвердить выбор, или любую другую кнопку, чтобы выбрать другого персонажа: ")
		fmt.Scanf("%s\n", &approveChoise)
		approveChoise = strings.ToLower(approveChoise)
	}
	return charClass
}

func main() {
	fmt.Println("Приветствую тебя, искатель приключений!")
	fmt.Println("Прежде чем начать игру...")

	var charName string
	fmt.Print("...назови себя: ")
	fmt.Scanf("%s\n", &charName)

	fmt.Printf("Здравствуй, %s\n", charName)
	fmt.Println("Сейчас твоя выносливость — 80, атака — 5 и защита — 10.")
	fmt.Println("Ты можешь выбрать один из трёх путей силы:")
	fmt.Println("Воитель, Маг, Лекарь")

	charClass := choiseCharClass()

	fmt.Println(startTraining(charName, charClass))
}
