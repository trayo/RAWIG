/*
Copyright (c) 2018 Tomasz "VedVid" Nowakowski

This software is provided 'as-is', without any express or implied
warranty. In no event will the authors be held liable for any damages
arising from the use of this software.

Permission is granted to anyone to use this software for any purpose,
including commercial applications, and to alter it and redistribute it
freely, subject to the following restrictions:

1. The origin of this software must not be misrepresented; you must not
   claim that you wrote the original software. If you use this software
   in a product, an acknowledgment in the product documentation would be
   appreciated but is not required.
2. Altered source versions must be plainly marked as such, and must not be
   misrepresented as being the original software.
3. This notice may not be removed or altered from any source distribution.
*/

package main

import (
	"errors"
	"unicode/utf8"
)

const (
	//colors
	colorCreature     = "green"
	colorCreatureDark = "dark green"
)

const (
	//special characters
	CorpseChar = "%"
)

type Creature struct {
	/*Creatures are living objects that
	  moves, attacks, dies, etc.*/
	BasicProperties
	VisibilityProperties
	CollisionProperties
	AIProperties
	FighterProperties
}

/*Creatures holds every creature on map.*/
type Creatures []*Creature

func NewCreature(layer, x, y int, character, color, colorDark string,
	alwaysVisible, blocked, blocksSight bool, ai, hp, attack,
	defense int) (*Creature, error) {
	/*Function NewCreture takes all values necessary by its struct,
	and creates then returns pointer to Creature*/
	var err error
	if layer < 0 {
		txt := LayerError(layer)
		err = errors.New("Creature layer is smaller than 0." + txt)
	}
	if x < 0 || x >= WindowSizeX || y < 0 || y >= WindowSizeY {
		txt := CoordsError(x, y)
		err = errors.New("Creature coords is out of window range." + txt)
	}
	if utf8.RuneCountInString(character) != 1 {
		txt := CharacterLengthError(character)
		err = errors.New("Creature character string length is not equal to 1." + txt)
	}
	if hp < 0 {
		txt := InitialHPError(hp)
		err = errors.New("Creature HPMax is smaller than 0." + txt)
	}
	if attack < 0 {
		txt := InitialAttackError(attack)
		err = errors.New("Creature attack value is smaller than 0." + txt)
	}
	if defense < 0 {
		txt := InitialDefenseError(defense)
		err = errors.New("Creature defense value is smaller than 0." + txt)
	}
	creatureBasicProperties := BasicProperties{layer, x, y, character, color,
		colorDark}
	creatureVisibilityPropeties := VisibilityProperties{alwaysVisible}
	creatureCollisionProperties := CollisionProperties{blocked, blocksSight}
	creatureAIProperties := AIProperties{ai}
	creatureFighterProperties := FighterProperties{hp, hp, attack, defense}
	creatureNew := &Creature{creatureBasicProperties,
		creatureVisibilityPropeties, creatureCollisionProperties,
		creatureAIProperties, creatureFighterProperties}
	return creatureNew, err
}

func (c *Creature) MoveOrAttack(tx, ty int, b Board, all Creatures) bool {
	/*Method MoveOrAttack decides if Creature will move or attack other Creature;
	It has *Creature receiver, and takes tx, ty (coords) integers as arguments,
	and map of current level, and list of all Creatures.
	Starts by target that is nil, then iterates through Creatures. If there is
	Creature on targeted tile, that Creature becomes new target for attack.
	Otherwise, Creature moves to specified Tile.
	It's supposed to take player as receiver (attack / moving enemies is
	handled differently - check ai.go and combat.go).*/
	var target *Creature
	var turnSpent bool
	for i, _ := range all {
		if all[i].X == c.X+tx && all[i].Y == c.Y+ty {
			if all[i].HPCurrent > 0 {
				target = all[i]
				break
			}
		}
	}
	if target != nil {
		c.AttackTarget(target)
		turnSpent = true
	} else {
		turnSpent = c.Move(tx, ty, b)
	}
	return turnSpent
}

func (c *Creature) Move(tx, ty int, b Board) bool {
	/*Move is method of Creature; it takes target x, y as arguments;
	check if next move won't put Creature off the screen, then updates
	Creature coords.*/
	turnSpent := false
	newX, newY := c.X+tx, c.Y+ty
	if newX >= 0 &&
		newX <= WindowSizeX-1 &&
		newY >= 0 &&
		newY <= WindowSizeX-1 {
		if b[newX][newY].Blocked == false {
			c.X = newX
			c.Y = newY
			turnSpent = true
		}
	}
	return turnSpent
}

func (c *Creature) Die() {
	/*Method Die is called, when Creature's HP drops below zero.
	Die() has *Creature as receiver.
	Receiver properties changes to fit better to corpse.*/
	c.Layer = DeadLayer
	c.Color = "dark red"
	c.ColorDark = "dark red"
	c.Char = CorpseChar
	c.Blocked = false
	c.BlocksSight = false
	c.AIType = NoAI
}
